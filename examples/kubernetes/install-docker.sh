#!/bin/bash
# Install Docker on Debian 10 (Buster)

set -ex

DOCKER_VERSION="19.03.5"
CONTAINERD_VERSION="1.2.10"

# Use iptables legacy (Only for Debian 10)
#
update-alternatives --set iptables /usr/sbin/iptables-legacy
update-alternatives --set ip6tables /usr/sbin/ip6tables-legacy

apt update

# Requirements
#
apt-get install -y apt-transport-https ca-certificates curl software-properties-common gnupg

# Install Docker
#
curl -fsSL https://download.docker.com/linux/debian/gpg | apt-key add -
echo "deb https://download.docker.com/linux/$(. /etc/os-release; echo "$ID") $(lsb_release -cs) stable" > /etc/apt/sources.list.d/docker-ce.list
apt update
PKG_DOCKER_VERSION=$(apt-cache madison docker-ce | grep ${DOCKER_VERSION} | head -1 | awk '{print $3}')
PKG_CONTAINERD_VERSION=$(apt-cache madison containerd.io | grep ${CONTAINERD_VERSION} | head -1 | awk '{print $3}')
cat << EOF > /etc/apt/preferences.d/docker-ce.pref
# Pinning to prevent docker upgrades

Package: docker-ce docker-ce-cli
Pin: version ${PKG_DOCKER_VERSION}
Pin-Priority: 1001

Package: containerd.io
Pin: version ${PKG_CONTAINERD_VERSION}
Pin-Priority: 1001
EOF
apt-get install -y docker-ce

# Configure docker mirror
cat << EOF > /etc/docker/daemon.json
{
          "registry-mirrors": ["https://docker-mirror.rennes.grid5000.fr"],
          "insecure-registries" : ["harbor.rennes.grid5000.fr","docker-mirror.rennes.grid5000.fr"]
}
EOF

# Restart docker
systemctl restart docker

# Wait until docker is available
until [ -S /var/run/docker.sock ]; do sleep 1; done