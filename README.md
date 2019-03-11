# KNI Installer

kni-install is a forked version of
[openshift-install](https://github.com/openshift/installer) that
serves as a staging area for a new 'baremetal' platform that is
intended to eventually be merged into openshift-install itself.

The 'baremetal' platform support will ultimately be implemented using
libvirt for the bootstrap VM and the [MetalKube
baremetal-operator](https://github.com/metalkube/bare-metal-operator)
for the masters and workers. However, many hacky short-cuts may used
in the interim as the support it prototyped!

The [facet project](https://github.com/metalkube/facet) is a closely
related project that provides a "day 1" UI (served by a REST API)
which gathers information from the user about the cluster
configuration and bare metal hosts, before using kni-install to
provision the cluster.

## Supported Platforms

* [Bare-metal](docs/dev/baremetal.md)

## Quick Start

First, install all [build dependencies](docs/dev/dependencies.md).

Clone this repository to `src/github.com/metalkube/kni-installer` in your [GOPATH](https://golang.org/cmd/go/#hdr-GOPATH_environment_variable). Then build the `kni-install` binary with:

```sh
hack/build.sh
```

This will create `bin/kni-install`. This binary can then be invoked to create an OpenShift cluster, like so:

```sh
bin/kni-install create cluster
```

The installer will show a series of prompts for user-specific information and use reasonable defaults for everything else.
In non-interactive contexts, prompts can be bypassed by [providing an `install-config.yaml`](docs/user/overview.md#multiple-invocations).

If you have trouble, refer to [the troubleshooting guide](docs/user/troubleshooting.md).

### Connect to the cluster

Details for connecting to your new cluster are printed by the `kni-install` binary upon completion, and are also available in the `.openshift_install.log` file.

Example output:

```sh
INFO Waiting 10m0s for the openshift-console route to be created...
INFO Install complete!
INFO Run 'export KUBECONFIG=/path/to/auth/kubeconfig' to manage the cluster with 'oc', the OpenShift CLI.
INFO The cluster is ready when 'oc login -u kubeadmin -p 5char-5char-5char-5char' succeeds (wait a few minutes).
INFO Access the OpenShift web-console here: https://console-openshift-console.apps.${CLUSTER_NAME}.${BASE_DOMAIN}:6443
INFO Login to the console with user: kubeadmin, password: 5char-5char-5char-5char
```

### Cleanup

Destroy the cluster and release associated resources with:

```sh
kni-install destroy cluster
```

Note that you almost certainly also want to clean up the installer state files too, including `auth/`, `terraform.tfstate`, etc.
The best thing to do is always pass the `--dir` argument to `install` and `destroy`.
And if you want to reinstall from scratch, `rm -rf` the asset directory beforehand.
