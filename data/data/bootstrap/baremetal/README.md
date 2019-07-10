# Bare Metal IPI Bootstrap Assets

The `baremetal` platform (IPI for Bare Metal hosts) includes some additional
assets on the bootstrap node for automating some infrastructure requirements
that would have normally been handled by some cloud infrastructure service.
This document explains these pieces and what they accomplish.

## API failover from bootstrap to masters

One problem being addressed is that the installation process expects the API to
first be reachable on the bootstrap VM, but later in the installation process,
the API comes up on the masters that have been deployed.

In the `baremetal` platform, the failover of the API server is done by using a
VIP (Virtual IP) that has been configured as part of `install-config.yaml` and
then managed by `keepalived`.

The API VIP first resides on the bootstrap VM.  Once the master nodes come up,
the VIP will move to the masters.  This happens because the masters will be
running `keepalived` with a higher priority set in their `keepalived`
configuration for the API VIP.

Relevant files:
* **files/etc/keepalived/keepalived.conf.tmpl** - `keepalived` configuration
  template
* **files/usr/local/bin/keepalived.sh** - This script runs before `keepalived`
  starts and generates the `keepalived` configuration file from the template.
* **systemd/units/keepalived.service** - systemd unit file for `keepalived`.
  This runs `keepalived.sh` to generate the proper configuration from the
  template and then runs podman to launch `keepalived`.

## Internal DNS

TODO
