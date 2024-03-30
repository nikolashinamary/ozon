package generating

import (
	"Ozon/Storages"
	"Ozon/payload"
	"crypto/md5"
	"crypto/sha256"
	"encoding/base64"
	"github.com/sirupsen/logrus"
	"time"
)

const (
	supportedLength = 10
	maxTrials       = 1000
)

func generateHash(url string) string {
	currentTime := time.Now()
	timeString := currentTime.GoString()
	hasher := md5.New()
	stringToEncode := url + timeString
	hasher.Write([]byte(stringToEncode))
	hashBytes := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(hashBytes)
}

func generateDoubleHash(url string) string {
	currentTime := time.Now()
	timeString := currentTime.GoString()
	hasher := sha256.New()
	stringToEncode := url + timeString
	hasher.Write([]byte(stringToEncode))
	hashBytes := hasher.Sum(nil)
	return base64.URLEncoding.EncodeToString(hashBytes)
}

func checkSymbols(url string) string {
	for i := 0; i < len(url); i++ {
		if url[i] <= 47 || (url[i] >= 58 && url[i] <= 64) || (url[i] >= 91 && url[i] <= 94) || url[i] == 96 || url[i] >= 123 {
			url = url[:i] + url[i+1:]
		}
	}
	return url

}

func checkExistance(url string, stor Storages.Storage, log1 *logrus.Logger) (bool, string) {
	exist := true
	for i := 0; i < len(url)-supportedLength+1; i++ {
		k := stor.ContainsByHash(url[i:i+supportedLength], log1)
		if !k {
			exist = false
			url = url[i : i+supportedLength]
			break
		}
	}
	return exist, url
}

func GenerateShortenURL(request payload.RequestEncode, stor Storages.Storage, log1 *logrus.Logger) (bool, string) {
	hashUrl := generateHash(request.Url)
	hashUrl = generateDoubleHash(hashUrl)
	successfulGenerated := false
	for i := 0; i < maxTrials; i++ {
		url := checkSymbols(hashUrl)
		if len(url) < supportedLength {
			hashUrl = generateDoubleHash(generateHash(request.Url))
			continue
		}
		flag, url := checkExistance(url, stor, log1)
		if flag {
			hashUrl = generateDoubleHash(generateHash(request.Url))
			continue
		}
		hashUrl = url
		successfulGenerated = true
		break

	}
	return successfulGenerated, hashUrl
}
