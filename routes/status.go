/**
 * Copyright (c) Mainflux
 *
 * Mainflux server is licensed under an Apache license, version 2.0.
 * All rights not explicitly granted in the Apache license, version 2.0 are reserved.
 * See the included LICENSE file for more details.
 */

package routes

import (
	"github.com/kataras/iris"
)

// GET /status
func GetStatus(ctx *iris.Context) {
	m := formJson("getStatus", "", ctx.Request.Body())
	ctx.Write(reqCore(m))
}
