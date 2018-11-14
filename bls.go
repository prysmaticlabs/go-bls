package bls

/*
#cgo CFLAGS:-DMCLBN_FP_UNIT_SIZE=6 -DMCL_DONT_USE_OPENSSL -I ${SRCDIR}/mcl/include/mcl -I ${SRCDIR}/mcl/lib
#cgo LDFLAGS:-lgmp -lstdc++
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

// GetLittleEndian --
func (sec *SecretKey) GetLittleEndian() []byte {
	return sec.v.serialize()
}

// SetLittleEndian --
func (sec *SecretKey) SetLittleEndian(buf []byte) error {
	return sec.v.setLittleEndian(buf)
}

// SerializeToHexStr --
func (sec *SecretKey) SerializeToHexStr() string {
	return sec.v.getString(IoSerializeHexStr)
}

// DeserializeHexStr --
func (sec *SecretKey) DeserializeHexStr(s string) error {
	return sec.v.setString(s, IoSerializeHexStr)
}

// GetHexString --
func (sec *SecretKey) GetHexString() string {
	return sec.v.getString(16)
}

// GetDecString --
func (sec *SecretKey) GetDecString() string {
	return sec.v.getString(10)
}

// SetHexString --
func (sec *SecretKey) SetHexString(s string) error {
	return sec.v.setString(s, 16)
}

// SetDecString --
func (sec *SecretKey) SetDecString(s string) error {
	return sec.v.setString(s, 10)
}

// IsEqual --
func (sec *SecretKey) IsEqual(rhs *SecretKey) bool {
	return sec.v.isEqual(&rhs.v)
}

// SetByCSPRNG --
func (sec *SecretKey) SetByCSPRNG() {
	sec.v.setByCSPRNG()
}

// Add --
func (sec *SecretKey) Add(rhs *SecretKey) {
	frAdd(&sec.v, &sec.v, &rhs.v)
}

// GetMasterSecretKey --
func (sec *SecretKey) GetMasterSecretKey(k int) (msk []SecretKey) {
	msk = make([]SecretKey, k)
	msk[0] = *sec
	for i := 1; i < k; i++ {
		msk[i].SetByCSPRNG()
	}
	return msk
}

// GetMasterPublicKey --
func GetMasterPublicKey(msk []SecretKey) (mpk []PublicKey) {
	n := len(msk)
	mpk = make([]PublicKey, n)
	for i := 0; i < n; i++ {
		mpk[i] = *msk[i].GetPublicKey()
	}
	return mpk
}

// PublicKey --
type PublicKey struct {
	v g2
}

// getPointer --
func (pub *PublicKey) getPointer() (p *C.blsPublicKey) {
	// #nosec
	return (*C.blsPublicKey)(unsafe.Pointer(pub))
}

// Serialize --
func (pub *PublicKey) Serialize() []byte {
	return pub.v.serialize()
}

// Deserialize --
func (pub *PublicKey) Deserialize(buf []byte) error {
	return pub.v.deserialize(buf)
}

// SerializeToHexStr --
func (pub *PublicKey) SerializeToHexStr() string {
	return pub.v.getString(IoSerializeHexStr)
}

// DeserializeHexStr --
func (pub *PublicKey) DeserializeHexStr(s string) error {
	return pub.v.setString(s, IoSerializeHexStr)
}

// GetHexString --
func (pub *PublicKey) GetHexString() string {
	return pub.v.getString(16)
}

// SetHexString --
func (pub *PublicKey) SetHexString(s string) error {
	return pub.v.setString(s, 16)
}

// IsEqual --
func (pub *PublicKey) IsEqual(rhs *PublicKey) bool {
	return pub.v.isEqual(&rhs.v)
}

// Add --
func (pub *PublicKey) Add(rhs *PublicKey) {
	g2Add(&pub.v, &pub.v, &rhs.v)
}

// Sign  --
type Sign struct {
	v g1
}

// getPointer --
func (sign *Sign) getPointer() (p *C.blsSignature) {
	// #nosec
	return (*C.blsSignature)(unsafe.Pointer(sign))
}

// Serialize --
func (sign *Sign) Serialize() []byte {
	return sign.v.serialize()
}

// Deserialize --
func (sign *Sign) Deserialize(buf []byte) error {
	return sign.v.deserialize(buf)
}

// SerializeToHexStr --
func (sign *Sign) SerializeToHexStr() string {
	return sign.v.getString(IoSerializeHexStr)
}

// DeserializeHexStr --
func (sign *Sign) DeserializeHexStr(s string) error {
	return sign.v.setString(s, IoSerializeHexStr)
}

// GetHexString --
func (sign *Sign) GetHexString() string {
	return sign.v.getString(16)
}

// SetHexString --
func (sign *Sign) SetHexString(s string) error {
	return sign.v.setString(s, 16)
}

// IsEqual --
func (sign *Sign) IsEqual(rhs *Sign) bool {
	return sign.v.isEqual(&rhs.v)
}

// GetPublicKey --
func (sec *SecretKey) GetPublicKey() (pub *PublicKey) {
	pub = new(PublicKey)
	C.blsGetPublicKey(pub.getPointer(), sec.getPointer())
	return pub
}

// Sign -- Constant Time version
func (sec *SecretKey) Sign(m string) (sign *Sign) {
	sign = new(Sign)
	buf := []byte(m)
	// #nosec
	C.blsSign(sign.getPointer(), sec.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	return sign
}

// Add --
func (sign *Sign) Add(rhs *Sign) {
	C.blsSignatureAdd(sign.getPointer(), rhs.getPointer())
}

// Verify --
func (sign *Sign) Verify(pub *PublicKey, m string) bool {
	buf := []byte(m)
	// #nosec
	return C.blsVerify(sign.getPointer(), pub.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf))) == 1
}
