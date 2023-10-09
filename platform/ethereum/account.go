package ethereum

import (
	"crypto/ecdsa"
	"encoding/hex"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
)

type EthereumAccount struct {
	privateKey *ecdsa.PrivateKey
	publicKey  *ecdsa.PublicKey
	address    string
}

func NewAccountFromPrivateKey(privateKeyHex string) (*EthereumAccount, error) {
	privateKeyECDSA, err := crypto.HexToECDSA(strings.TrimPrefix(privateKeyHex, "0x"))
	if err != nil {
		return nil, err
	}
	publicKey := privateKeyECDSA.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}
	address := crypto.PubkeyToAddress(*publicKeyECDSA).Hex()
	return &EthereumAccount{
		privateKey: privateKeyECDSA,
		publicKey:  publicKeyECDSA,
		address:    address,
	}, nil
}

func (account *EthereumAccount) GetPrivateKey() string {
	return hex.EncodeToString(crypto.FromECDSA(account.privateKey))
}

func (account *EthereumAccount) GetPublicKey() string {
	return hex.EncodeToString(crypto.FromECDSAPub(account.publicKey))
}

func (account *EthereumAccount) GetAddress() common.Address {
	return crypto.PubkeyToAddress(*account.publicKey)
}

func (account *EthereumAccount) Sign(chainid *big.Int, tx *types.Transaction) (*types.Transaction, error) {
	signedTx, err := types.SignTx(
		tx,
		types.NewCancunSigner(chainid),
		account.privateKey,
	)

	if err != nil {
		return nil, err
	}

	return signedTx, nil
}
