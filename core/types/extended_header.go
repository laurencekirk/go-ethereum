package types

import (
	"fmt"
	"math/big"
	"github.com/ethereum/go-ethereum/common/hexutil"
)

const (
	signatureLength = 65
)
//go:generate gencodec -type ExtendedHeader -field-override extendedHeaderMarshaling -out gen_extended_header_json.go

// A Signature is a 65 byte ECDSA signature in the [R || S || V] format where V is 0 or 1.
type Signature [signatureLength]byte

// Extended Header is a simple data container for storing extra data - that makes up part of the extended protocol
type ExtendedHeader struct {
	Seed  			*big.Int		`json:"seed"       gencodec:"required"`
	Signature 		Signature		`json:"signature"   gencodec:"required"`
}

// field type overrides for gencodec
type extendedHeaderMarshaling struct {
	Seed *hexutil.Big
}

func (eh ExtendedHeader) String() string {
	return fmt.Sprintf(`
[
	Seed:			%v
	Signature:		%v
]
`, eh.Seed, string(eh.Signature[:]))
}

func (b *Block) ExtendedHeader() *ExtendedHeader            { return b.header.ExtendedHeader }

// SetExtendedHeader converts a byte slice to a ExtendedHeade.
// It panics if b is not of suitable size.
func (h *Header) SetExtendedHeader(sig []byte) {
	if len(sig) != signatureLength {
		panic(fmt.Sprintf("The signature to be used in the Extended Header is not the correct size: expected %d; got %d", signatureLength, len(sig)))
	}

	copy(h.ExtendedHeader.Signature[:], sig[:])
}