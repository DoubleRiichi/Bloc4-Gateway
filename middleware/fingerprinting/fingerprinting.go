package fingerprinting

import (
	"crypto/sha512"
	"hash"
	"net"
	"net/http"
	"strings"
)

func stringToSHA512(input string) []byte {
	var hash hash.Hash = sha512.New()
	hash.Write([]byte(input))

	return hash.Sum(nil)
}

func FingerprintClient(r *http.Request) []byte {

	var identifying_string strings.Builder

	ip_addr, _, _ := net.SplitHostPort(r.RemoteAddr)
	real_ip := r.Header.Get("X-Real-Ip")
	forwarded_for := r.Header.Get("X-Forwarded-For")

	identifying_string.WriteString(r.UserAgent())
	identifying_string.WriteString(ip_addr)
	identifying_string.WriteString(real_ip)

	for _, ip := range forwarded_for {
		identifying_string.WriteRune(ip)
	}

	fingerprint_hash := stringToSHA512(identifying_string.String())

	return fingerprint_hash
}
