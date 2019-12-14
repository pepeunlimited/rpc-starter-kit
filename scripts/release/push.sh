#!/usr/bin/env bash

# This script pushes the docker images for the given version of Terraform,
# along with the "light", "full" and "latest" tags, up to docker hub.
#
# You must already be logged in to docker using "docker login" before running
# this script.

set -eu

VERSION="$1"
VERSION_SLUG="${VERSION#v}"

echo "-- Pushing tags $VERSION_SLUG and latest up to dockerhub --"
echo ""

docker push "pepeunlimited/grpc-starter-kit:$VERSION_SLUG"
docker push "pepeunlimited/grpc-starter-kit:latest"