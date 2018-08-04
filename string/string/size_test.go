package string

import (
	"fmt"
	"testing"
)

var str = `
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="dbnames"

中文字符测试表单提交
-----------------------------3354441481387117281646263294
Content-Disposition: form-data; name="file"; filename=""
Content-Type: application/octet-stream


-----------------------------3354441481387117281646263294--
`

func TestSize(t *testing.T) {
	fmt.Println(len(str))
}
