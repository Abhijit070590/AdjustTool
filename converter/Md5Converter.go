package converter

import (
	"crypto/md5"
	"fmt"
	"log"
)

type Md5Converter struct {
}

func (c *Md5Converter) GetHash(source []byte) string {
	if source == nil {
		log.Println("Response not Found")
		return ""
	}
	hash := md5.Sum(source)
	return fmt.Sprintf("%x", hash)
}
