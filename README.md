## bosh-softlayer-cpi-release

A [BOSH](https://github.com/cloudfoundry/bosh) release for `bosh-softlayer-cpi` written in Go. This is based on [cppforlife](http://github.com/cppforlife)'s [bosh-warden-cpi-release](http://github.com/cppforlife/bosh-warden-cpi-release) project.

### Example SoftLayer environment

`bosh-softlayer-cpi` release can be deployed with any BOSH Director 
just like any other BOSH release.


### Running tests

1. Follow instructions above to install the release to your BOSH director

1. Clone BOSH repository into `$HOME/workspace/bosh` to get BATS source code

1. Download SoftLayer stemcell #3 to `$HOME/Downloads/bosh-stemcell-softlayer-ubuntu-trusty-go_agent.tgz`
   from [BOSH Artifacts](https://s3.amazonaws.com/bosh-jenkins-artifacts/bosh-stemcell/softlayer/bosh-stemcell-softlayer-ubuntu-trusty-go_agent.tgz)

1. Run BOSH Acceptance Tests via `spec/run-bats.sh`


### Experimental `bosh-micro` usage

See [bosh-micro usage doc](docs/bosh-micro-usage.md)


### Todo

- Use standalone BATS and CPI lifecycle tests
- Use BATS errand for running tests
