package conf

type Config struct {
	System System `mapstructure:"system"`
	Mysql  Mysql  `mapstructure:"mysql"`
}

type System struct {
	Port int `mapstructure:"port"`
}

// Mysql mysql simple config
type Mysql struct {
	DSN string `mapstructure:"dsn"` // data source name.
}
