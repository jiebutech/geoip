package main

import (
	"github.com/jiebutech/app"
	"gitlab.jiebu.com/server/geoip/conf"
	"gitlab.jiebu.com/server/geoip/route"
)

func main() {
	app.RunApp(conf.GetConf().Configuration, route.RegisterRoute)
}
