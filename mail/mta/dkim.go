package mta

import (
	dkim "github.com/toorop/go-dkim"
)

const (
	// privateKey 公私钥使用 openssl 生成
	privateKey = `-----BEGIN RSA PRIVATE KEY-----
MIICXAIBAAKBgQDdhI8+sgSEagT0wD2Ir8wBXmBlXRhOU8wAdSDPUe7e4w2YI7Ql
5V0On0oxL3l9vIkCpaeE220dwf4CjBumaMPV4xpSt0y8Ujp4SHq1arw99hVQz5S4
GLs3qzFoD5qWiuTmiEldCqxIfK1gCRtfMnxhVNN6HcPE55a6I2W4aX+qgwIDAQAB
AoGAennPZX+xpbCkmtora4afSeZvb7vgM3Z7ZYldIaOQaeWp34NmGEnPgMUrlTRF
fPKf4jnK/FcB5qEamzfyFbj5BSK7yTI/Dao6oeN0Mkag7JLuydqRNilJJyYgpywr
3CgG5+t5qwh3ivwKT1Vnc4wjxVxz2OZpexb8A/E2WFucCDECQQD101fUogihEEra
G3e7n04auyOh4/ZV752ZpZYr71mmIDCxzVhfagaFELfCE5bi9IJT8CSZfiTPRFKt
xPYW+5upAkEA5q+pmjpIzcTQsDlnG3l3OR3D5ogJQx8inmUvtrMyotSUIpdppeDX
piv0NPsZ5J0s8zACByCH1m0L91pziKqQSwJBAOkyNc2WcI0qAXfqOqkXtGYTRPgc
YuCe0Gii9lRzWB4Jx2fEHqNU1x5//3HyV16xCLlLw8yAJ7cfXzdM8w5WXRECQD5R
Jdfr9s7fZCC24Qui/HoJeGpGRXpEZu2zF/ia4ArssjfF/1w4KQlSxl2pl40SiJoJ
VgLm3ssmGh1v6dX5fZECQGDihvOsdQNUI+ro37vV1RvEv5LABWf5DFGWdIsJzsY8
NGRIEZ9tbRbpqAkN9xKgow4mxbdQTwCpTD1MI6e2uS4=
-----END RSA PRIVATE KEY-----
`
)

// DKIM dkim 签名 email 为邮件内容
func DKIM(email *[]byte) error {
	options := dkim.NewSigOptions()
	options.PrivateKey = []byte(privateKey) // 生成的私钥
	options.Domain = "tbycq.com"            // 你的域名
	options.Selector = "dkim"               // DNS解析记录前缀 xxx._domainkey selector 就是这个 xxx
	options.SignatureExpireIn = 3600
	options.BodyLength = 50
	options.Headers = []string{"from", "date", "mime-version", "received", "received"}
	options.AddSignatureTimestamp = true
	options.Canonicalization = "relaxed/relaxed" // relaxed 表示使用宽松签名
	err := dkim.Sign(email, options)
	return err
}
