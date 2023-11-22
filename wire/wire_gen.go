// Code generated by Wire. DO NOT EDIT.

//go:generate go run github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package wire

import (
	"gitlab.jiebu.com/server/geoip/app/api"
	"gitlab.jiebu.com/server/geoip/app/service"
)

// Injectors from wire.go:

func NewGeoIpController() *api.GeoIpCtrl {
	geoIpService := service.NewGeoIpService()
	geoIpCtrl := api.NewGeoIpController(geoIpService)
	return geoIpCtrl
}