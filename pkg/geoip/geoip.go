package geoip

import (
	"github.com/oschwald/geoip2-golang"
	"net"
	"path"
	"runtime"
)

type service struct {
	ipDb *geoip2.Reader
}

type GeopIpSvc interface {
	GetIpCountry(ip string) (string, error)
}

func NewGeoIpService() GeopIpSvc {
	_, file, _, _ := runtime.Caller(0)
	dbPath := path.Dir(file) + "/GeoLite2-City.mmdb"
	db, err := geoip2.Open(dbPath)
	if err != nil {
		panic(err)
	}
	return &service{
		ipDb: db,
	}
}

// GetIpCountry 获取ip的国家信息
func (s *service) GetIpCountry(ip string) (string, error) {
	netIp := net.ParseIP(ip)
	record, err := s.ipDb.Country(netIp)
	if err != nil {
		return "", err
	}
	return record.Country.IsoCode, nil
}
