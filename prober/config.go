package prober

type Config struct {
	ReadinessPath string `default:"/ready"`
	LivenessPath  string `default:"/live"`
	Address       string `default:"0.0.0.0:8000"`
}
