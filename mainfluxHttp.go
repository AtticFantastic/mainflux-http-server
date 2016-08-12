/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
    "log"
    "fmt"
    "strconv"
    "os"

    "github.com/iris-contrib/middleware/logger"
    "github.com/kataras/iris"
    "github.com/nats-io/nats"

    "github.com/fatih/color"
    "github.com/spf13/viper"
)

var Nc *nats.Conn

func main() {

    /** Viper setup **/
    // We can use config.yml from different locations,
    // depending if we run from
    cfgDir := os.Getenv("MAINFLUX_HTTP_SERVER_CONFIG_DIR")
    if cfgDir == "" {
        // default cfg path to source dir, as we keep cfg.yml there
        cfgDir = os.Getenv("GOPATH") + "/src/github.com/mainflux/mainflux-http-server"
    }


    viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
    viper.SetConfigName("config") // name of config file (without extension)
    viper.AddConfigPath(cfgDir)   // path to look for the config file in
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    host := viper.GetString("server.host")
    port := viper.GetInt("server.port")

    ntshost := viper.GetString("nats.host")
    ntsport := viper.GetInt("nats.port")

    // Iris config
    iris.Config.DisableBanner = true

    // set the global middlewares
	  iris.Use(logger.New(iris.Logger))

    // set the custom errors
    iris.OnError(iris.StatusNotFound, func(ctx *iris.Context) {
        ctx.Render("errors/404.html", iris.Map{"Title": iris.StatusText(iris.StatusNotFound)})
    })

    iris.OnError(iris.StatusInternalServerError, func(ctx *iris.Context) {
        ctx.Render("errors/500.html", nil, iris.RenderOptions{"layout": iris.NoLayout})
    })

    // register the public API
    registerAPI()

    /** Connect to NATS broker */
    ncp, err := nats.Connect("nats://" + ntshost + ":" + strconv.Itoa(ntsport))
    if err != nil {
       log.Fatalf("Can't connect: %v\n", err)
    }
    Nc = ncp

    color.Cyan(banner)
    color.Cyan("Magic happens on port " + strconv.Itoa(port))

    // start the server
    iris.Listen(host + ":" + strconv.Itoa(port))
}

func registerAPI() {
    iris.API("/status", StatusAPI{})
    iris.API("/devices", DeviceAPI{})
    iris.API("/channels", ChannelAPI{})
}

var banner = `
_|      _|            _|                _|_|  _|                      
_|_|  _|_|    _|_|_|      _|_|_|      _|      _|  _|    _|  _|    _|  
_|  _|  _|  _|    _|  _|  _|    _|  _|_|_|_|  _|  _|    _|    _|_|    
_|      _|  _|    _|  _|  _|    _|    _|      _|  _|    _|  _|    _|  
_|      _|    _|_|_|  _|  _|    _|    _|      _|    _|_|_|  _|    _|  
                                                                     

                == Industrial IoT System ==
       
                Made with <3 by Mainflux Team
[w] http://mainflux.io
[t] @mainflux

                    ** HTTP SERVER **

`

