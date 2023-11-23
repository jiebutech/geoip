package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebutech/app"
	"github.com/jiebutech/geoip/app/dto"
	"github.com/jiebutech/geoip/errcode"
	"github.com/jiebutech/geoip/pkg/geoip"
	"github.com/jiebutech/log"
)

type GeoIpCtrl struct {
	svc geoip.GeopIpSvc
}

func NewGeoIpController(svc geoip.GeopIpSvc) *GeoIpCtrl {
	return &GeoIpCtrl{svc: svc}
}

func (ctrl *GeoIpCtrl) GetCountry(c *gin.Context) {
	ctx := app.ConvertToApiContext(c)
	req := new(dto.IpReq)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.ResponseJson(errcode.ErrBadRequest, "Bad Request", nil)
		return
	}
	country, err := ctrl.svc.GetIpCountry(req.Ip)
	if err != nil {
		log.Error(ctx, "get country from ip error: "+err.Error(), req.Ip)
		ctx.ResponseJson(errcode.ErrServerErr, "Server err", nil)
		return
	}
	ctx.SuccessData(dto.CountryResp{Country: country})
}
