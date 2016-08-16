/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package routes

import (
    "log"
    "time"
    "github.com/mainflux/mainflux-http-server/broker"
    "github.com/kataras/iris"
)

func reqCore(s string) string {
    msg, err := broker.Nc.Request("core_in", []byte(s), 1000*time.Millisecond)
    if err != nil {
		    if broker.Nc.LastError() != nil {
			      log.Fatalf("Error in Request: %v\n", broker.Nc.LastError())
		    }
		    log.Fatalf("Error in Request: %v\n", err)
	  }

    return string(msg.Data)
}

func formJson (m string, id string, b []byte) string {
    var t string

    if len(b) == 0 {
        t = `"body" : {}`
    } else {
        t = `"body" : ` + string(b)
    }

    return `{"method": "` + m + `", "id": "` + id + `", ` + t + `}`
}

// GET /devices
func GetDevices(ctx *iris.Context) {
    s := formJson("getDevices", "", ctx.Request.Body())
    ctx.Write(reqCore(s))
}

// POST /devices
func CreateDevice(ctx *iris.Context) {
    s := formJson("createDevice", "", ctx.Request.Body())
    ctx.Write(reqCore(s))
}

// GET /devices/:id
func GetDevice(ctx *iris.Context) {
    id := ctx.Param("id")
    s := formJson("getDevice", id, ctx.Request.Body())
    ctx.Write(reqCore(s))
}

// PUT /devices/:id
func UpdateDevice(ctx *iris.Context) {
    id := ctx.Param("id")
    s := formJson("updateDevice", id, ctx.Request.Body())
    ctx.Write(reqCore(s))
}

// DELETE /devices/:id
func DeleteDevice(ctx *iris.Context) {
    id := ctx.Param("id")
    s := formJson("deleteDevice", id, ctx.Request.Body())
    ctx.Write(reqCore(s))
}
