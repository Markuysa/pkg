package tracer

type Config struct {
	// ServiceName is the name of the service.
	ServiceName string `validate:"required"`
	// URl is the URL of the Jaeger agent.
	URL  string `validate:"required"`
	Auth *Auth
}

type Auth struct {
	// Username is the username for the Jaeger agent.
	Username string
	// Password is the password for the Jaeger agent.
	Password string
}
