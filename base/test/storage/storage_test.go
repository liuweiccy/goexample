package storage

import (
	"strings"
	"testing"
)

// 进行了白盒测试
// 对邮件发送函数进行了替换伪造，便于进行实际的测试
func TestCheckQuota(t *testing.T) {
	// 保留待恢复的notifyUser
	saved := notifyUser
	defer func() {
		notifyUser = saved
	}()

	var notifiedUser, notifiedMsg string

	notifyUser = func(username, msg string) {
		notifiedUser, notifiedMsg = username, msg
	}

	const user = "testing@qq.com"
	CheckQuota(user)
	if notifiedUser == "" && notifiedMsg == "" {
		t.Fatalf("notifyUser not called!")
	}

	if notifiedUser != user {
		t.Errorf("通知了错误的用户%s, 应该通知%s", notifiedUser, user)
	}

	const percentSrc = "98%"
	if !strings.Contains(notifiedMsg, percentSrc) {
		t.Errorf("期望包含%s,但实际字符串%s未包含", percentSrc, notifiedMsg)
	}

}
