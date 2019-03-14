#!/usr/bin/env bash
set -e

mkdir --parents /etc/keepalived

KEEPALIVED_IMAGE=quay.io/celebdor/keepalived:latest
if ! podman inspect "$KEEPALIVED_IMAGE" &>/dev/null; then
    echo "Pulling release image..."
    podman pull "$KEEPALIVED_IMAGE"
fi

API_DNS="$(sudo awk -F[/:] '/apiServerURL/ {print $5}' /opt/openshift/manifests/cluster-infrastructure-02-config.yml)"
API_VIP="$(dig +noall +answer "$API_DNS" | awk '{print $NF}')"
IFACE_CIDRS="$(ip addr show | grep -v "scope host" | grep -Po 'inet \K[\d.]+/[\d.]+' | xargs)"
SUBNET_CIDR="$(/usr/local/bin/get_vip_subnet_cidr "$API_VIP" "$IFACE_CIDRS")"
INTERFACE="$(ip -o addr show to "$SUBNET_CIDR" | awk '{print $2}')"
DNS_VIP="$(/usr/local/bin/nthhost "$SUBNET_CIDR" 2)"

export API_VIP
export INTERFACE
export DNS_VIP
envsubst < /etc/keepalived/keepalived.conf.tmpl | sudo tee /etc/keepalived/keepalived.conf

MATCHES="$(sudo podman ps -a --format "{{.Names}}" | awk '/keepalived$/ {print $0}')"
if [[ -z "$MATCHES" ]]; then
    # TODO(bnemec): Figure out how to run with less perms
    podman create \
            --name keepalived \
            --volume /etc/keepalived:/etc/keepalived:z \
            --network=host \
            --privileged \
            --cap-add=ALL \
            "${KEEPALIVED_IMAGE}" \
            /usr/sbin/keepalived -f /etc/keepalived/keepalived.conf \
                --dont-fork -D -l -P
fi
