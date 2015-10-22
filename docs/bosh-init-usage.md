## Experimental `bosh-init` usage

To start experimenting with bosh-softlayer-cpi release and new bosh-init cli:

1. Create a deployment directory

```
mkdir my-bosh-init
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
bosh-init deployment my-micro/manifest.yml
```

1. Kick off a deploy

```
bosh-init deploy ~/Downloads/bosh-softlayer-cpi-?.tgz ~/Downloads/stemcell.tgz
```
