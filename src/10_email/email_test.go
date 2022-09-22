package main

import (
	"net/smtp"
	"strings"
	"testing"
)

func TestSendMail(t *testing.T) {
	var userName = "1154278413@qq.com"
	var password = "uspqoarghosrfejc"
	var mailServer = "smtp.qq.com:465"
	var to = "yang.zhang-iog@daocloud.io"
	var subject = "犀点意象安全团队"
	contentType := "Content-Type: text/html; charset=UTF-8"
	msg := []byte("To: " + to + "\r\nFrom: " + userName + "\r\nSubject: " + subject + "\r\n" + contentType + "\r\n\r\n" + "xxx")
	hp := strings.Split(mailServer, ":")
	auth := smtp.PlainAuth("", userName, password, hp[0])
	err := smtp.SendMail(mailServer, auth, userName, []string{to}, msg)
	if err != nil {
		t.Error()
	}

}
