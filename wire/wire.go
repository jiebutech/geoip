//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"gitlab.jiebu.com/server/geoip/app/api"
	"gitlab.jiebu.com/server/geoip/app/service"
)

func NewGeoIpController() *api.GeoIpCtrl {
	panic(wire.Build(api.NewGeoIpController, service.NewGeoIpService))
}
