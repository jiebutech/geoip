package route

import (
	"gitlab.jiebu.com/base/core"
	"gitlab.jiebu.com/server/geoip/wire"
)

// RegisterRoute 路由注册
func RegisterRoute(app *core.App) {
	ctrl := wire.NewGeoIpController()
	app.Engine().GET("/country", ctrl.GetCountry)
}
