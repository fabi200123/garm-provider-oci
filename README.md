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
availability_domain = "mQqX:US-ASHBURN-AD-2"
compartment_id = "ocid1.compartment.oc1...fsbq"
subnet_id = "ocid1.subnet.oc1.iad....feoplaka"
ngs_id = "ocid1.networksecuritygroup....pfzya"
tenancy_id = "ocid1.tenancy.oc1..aaaaaaaajds7tbqbvrcaiavm2uk34t7wke7jg75aemsacljymbjxcio227oq"
user_id = "ocid1.user.oc1...ug6l37u6a"
region = "us-ashburn-1"
fingerprint = "38...6f:bb"
private_key_path = "/home/ubuntu/.oci/private_key.pem"
private_key_password = ""
```

## Creating a pool

After you [add it to garm as an external provider](https://github.com/cloudbase/garm/blob/main/doc/providers.md#the-external-provider), you need to create a pool that uses it. Assuming you named your external provider as ```oci``` in the garm config, the following command should create a new pool:

```bash
garm-cli pool create \
    --os-type linux \
    --os-arch amd64 \
    --enabled=true \
    --flavor VM.Standard.E4.Flex \
    --image ocid1.image.oc1.iad.aaaaaaaah4rpzimrmnqfaxcm2xe3hdtegn4ukqje66rgouxakhvkaxer24oa \
    --min-idle-runners 0 \
    --repo 26ae13a1-13e9-47ec-92c9-1526084684cf \
    --tags oci,linux \
    --provider-name oci
```

```bash
garm-cli pool create \
    --os-type windows \
    --os-arch amd64 \
    --enabled=true \
    --flavor VM.Standard.E4.Flex \
    --image ocid1.image.oc1.iad.aaaaaaaamf7b6c6kvz2itjyflse6ibax2dgmqts2jlahl2zl3mbxlakv4h5a \
    --min-idle-runners 1 \
    --repo 26ae13a1-13e9-47ec-92c9-1526084684cf \
    --tags oci,windows \
    --provider-name oci
```