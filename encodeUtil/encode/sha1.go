package encode

import (
	"crypto/sha1"
	"fmt"
	"hash"
	"io"
)

type SHA1Encode struct {
}

/*****************************
*  哈希算法  ----- 安全散列算法
*****************************/
func (this *SHA1Encode) GenerateAppKey(args ...string) string {

	var h hash.Hash = sha1.New()

	for _, str := range args {
		io.WriteString(h, str)
	}

	var result string
	result = fmt.Sprintf("%x", h.Sum(nil))

	return result
}
