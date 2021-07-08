package config

// PostgreSQL struct
type PostgreSQL struct {
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     int64  `mapstructure:"port" json:"port" yaml:"port"`
	Dbname   string `mapstructure:"dbname" json:"dbname" yaml:"dbname"`
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Sslmode  string `mapstructure:"sslmode" json:"sslmode" yaml:"sslmode"`
	Timezone string `mapstructure:"timezone" json:"timezone" yaml:"timezone"`
}
