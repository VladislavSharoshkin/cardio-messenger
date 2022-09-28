package crypto

import (
	"crypto/rand"
	"crypto/sha256"
	"crypto/x509"
	"encoding/base64"
	"encoding/hex"
	"encoding/pem"

	"github.com/mr-tron/base58"

	"io/ioutil"

	"github.com/zeebo/xxh3"

	"strconv"
	"strings"
)

func Sha(data []byte) []byte {
	h := sha256.New()
	h.Write(data)
	return h.Sum(nil)
}

func ShaString(data string) string {
	return hex.EncodeToString(Sha([]byte(data)))
}

func RandomString() string {
	bytes := make([]byte, 32)
	rand.Read(bytes)
	return hex.EncodeToString(bytes)
}

func RandomBytes(len int) []byte {
	data := make([]byte, len)
	rand.Read(data)
	return data
}

func RandomBytes256() []byte {
	return RandomBytes(32)
}

func ToBase64(data []byte) string {
	sEnc := base64.StdEncoding.EncodeToString(data)
	return sEnc
}

func FromBase64(data string) []byte {
	sDec, _ := base64.StdEncoding.DecodeString(data)
	return sDec
}

func PasswordHash(password string, salt []byte) string {
	h := sha256.New()
	h.Write([]byte(password))
	h.Write(salt)

	var hashedPassword = ToBase64(h.Sum(nil)) + "." + ToBase64(salt)
	return hashedPassword
}

func PasswordCheck(password string, hashedPassword string) bool {
	var hashAndSalt = strings.Split(hashedPassword, ".")
	if len(hashAndSalt) < 2 {
		return false
	}

	var salt = FromBase64(hashAndSalt[1])

	if PasswordHash(password, salt) == hashedPassword {
		return true
	}
	return false
}

func XxHash64(data string) string {
	return strconv.Itoa(int(xxh3.HashString(data)))
}

func CertFingerprint(path string) (string, error) {
	pemContent, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}

	block, _ := pem.Decode(pemContent)
	if block == nil {
		return "", err
	}

	// pass cert bytes
	cert, err := x509.ParseCertificate(block.Bytes)
	if err != nil {
		return "", err
	}

	fingerprint := Sha(cert.Raw)
	return hex.EncodeToString(fingerprint), nil
}

func NewToken() string {
	return base58.Encode(RandomBytes256())
}
