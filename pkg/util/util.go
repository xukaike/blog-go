package util

import (
	"crypto/md5"
	"fmt"
)

func Md5(str string) string{
	bytes := []byte(str)
	hash := md5.Sum(bytes)
	return fmt.Sprintf("%x",hash)
}