package main

import (
	"bytes"
	"compress/gzip"
	"net/http"
	"strings"
)

func removeDefaultHeaders(w http.ResponseWriter) {
	w.Header()["Date"] = nil
	w.Header()["Server"] = nil
}

func setEncoding(w http.ResponseWriter, r *http.Request) string {
	headers := r.Header

	acceptEncodingHeaderKey := http.CanonicalHeaderKey("Accept-Encoding")

	var clientSupportedEncoding []string

	if headers[acceptEncodingHeaderKey] != nil && len(headers[acceptEncodingHeaderKey]) > 0 {
		clientSupportedEncoding = strings.Split(headers[acceptEncodingHeaderKey][0], ", ") // splitting by comma and space
	} else {
		// encoding not required
		return ""
	}

	commonEncodingScheme := ""

	// find a supported encoding common with client
	for _, ce := range clientSupportedEncoding {
		for _, se := range serverConfig.supportedEncoding {
			if ce == se {
				commonEncodingScheme = ce
			}
		}
	}

	if commonEncodingScheme == "" {
		// no common encoding found
		// response won't be compressed
		return ""
	}

	w.Header().Set("Content-Encoding", commonEncodingScheme)
	return commonEncodingScheme
}

func compressWithGzip(content string) []byte {
	var b bytes.Buffer
	w := gzip.NewWriter(&b)
	w.Write([]byte(content))
	w.Close()

	return b.Bytes()
}
