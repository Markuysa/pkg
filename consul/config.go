package consul

type (
	Config struct {
		Address      string
		ServiceExtra ServiceExtra
	}
	ServiceExtra struct {
		ServiceID  string `default:"unknown"`
		Name       string `default:"unknown"`
		ExposePort int    `default:"8000"`
		Address    string
		Tags       []string
		Namespace  string `default:"default"`
		Probe      Probe
	}
	Probe struct {
		HealthCheck string
		Timeout     string `default:"30s"`
		Interval    string `default:"10s"`
	}
)
