package conf

import (
	"github.com/jiebutech/config"
)

type AppConf struct {
	*config.Configuration `json:",inline" yaml:",inline"`
}

var cfg *AppConf

func init() {
	c := new(AppConf)
	err := config.Init(c)
	if err != nil {
		panic(err)
	}
	cfg = c
}

func GetConf() *AppConf {
	return cfg
}
