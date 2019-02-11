
#include "include/mcl/bls.h"

void blsGetPublicKey(blsPublicKey *pub, const blsSecretKey *sec){}
int blsInit(int curve, int compiledTimeVar) { return 0; }
void blsSign(blsSignature *sig, const blsSecretKey *sec, const void *m, mclSize size){}
void blsSignatureAdd(blsSignature *sig, const blsSignature *rhs){}
int blsVerify(const blsSignature *sig, const blsPublicKey *pub, const void *m, mclSize size){ return 0; }
void mclBnFr_add(mclBnFr *z, const mclBnFr *x, const mclBnFr *y){}
void mclBnFr_clear(mclBnFr *x){}
mclSize mclBnFr_deserialize(mclBnFr *x, const void *buf, mclSize bufSize){ return 0; }
void mclBnFr_div(mclBnFr *z, const mclBnFr *x, const mclBnFr *y){}
mclSize mclBnFr_getStr(char *buf, mclSize maxBufSize, const mclBnFr *x, int ioMode){ return 0; }
int mclBnFr_setStr(mclBnFr *x, const char *buf, mclSize bufSize, int ioMode){ return 0; }
void mclBnFr_inv(mclBnFr *y, const mclBnFr *x){}
int mclBnFr_isEqual(const mclBnFr *x, const mclBnFr *y){ return 0; }
int mclBnFr_setLittleEndian(mclBnFr *x, const void *buf, mclSize bufSize){ return 0; }
int mclBnFr_isOne(const mclBnFr *x){ return 0; }
int mclBnFr_isZero(const mclBnFr *x){ return 0; }
void mclBnFr_mul(mclBnFr *z, const mclBnFr *x, const mclBnFr *y){}
void mclBnFr_neg(mclBnFr *y, const mclBnFr *x){}
mclSize mclBnFr_serialize(void *buf, mclSize maxBufSize, const mclBnFr *x){ return 0; }
int mclBnFr_setByCSPRNG(mclBnFr *x){ return 0; }
int mclBnFr_setHashOf(mclBnFr *x, const void *buf, mclSize bufSize){ return 0; }
void mclBnFr_setInt(mclBnFr *y, mclInt x){}
void mclBnFr_setInt32(mclBnFr *y, int x){}
void mclBnFr_sub(mclBnFr *z, const mclBnFr *x, const mclBnFr *y){}
void mclBnG1_add(mclBnG1 *z, const mclBnG1 *x, const mclBnG1 *y){}
void mclBnG1_clear(mclBnG1 *x){}
void mclBnG1_dbl(mclBnG1 *y, const mclBnG1 *x){}
mclSize mclBnG1_deserialize(mclBnG1 *x, const void *buf, mclSize bufSize){ return 0; }
mclSize mclBnG1_getStr(char *buf, mclSize maxBufSize, const mclBnG1 *x, int ioMode){ return 0; }
int mclBnG1_hashAndMapTo(mclBnG1 *x, const void *buf, mclSize bufSize){ return 0; }
int mclBnG1_isEqual(const mclBnG1 *x, const mclBnG1 *y){ return 0; }
int mclBnG1_isZero(const mclBnG1 *x){ return 0; }
void mclBnG1_mul(mclBnG1 *z, const mclBnG1 *x, const mclBnFr *y){}
void mclBnG1_mulCT(mclBnG1 *z, const mclBnG1 *x, const mclBnFr *y){}
void mclBnG1_neg(mclBnG1 *y, const mclBnG1 *x){}
mclSize mclBnG1_serialize(void *buf, mclSize maxBufSize, const mclBnG1 *x){ return 0; }
int mclBnG1_setStr(mclBnG1 *x, const char *buf, mclSize bufSize, int ioMode){ return 0; }
void mclBnG1_sub(mclBnG1 *z, const mclBnG1 *x, const mclBnG1 *y){}
void mclBnG2_add(mclBnG2 *z, const mclBnG2 *x, const mclBnG2 *y){}
void mclBnG2_clear(mclBnG2 *x){}
void mclBnG2_dbl(mclBnG2 *y, const mclBnG2 *x){}
mclSize mclBnG2_deserialize(mclBnG2 *x, const void *buf, mclSize bufSize){ return 0; } 
mclSize mclBnG2_getStr(char *buf, mclSize maxBufSize, const mclBnG2 *x, int ioMode){ return 0; }
int mclBnG2_hashAndMapTo(mclBnG2 *x, const void *buf, mclSize bufSize) { return 0; }
int mclBnG2_isEqual(const mclBnG2 *x, const mclBnG2 *y){ return 0; }
int mclBnG2_isZero(const mclBnG2 *x){ return 0; }
void mclBnG2_mul(mclBnG2 *z, const mclBnG2 *x, const mclBnFr *y){}
void mclBnG2_neg(mclBnG2 *y, const mclBnG2 *x){}
mclSize mclBnG2_serialize(void *buf, mclSize maxBufSize, const mclBnG2 *x){ return 0; }
int mclBnG2_setStr(mclBnG2 *x, const char *buf, mclSize bufSize, int ioMode){ return 0; }
void mclBnG2_sub(mclBnG2 *z, const mclBnG2 *x, const mclBnG2 *y){}
void mclBnGT_add(mclBnGT *z, const mclBnGT *x, const mclBnGT *y){}
void mclBnGT_clear(mclBnGT *x){}
mclSize mclBnGT_deserialize(mclBnGT *x, const void *buf, mclSize bufSize){ return 0; }
void mclBnGT_div(mclBnGT *z, const mclBnGT *x, const mclBnGT *y){}
mclSize mclBnGT_getStr(char *buf, mclSize maxBufSize, const mclBnGT *x, int ioMode) { return 0; }
void mclBnGT_inv(mclBnGT *y, const mclBnGT *x){}
int mclBnGT_isEqual(const mclBnGT *x, const mclBnGT *y){ return 0; }
int mclBnGT_isOne(const mclBnGT *x){ return 0; }
int mclBnGT_isZero(const mclBnGT *x){ return 0; }
void mclBnGT_mul(mclBnGT *z, const mclBnGT *x, const mclBnGT *y){}
void mclBnGT_neg(mclBnGT *y, const mclBnGT *x){}
void mclBnGT_pow(mclBnGT *z, const mclBnGT *x, const mclBnFr *y){}
mclSize mclBnGT_serialize(void *buf, mclSize maxBufSize, const mclBnGT *x){ return 0; }
void mclBnGT_setInt(mclBnGT *y, mclInt x){}
int mclBnGT_setStr(mclBnGT *x, const char *buf, mclSize bufSize, int ioMode){ return 0; }
void mclBnGT_sub(mclBnGT *z, const mclBnGT *x, const mclBnGT *y){}
int mclBn_FrEvaluatePolynomial(mclBnFr *out, const mclBnFr *cVec, mclSize cSize, const mclBnFr *x){ return 0; }
int mclBn_FrLagrangeInterpolation(mclBnFr *out, const mclBnFr *xVec, const mclBnFr *yVec, mclSize k){ return 0; }
int mclBn_G1EvaluatePolynomial(mclBnG1 *out, const mclBnG1 *cVec, mclSize cSize, const mclBnFr *x){ return 0; }
int mclBn_G1LagrangeInterpolation(mclBnG1 *out, const mclBnFr *xVec, const mclBnG1 *yVec, mclSize k){ return 0; }
int mclBn_G2EvaluatePolynomial(mclBnG2 *out, const mclBnG2 *cVec, mclSize cSize, const mclBnFr *x){ return 0; }
int mclBn_G2LagrangeInterpolation(mclBnG2 *out, const mclBnFr *xVec, const mclBnG2 *yVec, mclSize k){ return 0; }
void mclBn_finalExp(mclBnGT *y, const mclBnGT *x){}
mclSize mclBn_getCurveOrder(char *buf, mclSize maxBufSize){ return 0; }
mclSize mclBn_getFieldOrder(char *buf, mclSize maxBufSize){ return 0; }
int mclBn_getOpUnitSize(void) { return 0; }
int mclBn_getUint64NumToPrecompute(void){ return 0; }
void mclBn_millerLoop(mclBnGT *z, const mclBnG1 *x, const mclBnG2 *y){}
void mclBn_pairing(mclBnGT *z, const mclBnG1 *x, const mclBnG2 *y){}
void mclBn_precomputeG2(uint64_t *Qbuf, const mclBnG2 *Q){}
void mclBn_precomputedMillerLoop(mclBnGT *f, const mclBnG1 *P, const uint64_t *Qbuf){}
void mclBn_precomputedMillerLoop2(mclBnGT *f, const mclBnG1 *P1, const uint64_t *Q1buf, const mclBnG1 *P2, const uint64_t *Q2buf) {}

