package service

import (
	"github.com/levigross/grequests"
	"log"
	"monitoring_service/util"
	"strings"
)

var (
	apiArr []string = nil
)

func init() {

	iniCfg, err := util.NewIni("./ini/service.ini")

	if nil != err {
		log.Println(err)
		log.Panic(err)
	}

	apiStr := iniCfg.Section("service_api").Key("monitor_apis").String()

	log.Println(apiStr)
	if "" == apiStr || !strings.HasPrefix(apiStr, "http") {

		log.Println("service.ini in section service_api in key monitor_apis has no val")
		return
	}
	apiArr = strings.Split(apiStr, ",")

}

func Monitor() {
	log.Println("monitor begin")
	if nil == apiArr {
		log.Println("api arr is nil")
		return
	}
	for i := 0; i < len(apiArr); i++ {
		resp, err := grequests.Get(apiArr[i], nil)
		if err != nil || resp == nil {
			util.SendEmailPure("997747173@qq.com", apiArr[i], err.Error())
		}
	}

}
