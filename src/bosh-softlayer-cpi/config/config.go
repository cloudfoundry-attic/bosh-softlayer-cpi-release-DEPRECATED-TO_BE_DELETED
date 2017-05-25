package config

import (
	"encoding/json"

	bosherr "github.com/cloudfoundry/bosh-utils/errors"
	boshsys "github.com/cloudfoundry/bosh-utils/system"

	boslconfig "bosh-softlayer-cpi/softlayer/config"

	"bosh-softlayer-cpi/registry"
)

type Config struct {
	Cloud Cloud `json:"cloud"`
}

type Cloud struct {
	Plugin     string        `json:"plugin"`
	Properties CPIProperties `json:"properties"`
}

type CPIProperties struct {
	SoftLayer boslconfig.Config
	Agent     registry.AgentOptions
	Registry  registry.ClientOptions
}

func NewConfigFromPath(configFile string, fs boshsys.FileSystem) (Config, error) {
	var config Config

	if configFile == "" {
		return config, bosherr.Errorf("Must provide a config file")
	}

	bytes, err := fs.ReadFile(configFile)
	if err != nil {
		return config, bosherr.WrapErrorf(err, "Reading config file '%s'", configFile)
	}

	if err = json.Unmarshal(bytes, &config); err != nil {
		return config, bosherr.WrapError(err, "Unmarshalling config contents")
	}

	if err = config.Validate(); err != nil {
		return config, bosherr.WrapError(err, "Validating config")
	}

	return config, nil
}

func NewConfigFromString(configString string) (Config, error) {
	var config Config
	var err error
	if configString == "" {
		return config, bosherr.Errorf("Must provide a config")
	}

	if err = json.Unmarshal([]byte(configString), &config); err != nil {
		return config, bosherr.WrapError(err, "Unmarshalling config contents")
	}

	if err = config.Validate(); err != nil {
		return config, bosherr.WrapError(err, "Validating config")
	}

	return config, nil
}

func (c Config) Validate() error {
	if c.Cloud.Plugin != "softlayer" {
		return bosherr.Errorf("Unsupported cloud plugin type %q", c.Cloud.Plugin)
	}
	if err := c.Cloud.Properties.SoftLayer.Validate(); err != nil {
		return bosherr.WrapError(err, "Validating SoftLayer configuration")
	}
	if err := c.Cloud.Properties.Agent.Validate(); err != nil {
		return bosherr.WrapError(err, "Validating agent configuration")
	}
	//if err := c.Cloud.Properties.Registry.Validate(); err != nil {
	//	return bosherr.WrapError(err, "Validating registry configuration")
	//}

	return nil
}
