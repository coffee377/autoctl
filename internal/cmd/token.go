package cmd

import (
	"fmt"
	"time"

	"github.com/coffee377/autoctl/internal/token"
	"github.com/redis/go-redis/v9"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	tokenRedisAddr     string
	tokenRedisPassword string
	tokenRedisDB       int
	tokenAccount       string
	tokenExpire        time.Duration
	tokenFixSession    bool
)

func GetTokenCommand() *cobra.Command {
	tokenCmd := &cobra.Command{
		Use:   "token",
		Short: "Generate token for authentication",
		RunE: func(cmd *cobra.Command, args []string) error {
			// 参数优先级：命令行参数 > 环境变量 > 配置文件 > 默认值
			redisAddr := viper.GetString("redis.host")
			if redisAddr == "" {
				redisAddr = "127.0.0.1"
			}
			if viper.IsSet("redis.port") {
				redisAddr = fmt.Sprintf("%s:%d", redisAddr, viper.GetInt("redis.port"))
			}
			if tokenRedisAddr != "" {
				redisAddr = tokenRedisAddr
			}

			redisPassword := viper.GetString("redis.password")
			if tokenRedisPassword != "" {
				redisPassword = tokenRedisPassword
			}

			redisDB := viper.GetInt("redis.db")
			if tokenRedisDB > 0 {
				redisDB = tokenRedisDB
			} else if redisDB == 0 {
				redisDB = 5 // 默认值
			}

			account := viper.GetString("token.account")
			if tokenAccount != "" {
				account = tokenAccount
			}
			if account == "" {
				account = "coffee377"
			}

			expire := viper.GetDuration("token.expire")
			if tokenExpire > 0 {
				expire = tokenExpire
			}
			if expire == 0 {
				expire = time.Minute * 5
			}

			redisOptions := redis.Options{
				Addr:     redisAddr,
				Password: redisPassword,
				DB:       redisDB,
			}
			jwtGenerator := token.NewJWTGenerator(token.WithJinQi(), token.WithRedis(redisOptions))

			sessionOpts := []token.SessionOptions{
				token.WithAccount(account),
				token.WithExpire(expire),
			}
			if tokenFixSession {
				sessionOpts = append(sessionOpts, token.WithFixSession())
			}

			session := token.NewJinQiSession(sessionOpts...)
			result, err := jwtGenerator.Generate(session)
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	tokenCmd.PersistentFlags().StringVarP(&tokenRedisAddr, "redis-addr", "a", "", "Redis server address (overrides config)")
	tokenCmd.PersistentFlags().StringVarP(&tokenRedisPassword, "redis-password", "p", "", "Redis server password (overrides config)")
	tokenCmd.PersistentFlags().IntVarP(&tokenRedisDB, "redis-db", "i", 0, "Redis database index (overrides config)")
	tokenCmd.PersistentFlags().StringVarP(&tokenAccount, "account", "u", "", "Account username (overrides config)")
	tokenCmd.PersistentFlags().DurationVarP(&tokenExpire, "expire", "e", 0, "Token expiration time (overrides config)")
	tokenCmd.PersistentFlags().BoolVarP(&tokenFixSession, "fix-session", "f", false, "Use fixed session ID")

	return tokenCmd
}
