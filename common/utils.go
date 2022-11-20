package common

import (
	"crypto/md5"
	"hash"
	"os"
	"strings"
	"time"
)

// predicate that checks if the given trimmed string is empty
func IsEmpty(s string) bool {
	return len(strings.TrimSpace(s)) == 0
}

// sleep function in millis
func SleepForMillis(millis int64) {
	time.Sleep(time.Duration(millis) * time.Millisecond)
}

func CreateDirIfNotExist(dir string) {
	if _, err := os.Stat(dir); os.IsNotExist(err) {
		os.MkdirAll(dir, 0700)
	}
}

func AsHash(s string) hash.Hash {
	h := md5.New()
	h.Reset()
	h.Sum([]byte(s))
	return h
}
