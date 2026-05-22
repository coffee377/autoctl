package configs

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Loader Spring Boot 风格的配置加载器
type Loader struct {
	configFile string
	appName    string
	profile    string
	cmd        *cobra.Command
}

func NewLoader(cmd *cobra.Command, appName string, opts ...Option) *Loader {
	l := &Loader{appName: appName, cmd: cmd}
	for _, opt := range opts {
		opt(l)
	}
	return l
}

func (l *Loader) Load(v *viper.Viper) (*Config, error) {
	v.SetConfigType("yaml")
	v.AutomaticEnv()
	v.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	//if l.configFile != "" {
	//	file, _ := os.Open(l.configFile)
	//	v.ReadConfig(file)
	//}

	// 1. 加载默认配置: xxx.yml
	v.SetConfigName(l.appName)

	// v.AddConfigPath("$HOME/.config/" + l.appName)
	v.AddConfigPath("./configs")
	v.AddConfigPath("./conf")
	v.AddConfigPath(".")
	v.AddConfigPath("$HOME/.autoctl")

	_ = v.ReadInConfig() // 忽略不存在的情况

	// 2. 加载 profile 配置: xxx.{profile}.yml
	if l.profile != "" {
		profileConf := fmt.Sprintf("%s.%s", l.appName, l.profile)
		v.SetConfigName(profileConf)
		err := v.MergeInConfig()
		if err != nil {
			return nil, err
		}
	}

	var cfg Config
	if err := v.Unmarshal(&cfg); err != nil {
		return nil, fmt.Errorf("unmarshal config: %w", err)
	}

	// 4. 应用默认值（struct tag 中的 default）
	applyDefaults(&cfg)

	return &cfg, nil
}

func applyDefaults(cfg *Config) {
	// ... 其他默认值
}

type Option func(*Loader)

func WithProfile(profile string) Option {
	return func(l *Loader) {
		if profile != "" {
			l.profile = profile
		}
	}
}

func WithConfig(path string) Option {
	return func(l *Loader) {
		if path != "" {
			l.configFile = path
		}
	}
}
