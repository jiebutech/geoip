package main

import (
	"github.com/jiebutech/app"
	"github.com/jiebutech/geoip/conf"
	"github.com/jiebutech/geoip/route"
)

func main() {
	app.RunApp(conf.GetConf().Configuration, route.RegisterRoute)
}
