package route

import (
	"github.com/jiebutech/app"
	"gitlab.jiebu.com/server/geoip/wire"
)

// RegisterRoute 路由注册
func RegisterRoute(app *app.App) {
	ctrl := wire.NewGeoIpController()
	app.Engine().GET("/country", ctrl.GetCountry)
}
