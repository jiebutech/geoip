package dto

type CountryResp struct {
	Country string `json:"country" form:"country" query:"country"`
}

type IpReq struct {
	Ip string `json:"ip" form:"ip" query:"ip" binding:"required"`
}
