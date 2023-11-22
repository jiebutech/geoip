package main

import (
	"gitlab.jiebu.com/base/core"
	"gitlab.jiebu.com/server/geoip/conf"
	"gitlab.jiebu.com/server/geoip/route"
)

func main() {
	core.RunApp(conf.GetConf().Configuration, route.RegisterRoute)
}
