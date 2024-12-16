package consul

type (
	Config struct {
		Address      string       `yaml:"address"`
		ServiceExtra ServiceExtra `yaml:"service_extra"`
	}
	ServiceExtra struct {
		ServiceID  string   `default:"unknown" yaml:"service_id"`
		Name       string   `default:"unknown" yaml:"name"`
		ExposePort int      `default:"8000" yaml:"expose_port"`
		Address    string   `yaml:"address"`
		Tags       []string `yaml:"tags"`
		Namespace  string   `yaml:"namespace" default:"default"`
		Probe      Probe    `yaml:"probe"`
	}
	Probe struct {
		HealthCheck string `yaml:"health_check"`
		Timeout     string `default:"30s" yaml:"timeout"`
		Interval    string `default:"10s" yaml:"interval"`
	}
)
