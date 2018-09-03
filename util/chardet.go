package util

import (
	"github.com/sanguohot/chardet"
	"net/http"
	"strings"
)

var defaultCharset = "utf-8"
var gbk = "gbk"
func DetectCharsetWithOnlyUtf8OrGbk(data []byte) string {
	strs := chardet.Possible(data)
	if strs[0] == defaultCharset {
		return defaultCharset
	}
	foundGbk := false
	for _, value := range strs {
		if value == gbk {
			foundGbk = true
		}
	}
	if foundGbk {
		return gbk
	}
	return defaultCharset
}
func DetectContentType(data []byte) string {
	contentType := http.DetectContentType(data)
	if strings.Index(contentType, defaultCharset) >= 0 {
		contentType = strings.Replace(contentType, defaultCharset, DetectCharsetWithOnlyUtf8OrGbk(data), -1)
	}
	return contentType
}
