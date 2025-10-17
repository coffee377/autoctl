package app

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/viper"
)

type Configuration struct {
	Namespace string  `mapstructure:"namespace"`
	Id        string  `mapstructure:"id"`
	Name      *string `mapstructure:"name"`

	ClientId     string `mapstructure:"client_id"`
	ClientSecret string `mapstructure:"client_secret"`

	AgentId   *string `mapstructure:"agent_id"`
	RobotCode *string `mapstructure:"robot_code"`
}

type Configurations struct {
	App []Configuration
}

func ReadConfiguration() *Configurations {
	config, err := loadConfig()
	if err != nil {
		log.Fatalf("加载配置失败：%v", err)
	}
	return config
}

// 加载指定环境的配置
func loadConfig() (*Configurations, error) {
	// 1. 读取环境变量，获取当前环境（默认 dev）
	env := os.Getenv("ENV")
	if env == "" {
		env = "dev" // 默认开发环境
	}
	fmt.Printf("当前环境：%s\n", env)

	// 2. 初始化 Viper
	v := viper.New()

	// 3. 配置文件路径（根据实际项目调整，这里假设在 config 目录下）
	v.AddConfigPath(".")
	v.AddConfigPath("..")

	// 4. （可选）加载通用默认配置（app.yml）
	v.SetConfigName("app")
	v.SetConfigType("yaml")

	if err := v.ReadInConfig(); err != nil {
		// 非必需，若没有默认配置可忽略错误
		log.Printf("未找到通用配置文件 app.yml，跳过加载：%v", err)
	}

	// 5. 加载环境特定配置（如 app.dev.yml）
	v.SetConfigName(fmt.Sprintf("app.%s", env)) // 构造文件名：app.dev
	_ = v.MergeInConfig()

	// 6. 反序列化为结构体
	var config Configurations
	if err := v.Unmarshal(&config); err != nil {
		log.Fatalf("配置反序列化失败：%v", err)
	}

	return &config, nil
}
