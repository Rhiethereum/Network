package ethereum_test

import (
	"encoding/hex"
	"strings"
	"testing"

	"github.com/rhiethereum/network/platform/ethereum"
)

func TestFunc_GenerateSecret(t *testing.T) {
	byteLen := 32
	randomBytes, err := ethereum.GenerateSecret(byteLen)
	if err != nil {
		t.Error(err)
	}
	if len(randomBytes) != byteLen {
		t.Errorf("returned secret length is not valid")
	}
}

func TestFunc_GetSecretHashFrom(t *testing.T) {
	hexdata := "b556de08f3bd3f094a26b006efbe725dd13f26689b4470f52d8ab2d65acb2347"
	bytedata, err := hex.DecodeString(hexdata)
	if err != nil {
		t.Error(err)
	}
	secretHash := ethereum.GetSecretHashFrom(bytedata)
	if secretHash != "0xc996ef7355cda61a0c06938db2af2d314d73bc128186353e43b8df05f7849e4d" {
		t.Errorf("unexpected secret hash")
	}
}

func TestFunc_GetSecretHashByte32From(t *testing.T) {
	secretHash := "0xc996ef7355cda61a0c06938db2af2d314d73bc128186353e43b8df05f7849e4d"
	secretHashBytes, err := hex.DecodeString(strings.Replace(secretHash, "0x", "", 1)) // slice:[]byte
	if err != nil {
		t.Error(err)
	}

	secretHashByte32, err := ethereum.GetSecretHashByte32From(secretHash) // array:[32]byte
	if err != nil {
		t.Error(err)
	}

	if len(secretHashBytes) != len(secretHashByte32) {
		t.Errorf("unexpected secret hash bytes length")
	}

	for i := 0; i < 32; i++ {
		if secretHashBytes[i] != secretHashByte32[i] {
			t.Errorf("both of bytes are not same each other")
		}
	}

}
