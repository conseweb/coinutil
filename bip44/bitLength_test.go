package bip44

import (
	"testing"

	. "github.com/conseweb/coinutil/bip44"
)

func TestBitLength(t *testing.T) {
	child, err := NewKeyFromMnemonic(
		"element fence situate special wrap snack method volcano busy ribbon neck sphere",
		TypeFactomFactoids,
		2147483648,
		0,
		19,
	)

	if err != nil {
		t.Errorf("%v", err)
	}
	k := child.Key()
	if len(k) != 32 {
		t.Errorf("len: %d, child.Key:%x\n", len(k), k)
		t.Errorf("%v", child.String())
	}

	child, err = NewKeyFromMnemonic(
		"element fence situate special wrap snack method volcano busy ribbon neck sphere",
		TypeFactomFactoids,
		2147483648,
		1,
		19,
	)

	if err != nil {
		t.Errorf("%v", err)
	}
	k = child.Key()
	if len(k) != 32 {
		t.Errorf("len: %d, child.Key:%x\n", len(k), k)
		t.Errorf("%v", child.String())
	}
}
