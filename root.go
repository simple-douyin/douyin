package main

import (
	"douyin/pkg/logx"
	"github.com/cloudwego/hertz/pkg/app/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var rootCmd = &cobra.Command{
	Use:                "douyin",
	DisableSuggestions: false,
	Version:            "1.0.0",
	Run: func(cmd *cobra.Command, args []string) {
		err := logx.New()
		cobra.CheckErr(err)
		h := server.Default()
		zap.L().Info("Server start ...", zap.Int("port", viper.GetInt("service.port")))
		register(h)
		h.Spin()
	},
}

func Execute() {
	err := rootCmd.Execute()
	cobra.CheckErr(err)
}

func init() {
	rootCmd.Flags().StringP("conf", "f", "service.toml", "config file path")
	err := viper.BindPFlags(rootCmd.Flags())
	cobra.CheckErr(err)

	setDefault()
	viper.SetConfigFile(viper.GetString("conf"))
	viper.WatchConfig()
}

func setDefault() {
	viper.SetDefault("service.port", 8888)
	viper.SetDefault("service.timeout", 10)
	viper.SetDefault("service.mode", "pro")
	viper.SetDefault("service.log.path", "./log/service/")
	viper.SetDefault("service.log.rotateDays", 7)
	viper.SetDefault("service.log.rotateSize", 7)
	viper.SetDefault("service.log.backups", 7)
}
