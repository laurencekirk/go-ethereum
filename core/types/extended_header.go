package types

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
)

const (
	signatureLength = 65
)
//go:generate gencodec -type ExtendedHeader -field-override extendedHeaderMarshaling -out gen_extended_header_json.go

// A Signature is a 65 byte ECDSA signature in the [R || S || V] format where V is 0 or 1.
type Signature [signatureLength]byte

// Extended Header is a simple data container for storing extra data - that makes up part of the extended protocol
type ExtendedHeader struct {
	Seed  			Signature		`json:"seed"       gencodec:"required"`
	Signature 		Signature		`json:"signature"   gencodec:"required"`
}

// field type overrides for gencodec
type extendedHeaderMarshaling struct {
}

func (eh ExtendedHeader) String() string {
	return fmt.Sprintf(`
[
	Seed:			%v
	Signature:		%v
]
`, string(eh.Seed[:]), string(eh.Signature[:]))
}

func (b *Block) ExtendedHeader() *ExtendedHeader            { return b.header.ExtendedHeader }

// SetExtendedHeader converts a byte slice to an ExtendedHeader.
// It panics if b is not of suitable size.
func (h *Header) SetExtendedHeader(sig []byte) {
	if h.ExtendedHeader != nil {
		h.ExtendedHeader.Signature.SetBytes(sig)
	}
}

func HexToSignature(s string) *Signature {
	return BytesToSignature(common.FromHex(s))
}

func BytesToSignature(b []byte) *Signature {
	var s Signature
	s.SetBytes(b)
	return &s
}

func (sig *Signature) SetBytes(b []byte) {
	if len(b) != signatureLength {
		panic(fmt.Sprintf("The signature to be used in the Extended Header is not the correct size: expected %d; got %d", signatureLength, len(b)))
	}

	copy(sig[:], b[:])
}

func (sig *Signature) String() string {
	return common.ToHex(sig[:])
}
