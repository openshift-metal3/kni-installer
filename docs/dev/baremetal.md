# Bare metal

## What is kni-install?

See the top-level [README file](../../README.md).

## Why a fork?

Forking is painful, but ...

- Bare metal provisioning support is implemented as a platform in
  openshift-install, so openshift-install will become the interface
  that anything else (especially facet) will be built on. kni-install
  allows to prototype - and build against - that interface, and avoid
  making incorrect assumptions about what will be possible in future.
- A wrapper program/script, or a golang program that uses
  openshift/installer as a library, are possible alternative
  approaches. However, it would be difficult to emulate the interface
  of a baremetal platform for openshift-install.
- Since everything we're prototyping now will become patches against
  openshift/installer, it's quite beneficial for us to be working in
  that codebase from the start.

## What's done?

### Launching bootstrap node

The bootstrap node is launched via the libvirt terraform provider.

### Provisioning Master Nodes

Master nodes are deployed with Ironic using [a terraform provider](https://github.com/metalkube/terraform-provider-ironic),
so the baremetal platform works the same as the cloud-based providers. Users
need to provide information about the hardware nodes.  This is currently
provided in the format of ironic_nodes.json, which is flattened due to
the [limitations of terraform HCL](https://blog.gruntwork.io/terraform-tips-tricks-loops-if-statements-and-gotchas-f739bbae55f9?gi=1a2b92150144), and then converted to YAML.

This data is used by the Terraform templates to deploy the masters.

We could potentially switch to using the baremetal-operator, but would
all other platforms have to switch from terraform at the same time, or
could it be bare metal specific initially?

## What's next?

### How do we get hardware information?

Today the installer takes a blob of data. The installer should be more
user-friendly for working with the hardware, and operate in a more
interactive way. One proposal would be to query Ironic for discovered
baremetal nodes. Hardware would be powered on early and boot the Ironic
discovery image.  This also has the benefit of reducing the number of
reboots required during deployment and shortening deployment time.

When kni-installer launches, the installer will show hardware available
in Ironic's inventory.

The user can select which three become the masters, and will be prompted
to enter the BMC credentials for the nodes.  After ensuring the nodes
validate successfully in Ironic, the installer will create the templated
terraform configuration and deploy the masters.

The terraform provider would need to be refactored to support this,
either by allowing the node resource to consume an existing baremetal
node, or by separating the hardware defintion from the concept of a
deployment.  This is tracked by [this issue](https://github.com/openshift-metalkube/terraform-provider-ironic/issues/6).


### Bootstrap Ignition Customizations

We need to make bare metal specific bootstrap ignition customizations.
How best to do it? With per platform directories for assets?

### Using External Ironic

Users may use Ironic to discover and introspect nodes before launching
kni-install. Does it make sense for kni-install to reuse this rather
than re-running Ironic on the bootstrap VM? The terraform provider would
need changes to accomodate consuming existing nodes.

## How to rebase?

The kni-installer `rebasing` branch regularly rebases to releases of
openshift/installer whereas the master branch will never rebase and
only merges from openshift/installer.

### Prepare the repo

Checkout this repo and add `github.com/openshift/installer` as a
remote that we'll call `upstream`:

```sh
git clone git@github.com:openshift-metalkube/kni-installer.git
cd kni-installer
git remote add -f upstream git@github.com:openshift/installer.git
```

### Identify the openshift release commit

From the release payload, identify the commit that openshift/installer
was built from. Do not rely on the "built from commit" information from
the installer, but rather extract the release info as shown below:

```console
$ ./openshift-install version
./openshift-install v4.1.0-201905061832-dirty
built from commit 9b486eedf57c20ce446b25734afc1187259a599e
release image quay.io/openshift-release-dev/ocp-release@sha256:be61a9ec132118e41a417b347242361d9cc96b1a73753e121dc7a74a1905baea
$ oc adm release info -a <PULL SECRET> -o json quay.io/openshift-release-dev/ocp-release@sha256:be61a9ec132118e41a417b347242361d9cc96b1a73753e121dc7a74a1905baea | jq -r '.references.spec.tags[] | select(.name == "installer") | .annotations["io.openshift.build.commit.id"] '
7adf0ade1bedbb619d595927eaf8b25db1e19d02
$ RELEASE_COMMIT=7adf0ade1bedbb619d595927eaf8b25db1e19d02
```

### Update the rebasing branch

Identify the last upstream commit that we merged:

```sh
BASE_COMMIT=$(git merge-base $RELEASE_COMMIT origin/master)
```

Create a rebasing branch if you have not already done so.

```sh
git checkout -b rebasing origin/rebasing
```

Identify the changes that have been recently merged into `master`
which need to be incorporated into the `rebasing` branch. The
`rebasing` branch should always be tagged with a name that identifies
its corresponding commit on master, so:

```sh
MASTER_COMMIT=$(git describe --tags rebasing | sed 's/.*-//')
git checkout -b tmp-rebase origin/master
git rebase --onto rebasing $MASTER_COMMIT
git branch -M rebasing
```

Tidy up the `rebasing` branch into a coherent patch series:

```sh
git rebase -i $BASE_COMMIT
git diff origin/master         # confirm there are no differences
```

Force-push the updated `rebasing` branch (and its corresponding tag)
which matches the latest `master` branch:

```sh
git tag rebasing-$(date -I)-$(git rev-parse --short origin/master) rebasing
git push origin tag rebasing-$(date -I)-$(git rev-parse --short origin/master)
git push origin +rebasing:rebasing
```

### Rebase to an openshift release

Still on the `rebasing` branch, identify the "Rename to kni-install"
commit:

```sh
RENAME_COMMIT=$(git log --oneline $BASE_COMMIT..HEAD | tail -1 | awk '{print $1}')
```

Fetch latest openshift/installer. It's a good idea to familiarize
yourself with the changes since the last rebase.

```sh
git fetch -t upstream
git log --graph $BASE_COMMIT..$RELEASE_COMMIT
```

Create a temporary branch from the release commit and redo the
rename commit:

```sh
git checkout -b tmp-rebase $RELEASE_COMMIT
sed -i 's|openshift/installer|openshift-metalkube/kni-installer|g' $(git grep -l openshift/installer | grep '\(cmd\|build.sh\|pkg\|assets_generate.go\)')
sed -i 's|openshift-install|kni-install|g' $(git grep -l openshift-install | grep '\(cmd\|build.sh\|pkg\)')
git mv cmd/openshift-install cmd/kni-install
gofmt -w $(git grep -l github.com/openshift-metalkube/kni-installer)
TAGS=libvirt ./hack/build.sh
git commit -a -c $RENAME_COMMIT
```

Now switch back to your rebasing branch and rebase all the changes
onto this new commit. This is where you will have to carefully resolve
merge conflicts!

```sh
git checkout rebasing
git rebase --onto tmp-rebase $RENAME_COMMIT
git branch -D tmp-rebase
BASE_COMMIT=$(git merge-base $RELEASE_COMMIT HEAD)
```

It's a good idea at this point to check that each commit in the series
builds:

```sh
git rebase -i $BASE_COMMIT
# “edit” each commit

while TAGS=libvirt ./hack/build.sh ; do git rebase --continue || break ; done
# If all goes well, this will end on a clean build and no rebase left to continue.
```

### Merge release

With all the hard rebasing work done, we can now update the master
branch to incorporate this work as a merge commit!

```sh
git checkout master
git merge --ff-only origin/master
git merge $RELEASE_COMMIT
```

This merge will more than likely fail with conflicts, but we can
resolve those by using the source tree from the `rebasing` branch.

```sh
git checkout rebasing .
git commit -a -m 'Merge latest openshift/installer release'
git diff rebasing
```

There should be no differences between the `master` and `rebasing`
branches at this point, but watch out for weirdness caused by things
like merge conflicts on files that were deleted.

### Push your changes

Now create a new `rebasing` tag, push both branches to your personal
remote, and create a PR for the `latest-upstream` branch against
`master` of `openshift-metalkube/kni-installer`:

```
REBASING_TAG=rebasing-$(date -I)-$(git rev-parse --short master)
git tag $REBASING_TAG rebasing
git push $MYREMOTE tag $REBASING_TAG
git push $MYREMOTE +master:latest-upstream
git push $MYREMOTE +rebasing:latest-upstream-rebasing
```

When the PR has merged, you'll need to push the `rebasing` branch and
tag to the main repo:

```
git push origin +rebasing:rebasing
git push origin tag $REBASING_TAG
```

## Handling Conflicts

### Gopkg.lock

We have vendored some additional dependencies (namely,
terraform-provider-ironic).  This may result in a conflict in
Gopkg.lock.

It is best not to try to resolve those conflicts by hand.  Handle the
merge in Gopkg.toml, and then run `dep ensure` to create a correct
Gopkg.lock that includes all of the updates from openshift, as well as
our changes to vendoring.
