package config

// Server struct
type Server struct {
	JWT        JWT        `mapstructure:"jwt" json:"jwt" yaml:"jwt"`
	Redis      Redis      `mapstructure:"redis" json:"redis" yaml:"redis"`
	PostgreSQL PostgreSQL `mapstructure:"postgresql" json:"postgresql" yaml:"postgresql"`
}
