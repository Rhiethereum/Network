package hdwallet

import (
	"encoding/hex"
	"fmt"
	"strconv"
	"strings"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
)

const (
	MNEMONIC_WORD_LEN_12 = 128
	MNEMONIC_WORD_LEN_24 = 256
)

// https://github.com/bitcoin/bips/blob/master/bip-0044.mediawiki
// m / purpose' / coin_type' / account' / change / address_index
const APOSTROPHE uint32 = 0x80000000 // 0', ', 2147483648

const (
	ETHEREUM_HD_PATH = "m'/44'/60'/0'/0"
)

type HDWallet struct {
	mnemonic   string
	passphrase string
	seed       []byte
	masterKey  *bip32.Key
	baseHDPath string
}

func NewMnemonicFromEntropy(bitSize int) (string, error) {
	// bitSize has to be a multiple 32 and be within the inclusive range of {128, 256}
	entropy256, err := bip39.NewEntropy(bitSize)

	if err != nil {
		return "", fmt.Errorf("fail to create entropy: %v", err)
	}

	mnemonic, err := bip39.NewMnemonic(entropy256)

	if err != nil {
		return "", fmt.Errorf("fail to create mnemonic: %v", err)
	}

	return mnemonic, nil
}

func NewHDWallet(mnemnonic string) (*HDWallet, error) {
	return NewHDWalletWithOptions(mnemnonic, "", ETHEREUM_HD_PATH)
}

func NewHDWalletWithOptions(mnemonic string, password string, hdPath string) (*HDWallet, error) {
	isValidMnemonic := bip39.IsMnemonicValid(mnemonic)

	if !isValidMnemonic {
		return nil, fmt.Errorf("invalid mnemonic string")
	}

	seed := bip39.NewSeed(mnemonic, password)
	masterKey, err := bip32.NewMasterKey(seed)
	if err != nil {
		return nil, fmt.Errorf("failed to create master key: %v", err)
	}

	return &HDWallet{
		mnemonic:   mnemonic,
		passphrase: password,
		seed:       seed,
		masterKey:  masterKey,
		baseHDPath: hdPath,
	}, nil
}

func (hdwallet *HDWallet) DerivedPathToPrivateKey(index uint32) (string, error) {
	parsedPath, err := parseHDPath(hdwallet.baseHDPath)

	if err != nil {
		return "", fmt.Errorf("fail to parse HD path: %v", err)
	}

	key := hdwallet.masterKey

	for i := 0; i < len(parsedPath); i++ {
		key, err = key.NewChildKey(parsedPath[i])
		if err != nil {
			return "", fmt.Errorf("invalid child key index: %v", err)
		}
	}

	key, err = key.NewChildKey(index)
	if err != nil {
		return "", fmt.Errorf("invalid child key index: %v", err)
	}

	return hex.EncodeToString(key.Key), nil
}

func parseHDPath(path string) ([]uint32, error) {
	parsedPath := []uint32{}
	pathElems := strings.Split(path, "/")

	if pathElems[0] != "m'" {
		return nil, fmt.Errorf("master key is not included in HD path")
	}

	for i := 1; i < len(pathElems); i++ {
		unum32, err := parseHDPathElement(pathElems[i])
		if err != nil {
			return nil, fmt.Errorf("fail to parse HD path element: %v", err)
		}
		parsedPath = append(parsedPath, unum32)
	}

	return parsedPath, nil
}

func parseHDPathElement(pathElem string) (uint32, error) {

	var isApostropheExist bool

	if strings.Contains(pathElem, "'") {
		pathElem = strings.Replace(pathElem, "'", "", 1)
		isApostropheExist = true
	}

	unum64, err := strconv.ParseUint(pathElem, 10, 32)
	if err != nil {
		return 0, fmt.Errorf("fail to convert string to uint64: %v", err)
	}

	unum32 := uint32(unum64)

	if isApostropheExist {
		return APOSTROPHE + unum32, nil
	}

	return unum32, nil
}
