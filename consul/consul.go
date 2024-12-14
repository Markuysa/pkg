package consul

import (
	"os"

	consulapi "github.com/hashicorp/consul/api"
)

func RegisterService(cfg Config) error {
	config := consulapi.DefaultConfig()

	config.Address = cfg.Address

	consul, err := consulapi.NewClient(config)
	if err != nil {
		return err
	}

	if cfg.Address == "" {
		cfg.Address, err = os.Hostname()
		if err != nil {
			return err
		}
	}

	reg := &consulapi.AgentServiceRegistration{
		ID:      cfg.ServiceExtra.ServiceID,
		Name:    cfg.ServiceExtra.Name,
		Port:    cfg.ServiceExtra.ExposePort,
		Address: cfg.ServiceExtra.Address,
		Check: &consulapi.AgentServiceCheck{
			HTTP:     cfg.ServiceExtra.Probe.HealthCheck,
			Interval: cfg.ServiceExtra.Probe.Interval,
			Timeout:  cfg.ServiceExtra.Probe.Timeout,
		},
		Tags:      cfg.ServiceExtra.Tags,
		Namespace: cfg.ServiceExtra.Namespace,
	}

	regiErr := consul.Agent().ServiceRegister(reg)
	if regiErr != nil {
		return err
	}

	return err
}
