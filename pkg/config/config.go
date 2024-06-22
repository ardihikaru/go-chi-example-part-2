package config

import (
	"time"

	"github.com/lestrrat-go/jwx/jwa"
	"github.com/spf13/viper"
)

type Reader interface {
	Get(key string) string
}

type ViperConfigReader struct {
	viper *viper.Viper
}

type General struct {
	BuildMode string `mapstructure:"buildMode"`
}

type Log struct {
	Level  string `mapstructure:"level"`
	Format string `mapstructure:"format"`
}

type LogPublisher struct {
	Notice bool `mapstructure:"notice"`
	Error  bool `mapstructure:"error"`
	Info   bool `mapstructure:"info"`
	Warn   bool `mapstructure:"warn"`
	Debug  bool `mapstructure:"debug"`
}

type Http struct {
	Address        string        `mapstructure:"address"`
	Port           int           `mapstructure:"port"`
	RequestTimeout time.Duration `mapstructure:"requestTimeout"`
	ReadTimeout    time.Duration `mapstructure:"readTimeout"`  // The maximum time to wait while writing data to the server
	WriteTimeout   time.Duration `mapstructure:"writeTimeout"` // The maximum time to wait while reading data from the server
	HttpClientTLS  bool          `mapstructure:"httpClientTLS"`
	LogHttpRequest bool          `mapstructure:"logHttpRequest"`
}

// Cors defines cors-related configurations
type Cors struct {
	AllowedOrigins   []string `mapstructure:"allowedOrigins"`
	AllowedMethods   []string `mapstructure:"allowedMethods"`
	AllowedHeaders   []string `mapstructure:"allowedHeaders"`
	ExposedHeaders   []string `mapstructure:"exposedHeaders"`
	AllowCredentials bool     `mapstructure:"allowCredentials"`
	MaxAge           int      `mapstructure:"maxAge"`
	Debug            bool     `mapstructure:"debug"`
}

// JwtAuth defines JWT authentication related
type JwtAuth struct {
	JWTSecret       string                 `mapstructure:"jwtSecret"`
	JWTAlgorithm    jwa.SignatureAlgorithm `mapstructure:"jwtAlgorithm"`
	JWTExpiredInSec int                    `mapstructure:"jwtExpiredInSec"`
}

type Config struct {
	General      General      `mapstructure:"general"`
	Log          Log          `mapstructure:"log"`
	LogPublisher LogPublisher `mapstructure:"logPublisher"`
	Http         Http         `mapstructure:"http"`
	Cors         Cors         `mapstructure:"cors"`
	JwtAuth      JwtAuth      `mapstructure:"jwtAuth"`
}

// Validate validates any miss configurations or missing configs
func (*Config) Validate() error {
	// TODO: implements this logic

	return nil
}

// Get gets config object
func Get() (*Config, error) {
	cfg, err := Load()
	if err != nil {
		return nil, err
	}

	return cfg, nil
}

// Load loads config from the config.yaml
func Load() (*Config, error) {
	v := viper.New()
	v.AddConfigPath(".")
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AutomaticEnv()

	if err := v.ReadInConfig(); err != nil {
		return nil, err
	}

	var config Config
	err := v.Unmarshal(&config)
	if err != nil {
		return nil, err
	}

	return &config, err
}
