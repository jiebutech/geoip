package service

import (
	"context"
	"github.com/oschwald/geoip2-golang"
	"gitlab.jiebu.com/base/log"
	"gitlab.jiebu.com/server/geoip/conf"
	"net"
)

type GeoIpService struct {
	ipDb *geoip2.Reader
}

func NewGeoIpService() *GeoIpService {
	db, err := geoip2.Open(conf.GetConf().Ip.Path)
	if err != nil {
		panic(err)
	}
	return &GeoIpService{
		ipDb: db,
	}
}

// GetIpCountry 获取ip的国家信息
func (s GeoIpService) GetIpCountry(ctx context.Context, ip string) (string, error) {
	netIp := net.ParseIP(ip)
	record, err := s.ipDb.Country(netIp)
	if err != nil {
		log.Error(ctx, "get country from ip error : "+err.Error(), ip)
		return "", err
	}
	return record.Country.IsoCode, nil
}
