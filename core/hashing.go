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

func asHexString(h hash.Hash, bytes []byte) string {
	if h == nil {
		s := ""
		if bytes != nil {
			s = string(bytes)
		}
		h = asHash(s)
	}

	return hex.EncodeToString(h.Sum(nil))
}
