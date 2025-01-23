#!/bin/bash

export GOFLAGS="-mod=vendor"
export GO111MODULE=""
export GOARCH="amd64"
export GOOS="linux"
export CGO_ENABLED="0"


go build -v -ldflags='-s' -o bin/terraform-provider-oci_v1.0.0

if [ $? != 0 ]; then
    exit $?
fi

mkdir -p $HOME/.terraform.d/plugins/terraform.local/local/oci/1.0.0/linux_amd64/

ln -sf $(pwd)/bin/terraform-provider-oci_v1.0.0 $HOME/.terraform.d/plugins/terraform.local/local/oci/1.0.0/linux_amd64/terraform-provider-oci_v1.0.0


rm -rf .terraform.lock.hcl
rm -rf .terraform

terraform init