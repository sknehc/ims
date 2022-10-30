package imsutils

import (
	"regexp"
	"strings"
)

// 消息分类，msg_type={1:分配用户名}
func DealMsg(str string) (int, string) {
	reg := regexp.MustCompile(`^newclient:.*`)
	if reg != nil {
		s := reg.FindAllStringSubmatch(str, -1)
		msg := strings.Split(s[0][0], "newclient:")
		return 1, msg[1]
	}
	return 2, str
}
