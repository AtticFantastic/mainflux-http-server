package config

import (
	"fmt"
	"github.com/spf13/viper"
	"os"
)

type Config struct {
	// HTTP
	HttpHost string
	HttpPort int

	// NATS
	NatsHost string
	NatsPort int
}

func (this *Config) Parse() {
	/** Viper setup **/
	if len(os.Args) > 1 {
		// We provided config file as an argument
		viper.SetConfigFile(os.Args[1])
	} else {
		cfgDir := os.Getenv("GOPATH") + "/src/github.com/mainflux/mainflux-http-server/config"
		viper.SetConfigType("yaml")   // or viper.SetConfigType("YAML")
		viper.SetConfigName("config") // name of config file (without extension)
		viper.AddConfigPath(cfgDir)   // path to look for the config file in
	}
	err := viper.ReadInConfig()   // Find and read the config file
	if err != nil {               // Handle errors reading the config file
		panic(fmt.Errorf("Fatal error config file: %s \n", err))
	}

	this.HttpHost = viper.GetString("server.host")
	this.HttpPort = viper.GetInt("server.port")

	this.NatsHost = viper.GetString("nats.host")
	this.NatsPort = viper.GetInt("nats.port")
}
