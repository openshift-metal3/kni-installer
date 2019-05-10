#!/usr/bin/env bash
set -e


mkdir --parents /etc/keepalived

API_DNS="$(sudo awk -F[/:] '/apiServerURL/ {print $5}' /opt/openshift/manifests/cluster-infrastructure-02-config.yml)"
CLUSTER_DOMAIN="${API_DNS#*.}"
read -d '.' -a CLUSTER_ARR <<< $CLUSTER_DOMAIN
CLUSTER_NAME=${CLUSTER_ARR[0]}
API_VIP="$(dig +noall +answer "$API_DNS" | awk '{print $NF}')"
IFACE_CIDRS="$(ip addr show | grep -v "scope host" | grep -Po 'inet \K[\d.]+/[\d.]+' | xargs)"
SUBNET_CIDR="$(/usr/local/bin/get_vip_subnet_cidr "$API_VIP" "$IFACE_CIDRS")"
DNS_VIP="$(dig +noall +answer "ns1.${CLUSTER_DOMAIN}" | awk '{print $NF}')"
grep -Ev "${DNS_VIP}|127.0.0.1" /etc/resolv.conf | tee /etc/coredns/resolv.conf
NUM_DNS_MEMBERS=$(grep -A 5 'controlPlane' /opt/openshift/manifests/cluster-config.yaml | awk '/replicas/ {print $2}')
export API_VIP CLUSTER_DOMAIN
envsubst < /etc/coredns/api-int.hosts.env > /etc/coredns/api-int.hosts

COREDNS_IMAGE="quay.io/openshift-metalkube/coredns-mdns:latest"
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
        --env CLUSTER_NAME="$CLUSTER_NAME" \
        --env NUM_DNS_MEMBERS="$NUM_DNS_MEMBERS" \
        "${COREDNS_IMAGE}" \
            --conf /etc/coredns/Corefile
fi
