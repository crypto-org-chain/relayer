package client

import (
	cmldsa65 "github.com/cometbft/cometbft/crypto/mldsa65"
	cpc "github.com/cometbft/cometbft/proto/tendermint/crypto"
	slcrypto "github.com/strangelove-ventures/cometbft-client/crypto"
	sljson "github.com/strangelove-ventures/cometbft-client/libs/json"
	slpc "github.com/strangelove-ventures/cometbft-client/proto/tendermint/crypto"
)

// cometbft-client has no ML-DSA-65 types, so its JSON decoder rejects any response carrying such a key.
func init() {
	sljson.RegisterType(mldsa65PubKey{}, cmldsa65.PubKeyName)
	sljson.RegisterType((*protoMldsa65PubKey)(nil), "tendermint.crypto.PublicKey_Mldsa65")
}

var _ slcrypto.PubKey = mldsa65PubKey{}

// mldsa65PubKey adapts cometbft's ML-DSA-65 key to cometbft-client's PubKey interface.
type mldsa65PubKey struct {
	cmldsa65.PubKey
}

func (k mldsa65PubKey) Address() slcrypto.Address {
	return slcrypto.AddressHash(k.Bytes())
}

func (k mldsa65PubKey) Equals(other slcrypto.PubKey) bool {
	o, ok := other.(mldsa65PubKey)
	return ok && k.PubKey.Equals(o.PubKey)
}

// protoMldsa65PubKey decodes the ML-DSA-65 case of the proto PublicKey oneof in validator updates.
// PublicKey_Ed25519 is embedded only for the oneof's unexported marker method; every method that
// reads the key is overridden below so it is never mistaken for an empty ed25519 key.
type protoMldsa65PubKey struct {
	slpc.PublicKey_Ed25519 `json:"-"`

	Mldsa65 []byte `json:"mldsa65,omitempty"`
}

func (k *protoMldsa65PubKey) cometbft() *cpc.PublicKey_Mldsa65 {
	if k == nil {
		return nil
	}
	return &cpc.PublicKey_Mldsa65{Mldsa65: k.Mldsa65}
}

func (k *protoMldsa65PubKey) Equal(that interface{}) bool {
	if o, ok := that.(protoMldsa65PubKey); ok {
		that = &o
	}
	o, ok := that.(*protoMldsa65PubKey)
	if !ok {
		return that == nil && k == nil
	}
	return k.cometbft().Equal(o.cometbft())
}

func (k *protoMldsa65PubKey) Compare(that interface{}) int {
	if o, ok := that.(protoMldsa65PubKey); ok {
		that = &o
	}
	o, ok := that.(*protoMldsa65PubKey)
	if !ok {
		if that == nil && k == nil {
			return 0
		}
		return 1
	}
	return k.cometbft().Compare(o.cometbft())
}

func (k *protoMldsa65PubKey) MarshalTo(dAtA []byte) (int, error) {
	return k.cometbft().MarshalTo(dAtA)
}

func (k *protoMldsa65PubKey) MarshalToSizedBuffer(dAtA []byte) (int, error) {
	return k.cometbft().MarshalToSizedBuffer(dAtA)
}

func (k *protoMldsa65PubKey) Size() int {
	return k.cometbft().Size()
}
