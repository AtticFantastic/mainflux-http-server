/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package main

import (
    "github.com/kataras/iris"
)

type StatusAPI struct {
	  *iris.Context
}

// GET /status
func (s StatusAPI) Get() {
    m := formJson("getStatus", "", s.RequestCtx.Request.Body())
    s.Write(reqCore(m))
}

