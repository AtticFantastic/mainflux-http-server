/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package routes

import (
	//"log"
	//"time"
	"github.com/kataras/iris"
)

/**
 * InfluxDB
 */
/*
// GET /channels/:id/ts
func (c ChannelAPI) GetBy() {
    s := formJson("queryTs", "", c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// GET /channels/:id/msg
func (c ChannelAPI) Get() {
    s := formJson("queryMsg", "", c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// POST /channels/:id/ts
func (c ChannelAPI) Get() {
    s := formJson("insertTs", "", c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// POST /channels/:id/msg
func (c ChannelAPI) Get() {
    s := formJson("insertMsg", "", c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}
*/

// GET /channels
func GetChannels(ctx *iris.Context) {
	s := formJson("getChannels", "", ctx.Request.Body())
	ctx.Write(reqCore(s))
}

// POST /channels
func CreateChannel(ctx *iris.Context) {
	s := formJson("createChannel", "", ctx.Request.Body())
	ctx.Write(reqCore(s))
}

// GET /channels/:id
func GetChannel(ctx *iris.Context) {
	id := ctx.Param("id")
	s := formJson("getChannel", id, ctx.Request.Body())
	ctx.Write(reqCore(s))
}

// PUT /channels/:id
func UpdateChannel(ctx *iris.Context) {
	id := ctx.Param("id")
	s := formJson("updateChannel", id, ctx.Request.Body())
	ctx.Write(reqCore(s))
}

// DELETE /channels/:id
func DeleteChannel(ctx *iris.Context) {
	id := ctx.Param("id")
	s := formJson("deleteChannel", id, ctx.Request.Body())
	ctx.Write(reqCore(s))
}
