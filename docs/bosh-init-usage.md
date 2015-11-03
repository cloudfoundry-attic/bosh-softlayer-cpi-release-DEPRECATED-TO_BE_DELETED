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
  url: https://bosh.io/d/github.com/cloudfoundry/bosh?v=219
  sha1: bbd03790a2839aab26d3fa4cfe1493d361872f33
- name: bosh-softlayer-cpi
  url: file://./bosh-softlayer-cpi-0+dev.1.tgz

resource_pools:
- name: vms
  network: default
  stemcell:
    url: file://./light-bosh-stemcell-3031-softlayer-esxi-ubuntu-trusty-go_agent.tgz
  cloud_properties:
    Domain: SOFTLAYER.COM      # <--- domain name for new director
    VmNamePrefix: VMNAMEPREFIX # <--- indicate a hostname prefix for new director
    StartCpus: 4
    MaxMemory: 8192
    Datacenter:
       Name: lon02
    HourlyBillingFlag: true
    PrimaryNetworkComponent:
       NetworkVlan:
          Id: 524956
    PrimaryBackendNetworkComponent:
       NetworkVlan:
          Id: 524954
    NetworkComponents:
    - MaxSpeed: 1000

disk_pools:
- name: disks
  disk_size: 40_000
  cloud_properties:
    consistent_performance_iscsi: true

networks:
- name: default
  type: dynamic
  dns:
  - 8.8.8.8
  
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
  - {name: powerdns, release: bosh}
  - {name: cpi, release: bosh-softlayer-cpi}

  resource_pool: vms
  persistent_disk_pool: disks

  networks:
  - name: default

  properties:
    nats:
      user: nats
      password: nats
      auth_timeout: 3
      address: 127.0.0.1
      listen_address: 0.0.0.0
      port: 4222
      no_epoll: false
      no_kqueue: true
      ping_interval: 5
      ping_max_outstanding: 2
      http:
        port: 9222
    redis:
      address: 127.0.0.1
      password: redis
      port: 25255
      loglevel: info
    postgres: &20585760
      user: postgres
      password: postges
      host: 127.0.0.1
      database: bosh
      adapter: postgres
    blobstore:
      address: 127.0.0.1
      director:
        user: director
        password: director
      agent:
        user: agent
        password: agent
      port: 25250
      provider: dav
    director:
      cpi_job: cpi
      address: 127.0.0.1
      name: bosh
      db:
        adapter: postgres
        database: bosh
        host: 127.0.0.1
        password: postges
        user: postgres
    hm:
      http:
        user: hm
        password: hm
        port: 25923
      director_account:
        user: admin
        password: Cl0udyWeather
      intervals:
        log_stats: 300
        agent_timeout: 180
        rogue_agent_alert: 180
        prune_events: 30
        poll_director: 60
        poll_grace_period: 30
        analyze_agents: 60
      pagerduty_enabled: false
      resurrector_enabled: false
    dns:
      address: 127.0.0.1
      domain_name: microbosh
      db: *20585760
      webserver:
        port: 8081
        address: 0.0.0.0
    softlayer: &softlayer
      username: fake-username # <--- Replace with username
      apiKey: fake-api-key    # <--- Replace with api-key

    cpi:
      agent:
        mbus: nats://nats:nats@127.0.0.1:4222
        ntp: []
        blobstore:
          provider: dav
          options:
            endpoint: http://127.0.0.1:25250
            user: agent
            password: agent

    ntp: &ntp []

cloud_provider:
  template: {name: cpi, release: bosh-softlayer-cpi}
  mbus: "https://admin:admin@VMNAMEPREFIX.SOFTLAYER.COM:6868" # <--- Replace with a hostname of new director in advance

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
