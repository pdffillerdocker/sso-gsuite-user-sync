package config

// Config ...
type Config struct {
	// Verbose toggles the verbosity
	Debug bool
	// LogLevel is the level with with to log for this config
	LogLevel string `mapstructure:"log_level"`
	// LogFormat is the format that is used for logging
	LogFormat string `mapstructure:"log_format"`
    // SecretBucket ...
    SecretBucket string `mapstructure:"secret_bucket"`
    // SecretBucketRegion ...
    SecretBucketRegion string `mapstructure:"secret_bucket_region"`
	// GoogleCredentials ...
	GoogleCredentials string `mapstructure:"google_credentials"`
	// GoogleAdmin ...
	GoogleAdmin string `mapstructure:"google_admin"`
	// SCIMEndpoint ....
	SCIMEndpoint string `mapstructure:"scim_endpoint"`
	// SCIMAccessToken ...
	SCIMAccessToken string `mapstructure:"scim_access_token"`
	// IsLambda ...
	IsLambda bool
	// Ignore users ...
	IgnoreUsers []string `mapstructure:"ignore_users"`
}

const (
	// DefaultLogLevel is the default logging level.
	DefaultLogLevel = "warn"
	// DefaultLogFormat is the default format of the logger
	DefaultLogFormat = "text"
	// DefaultDebug is the default debug status.
	DefaultDebug = false
)

// New returns a new Config
func New() *Config {
	return &Config{
		Debug:             DefaultDebug,
		LogLevel:          DefaultLogLevel,
		LogFormat:         DefaultLogFormat,
	}
}
