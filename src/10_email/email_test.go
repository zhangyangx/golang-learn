package main

import (
	"net/smtp"
	"strings"
	"testing"
)

func TestSendMail(t *testing.T) {
	var userName = "yang.zhang-iog@daocloud.io"
	var password = "zy214214~"
	var mailServer = "smtp.partner.outlook.cn:25"
	var to = "1154278413@qq.com"
	var subject = "犀点意象安全团队"
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + userName + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + "xxx")
	hp := strings.Split(mailServer, ":")
	auth := smtp.PlainAuth("xxx", userName, password, hp[0])
	err := smtp.SendMail(mailServer, auth, userName, []string{to}, msg)
	if err != nil {
		t.Error()
	}

}
