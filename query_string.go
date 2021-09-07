package rhelper

import (
	"net/http"
	"strconv"
	"strings"
)

// QueryStringToString definition
func QueryString(r *http.Request, key string) string {
	return strings.TrimSpace(r.URL.Query().Get(key))
}

func QueryStrings(r *http.Request, key string) []string {
	return r.URL.Query()[key]
}

// QueryStringToInt definition
func QueryStringToInt(r *http.Request, key string, def int) int {
	s := QueryString(r, key)
	if s == "" {
		return def
	}

	v, err := strconv.Atoi(s)
	if err != nil {
		return def
	}

	return v
}

// QueryStringToInt definition
func QueryStringToInt64(r *http.Request, key string, def int64) int64 {
	s := QueryString(r, key)
	if s == "" {
		return def
	}

	v, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return def
	}

	return v
}

// QueryStringToFloat64 definition
func QueryStringToFloat64(r *http.Request, key string, def float64) float64 {
	s := QueryString(r, key)
	if s == "" {
		return def
	}

	v, err := strconv.ParseFloat(s, 64)
	if err != nil {
		return def
	}

	return v
}

// QueryStringToFloat64 definition
func QueryStringToBool(r *http.Request, key string) bool {
	s := QueryString(r, key)
	if s == "" {
		return false
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return false
	}

	return v
}

// QueryStringToFloat64 definition
func QueryStringToPointerBool(r *http.Request, key string) *bool {
	s := QueryString(r, key)
	if s == "" {
		return nil
	}

	v, err := strconv.ParseBool(s)
	if err != nil {
		return nil
	}

	return &v
}
