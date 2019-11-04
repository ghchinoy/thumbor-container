package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"
)

func main() {
	// guards
	// must have parseable URL
	if len(os.Args) < 2 {
		log.Fatalln("Please provide an URL")
	}
	unsafeURL, err := url.Parse(os.Args[1])
	if err != nil {
		log.Fatalln(err)
	}

	// must have SECURITY_KEY
	hmacSecurity := os.Getenv("SECURITY_KEY")
	if hmacSecurity == "" {
		log.Fatalln("Please provide an env var SECURITY_KEY")
	}

	// must have an origin
	safeOrigin, err := getSafeOrigin(unsafeURL)

	// create an HMAC signature for the path
	unsafePath := strings.Replace(unsafeURL.Path, "/unsafe/", "", 1)
	log.Println("HMAC of the following:", unsafePath)
	mac := hmac.New(sha1.New, []byte(hmacSecurity))
	mac.Write([]byte(unsafePath))
	signature := mac.Sum(nil)
	encodedSignature := base64.StdEncoding.EncodeToString([]byte(signature))

	encodedSignature = urlsafe(encodedSignature)
	log.Println(encodedSignature)

	fmt.Printf("%s/%s/%s\n", safeOrigin, encodedSignature, unsafePath)

}

// getSafeOrigin returns the proper scheme://host:port for either
// the env var or the given url
func getSafeOrigin(unsafeURL *url.URL) (string, error) {
	safeOrigin := os.Getenv("THUMBOR_ORIGIN")
	if safeOrigin == "" {
		log.Println("No alternative origin provided. Using provided origin:", unsafeURL.Hostname())
		// reconstruct scheme://host:port
		var port string
		if unsafeURL.Port() != "" {
			port = fmt.Sprintf(":%s", unsafeURL.Port())
		}
		safeOrigin = fmt.Sprintf("%s://%s%s", unsafeURL.Scheme, unsafeURL.Hostname(), port)
	}
	return safeOrigin, nil
}

// urlsafe is an implementation of python's urlsafe_b64encode, but on a string
// // https://docs.python.org/2/library/base64.html
// // Encode string s using the URL- and filesystem-safe alphabet, which
// // substitutes - instead of + and _ instead of / in the standard Base64
// // alphabet. The result can still contain =.
func urlsafe(in string) string {
	return strings.Replace(
		strings.Replace(
			// substitute - instead of +
			in, "+", "-", -1,
		),
		// substitute _ instead of /
		"/", "_", -1)
}
