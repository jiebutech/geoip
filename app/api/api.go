package api

import (
	"github.com/gin-gonic/gin"
	"github.com/jiebutech/app"
	"gitlab.jiebu.com/server/geoip/app/dto"
	"gitlab.jiebu.com/server/geoip/app/service"
	"gitlab.jiebu.com/server/geoip/errcode"
)

type GeoIpCtrl struct {
	svc *service.GeoIpService
}

func NewGeoIpController(svc *service.GeoIpService) *GeoIpCtrl {
	return &GeoIpCtrl{svc: svc}
}

func (ctrl *GeoIpCtrl) GetCountry(c *gin.Context) {
	ctx := app.ConvertToApiContext(c)
	req := new(dto.IpReq)
	if err := ctx.ShouldBind(req); err != nil {
		ctx.ResponseJson(errcode.ErrBadRequest, "Bad Request", nil)
		return
	}
	country, err := ctrl.svc.GetIpCountry(ctx, req.Ip)
	if err != nil {
		ctx.ResponseJson(errcode.ErrServerErr, "Server err", nil)
		return
	}
	ctx.SuccessData(dto.CountryResp{Country: country})
}
