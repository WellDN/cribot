package config

type PostgreSQLConfig struct {
    DSN     string
}

type SQLConfig struct {
    DSN     string
}

type Config struct {
    PostgreSQL PostgreSQLConfig

    SQL SQLConfig
}

type CribotConfig struct {
    PostgreSQLConfig
}

var defaultConfig = Config{
	PostgreSQL: PostgreSQLConfig{
		DSN: "host=/var/run/postgresql database=cribot sslmode=disable",
	},
}
