package cmd

import (
	"fmt"
	"time"

	"github.com/coffee377/autoctl/configs"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	tokenRedisAddr     string
	tokenRedisPassword string
	tokenRedisDB       int
	tokenAccount       string
	tokenExpire        time.Duration
	tokenSessionFixed  bool
)

func GetTokenCommand() *cobra.Command {
	tokenCmd := &cobra.Command{
		Use:   "token",
		Short: "Generate token for authentication",
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return bind(cmd)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			loader := configs.NewLoader(
				cmd, "auto",
				configs.WithProfile(viper.GetString("profile")),
				configs.WithConfig(viper.GetString("file")),
			)
			config, err := loader.Load(viper.GetViper())
			if err != nil {
				return err
			}

			executor := NewTokenExecutor(config)
			result, err := executor.Execute()
			if err != nil {
				return err
			}
			fmt.Println(result)
			return nil
		},
	}

	tokenCmd.PersistentFlags().StringVarP(&tokenRedisAddr, "redis-addr", "", "", "Redis server address (overrides config)")
	tokenCmd.PersistentFlags().StringVarP(&tokenRedisPassword, "redis-password", "", "", "Redis server password (overrides config)")
	tokenCmd.PersistentFlags().IntVarP(&tokenRedisDB, "redis-db", "i", 0, "Redis database index (overrides config)")
	tokenCmd.PersistentFlags().StringVarP(&tokenAccount, "token-account", "a", "", "Account username (overrides config)")
	tokenCmd.PersistentFlags().DurationVarP(&tokenExpire, "token-expire", "e", 0, "Token expiration time (overrides config)")
	tokenCmd.PersistentFlags().BoolVarP(&tokenSessionFixed, "session-fixed", "s", false, "Use fixed session ID")

	return tokenCmd
}

func bind(cmd *cobra.Command) error {
	_ = BindFlags(cmd, "redis-addr", "redis.address")
	_ = BindFlags(cmd, "redis-password", "redis.password")
	_ = BindFlags(cmd, "redis-db", "redis.db")
	_ = BindFlags(cmd, "token-account", "token.account")
	_ = BindFlags(cmd, "token-expire", "token.expire")
	_ = BindFlags(cmd, "session-fixed", "token.session-fixed")
	return nil
}

func BindFlags(cmd *cobra.Command, flagName, bindKey string) error {
	if flagName != "" && bindKey != "" {
		flag := cmd.Flags().Lookup(flagName)
		if flag != nil {
			return viper.BindPFlag(bindKey, flag)
		}
	}
	return nil
}
