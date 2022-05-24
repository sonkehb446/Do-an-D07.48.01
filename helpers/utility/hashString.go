package utility

import (
	"crypto/md5"
	"encoding/hex"
)

func HashStringMD5(text string) string {
	hasher := md5.New()
	hasher.Write([]byte(text))
	return hex.EncodeToString(hasher.Sum(nil))
}

func CheckStringHashMD5(text string, hash string) bool {
	in := HashStringMD5(text)
	if in == hash {
		return true
	}
	return false
}
