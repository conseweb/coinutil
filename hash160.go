// Copyright (c) 2013-2014 The btcsuite developers
// Use of this source code is governed by an ISC
// license that can be found in the LICENSE file.

package coinutil

import (
	"hash"

<<<<<<< HEAD
	"github.com/btcsuite/golangcrypto/ripemd160"
	"github.com/conseweb/fastsha256"
=======
	"github.com/btcsuite/fastsha256"
	"github.com/btcsuite/golangcrypto/ripemd160"
>>>>>>> 9f1e172e65222676bc8b0dd85a5ea803f1dcecd8
)

// Calculate the hash of hasher over buf.
func calcHash(buf []byte, hasher hash.Hash) []byte {
	hasher.Write(buf)
	return hasher.Sum(nil)
}

// Hash160 calculates the hash ripemd160(sha256(b)).
func Hash160(buf []byte) []byte {
	return calcHash(calcHash(buf, fastsha256.New()), ripemd160.New())
}
