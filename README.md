## bosh-softlayer-cpi-release

A [BOSH](https://github.com/cloudfoundry/bosh) release for the [`bosh-softlayer-cpi`](http://github.com/maximilien/bosh-softlayer-cpi) written in Go. This is based on [cppforlife](http://github.com/cppforlife)'s [bosh-warden-cpi-release](http://github.com/cppforlife/bosh-warden-cpi-release) project.

### Example SoftLayer environment

`bosh-softlayer-cpi` release can be deployed with any BOSH Director 
just like any other BOSH release.

1. Install Vagrant dependencies

```
vagrant plugin install vagrant-bosh
gem install bosh_cli --no-ri --no-rdoc
```

1. Create a new VM with BOSH Director and BOSH Warden CPI releases

```
vagrant up
```

Note: See [deployment manifest](manifests/softlayer-bosh.yml) 
to see how bosh and bosh-warden-cpi releases are collocated.

1. Target deployed BOSH Director

```
bosh target localhost:25555
bosh status
```

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
