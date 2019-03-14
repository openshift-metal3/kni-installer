#!/usr/bin/env bash
set -e


mkdir --parents /etc/keepalived

API_DNS="$(sudo awk -F[/:] '/apiServerURL/ {print $5}' /opt/openshift/manifests/cluster-infrastructure-02-config.yml)"
CLUSTER_DOMAIN="${API_DNS#*.}"
API_VIP="$(dig +noall +answer "$API_DNS" | awk '{print $NF}')"
IFACE_CIDRS="$(ip addr show | grep -v "scope host" | grep -Po 'inet \K[\d.]+/[\d.]+' | xargs)"
SUBNET_CIDR="$(/usr/local/bin/get_vip_subnet_cidr "$API_VIP" "$IFACE_CIDRS")"
DNS_VIP="$(/usr/local/bin/nthhost "$SUBNET_CIDR" 2)"
grep -v "${DNS_VIP}" /etc/resolv.conf | tee /etc/coredns/resolv.conf

COREDNS_IMAGE="quay.io/openshift-metalkube/coredns:latest"
if ! podman inspect "$COREDNS_IMAGE" &>/dev/null; then
    echo "Pulling release image..."
    podman pull "$COREDNS_IMAGE"
fi
MATCHES="$(sudo podman ps -a --format "{{.Names}}" | awk '/coredns$/ {print $0}')"
if [[ -z "$MATCHES" ]]; then
    /usr/bin/podman create \
        --name coredns \
        --volume /etc/coredns:/etc/coredns:z \
        --network host \
        --env CLUSTER_DOMAIN="$CLUSTER_DOMAIN" \
        "${COREDNS_IMAGE}" \
            --conf /etc/coredns/Corefile
fi
