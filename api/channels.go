package api

import (
    //"log"
    //"time"
    "github.com/kataras/iris"
)

type ChannelAPI struct {
	  *iris.Context
}

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

/**
 * MongoDB
 */
// GET /channels
func (c ChannelAPI) Get() {
    s := formJson("getChannels", "", c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// POST /channels
func (c ChannelAPI) Post() {
    s := formJson("createChannel", "", c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// GET /channels/:id
func (c ChannelAPI) GetBy(id string) {
    s := formJson("getChannel", id, c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// PUT /devices/:id
func (c ChannelAPI) PutBy(id string) {
    s := formJson("updateChannel", id, c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}

// DELETE /devices/:id
func (c ChannelAPI) DeleteBy(id string) {
    s := formJson("deleteChannel", id, c.RequestCtx.Request.Body())
    c.Write(reqCore(s))
}
