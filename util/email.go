package util

import (
	_log "github.com/Sirupsen/logrus"
	"gopkg.in/gomail.v2"
)

var (
	err    error
	dialer *gomail.Dialer
	msg    *gomail.Message
)

func SendEmail(from, to, cc, subject, content, username, userpwd, host string, port int) {

	dialer = gomail.NewDialer(host, port, username, userpwd)

	msg = gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetAddressHeader("Cc", cc, cc)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", content)

	err = dialer.DialAndSend(msg)

	go _log.WithFields(_log.Fields{
		"err": err,
	}).Errorln("util emailutil senemail")

}

func SendEmailNoCc(from, to, subject, content, username, userpwd, host string, port int) {

	dialer = gomail.NewDialer(host, port, username, userpwd)

	msg = gomail.NewMessage()
	msg.SetHeader("From", from)
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", content)

	err = dialer.DialAndSend(msg)

	go _log.WithFields(_log.Fields{
		"err": err,
	}).Errorln("util emailutil senemail")
}

func SendEmailPure(to, subject, content string) {

	iniCfg, err := NewIni("./ini/email.ini")
	// todo：暂时先如此
	iniSec := iniCfg.Section("email")

	e_server := iniSec.Key("e_server").String()
	// e_port := iniSec.Key("e_port").Int()
	u_name := iniSec.Key("u_name").String()
	u_pwd := iniSec.Key("u_pwd").String()

	//  todo:暂时先这样
	dialer = gomail.NewDialer(e_server, 465, u_name, u_pwd)

	msg = gomail.NewMessage()
	msg.SetHeader("From", "xx")
	msg.SetHeader("To", to)
	msg.SetHeader("Subject", subject)
	msg.SetBody("text/html", content)

	err = dialer.DialAndSend(msg)

	go _log.WithFields(_log.Fields{
		"err": err,
	}).Errorln("util emailutil senemail")
}
