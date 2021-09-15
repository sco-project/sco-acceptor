/**
    package: sco_tracers
    filename: utils
    author: diogo@gmail.com
    time: 2021/9/14 11:36
**/
package utils


import (
	"crypto/rand"
	"encoding/hex"
)

// 生成一下UUID
func GenerateUUID() string {
	uuid := make([]byte, 16)
	n, err := rand.Read(uuid)
	if n != len(uuid) || err != nil {
		return ""
	}
	uuid[8] = 0x80 // variant bits see page 5
	uuid[4] = 0x40 // version 4 Pseudo Random, see page 7
	return hex.EncodeToString(uuid)
}

