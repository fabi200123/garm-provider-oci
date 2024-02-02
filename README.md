# Garm External Provider For OCI

The OCI external provider allows [garm](https://github.com/cloudbase/garm) to create Linux and Windows runners on top of OCI virtual machines.

## Build

Clone the repo:

```bash
git clone https://github.com/cloudbase/garm-provider-oci
```

Build the binary:

```bash
cd garm-provider-oci
go build .
```

Copy the binary on the same system where garm is running, and [point to it in the config](https://github.com/cloudbase/garm/blob/main/doc/providers.md#the-external-provider).

## Configure

The config file for this external provider is a simple toml used to configure the OCI credentials it needs to spin up virtual machines.

```bash
tenancy_id = "sample_tenancy_id"
user_id = "sample_user_id"
region = "sample_region"
fingerprint = "sample_fingerprint"
private_key_path = "sample_private_key_path"
private_key_password = "sample_private_key_password"
```

## Creating a pool

After you [add it to garm as an external provider](https://github.com/cloudbase/garm/blob/main/doc/providers.md#the-external-provider), you need to create a pool that uses it. Assuming you named your external provider as ```oci``` in the garm config, the following command should create a new pool:

```bash
garm-cli pool create \
    --os-type windows \
    --os-arch amd64 \
    --enabled=true \
    --flavor  \
    --image  \
    --min-idle-runners 0 \
    --repo 5b4f2fb0-3485-45d6-a6b3-545bad933df3 \
    --tags oci,windows \
    --provider-name oci
```