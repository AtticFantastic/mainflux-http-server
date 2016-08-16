package config

import (
    "os"
    "fmt"
    "github.com/spf13/viper"
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
    // We can use config.yml from different locations,
    // depending if we run from
    cfgDir := os.Getenv("MAINFLUX_HTTP_SERVER_CONFIG_DIR")
    if cfgDir == "" {
        // default cfg path to source dir, as we keep cfg.yml there
        cfgDir = os.Getenv("GOPATH") + "/src/github.com/mainflux/mainflux-http-server/config"
    }


    viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
    viper.SetConfigName("config") // name of config file (without extension)
    viper.AddConfigPath(cfgDir)   // path to look for the config file in
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    this.HttpHost = viper.GetString("server.host")
    this.HttpPort = viper.GetInt("server.port")

    this.NatsHost = viper.GetString("nats.host")
    this.NatsPort = viper.GetInt("nats.port")
}
