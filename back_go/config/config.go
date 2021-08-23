package config

type HTTPConfig struct {
	// URL to bind
	URL string

	// fasthttp configuration
	ReadBufferSize     int
	WriteBufferSize    int
	ServerName         string
	ReadTimeout        int64
	WriteTimeout       int64
	MaxRequestBodySize int
	ReduceMemoryUsage  bool
	MaxConnsPerIP      int

	// debug and profiling
	Debug     bool
	Profiling bool
}

type DatabaseConfig struct {
	ConnectionString        string
	EnableMigration         bool
	MigrationsPath          string
	ConnectionMaxLifetimeMs int64
	MaxIdleConnections      int
	MaxOpenConnections      int
	Debug                   bool
}



type Config struct {
	HTTP        HTTPConfig
	Database    DatabaseConfig
}

