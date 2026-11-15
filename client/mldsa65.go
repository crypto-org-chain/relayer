package client

import (
	cmldsa65 "github.com/cometbft/cometbft/crypto/mldsa65"
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
// The embedded type only supplies the oneof's unexported marker method; RPCClient drops validator updates.
type protoMldsa65PubKey struct {
	slpc.PublicKey_Ed25519 `json:"-"`

	Mldsa65 []byte `json:"mldsa65,omitempty"`
}
