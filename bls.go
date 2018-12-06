package bls

/*
#cgo CFLAGS:-DMCLBN_FP_UNIT_SIZE=6 -DMCL_DONT_USE_OPENSSL -I ${SRCDIR}/mcl/include/mcl -I ${SRCDIR}/mcl/lib
#cgo LDFLAGS:-L${SRCDIR}/mcl/lib -lbls384_dy -lstdc++
#include "mcl/include/mcl/bls.h"
*/
import "C"
import "fmt"
import "unsafe"
import "log"

func init() {
	if err := initializeBLS(BLS12_381); err != nil {
		log.Fatalf("Could not initialize BLS12-381 curve: %v", err)
	}
}

// initializeBLS --
// call this function before calling all the other operations
// this function is not thread safe
func initializeBLS(curve int) error {
	err := C.blsInit(C.int(curve), C.MCLBN_COMPILED_TIME_VAR)
	if err != 0 {
		return fmt.Errorf("ERR Init curve=%d", curve)
	}
	return nil
}

// SecretKey --
type SecretKey struct {
	v fr
}

// getPointer --
func (sec *SecretKey) getPointer() (p *C.blsSecretKey) {
	// #nosec
	return (*C.blsSecretKey)(unsafe.Pointer(sec))
}

// LittleEndian returns the serialized, little-endian formatted byte slice
// of the secret key.
func (sec *SecretKey) LittleEndian() []byte {
	return sec.v.serialize()
}

// SetLittleEndian sets a secret key based on a little-endian formatted
// byte slice.
func (sec *SecretKey) SetLittleEndian(buf []byte) error {
	return sec.v.setLittleEndian(buf)
}

// SerializeToHexStr returns a hex string representation of a private key.
func (sec *SecretKey) SerializeToHexStr() string {
	return sec.v.getString(IoSerializeHexStr)
}

// DeserializeHexStr deserializes a hex string into a private key.
func (sec *SecretKey) DeserializeHexStr(s string) error {
	return sec.v.setString(s, IoSerializeHexStr)
}

// HexString returns the hex string of the private key.
func (sec *SecretKey) HexString() string {
	return sec.v.getString(16)
}

// DecString returns a decimal string representation of the private key.
func (sec *SecretKey) DecString() string {
	return sec.v.getString(10)
}

// SetHexString sets a private key based on a hex string.
func (sec *SecretKey) SetHexString(s string) error {
	return sec.v.setString(s, 16)
}

// SetDecString sets a private key based on a decimal string.
func (sec *SecretKey) SetDecString(s string) error {
	return sec.v.setString(s, 10)
}

// IsEqual compares two private keys.
func (sec *SecretKey) IsEqual(rhs *SecretKey) bool {
	return sec.v.isEqual(&rhs.v)
}

// SetByCSPRNG sets a private key's internal representation using a
// cryptographically-secure, pseudorandom number generator.
func (sec *SecretKey) SetByCSPRNG() {
	sec.v.setByCSPRNG()
}

// Add two private keys together.
func (sec *SecretKey) Add(rhs *SecretKey) {
	frAdd(&sec.v, &sec.v, &rhs.v)
}

// GetMasterSecretKey creates a series of k secret keys and using a
// pseudorandom number generator and returns them as a slice.
func (sec *SecretKey) GetMasterSecretKey(k int) (msk []SecretKey) {
	msk = make([]SecretKey, k)
	msk[0] = *sec
	for i := 1; i < k; i++ {
		msk[i].SetByCSPRNG()
	}
	return msk
}

// GetMasterPublicKey returns a list of public keys for a slice of
// private keys.
func GetMasterPublicKey(msk []SecretKey) (mpk []PublicKey) {
	n := len(msk)
	mpk = make([]PublicKey, n)
	for i := 0; i < n; i++ {
		mpk[i] = *msk[i].GetPublicKey()
	}
	return mpk
}

// PublicKey definition in the BLS scheme.
type PublicKey struct {
	v g2
}

// getPointer --
func (pub *PublicKey) getPointer() (p *C.blsPublicKey) {
	// #nosec
	return (*C.blsPublicKey)(unsafe.Pointer(pub))
}

// Serialize a public key into a byte slice.
func (pub *PublicKey) Serialize() []byte {
	return pub.v.serialize()
}

// Deserialize converts a byte slice into a public key.
func (pub *PublicKey) Deserialize(buf []byte) error {
	return pub.v.deserialize(buf)
}

// SerializeToHexStr returns a hex string serialization of a public key.
func (pub *PublicKey) SerializeToHexStr() string {
	return pub.v.getString(IoSerializeHexStr)
}

// DeserializeHexStr sets a public key from a hex string.
func (pub *PublicKey) DeserializeHexStr(s string) error {
	return pub.v.setString(s, IoSerializeHexStr)
}

// HexString returns the hex string representation of a public key.
func (pub *PublicKey) HexString() string {
	return pub.v.getString(16)
}

// SetHexString sets a public key from a hex string.
func (pub *PublicKey) SetHexString(s string) error {
	return pub.v.setString(s, 16)
}

// IsEqual compares two BLS public keys.
func (pub *PublicKey) IsEqual(rhs *PublicKey) bool {
	return pub.v.isEqual(&rhs.v)
}

// Add two BLS public keys together
func (pub *PublicKey) Add(rhs *PublicKey) {
	g2Add(&pub.v, &pub.v, &rhs.v)
}

// Sign represents a signature in the BLS signature aggregation scheme.
type Sign struct {
	v g1
}

// getPointer --
func (sign *Sign) getPointer() (p *C.blsSignature) {
	// #nosec
	return (*C.blsSignature)(unsafe.Pointer(sign))
}

// Serialize a signature into a byte array.
func (sign *Sign) Serialize() []byte {
	return sign.v.serialize()
}

// Deserialize a signature from a byte array.
func (sign *Sign) Deserialize(buf []byte) error {
	return sign.v.deserialize(buf)
}

// SerializeToHexStr serializes a signature into a hex string.
func (sign *Sign) SerializeToHexStr() string {
	return sign.v.getString(IoSerializeHexStr)
}

// DeserializeHexStr creates a signature from a hex string serialization.
func (sign *Sign) DeserializeHexStr(s string) error {
	return sign.v.setString(s, IoSerializeHexStr)
}

// HexString representation of a signature.
func (sign *Sign) HexString() string {
	return sign.v.getString(16)
}

// SetHexString sets a signature from a hex string.
func (sign *Sign) SetHexString(s string) error {
	return sign.v.setString(s, 16)
}

// IsEqual compares two signatures.
func (sign *Sign) IsEqual(rhs *Sign) bool {
	return sign.v.isEqual(&rhs.v)
}

// GetPublicKey returns the public key corresponding to a BLS private key.
func (sec *SecretKey) GetPublicKey() (pub *PublicKey) {
	pub = new(PublicKey)
	C.blsGetPublicKey(pub.getPointer(), sec.getPointer())
	return pub
}

// Sign a string message using a BLS private key in constant time.
func (sec *SecretKey) Sign(m string) (sign *Sign) {
	sign = new(Sign)
	buf := []byte(m)
	// #nosec
	C.blsSign(sign.getPointer(), sec.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	return sign
}

// Add two BLS signatures together.
func (sign *Sign) Add(rhs *Sign) {
	C.blsSignatureAdd(sign.getPointer(), rhs.getPointer())
}

// Verify a signature using a BLS public key and a message string.
func (sign *Sign) Verify(pub *PublicKey, m string) bool {
	buf := []byte(m)
	// #nosec
	return C.blsVerify(sign.getPointer(), pub.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf))) == 1
}
