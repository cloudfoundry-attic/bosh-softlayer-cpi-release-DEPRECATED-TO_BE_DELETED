package action

import (
	instance "bosh-softlayer-cpi/softlayer/virtual_guest_service"
	bosherr "github.com/cloudfoundry/bosh-utils/errors"
)

type HasVM struct {
	vmService instance.Service
}

func NewHasVM(
	vmService instance.Service,
) HasVM {
	return HasVM{
		vmService: vmService,
	}
}

func (hv HasVM) Run(vmCID VMCID) (bool, error) {
	_, found, err := hv.vmService.Find(vmCID.Int())
	if err != nil {
		return false, bosherr.WrapErrorf(err, "Finding vm '%s'", vmCID)
	}

	return found, nil
}
