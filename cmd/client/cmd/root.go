package cmd

import (
	"strings"

	"github.com/joho/godotenv"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
	"github.com/spf13/viper"

	"github.com/go-developer-ya-practicum/gophkeeper/pkg/token"
	"github.com/go-developer-ya-practicum/gophkeeper/pkg/version"
)

var (
	cfgFile string

	defaults = map[string]interface{}{
		"grpc.address": "127.0.0.1:9090",
	}

	tokenStorage token.Storage
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:     "gophkeeper-cli",
	Short:   "GophKeeper client",
	Version: version.Info(),
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		log.Error().Err(err)
	}
}

func init() {
	tokenStorage = token.NewFileStorage("token.txt")

	if err := godotenv.Load(); err != nil {
		log.Debug().Msg("No .env file found")
	}

	rootCmd.PersistentFlags().StringVarP(
		&cfgFile, "config", "c", "", "Client config filepath")

	rootCmd.PersistentFlags().StringP(
		"grpc-address", "g", "", "Server grpc address")

	rootCmd.PersistentFlags().StringP(
		"encryption-key", "k", "", "Secret encryption key")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	for key, value := range defaults {
		viper.SetDefault(key, value)
	}

	if cfgFile != "" {
		viper.SetConfigFile(cfgFile)
	} else {
		viper.AddConfigPath("/etc/gophkeeper")
		viper.AddConfigPath("$HOME")
		viper.AddConfigPath(".")

		viper.SetConfigName("client-config")
	}
	if err := viper.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); !ok {
			log.Fatal().Err(err).Msg("Failed to load server config")
		}
		log.Debug().Msg("Config file not found")
	} else {
		log.Debug().Msgf("Using config file: %s", viper.ConfigFileUsed())
	}

	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer("-", "_", ".", "_"))

	rootCmd.Flags().VisitAll(func(flag *pflag.Flag) {
		key := strings.ReplaceAll(flag.Name, "-", ".")
		if err := viper.BindPFlag(key, flag); err != nil {
			log.Fatal().Err(err).Msg("Failed to bind flag")
		}
	})
}
