package main

import (
	"./gorsa"
	"errors"
	"log"
	"time"
)

var (
	//server keys
	server_private_key = `-----BEGIN RSA PRIVATE KEY-----
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
-----END RSA PRIVATE KEY-----
`
	server_public_key = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA2dKBz9XY7NUKlOnTHne0
8edeuZmi+P7CGk7M75Bl9k4ddkI+s7lLzB9fAUr34OeLtPLEO6A2sj/CCFORzovA
95/bJdoPhCd61Yd9aZ4g9q0jW1/qy5z2bjdz4Q/HI5fjuiqdW/ga6pQp/UYRMMKU
i6KFzsZdFHTBHnJkBA4QyOaSzj0KGbWrUOhXvBe1T4xsyGisDQequ7u4IBm8BWF5
laZsMiqJlDucVwWbuOS2JLF7hl50w3wvV8fOLnb2evZaHhbf1lcxUMmcPyVIX6YE
AcW8btOqWLIYp2YwsbufCWtHoZ3pyZ4DNNbtiTA0kbYnf98eADq9yOvwuGPbkr5g
AECjg1GycEk+oSLPxbrATnBulXFa/OmC9o5MjsjGoO1RSxqPQV4tcOjOStc5lswk
8RTrjYQigzOqfUAQMjH95W4wjnupd/WF4gMz3tMCF9DIbLLFijk+kN5L8EI3xjFe
UrLYjoVSmfrq1n16DPHtUIAqd6ZZg6QLiCKtjvVeABjb9ORdJ/33MKrVl+qtrDrE
8Sso3Fvxot7OaMFQI836FcQDfqthoo7mpyzW5vdKhGPHDSOC1IErDYi/HX1iJNqW
II8abvrZQ8skxzUxXTZn8yNqr+CN+tHUE1hjN+e6M/wyrk5QLiEzlaM6SCqQ1zo/
hGFwAv3vEMN5TheoP9ru9JUCAwEAAQ==
-----END PUBLIC KEY-----
`

	//client keys
	client_public_key = `-----BEGIN PUBLIC KEY-----
MIIBIjANBgkqhkiG9w0BAQEFAAOCAQ8AMIIBCgKCAQEAn+T2Te0MGRpr4OtdPEyt
iwGY3nSL+5kjCgU3Z8n3CpINXBctnq0cUB8JEoh2J0Q1PTlM2ljmNkrg4nf9IFLD
SUZBjKGeUF/ssOcuosamcyV9m9Ajb6yVT86wHqQnZolVxa1rLXwW5WPyXKR1Zj9G
PcrL+4UjnXn5th39IXn6iuyHIdGFVBkmVhw1K+Py73aOJafdz3IiXQmj4MBChq3Z
jjuM2ieMaYt9NdCCU+RE5+uxFnb55rHEpI3WQMedksgaARo9OixHiLj81+eCrYOd
DAYiQb/h6yvLQ633n269gqF/wKLfR5HvcG4cT86W8DiPDK6kFFUFB9MFNsO5PIVX
MwIDAQAB
-----END PUBLIC KEY-----
`
	client_private_key = `-----BEGIN RSA PRIVATE KEY-----
MIIEowIBAAKCAQEAn+T2Te0MGRpr4OtdPEytiwGY3nSL+5kjCgU3Z8n3CpINXBct
nq0cUB8JEoh2J0Q1PTlM2ljmNkrg4nf9IFLDSUZBjKGeUF/ssOcuosamcyV9m9Aj
b6yVT86wHqQnZolVxa1rLXwW5WPyXKR1Zj9GPcrL+4UjnXn5th39IXn6iuyHIdGF
VBkmVhw1K+Py73aOJafdz3IiXQmj4MBChq3ZjjuM2ieMaYt9NdCCU+RE5+uxFnb5
5rHEpI3WQMedksgaARo9OixHiLj81+eCrYOdDAYiQb/h6yvLQ633n269gqF/wKLf
R5HvcG4cT86W8DiPDK6kFFUFB9MFNsO5PIVXMwIDAQABAoIBAGXcIKSn/GsJj+vr
RKwjVPMXA2HD8WFIvy+k/xTZbs2HGYevJFrPgRfxjYAG+u4s35WvKSx4McAol3bu
ZYzv7ISDg0KpuwcM9OYFgIol7uev3IMkZ06/LSd4Tm5WuNrzZnbV8U86nANqVn16
L2gBqye2R5hlyriJefDwzc9Du74/yk+YwCEWUfCNtFZtEpgngVyg4Awi0mKBcNCW
gmPBmMEUjw9tEGrmfFsXlv+65xuyWjKpiDnyeiwhPwZ+iwk2v4UDauxYo5ItTJxA
SRDQJLPuRTl5GZHCYVTk0ZqUEYB3tQ8KHZK8UIhIekVt3xDQS2ezDOEZYogLqfDZ
8WILRbECgYEAzX6iqIAfoVeTCIurEPfkzX1o9NsCNQ8YmPCiUwcw2pZIvBoBgmx+
IxG3HCmP/lLqDgnmKRE81ywrPnrM+H6uM2zG1+4rqVB7ZfuBOgbmBuXO+ykWDwaJ
+N6XqAaX7atszMDepdNZ1looIBXVlyGgPEARJoob4qGutA43MpPea48CgYEAxzE8
kymPB4u9WZU83K+uJT3MHqZ/XQ2oUxJPx9L0G6Lq9CSfWaob1v9+OpKnuTeit+Dl
qUmM/OG4O99a16eSHXKYcH6wpgCCkQa6QYgCw10/ePgHPZR9meOlMImPkqenKDiW
lpojc2dgt1r5CHLU1H9wOOzFVSk3OjXUu14GWB0CgYBctKqxFcwvP5DnpAryxsME
KXW2MS+XvE2+gaDZkD7r/iymH51at5NYQt+25tE9TS6mUrTxBgPKlvSTZVcfLBdu
v74dICe2ZMamxYYRJZeUuQxdprV65dpM2i7U8NSNtm6JGOHZMbYvy9CaWP9ZBSSC
yexy1I+r+IEWXets9+S55QKBgQCpOI75hojYJEMMY9h8s1MCmGlccJ4q83tYd2oH
nJN9Fv77wlEN35IfKJaYNBcBJt65z/nLw4xfGXNmie2m39kl40DV8QspO62wf0Kk
gWO3jPRrVGryScAaiGbBwVuTf5JmfQyRjQko9V4Y5tRB9SwSCuHPunY6TvV4IltD
gjFpkQKBgF1LjRZ5PiUnoFDnhwI7QhjHj+DpEfV0eYnEooGa7yyKCT+KgqcOLZHk
/qOhn6zC0xEp+5FFKTCe92OsfTyMHNljipMoAWp8ejFQ9cAzUVs2KhDJKQgzkvUr
YJRivCTlXWSFFjhshM1PssSo0caOTiaczlqYMHVQq98N/8GHg/iE
-----END RSA PRIVATE KEY-----
`
)

func main() {

	data := `{
  "start_time":"` + time.Now().UTC().Format("2006-01-02 15:04:05") + `",
  "usable_time":10,
  "user_num":10,
  "total_size":10737418240,
  "company":"company",
  "department":"department",
  "version":"profession",
  "version_number":"1.0.0",
  "uuid":"8687f766-0178-4f9d-9639-004c8a7dfd3d"
}`

	// 私钥加密公钥解密
	if enctyptedStr, err := applyPriEPubD(data); err != nil {
		log.Println(err)
	} else {
		println("私钥加密结果：\n" + enctyptedStr)
	}

	// 公钥加密私钥解密
	if enctyptedStr, err := applyPubEPriD(data); err != nil {
		log.Println(err)
	} else {
		println("公钥加密结果：\n" + enctyptedStr)
	}

}

// 私钥加密公钥解密
func applyPriEPubD(data string) (string, error) {
	prienctypt, err := gorsa.PriKeyEncrypt(data, server_private_key)
	if err != nil {
		return "", err
	}

	pubdecrypt, err := gorsa.PublicDecrypt(prienctypt, server_public_key)
	if err != nil {
		return "", err
	}
	if string(pubdecrypt) != data {
		return "", errors.New(`解密测试失败`)
	}
	return prienctypt, nil
}

// 公钥加密私钥解密
func applyPubEPriD(data string) (string, error) {

	pubenctypt, err := gorsa.PublicEncrypt(data, client_public_key)
	if err != nil {
		return "", err
	}

	pridecrypt, err := gorsa.PriKeyDecrypt(pubenctypt, client_private_key)
	if err != nil {
		return "", err
	}
	if string(pridecrypt) != data {
		return "", errors.New(`解密测试失败`)
	}
	return pubenctypt, nil
}
