package hdwallet_test

import (
	"strings"
	"testing"

	hdwallet "github.com/rhiethereum/network/platform/ethereum/ethereum-hdwallet"
)

const mnemonic = "test test test test test test test test test test test junk"

func TestFunc_NewMnemonicFromEntropy(t *testing.T) {
	mnemnonic, err := hdwallet.NewMnemonicFromEntropy(hdwallet.MNEMONIC_WORD_LEN_12)
	if err != nil {
		t.Error(err)
	}
	if len(strings.Split(mnemnonic, " ")) != 12 {
		t.Error("mnemonic words length invalid")
	}
}

func TestFlow_HDWalletHandle(t *testing.T) {
	ethwallet, err := hdwallet.NewHDWallet(mnemonic)
	if err != nil {
		t.Error(err)
	}

	priv0, err := ethwallet.DerivedPathToPrivateKey(0)
	if err != nil {
		t.Error(err)
	}

	priv1, err := ethwallet.DerivedPathToPrivateKey(1)
	if err != nil {
		t.Error(err)
	}

	if priv0 != "ac0974bec39a17e36ba4a6b4d238ff944bacb478cbed5efcae784d7bf4f2ff80" {
		t.Error("HDWallet index 0 invalid")
	}

	if priv1 != "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d" {
		t.Error("HDWallet index 1 invalid")
	}
}
