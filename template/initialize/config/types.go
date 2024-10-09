package config

// Config 包含所有应用的配置项
type Config struct {
	Listen        int                 `mapstructure:"LISTEN"`
	MODE          string              `mapstructure:"MODE"`
	ROUTER_PREFIX string              `mapstructure:"ROUTER_PREFIX"`
	AllowedHosts  []string            `mapstructure:"ALLOWED_HOSTS"`
	SecretKey     string              `mapstructure:"SECRET_KEY"`
	Databases     map[string]DBConfig `mapstructure:"DATABASES"`
	Redis         RedisConfig         `mapstructure:"REDIS"`
}

// DBConfig 定义数据库配置
type DBConfig struct {
	Name     string `mapstructure:"NAME"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
	Host     string `mapstructure:"HOST"`
	Port     string `mapstructure:"PORT"`
	Config   string `mapstructure:"CONFIG"`
}

func (m *DBConfig) Dsn() string {
	return m.User + ":" + m.Password + "@tcp(" + m.Host + ":" + m.Port + ")/" + m.Name + "?" + m.Config
}

// RedisConfig 定义 Redis 配置
type RedisConfig struct {
	Host     string `mapstructure:"HOST"`
	User     string `mapstructure:"USER"`
	Password string `mapstructure:"PASSWORD"`
}
