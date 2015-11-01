## Experimental `bosh-init` usage

At present, it is able to creat and update bosh environment on SoftLayer by bosh-init and bosh-softlayer-cpi

This document shows how to initialize new environment on Softlayer.

0. Create a deployment directory, install bosh-init

http://bosh.io/docs/install-bosh-init.html

```
mkdir my-bosh-init
```

1. Create a deployment manifest file named bosh.yml in the deployment directory based on the template below.

```
---
name: bosh

releases:
- name: bosh
  url: https://bosh.io/d/github.com/cloudfoundry/bosh?v=218
  sha1: 7ad794897468a453e81b018da53c8475a23cef2b
- name: bosh-softlayer-cpi
  url: file://./bosh-softlayer-cpi-0+dev.1.tgz #<--- Replace with SL eCPI release

resource_pools:
- name: vms
  network: default
  stemcell:
    url: file://./light-bosh-stemcell-3031-softlayer-esxi-ubuntu-trusty-go_agent.tgz # <--- Stemcell and bosh-agent should need some changes due to constrains of Softlayer
  cloud_properties:
    Domain: softlayer.com
    VmNamePrefix: bosh-experimental  # <--- It is better to use a catchy name which will be used in the following section
    StartCpus: 1
    MaxMemory: 1024
    Datacenter:
       Name: lon02
    HourlyBillingFlag: true
disk_pools:
- name: disks
  disk_size: 40_000
  cloud_properties:
    consistent_performance_iscsi: true


networks:
- name: default
  type: dynamic
  dns: [8.8.8.8] # <--- Replace with your DNS
  preconfigured: true


jobs:
- name: bosh
  instances: 1

  templates:
  - {name: nats, release: bosh}
  - {name: redis, release: bosh}
  - {name: postgres, release: bosh}
  - {name: blobstore, release: bosh}
  - {name: director, release: bosh}
  - {name: health_monitor, release: bosh}
  - {name: cpi, release: bosh-softlayer-cpi}

  resource_pool: vms
  persistent_disk_pool: disks

  networks:
  - name: default

  properties:
    nats:
      address: 127.0.0.1
      user: nats
      password: nats-password

    redis:
      listen_addresss: 127.0.0.1
      address: 127.0.0.1
      password: redis-password

    postgres: &db
      host: 127.0.0.1
      user: postgres
      password: postgres-password
      database: bosh
      adapter: postgres

    blobstore:
      address: 127.0.0.1
      port: 25250
      provider: dav
      director: {user: director, password: director-password}
      agent: {user: agent, password: agent-password}

    director:
      address: 127.0.0.1
      name: my-bosh
      db: *db
      cpi_job: cpi
      max_threads: 3

    hm:
      director_account: {user: admin, password: admin}
      resurrector_enabled: true

    softlayer: &softlayer
      username: fake-username # <--- Replace with username of your SL account
      apiKey: fake-api-key    # <--- Replace with password of your SL account
      public_vlan_id: fake-public-vlan   # <--- Replace with proper private vlan if needed
      private_vlan_id: fake-private-vlan # <--- Replace with proper private vlan if needed
      data_center: lon02

    cpi:
      agent: {mbus: "nats://nats:nats-password@127.0.0.1:4222"}

    ntp: &ntp []

cloud_provider:
  template: {name: cpi, release: bosh-softlayer-cpi}
  mbus: "https://admin:admin@bosh-experimental.softlayer.com:6868" # <--- Replace with VmNamePrefix + Domain indicated in cloud_properties of resource_pools section, as bosh-init does not support dynamic ip, it is only supporting static/floating ip, so we are using predined hostname in mbus. Please don`t use IP here.
  properties:
    softlayer: *softlayer
    cpi:
      agent:
        mbus: https://admin:admin@127.0.0.1:6868
        ntp: *ntp
        blobstore:
          provider: local
          options:
            blobstore_path: /var/vcap/micro_bosh/data/cache
```
3. Prepare a new bosh-agent and stemcell for eCPI
   please use the version of http://github.com/mattcui/bosh-agent bm_beta branch in softlayer stemcell
   agent.json should be as follows in softlayer stemcell

```    
{
  "Platform": {
    "Linux": {
      "CreatePartitionIfNoEphemeralDisk": false
    }
  },
  "Infrastructure": {
    "Settings": {
      "DevicePathResolutionType": "virtio",
      "NetworkingType": "manual",
      "Sources": [
        {
          "Type": "File",
          "SettingsPath": "/var/vcap/bosh/user_data.json"
        }
      ],
      "UseRegistry": true
    }
  }
} 
```
 create a directory /var/vcap/bosh/micro_bosh in stemcell for local blobstore of bosh-agent

4. Set deployment

```
bosh-init deployment bosh.yml
```

5. Kick off a deploy

```
bosh-init deploy bosh.yml
```
