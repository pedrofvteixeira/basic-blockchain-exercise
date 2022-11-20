package core

import (
	"crypto/sha1"
	"encoding/hex"
	"hash"
)

func asHash(s string) hash.Hash {
	h := sha1.New()
	h.Reset()
	h.Write([]byte(s))
	return h
}

func asString(h hash.Hash, bytes []byte) string {
	if h == nil {
		h = asHash("")
	}

	h.Write(bytes)
	return hex.EncodeToString(h.Sum(nil))
}
