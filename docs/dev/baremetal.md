# Bare metal

## What is kni-install?

kni-install is a forked version of openshift-install that servers as a
staging area for a new 'baremetal' platform that is intended to
eventually be merged into openshift-install itself.

The 'baremetal' platform support will ultimately be implemented using
libvirt for the bootstrap VM and the [MetalKube
baremetal-operator](https://github.com/metalkube/bare-metal-operator)
for the masters and workers. However, many hacky short-cuts may used
as the support it prototyped!

The [facet project](https://github.com/metalkube/facet) is a closely
related project that provides a "day 1" UI (served by a REST API)
which gathers information from the user about the cluster
configuration and bare metal hosts, before using kni-install to
provision the cluster.

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

The kni-installer rebasing branch regularly rebases to latest
openshift/installer whereas the master branch will never rebase and
only merges from openshift/installer.

So, the procedure is roughly:

FIXME: add instructions for rebasing any changes on kni-installer
master onto the rebasing branch, and then tidying up the patch
series.

- Checkout the rebasing branch and find the "Rename to kni-install"
  commit. Let's call that $RENAME_COMMIT.
- Fetch latest openshift/installer - let's say you've named its
  git remote "upstream". It's a good idea to familiarize yourself
  with the changes since your last rebase.
- Create a temporary branch from latest openshift/install and
  redo the rename:

```sh
git checkout -b tmp-rebase upstream/master
sed -i 's|openshift/installer|metalkube/kni-installer|' $(git grep -l openshift/installer | grep '\(cmd\|build.sh\|pkg\|assets_generate.go\)')
sed -i 's|openshift-install|kni-install|' $(git grep -l openshift-install | grep '\(cmd\|build.sh\|pkg\)')
git mv cmd/openshift-install cmd/kni-install
gofmt -w $(git grep -l github.com/metalkube/kni-installer)
TAGS=libvirt ./hack/build.sh
git commit -a -c $RENAME_COMMIT
```

- Now switch back to your rebasing branch and rebase all the
  changes onto this new commit.

```sh
git rebase --onto tmp-rebase $RENAME_COMMIT
TAGS=libvirt ./hack/build.sh
```

- Obviously, fix any merge conflicts. It's a good idea to check
  that individual commit builds.

FIXME: add a procedure for merging latest upstream/master into
kni-installer master, re-using the tree from the rebasing
branch.
