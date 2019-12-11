package main

import (
	"./gorsa"
	"errors"
	"log"
	"time"
)

var (
	//server keys
	server_private_key = `-----BEGIN PRIVATE KEY-----
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
6OQ+6oo+1RR3mZOjksu9Cwm40ND/lEn2
-----END PRIVATE KEY-----`
	server_public_key = `-----BEGIN PUBLIC KEY-----
MIICIjANBgkqhkiG9w0BAQEFAAOCAg8AMIICCgKCAgEA28OcOthPM25rk+6wzv+c
xk6YGUkkcUyQjv+Z+IjxT4t/JvwWcEGyAwAIbjTyIyWHyniNj8vI6qEuWIQVlZHo
FSTsmLWSyNGn+FrkZtp2ZBUUUMce7cEeQOxTzNEI3NiuuqYvr2SqP3j0FdQFIehz
jicxWzALbrtMe9Odk/gnrRrRWoNjZi6Nox79iDuDqH0TWyqLRuTyaCwDxIEIKFTd
rwOv/pXohrwkYyQE/NAJuNOO2GSJMGKoP09R6+/aJLhw0EKX/MKroGc9+Ynb00Oz
FxfTIli1riheE4enQmA4TQi7b1V0WSneUsJMEmncq+GRaLoCPiXOjWm+CeyJNLU2
IsnYioXKrNGnquvq1p9s9ypWqKlMLVYOjKk1Izpyfk7Difan/9gmmdVjklgEwID9
/F642qzWVjivx3nx7ViJd1cR2gqOg7xRQ/63LPu67ZE9Z5+NdRdSlwIC+BiGF7m4
hqx5154jo2TVuEjy0SRymLab/ZeYmnrrMxfOzVuH6T0aEo1gE06r/SeEDRRnugw/
Y/9o1d5dWJ4Ebr5+gdwGt3y1/ua/GC70XjOV6t8lrOkPBAqlx1WRL0+ZbWdpifc/
zxyvPPmHsc7KFP6+AvpW+1Vtv9vgLlFKtP+BmpifanO85iv29s/Ql8Nd10SFlAhP
mSDdJhlv5yW75S60/vzXFKsCAwEAAQ==
-----END PUBLIC KEY-----`

	//client keys
	client_public_key  = ``
	client_private_key = ``
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
		print(enctyptedStr)
	}

	// 公钥加密私钥解密
	//if enctyptedStr, err := applyPubEPriD(data); err != nil {
	//	log.Println(err)
	//} else {
	//	print(enctyptedStr)
	//}

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
