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
    "time"
    "github.com/kataras/iris"
)

type DeviceAPI struct {
	  *iris.Context
}

func reqCore(s string) string {
    msg, err := Nc.Request("core_in", []byte(s), 1000*time.Millisecond)
    if err != nil {
		    if Nc.LastError() != nil {
			      log.Fatalf("Error in Request: %v\n", Nc.LastError())
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
func (d DeviceAPI) Get() {
    s := formJson("getDevices", "", d.RequestCtx.Request.Body())
    d.Write(reqCore(s))
}

// POST /devices
func (d DeviceAPI) Post() {
    s := formJson("createDevice", "", d.RequestCtx.Request.Body())
    d.Write(reqCore(s))
}

// GET /devices/:id
func (d DeviceAPI) GetBy(id string) {
    s := formJson("getDevice", id, d.RequestCtx.Request.Body())
    d.Write(reqCore(s))
}

// PUT /devices/:id
func (d DeviceAPI) PutBy(id string) {
    s := formJson("updateDevice", id, d.RequestCtx.Request.Body())
    d.Write(reqCore(s))
}

// DELETE /devices/:id
func (d DeviceAPI) DeleteBy(id string) {
    s := formJson("deleteDevice", id, d.RequestCtx.Request.Body())
    d.Write(reqCore(s))
}
