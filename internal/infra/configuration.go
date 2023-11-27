package infra

type SSLEnabled bool

func (c SSLEnabled) Get() string {
	if c {
		return "enable"
	}
	return "disable"
}

type ApplicationConfiguration struct {
	DatabaseConfiguration DatabaseConfiguration `yaml:"databaseConfiguration"`
	ServerConfiguration   `yaml:"serverConfiguration"`
}

type DatabaseConfiguration struct {
	DatabaseHost     string     `yaml:"databaseHost"`
	DatabasePort     int        `yaml:"databasePort"`
	DatabaseUser     string     `yaml:"databaseUser"`
	DatabasePassword string     `yaml:"databasePassword"`
	DatabaseName     string     `yaml:"databaseName"`
	SslMode          SSLEnabled `yaml:"sslEnabled"`
}

type ServerConfiguration struct {
	port string `yaml:"port"`
}
