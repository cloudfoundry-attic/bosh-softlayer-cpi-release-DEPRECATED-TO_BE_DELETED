# BOSH Softlayer CPI Release for BOSH V1

* Documentation: [bosh.io/docs](https://bosh.io/docs)
* BOSH Slack channel: [#bosh](https://cloudfoundry.slack.com/archives/bosh)
* BOSH SoftLayer CPI Slack channel: [#bosh-softlayer-cpi](https://cloudfoundry.slack.com/archives/bosh-softlayer-cpi)
* Mailing list: [cf-bosh](https://lists.cloudfoundry.org/pipermail/cf-bosh)
* CI: <http://159.122.236.143:8080/pipelines/bosh-softlayer-cpi-release-v1>
* Roadmap: [Pivotal Tracker](https://www.pivotaltracker.com/n/projects/1344876)

## Releases

This is a BOSH release dedicated for the Softlayer CPI Release V1: [[BOSH 236](https://s3.amazonaws.com/bosh-softlayer-cpi-stemcells/bosh-236%2Bdev.12.tgz)]

The latest version for the SoftlLayer CPI release is here. Also, find it on [bosh.io](http://bosh.io).

To use this CPI you will need to use the SoftLayer light stemcell. Latest version: [[3169](https://s3.amazonaws.com/bosh-softlayer-cpi-stemcells/light-bosh-stemcell-3169-softlayer-esxi-ubuntu-trusty-go_agent.tgz))]. Also, find it on [bosh.io](http://bosh.io).

## Bootstrap on SoftLayer

There is a bosh-init release specific for Softlayer CPI Release V1: [[bosh-init-0.0.81-linux-amd64-for-softlayer-cpi](https://s3.amazonaws.com/bosh-softlayer-cpi-stemcells/bosh-init-0.0.81-linux-amd64]

See [bosh-init-usage](docs/bosh-init-usage.md). Use the CPI and stemcells releases above to do so.

## Deployment Manifests Samples

For concourse, see [concourse-deployment-manifest](docs/concourse_sample_v1_schema.yml).
