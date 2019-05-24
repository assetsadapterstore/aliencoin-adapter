package aliencoin_addrdec

import (
	"github.com/blocktree/go-owcdrivers/addressEncoder"
)

const (
	btcAlphabet = "123456789ABCDEFGHJKLMNPQRSTUVWXYZabcdefghijkmnopqrstuvwxyz"
)

var (
	ALC_mainnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x17}, Suffix: nil}
	ALC_testnetAddressP2PKH         = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x74}, Suffix: nil}
	ALC_mainnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0x2e}, Suffix: []byte{0x01}}
	ALC_testnetPrivateWIFCompressed = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "", HashLen: 32, Prefix: []byte{0xc4}, Suffix: []byte{0x01}}
	ALC_mainnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0x30}, Suffix: nil}
	ALC_testnetAddressP2SH          = addressEncoder.AddressType{EncodeType: "base58", Alphabet: btcAlphabet, ChecksumType: "doubleSHA256", HashType: "h160", HashLen: 20, Prefix: []byte{0xef}, Suffix: nil}
	ALC_mainnetAddressBech32V0      = addressEncoder.AddressType{"bech32", btcAlphabet, "alc", "h160", 20, nil, nil}
	ALC_testnetAddressBech32V0      = addressEncoder.AddressType{"bech32", btcAlphabet, "tb", "h160", 20, nil, nil}

	Default = AddressDecoderV2{}
)

//AddressDecoderV2
type AddressDecoderV2 struct {
	IsTestNet bool
}

//AddressDecode 地址解析
func (dec *AddressDecoderV2) AddressDecode(addr string, opts ...interface{}) ([]byte, error) {

	cfg := ALC_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = ALC_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	return addressEncoder.AddressDecode(addr, cfg)
}

//AddressEncode 地址编码
func (dec *AddressDecoderV2) AddressEncode(hash []byte, opts ...interface{}) (string, error) {

	cfg := ALC_mainnetAddressP2PKH
	if dec.IsTestNet {
		cfg = ALC_testnetAddressP2PKH
	}

	if len(opts) > 0 {
		for _, opt := range opts {
			if at, ok := opt.(addressEncoder.AddressType); ok {
				cfg = at
			}
		}
	}

	address := addressEncoder.AddressEncode(hash, cfg)
	return address, nil
}
