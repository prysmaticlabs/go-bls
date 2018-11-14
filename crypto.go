package bls

/*
#cgo CFLAGS:-DMCLBN_FP_UNIT_SIZE=6
#include "mcl/include/mcl/bn.h"
*/
import "C"
import "fmt"
import "unsafe"

// CurveFp254BNb -- 254 bit curve.
const CurveFp254BNb = C.mclBn_CurveFp254BNb

// CurveFp382_1 -- 382 bit curve 1.
const CurveFp382_1 = C.mclBn_CurveFp382_1

// CurveFp382_2 -- 382 bit curve 2.
const CurveFp382_2 = C.mclBn_CurveFp382_2

// BLS12_381 Curve.
const BLS12_381 = C.MCL_BLS12_381

// IoSerializeHexStr -- serialization parameter for the underlying cryptographic
// library created by @herumi.
const IoSerializeHexStr = C.MCLBN_IO_SERIALIZE_HEX_STR

// GetMaxOpUnitSize --
func GetMaxOpUnitSize() int {
	return int(C.MCLBN_FP_UNIT_SIZE)
}

// GetOpUnitSize --
// the length of Fr is GetOpUnitSize() * 8 bytes
func GetOpUnitSize() int {
	return int(C.mclBn_getOpUnitSize())
}

// curveOrder returns the order of the group G1.
func curveOrder() string {
	buf := make([]byte, 1024)
	// #nosec
	n := C.mclBn_getCurveOrder((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return string(buf[:n])
}

// fieldOrder returns the characteristic of the field defining
// a particular BLS curve.
func fieldOrder() string {
	buf := make([]byte, 1024)
	// #nosec
	n := C.mclBn_getFieldOrder((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)))
	if n == 0 {
		panic("implementation err. size of buf is small")
	}
	return string(buf[:n])
}

// fr represents a cryptographic group defined by the underlying
// herumi MCL library.
type fr struct {
	v C.mclBnFr
}

// getPointer() returns the underlying cast getPointer() to a C struct.
func (x *fr) getPointer() (p *C.mclBnFr) {
	// #nosec
	return (*C.mclBnFr)(unsafe.Pointer(x))
}

// clear the underlying C struct.
func (x *fr) clear() {
	// #nosec
	C.mclBnFr_clear(x.getPointer())
}

// setInt64 for the value of the underlying C struct's data.
func (x *fr) setInt64(v int64) {
	// #nosec
	C.mclBnFr_setInt(x.getPointer(), C.int64_t(v))
}

// setString of the underlying value in the C struct.
func (x *fr) setString(s string, base int) error {
	buf := []byte(s)
	// #nosec
	err := C.mclBnFr_setStr(x.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(base))
	if err != 0 {
		return fmt.Errorf("err mclBnFr_setStr %x", err)
	}
	return nil
}

// deserialize the underlying C struct.
func (x *fr) deserialize(buf []byte) error {
	// #nosec
	err := C.mclBnFr_deserialize(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err == 0 {
		return fmt.Errorf("err mclBnFr_deserialize %x", buf)
	}
	return nil
}

// setLittleEndian sets the value of the C struct's data from a little-endian
// format bytes array.
func (x *fr) setLittleEndian(buf []byte) error {
	// #nosec
	err := C.mclBnFr_setLittleEndian(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err != 0 {
		return fmt.Errorf("err mclBnFr_setLittleEndian %x", err)
	}
	return nil
}

// isEqual compares two fr values.
func (x *fr) isEqual(rhs *fr) bool {
	return C.mclBnFr_isEqual(x.getPointer(), rhs.getPointer()) == 1
}

// isZero checks if the underlying C data is empty.
func (x *fr) isZero() bool {
	return C.mclBnFr_isZero(x.getPointer()) == 1
}

// IsOne checks if the value of the underlying C data is non-empty.
func (x *fr) isOne() bool {
	return C.mclBnFr_isOne(x.getPointer()) == 1
}

// setByCSPRNG sets the underlying value using a cryptographically-secure
// pseudorandom number generator.
func (x *fr) setByCSPRNG() {
	err := C.mclBnFr_setByCSPRNG(x.getPointer())
	if err != 0 {
		panic("err mclBnFr_setByCSPRNG")
	}
}

// setHashOf using a byte slice.
func (x *fr) setHashOf(buf []byte) bool {
	// #nosec
	return C.mclBnFr_setHashOf(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf))) == 0
}

// getString given an integer base.
func (x *fr) getString(base int) string {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnFr_getStr((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), x.getPointer(), C.int(base))
	if n == 0 {
		panic("err mclBnFr_getStr")
	}
	return string(buf[:n])
}

// serialize into a byte-slice.
func (x *fr) serialize() []byte {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnFr_serialize(unsafe.Pointer(&buf[0]), C.size_t(len(buf)), x.getPointer())
	if n == 0 {
		panic("err mclBnFr_serialize")
	}
	return buf[:n]
}

// frNeg negates the current value.
func frNeg(out *fr, x *fr) {
	C.mclBnFr_neg(out.getPointer(), x.getPointer())
}

// frInv calculates the inverse.
func frInv(out *fr, x *fr) {
	C.mclBnFr_inv(out.getPointer(), x.getPointer())
}

// frAdd two fr values.
func frAdd(out *fr, x *fr, y *fr) {
	C.mclBnFr_add(out.getPointer(), x.getPointer(), y.getPointer())
}

// frSub two fr values.
func frSub(out *fr, x *fr, y *fr) {
	C.mclBnFr_sub(out.getPointer(), x.getPointer(), y.getPointer())
}

// frMul two fr values.
func frMul(out *fr, x *fr, y *fr) {
	C.mclBnFr_mul(out.getPointer(), x.getPointer(), y.getPointer())
}

// frDiv two fr values.
func frDiv(out *fr, x *fr, y *fr) {
	C.mclBnFr_div(out.getPointer(), x.getPointer(), y.getPointer())
}

// g1 represents a group in the BLS signature scheme.
type g1 struct {
	v C.mclBnG1
}

// getPointer() returns an underlying cast C-struct for the g1 group.
func (x *g1) getPointer() (p *C.mclBnG1) {
	// #nosec
	return (*C.mclBnG1)(unsafe.Pointer(x))
}

// clear the g1 values.
func (x *g1) clear() {
	// #nosec
	C.mclBnG1_clear(x.getPointer())
}

// setString --
func (x *g1) setString(s string, base int) error {
	buf := []byte(s)
	// #nosec
	err := C.mclBnG1_setStr(x.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(base))
	if err != 0 {
		return fmt.Errorf("err mclBnG1_setStr %x", err)
	}
	return nil
}

// deserialize --
func (x *g1) deserialize(buf []byte) error {
	// #nosec
	err := C.mclBnG1_deserialize(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err == 0 {
		return fmt.Errorf("err mclBnG1_deserialize %x", buf)
	}
	return nil
}

// isEqual --
func (x *g1) isEqual(rhs *g1) bool {
	return C.mclBnG1_isEqual(x.getPointer(), rhs.getPointer()) == 1
}

// isZero --
func (x *g1) isZero() bool {
	return C.mclBnG1_isZero(x.getPointer()) == 1
}

// hashAndMapTo --
func (x *g1) hashAndMapTo(buf []byte) error {
	// #nosec
	err := C.mclBnG1_hashAndMapTo(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err != 0 {
		return fmt.Errorf("err mclBnG1_hashAndMapTo %x", err)
	}
	return nil
}

// getString --
func (x *g1) getString(base int) string {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnG1_getStr((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), x.getPointer(), C.int(base))
	if n == 0 {
		panic("err mclBnG1_getStr")
	}
	return string(buf[:n])
}

// serialize --
func (x *g1) serialize() []byte {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnG1_serialize(unsafe.Pointer(&buf[0]), C.size_t(len(buf)), x.getPointer())
	if n == 0 {
		panic("err mclBnG1_serialize")
	}
	return buf[:n]
}

// g1Neg --
func g1Neg(out *g1, x *g1) {
	C.mclBnG1_neg(out.getPointer(), x.getPointer())
}

// g1Dbl --
func g1Dbl(out *g1, x *g1) {
	C.mclBnG1_dbl(out.getPointer(), x.getPointer())
}

// g1Add --
func g1Add(out *g1, x *g1, y *g1) {
	C.mclBnG1_add(out.getPointer(), x.getPointer(), y.getPointer())
}

// g1Sub --
func g1Sub(out *g1, x *g1, y *g1) {
	C.mclBnG1_sub(out.getPointer(), x.getPointer(), y.getPointer())
}

// g1Mul --
func g1Mul(out *g1, x *g1, y *fr) {
	C.mclBnG1_mul(out.getPointer(), x.getPointer(), y.getPointer())
}

// g1MulCT -- constant time (depending on bit lengh of y)
func g1MulCT(out *g1, x *g1, y *fr) {
	C.mclBnG1_mulCT(out.getPointer(), x.getPointer(), y.getPointer())
}

// g2 --
type g2 struct {
	v C.mclBnG2
}

// getPointer() --
func (x *g2) getPointer() (p *C.mclBnG2) {
	// #nosec
	return (*C.mclBnG2)(unsafe.Pointer(x))
}

// clear --
func (x *g2) Clear() {
	// #nosec
	C.mclBnG2_clear(x.getPointer())
}

// setString --
func (x *g2) setString(s string, base int) error {
	buf := []byte(s)
	// #nosec
	err := C.mclBnG2_setStr(x.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(base))
	if err != 0 {
		return fmt.Errorf("err mclBnG2_setStr %x", err)
	}
	return nil
}

// deserialize --
func (x *g2) deserialize(buf []byte) error {
	// #nosec
	err := C.mclBnG2_deserialize(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err == 0 {
		return fmt.Errorf("err mclBnG2_deserialize %x", buf)
	}
	return nil
}

// isEqual --
func (x *g2) isEqual(rhs *g2) bool {
	return C.mclBnG2_isEqual(x.getPointer(), rhs.getPointer()) == 1
}

// isZero --
func (x *g2) isZero() bool {
	return C.mclBnG2_isZero(x.getPointer()) == 1
}

// hashAndMapTo --
func (x *g2) hashAndMapTo(buf []byte) error {
	// #nosec
	err := C.mclBnG2_hashAndMapTo(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err != 0 {
		return fmt.Errorf("err mclBnG2_hashAndMapTo %x", err)
	}
	return nil
}

// getString --
func (x *g2) getString(base int) string {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnG2_getStr((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), x.getPointer(), C.int(base))
	if n == 0 {
		panic("err mclBnG2_getStr")
	}
	return string(buf[:n])
}

// serialize --
func (x *g2) serialize() []byte {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnG2_serialize(unsafe.Pointer(&buf[0]), C.size_t(len(buf)), x.getPointer())
	if n == 0 {
		panic("err mclBnG2_serialize")
	}
	return buf[:n]
}

// g2Neg --
func g2Neg(out *g2, x *g2) {
	C.mclBnG2_neg(out.getPointer(), x.getPointer())
}

// g2Dbl --
func g2Dbl(out *g2, x *g2) {
	C.mclBnG2_dbl(out.getPointer(), x.getPointer())
}

// g2Add --
func g2Add(out *g2, x *g2, y *g2) {
	C.mclBnG2_add(out.getPointer(), x.getPointer(), y.getPointer())
}

// g2Sub --
func g2Sub(out *g2, x *g2, y *g2) {
	C.mclBnG2_sub(out.getPointer(), x.getPointer(), y.getPointer())
}

// g2Mul --
func g2Mul(out *g2, x *g2, y *fr) {
	C.mclBnG2_mul(out.getPointer(), x.getPointer(), y.getPointer())
}

// gT --
type gT struct {
	v C.mclBnGT
}

// getPointer() --
func (x *gT) getPointer() (p *C.mclBnGT) {
	// #nosec
	return (*C.mclBnGT)(unsafe.Pointer(x))
}

// clear --
func (x *gT) clear() {
	// #nosec
	C.mclBnGT_clear(x.getPointer())
}

// setInt64 --
func (x *gT) setInt64(v int64) {
	// #nosec
	C.mclBnGT_setInt(x.getPointer(), C.int64_t(v))
}

// setString --
func (x *gT) setString(s string, base int) error {
	buf := []byte(s)
	// #nosec
	err := C.mclBnGT_setStr(x.getPointer(), (*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), C.int(base))
	if err != 0 {
		return fmt.Errorf("err mclBnGT_setStr %x", err)
	}
	return nil
}

// deserialize --
func (x *gT) deserialize(buf []byte) error {
	// #nosec
	err := C.mclBnGT_deserialize(x.getPointer(), unsafe.Pointer(&buf[0]), C.size_t(len(buf)))
	if err == 0 {
		return fmt.Errorf("err mclBnGT_deserialize %x", buf)
	}
	return nil
}

// isEqual --
func (x *gT) isEqual(rhs *gT) bool {
	return C.mclBnGT_isEqual(x.getPointer(), rhs.getPointer()) == 1
}

// isZero --
func (x *gT) isZero() bool {
	return C.mclBnGT_isZero(x.getPointer()) == 1
}

// isOne --
func (x *gT) isOne() bool {
	return C.mclBnGT_isOne(x.getPointer()) == 1
}

// getString --
func (x *gT) getString(base int) string {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnGT_getStr((*C.char)(unsafe.Pointer(&buf[0])), C.size_t(len(buf)), x.getPointer(), C.int(base))
	if n == 0 {
		panic("err mclBnGT_getStr")
	}
	return string(buf[:n])
}

// serialize --
func (x *gT) serialize() []byte {
	buf := make([]byte, 2048)
	// #nosec
	n := C.mclBnGT_serialize(unsafe.Pointer(&buf[0]), C.size_t(len(buf)), x.getPointer())
	if n == 0 {
		panic("err mclBnGT_serialize")
	}
	return buf[:n]
}

// gTNeg --
func gTNeg(out *gT, x *gT) {
	C.mclBnGT_neg(out.getPointer(), x.getPointer())
}

// gTInv --
func gTInv(out *gT, x *gT) {
	C.mclBnGT_inv(out.getPointer(), x.getPointer())
}

// gTAdd --
func gTAdd(out *gT, x *gT, y *gT) {
	C.mclBnGT_add(out.getPointer(), x.getPointer(), y.getPointer())
}

// gTSub --
func gTSub(out *gT, x *gT, y *gT) {
	C.mclBnGT_sub(out.getPointer(), x.getPointer(), y.getPointer())
}

// gTMul --
func gTMul(out *gT, x *gT, y *gT) {
	C.mclBnGT_mul(out.getPointer(), x.getPointer(), y.getPointer())
}

// gTDiv --
func gTDiv(out *gT, x *gT, y *gT) {
	C.mclBnGT_div(out.getPointer(), x.getPointer(), y.getPointer())
}

// gTPow --
func gTPow(out *gT, x *gT, y *fr) {
	C.mclBnGT_pow(out.getPointer(), x.getPointer(), y.getPointer())
}

// pairing --
func pairing(out *gT, x *g1, y *g2) {
	C.mclBn_pairing(out.getPointer(), x.getPointer(), y.getPointer())
}

// finalExp --
func finalExp(out *gT, x *gT) {
	C.mclBn_finalExp(out.getPointer(), x.getPointer())
}

// millerLoop --
func millerLoop(out *gT, x *g1, y *g2) {
	C.mclBn_millerLoop(out.getPointer(), x.getPointer(), y.getPointer())
}

// getUint64NumToPrecompute --
func getUint64NumToPrecompute() int {
	return int(C.mclBn_getUint64NumToPrecompute())
}

// precomputeG2 --
func precomputeG2(Qbuf []uint64, Q *g2) {
	// #nosec
	C.mclBn_precomputeG2((*C.uint64_t)(unsafe.Pointer(&Qbuf[0])), Q.getPointer())
}

// precomputedMillerLoop --
func precomputedMillerLoop(out *gT, P *g1, Qbuf []uint64) {
	// #nosec
	C.mclBn_precomputedMillerLoop(out.getPointer(), P.getPointer(), (*C.uint64_t)(unsafe.Pointer(&Qbuf[0])))
}

// precomputedMillerLoop2 --
func precomputedMillerLoop2(out *gT, P1 *g1, Q1buf []uint64, P2 *g1, Q2buf []uint64) {
	// #nosec
	C.mclBn_precomputedMillerLoop2(out.getPointer(), P1.getPointer(), (*C.uint64_t)(unsafe.Pointer(&Q1buf[0])), P1.getPointer(), (*C.uint64_t)(unsafe.Pointer(&Q1buf[0])))
}

// frEvaluatePolynomial -- y = c[0] + c[1] * x + c[2] * x^2 + ...
func frEvaluatePolynomial(y *fr, c []fr, x *fr) error {
	// #nosec
	err := C.mclBn_FrEvaluatePolynomial(y.getPointer(), (*C.mclBnFr)(unsafe.Pointer(&c[0])), (C.size_t)(len(c)), x.getPointer())
	if err != 0 {
		return fmt.Errorf("err mclBn_FrEvaluatePolynomial")
	}
	return nil
}

// g1EvaluatePolynomial -- y = c[0] + c[1] * x + c[2] * x^2 + ...
func g1EvaluatePolynomial(y *g1, c []g1, x *fr) error {
	// #nosec
	err := C.mclBn_G1EvaluatePolynomial(y.getPointer(), (*C.mclBnG1)(unsafe.Pointer(&c[0])), (C.size_t)(len(c)), x.getPointer())
	if err != 0 {
		return fmt.Errorf("err mclBn_G1EvaluatePolynomial")
	}
	return nil
}

// g2EvaluatePolynomial -- y = c[0] + c[1] * x + c[2] * x^2 + ...
func g2EvaluatePolynomial(y *g2, c []g2, x *fr) error {
	// #nosec
	err := C.mclBn_G2EvaluatePolynomial(y.getPointer(), (*C.mclBnG2)(unsafe.Pointer(&c[0])), (C.size_t)(len(c)), x.getPointer())
	if err != 0 {
		return fmt.Errorf("err mclBn_G2EvaluatePolynomial")
	}
	return nil
}

// frLagrangeInterpolation --
func frLagrangeInterpolation(out *fr, xVec []fr, yVec []fr) error {
	if len(xVec) != len(yVec) {
		return fmt.Errorf("err FrLagrangeInterpolation:bad size")
	}
	// #nosec
	err := C.mclBn_FrLagrangeInterpolation(out.getPointer(), (*C.mclBnFr)(unsafe.Pointer(&xVec[0])), (*C.mclBnFr)(unsafe.Pointer(&yVec[0])), (C.size_t)(len(xVec)))
	if err != 0 {
		return fmt.Errorf("err FrLagrangeInterpolation")
	}
	return nil
}

// g1LagrangeInterpolation --
func g1LagrangeInterpolation(out *g1, xVec []fr, yVec []g1) error {
	if len(xVec) != len(yVec) {
		return fmt.Errorf("err G1LagrangeInterpolation:bad size")
	}
	// #nosec
	err := C.mclBn_G1LagrangeInterpolation(out.getPointer(), (*C.mclBnFr)(unsafe.Pointer(&xVec[0])), (*C.mclBnG1)(unsafe.Pointer(&yVec[0])), (C.size_t)(len(xVec)))
	if err != 0 {
		return fmt.Errorf("err G1LagrangeInterpolation")
	}
	return nil
}

// g2LagrangeInterpolation --
func g2LagrangeInterpolation(out *g2, xVec []fr, yVec []g2) error {
	if len(xVec) != len(yVec) {
		return fmt.Errorf("err G2LagrangeInterpolation:bad size")
	}
	// #nosec
	err := C.mclBn_G2LagrangeInterpolation(out.getPointer(), (*C.mclBnFr)(unsafe.Pointer(&xVec[0])), (*C.mclBnG2)(unsafe.Pointer(&yVec[0])), (C.size_t)(len(xVec)))
	if err != 0 {
		return fmt.Errorf("err G2LagrangeInterpolation")
	}
	return nil
}
