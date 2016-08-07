package main

import (
    "./api"
    "./mfconns"

    //"log"
    //"fmt"

    "github.com/iris-contrib/middleware/logger"
    "github.com/kataras/iris"
    "github.com/nats-io/nats"

    "github.com/fatih/color"
)

var nc nats.Conn

func main() {

    // Iris config
    //irisConfig := config.Iris{ DisableBanner: true }
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

    // start the server
    iris.Listen("127.0.0.1:7070")
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

Magic happens on port 7070
`

