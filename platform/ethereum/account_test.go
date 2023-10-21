package ethereum_test

import (
	"testing"

	"github.com/rhiethereum/network/platform/ethereum"
)

const testPrivateKey = "59c6995e998f97a5a0044966f0945389dc9e86dae88c7a8412f4603b6b78690d"

func TestFlow_AccountHandle(t *testing.T) {
	account, err := ethereum.NewAccountFromPrivateKey(testPrivateKey)
	if err != nil {
		t.Error(err)
	}

	accountAddress := account.GetAddress()
	accountPubKey := account.GetPublicKey()
	accountPrivKey := account.GetPrivateKey()

	if accountAddress.String() != "0x70997970C51812dc3A010C7d01b50e0d17dc79C8" {
		t.Errorf("unexpected address")
	}
	if accountPubKey != "04ba5734d8f7091719471e7f7ed6b9df170dc70cc661ca05e688601ad984f068b0d67351e5f06073092499336ab0839ef8a521afd334e53807205fa2f08eec74f4" {
		t.Errorf("unexpected public key")
	}
	if accountPrivKey != testPrivateKey {
		t.Errorf("unexpected private key")
	}

}
