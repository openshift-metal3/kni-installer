# Bare metal

## What is kni-install?

See the top-level [README file](../../README.md).

## Why a fork?

Forking is painful, but ...

- Bare metal provisioning support will be implemented as a platform in
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

## What's next?

### Bootstrap Ignition Customizations

We need to make bare metal specific bootstrap ignition customizations.
How best to do it? With per platform directories for assets?

### Provisioning Master Nodes

How do we prototype the Ironic based provisioning initially? With a
path that is completely different from the other platforms?

Would we move to terraform next as something more workable in the
medium term, since that's what other platforms do?

When we would switch to using baremetal-operator? Would all other
platforms have to switch from terraform at the same time, or could it
be bare metal specific initially?

### Using External Ironic

Users may use Ironic to discover and introspect nodes before launching
kni-install. Does it make sense for kni-install to reuse this rather
than re-running Ironic on the bootstrap VM.

## How to rebase?

The kni-installer `rebasing` branch regularly rebases to latest
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

### Update the rebasing branch

Identify the last upstream commit that we merged:

```sh
BASE_COMMIT=$(git merge-base upstream/master origin/master)
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

### Rebase to latest upstream

Still on the `rebasing` branch, identify the "Rename to kni-install"
commit:

```sh
RENAME_COMMIT=$(git log --oneline $BASE_COMMIT..HEAD | tail -1 | awk '{print $1}')
```

Fetch latest openshift/installer. It's a good idea to familiarize
yourself with the changes since the last rebase.

```sh
git fetch -t upstream
git log --graph $BASE_COMMIT..upstream/master
```

Create a temporary branch from latest openshift/install and redo the
rename commit:

```sh
git checkout -b tmp-rebase upstream/master
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
BASE_COMMIT=$(git merge-base upstream/master HEAD)
```

It's a good idea at this point to check that each commit in the series
builds:

```sh
git rebase -i $BASE_COMMIT
TAGS=libvirt ./hack/build.sh
```

### Merge latest upstream

With all the hard rebasing work done, we can now update the master
branch to incorporate this work as a merge commit!

```sh
git checkout master
git merge --ff-only origin/master
git merge upstream/master
```

This merge will more than likely fail with conflicts, but we can
resolve those by using the source tree from the `rebasing` branch.

```sh
git checkout rebasing .
git commit -a -m 'Merge latest openshift/installer'
git diff rebasing
```

There should be no differences between the `master` and `rebasing`
branches at this point, but watch out for weirdness caused by things
like merge conflicts on files that were deleted.

### Push your changes

Now create a new `rebasing` tag, push both branches to your personal
remote, and creating a PR:

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
