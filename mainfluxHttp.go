package main

import (
    "./api"
    "./mfconns"

    //"log"
    "fmt"
    "strconv"

    "github.com/iris-contrib/middleware/logger"
    "github.com/kataras/iris"
    "github.com/nats-io/nats"

    "github.com/fatih/color"
    "github.com/spf13/viper"
)

var nc nats.Conn

func main() {

    // Viper setup
    viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
    viper.SetConfigName("config") // name of config file (without extension)
    viper.AddConfigPath(".")   // path to look for the config file in
    err := viper.ReadInConfig() // Find and read the config file
    if err != nil { // Handle errors reading the config file
        panic(fmt.Errorf("Fatal error config file: %s \n", err))
    }

    host := viper.GetString("host")
    port := viper.GetInt("port")

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
    mfconns.Nc, _ = nats.Connect(nats.DefaultURL)

    color.Cyan(banner)
    color.Cyan("Magic happens on port " + strconv.Itoa(port))

    // start the server
    iris.Listen(host + ":" + strconv.Itoa(port))
}

func registerAPI() {
    iris.API("/status", api.StatusAPI{})
    iris.API("/devices", api.DeviceAPI{})
    iris.API("/channels", api.ChannelAPI{})
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

