## Experimental `bosh-micro` usage

!!! `bosh-micro` CLI is still being worked on !!!

To start experimenting with bosh-warden-cpi release and new bosh-micro cli:

1. Create a deployment directory

```
mkdir my-micro
```

1. Create `manifest.yml` inside deployment direcrtory with following contents

```
cloud_provider:
  properties:
    cpi:
      host_ip: 10.254.50.4
      softlayer:
        username: your-softlayer-username@your-company.com
        api_key: your-softlayer-api-key
      agent:
        mbus: nats://nats:nats-password@10.254.50.4:4222
        blobstore:
          provider: dav
          options:
            endpoint: http://10.254.50.4:25251
            user: agent
            password: agent-password
```

1. Set deployment

```
bosh-micro deployment my-micro/manifest.yml
```

1. Kick off a deploy

```
bosh-micro deploy ~/Downloads/bosh-softlayer-cpi-?.tgz ~/Downloads/stemcell.tgz
```

Currently bosh-micro CLI does not anything after creating a stemcell in IaaS.
