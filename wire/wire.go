//go:build wireinject

package wire

import (
	"github.com/google/wire"
	"github.com/jiebutech/geoip/app/api"
	"github.com/jiebutech/geoip/pkg/geoip"
)

func NewGeoIpController() *api.GeoIpCtrl {
	panic(wire.Build(api.NewGeoIpController, geoip.NewGeoIpService))
}
