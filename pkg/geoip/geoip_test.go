package geoip

import "testing"

func TestGeoIP(t *testing.T) {
	svc := NewGeoIpService()
	iso, err := svc.GetIpCountry("47.243.117.37")
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(iso)
}
