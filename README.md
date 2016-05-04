# BOSH Softlayer CPI Release for BOSH V2

* Documentation: [bosh.io/docs](https://bosh.io/docs)
* IRC: [`#bosh` on freenode](https://webchat.freenode.net/?channels=bosh)
* Mailing list: [cf-bosh](https://lists.cloudfoundry.org/pipermail/cf-bosh)
* CI: <https://main.bosh-ci.cf-app.com/pipelines/bosh-openstack-cpi>
* Roadmap: [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/1344876)

This is a BOSH release for the Softlayer CPI.

## Bootstrap on SoftLayer

see [bosh-init-usage](docs/bosh-init-usage.md).

## Deployment Manifest Schema Sample:

see [cf-deployment-manifest](docs/cf_deployment_sl_sample.yml)

Using BOSH to deploy concousre

cloud config sample:

azs:
- name: lon02
  cloud_properties:
    Datacenter: { Name: lon02  }
    PrimaryNetworkComponent: { NetworkVlan: { Id: 524956  } }
    PrimaryBackendNetworkComponent: { NetworkVlan: { Id: 524954  } }

vm_types:
- name: compilation
  cloud_properties:
    Bosh_ip: 10.113.189.94
    StartCpus:  4
    MaxMemory:  8192
    EphemeralDiskSize: 100
    HourlyBillingFlag: true
    VmNamePrefix: concourse-compilation-worker-
- name: concourse-server
  cloud_properties:
    Bosh_ip: 10.113.189.94
    StartCpus:  4
    MaxMemory:  8192
    EphemeralDiskSize: 100
    HourlyBillingFlag: true
    VmNamePrefix: concourse-server-

vm_extensions:
- name: lon02-vlan-id
  cloud_properties:
    Datacenter: { Name: lon02  }
    PrimaryNetworkComponent: { NetworkVlan: { Id: 524956  } }
    PrimaryBackendNetworkComponent: { NetworkVlan: { Id: 524954  } }

disk_types:
- name: small
  disk_size: 50_000
- name: large
  disk_size: 500_000

networks:
- name: concourse
  type: dynamic
  subnets:
  - {az: lon02, dns: [10.113.189.94, 8.8.8.8, 10.0.80.11, 10.0.80.12]}

compilation:
  workers: 5
  reuse_compilation_vms: true
  az: lon02
  vm_type: compilation
  network: concourse

Concourse server deployment manifest:

---
name: concourse
director_uuid: 88b385f0-1d9c-44a0-83c8-4e5f559997ca

releases:
- name: concourse
  version: latest
- name: garden-linux
  version: latest

stemcells:
- alias: default
  name: light-bosh-stemcell-3169-softlayer-esxi-ubuntu-trusty-go_agent
  version: 3169

instance_groups:
- name: concourse
  instances: 1
  azs: [lon02]
  jobs:
  - name: atc
    release: concourse
    properties:
      basic_auth_username: &atc_username admin
      basic_auth_password: &atc_password c1oudc0w
      postgresql_database: &atc_db atc
      external_url: https://ci.foo.com
      publicly_viewable: true
  - name: postgresql
    release: concourse
    properties:
      databases:
      - name: *atc_db
        role: atc
        password: atc
  - name: tsa
    release: concourse
    properties:
      forward_host: 127.0.0.1
      atc:
        address: 127.0.0.1:8080
  - name: garden
    release: garden-linux
    properties:
      listen_network: tcp
      listen_address: 0.0.0.0:7777
      allow_host_access: true
      disk_quota_enabled: false
      log_level: debug
  - name: groundcrew
    release: concourse
    properties:
      tsa:
        host: 127.0.0.1
      garden:
        address: 127.0.0.1:7777
  vm_type: concourse-server
  stemcell: default
  persistent_disk_type: small
  networks:
  - name: concourse

update:
  canaries: 0
  canary_watch_time: 1000-60000
  update_watch_time: 1000-60000
  max_in_flight: 10
