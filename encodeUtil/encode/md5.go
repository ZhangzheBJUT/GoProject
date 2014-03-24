package encode

import (
	"crypto/md5"
	"fmt"
	"hash"
	"io"
)

type MD5Encode int

/*****************************
*  哈希算法  ----- MD5  摘要算法
*****************************/
func (this *MD5Encode) GenerateAppKey(args ...string) string {

	var h hash.Hash = md5.New()

	for _, str := range args {
		io.WriteString(h, str)
	}

	var result string
	result = fmt.Sprintf("%x", h.Sum(nil))
	return result
}
