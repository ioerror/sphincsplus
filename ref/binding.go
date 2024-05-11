//go:build ((linux && amd64) || (darwin && amd64)) && (sphincs_haraka_128f || sphincs_haraka_128s || sphincs_haraka_192f || sphincs_haraka_192s || sphincs_haraka_256f || sphincs_haraka_256s || sphincs_sha2_128f || sphincs_sha2_128s || sphincs_sha2_192f || sphincs_sha2_192s || sphincs_sha2_256f || sphincs_sha2_256s || sphincs_shake_128f || sphincs_shake_128s || sphincs_shake_192f || sphincs_shake_192s || sphincs_shake_256f || sphincs_shake_256s)

package sphincsplus

//#cgo amd64 LDFLAGS: "-L./"
//#cgo amd64 CFLAGS: -O3 -std=c99 -Wconversion -Wpedantic -D CGO=1
//#cgo linux/amd64 LDFLAGS: "-L./ -L/usr/lib/x86_64-linux-gnu/"
//#cgo linux/amd64 CFLAGS: -O3 -std=c99 -Wconversion -Wpedantic -D PARAMS=sphincs_shake_256f -D CGO=1
//#include "api.h"
import "C"
import (
	"fmt"
	"unsafe"

	"github.com/katzenpost/hpqc/util"

	"github.com/katzenpost/sphincsplus/ref/params"
)

var (
	_ = sphincsplus.A

	// PublicKeySize is the size in bytes of the public key.
	PublicKeySize int = C.CRYPTO_PUBLICKEYBYTES

	// PrivateKeySize is the size in bytes of the private key.
	PrivateKeySize int = C.CRYPTO_SECRETKEYBYTES

	// SignatureSize is the size in bytes of the signature.
	SignatureSize int = C.CRYPTO_BYTES

	// PARAMS are the signature parameters
	SignatureName int = C.PARAMS

	// ErrPublicKeySize indicates the raw data is not the correct size for a public key.
	ErrPublicKeySize error = fmt.Errorf("%s: raw public key data size is wrong", Name())

	// ErrPrivateKeySize indicates the raw data is not the correct size for a private key.
	ErrPrivateKeySize error = fmt.Errorf("%s: raw private key data size is wrong", Name())
)

// Name returns the string naming of the current
// Sphincs+ that this binding is being used with.
func Name() string {
	return fmt.Sprint(SignatureName)
}

// NewKeypair generates a new Sphincs+ keypair.
func NewKeypair() (*PrivateKey, *PublicKey) {
	privKey := &PrivateKey{
		privateKey: make([]byte, C.CRYPTO_SECRETKEYBYTES),
	}
	pubKey := &PublicKey{
		publicKey: make([]byte, C.CRYPTO_PUBLICKEYBYTES),
	}
	C.crypto_sign_keypair((*C.uchar)(unsafe.Pointer(&pubKey.publicKey[0])),
		(*C.uchar)(unsafe.Pointer(&privKey.privateKey[0])))
	return privKey, pubKey
}

// PublicKey is a public Sphincs+ key.
type PublicKey struct {
	publicKey []byte
}

// Reset overwrites the key with zeros.
func (p *PublicKey) Reset() {
	util.ExplicitBzero(p.publicKey)
}

// Verify checks whether the given signature is valid.
func (p *PublicKey) Verify(signature, message []byte) bool {
	ret := C.crypto_sign_verify((*C.uchar)(unsafe.Pointer(&signature[0])),
		C.size_t(len(signature)),
		(*C.uchar)(unsafe.Pointer(&message[0])),
		C.size_t(len(message)),
		(*C.uchar)(unsafe.Pointer(&p.publicKey[0])))
	if ret == 0 {
		return true
	}
	return false
}

// Bytes returns the PublicKey as a byte slice.
func (p *PublicKey) Bytes() []byte {
	return p.publicKey
}

// FromBytes loads a PublicKey from the given byte slice.
func (p *PublicKey) FromBytes(data []byte) error {
	if len(data) != PublicKeySize {
		return ErrPublicKeySize
	}

	p.publicKey = data
	return nil
}

// Verify checks whether the given signature is valid.
/*
func (p *PublicKey) Verify(signature, message []byte) bool {
	ret := C.crypto_sign_verify((*C.uchar)(unsafe.Pointer(&signature[0])),
		C.ulong(len(signature)),
		(*C.uchar)(unsafe.Pointer(&message[0])),
		C.ulong(len(message)),
		(*C.uchar)(unsafe.Pointer(&p.publicKey[0])))
	if ret == 0 {
		return true
	}
	return false
}
*/

// PrivateKey is a private Sphincs+ key.
type PrivateKey struct {
	privateKey []byte
}

// Reset overwrites the key with zeros.
func (p *PrivateKey) Reset() {
	util.ExplicitBzero(p.privateKey)
}

// Sign signs the given message and returns the signature.
func (p *PrivateKey) Sign(message []byte) []byte {
	signature := make([]byte, C.CRYPTO_BYTES)
	sigLen := C.size_t(C.CRYPTO_BYTES)
	C.crypto_sign_signature((*C.uchar)(unsafe.Pointer(&signature[0])),
		&sigLen,
		(*C.uchar)(unsafe.Pointer(&message[0])),
		(C.size_t)(len(message)),
		(*C.uchar)(unsafe.Pointer(&p.privateKey[0])))
	return signature
}

// Bytes returns the PrivateKey as a byte slice.
func (p *PrivateKey) Bytes() []byte {
	return p.privateKey
}

// FromBytes loads a PrivateKey from the given byte slice.
func (p *PrivateKey) FromBytes(data []byte) error {
	if len(data) != PrivateKeySize {
		return ErrPrivateKeySize
	}

	p.privateKey = data
	return nil
}
