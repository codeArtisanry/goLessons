package utils

import (
	"crypto"
	"crypto/hmac"
	"crypto/rand"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"log"
	mrand "math/rand"
	"time"
	//"crypto/sha1"
	"strconv"
	"strings"

	"encoding/base64"

	"github.com/gorilla/feeds"
)

const OneDayMilliSec = int64(24 * 3600 * 1000)
const millisecondUit = int64(time.Millisecond / time.Nanosecond)

var passwordEncryptKeyUser = []byte("appwillgoogle@bj001001")
var passwordEncryptKeyAdmin = []byte("appwillgoogle@blaxianyou")
var noProtectForTask bool

// only for other pkg's init func
func InitConf(_noProtectForTask bool) {
	noProtectForTask = _noProtectForTask
}

func GenerateKey() string {
	u := feeds.NewUUID()
	return fmt.Sprintf("%x%x%x%x%x", u[:4], u[4:6], u[6:8], u[8:10], u[10:])
}

func GetNowSecond() int64 {
	return int64(time.Now().Unix())
}

func GetNowMillisecond() int64 {
	return time.Now().UnixNano() / millisecondUit
}

func GetNowNanosecond() int64 {
	return time.Now().UnixNano()
}

func GetNowStringYMD() string {
	return time.Now().Format("2006-01-02")
}

func GetNowStringYMDHMS() string {
	return time.Now().Format("20060102150405")
}

func GetDateFromUTC(msUTC int64) string {
	return time.Unix(msUTC/1000, (msUTC%1000)*millisecondUit).Format("2006-01-02")
}

func GenDateListFromUTCDur(msStartUTC, msEndUTC int64) []string {
	if msStartUTC > msEndUTC {
		panic("start utc > end_utc")
	}

	var res []string
	endDate := GetDateFromUTC(msEndUTC)
	lastDate := ""
	for t := msStartUTC; ; t += 3600 * 24 * 1000 {
		_date := GetDateFromUTC(t)
		if len(res) == 0 {
			res = append(res, _date)
			continue
		}
		if _date > endDate || (_date == endDate && _date == lastDate) {
			break
		}

		res = append(res, _date)
		lastDate = _date
	}

	return res
}

func GetRandomNumberStr() []byte {
	len := 6
	res := make([]byte, len)
	r := mrand.New(mrand.NewSource(time.Now().UnixNano()))

	for i := 0; i < len; i++ {
		num := r.Intn(10)
		numStr := strconv.Itoa(num)
		res[i] = []byte(numStr)[0]
	}
	return res[:6]
}

func EncryptUserPassCode(code string) string {
	hash := hmac.New(sha256.New, passwordEncryptKeyUser)
	hash.Write([]byte(code))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func EncryptAdminPassCode(code string) string {
	hash := hmac.New(sha256.New, passwordEncryptKeyUser)
	hash.Write([]byte(code))
	return fmt.Sprintf("%x", hash.Sum(nil))
}

func SendChanNonblock(ch chan interface{}, i interface{}) {
	select {
	case ch <- i: //
	default: //
	}
}

func DoEverTask(f func() time.Duration, task_name string, defaultInterval time.Duration) {
	if defaultInterval < time.Second {
		defaultInterval = time.Second
	}

	go func() {
		for {
			func() {
				if !noProtectForTask {
					defer func() {
						i := recover()
						log.Println("task panic: ", i, task_name)
					}()
				}
				for {
					if interval := f(); interval > 0 {
						time.Sleep(interval)
					} else {
						time.Sleep(defaultInterval)
					}
				}
			}()
			time.Sleep(defaultInterval)
		}
	}()
}

func GenRsaPrivateKey(privKeyBS []byte) (*rsa.PrivateKey, error) {
	block, _ := pem.Decode(privKeyBS)
	if block == nil {
		return nil, fmt.Errorf("key <%s> error", string(privKeyBS))
	}

	return x509.ParsePKCS1PrivateKey(block.Bytes)
}

func RsaSignWithPrivateKey(hashed []byte, privKey *rsa.PrivateKey, hashKind crypto.Hash) ([]byte, error) {
	sign, err := rsa.SignPKCS1v15(rand.Reader, privKey, hashKind, hashed)
	if err != nil {
		return nil, err
	}

	return []byte(base64.StdEncoding.EncodeToString(sign)), nil
}

func RsaVerify(hashed []byte, sign []byte, pubKey []byte, hashKind crypto.Hash) error {
	block, _ := pem.Decode(pubKey)
	pub, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		return err
	}

	rsaPub, _ := pub.(*rsa.PublicKey)
	data, _ := base64.StdEncoding.DecodeString(string(sign))

	return rsa.VerifyPKCS1v15(rsaPub, hashKind, hashed, data)
}

func GenRandomString(len int) string {
	strBS := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}
	var res []string
	for i := 0; i < len; i++ {
		if mrand.Intn(100)%3 == 0 {
			res = append(res, strconv.Itoa(1+mrand.Intn(9)))
		} else {
			if mrand.Intn(100)%2 == 0 {
				res = append(res, strings.ToUpper(strBS[mrand.Intn(26)]))
			} else {
				res = append(res, strBS[mrand.Intn(26)])
			}
		}
	}

	return strings.Join(res, "")
}
