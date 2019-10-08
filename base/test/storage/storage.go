package storage

import (
	"fmt"
	"log"
	"net/smtp"
)

const (
	sender   = "test@qq.com"
	password = "123456"
	hostname = "smtp.qq.com"

	template = `注意：你已经使用了%d的字节了，达到总量的%d%%`
)

func bytesInUse(username string) int64 {
	return 980000000
}

var notifyUser = func(username, msg string) {
	auth := smtp.PlainAuth("", sender, password, hostname)
	err := smtp.SendMail(hostname+":578", auth, sender, []string{username}, []byte(msg))
	if err != nil {
		log.Printf("smtp.SendMail(%s), failed:%s", username, err)
	}
}

func CheckQuota(username string) {
	used := bytesInUse(username)
	const quota = 1000000000
	percent := 100 * used / quota

	if percent < 90 {
		return
	}

	msg := fmt.Sprintf(template, used, percent)
	notifyUser(username, msg)
}
