package util

import (
	"crypto/md5"
	"fmt"
	"strings"
)

func MD5(str string) string {
	md5Byte := md5.Sum([]byte(str))
	md5Str := fmt.Sprintf("%x", md5Byte)
	return strings.ToLower(md5Str)
}
