#include <stdint.h> // for uint64_, uint8_t
#include <stdlib.h> // for size_t

#define MCLBN_COMPILED_TIME_VAR 0
#define MCLBN_IO_SERIALIZE_HEX_STR 0
#define mclSize size_t
#define mclInt int64_t

enum {
  MCL_BN254 = 0,
  MCL_BN381_1 = 1,
  MCL_BN381_2 = 2,
  MCL_BN462 = 3,
  MCL_BN_SNARK1 = 4,
  MCL_BLS12_381 = 5,
  MCL_BN160 = 6
};

enum {
  mclBn_CurveFp254BNb = 0,
  mclBn_CurveFp382_1 = 1,
  mclBn_CurveFp382_2 = 2,
  mclBn_CurveFp462 = 3,
  mclBn_CurveSNARK1 = 4,
  mclBls12_CurveFp381 = 5
};

typedef struct {
} blsPublicKey;
typedef struct {
} blsSecretKey;
typedef struct {
} blsSignature;
typedef struct {
} mclBnFr;

void blsGetPublicKey(blsPublicKey *pub, const blsSecretKey *sec);
int blsInit(int curve, int compiledTimeVar);
void blsSign(blsSignature *sig, const blsSecretKey *sec, const void *m, mclSize size);
void blsSignatureAdd(blsSignature *sig, const blsSignature *rhs);
int blsVerify(const blsSignature *sig, const blsPublicKey *pub, const void *m, mclSize size);
void mclBnFr_add(mclBnFr *z, const mclBnFr *x, const mclBnFr *y);
void mclBnFr_clear(mclBnFr *x);
mclSize mclBnFr_deserialize(mclBnFr *x, const void *buf, mclSize bufSize);
void mclBnFr_div(mclBnFr *z, const mclBnFr *x, const mclBnFr *y);
mclSize mclBnFr_getStr(char *buf, mclSize maxBufSize, const mclBnFr *x, int ioMode);
int mclBnFr_setStr(mclBnFr *x, const char *buf, mclSize bufSize, int ioMode);
void mclBnFr_inv(mclBnFr *y, const mclBnFr *x);
int mclBnFr_isEqual(const mclBnFr *x, const mclBnFr *y);
int mclBnFr_setLittleEndian(mclBnFr *x, const void *buf, mclSize bufSize);
int mclBnFr_isOne(const mclBnFr *x);
int mclBnFr_isZero(const mclBnFr *x);
void mclBnFr_mul(mclBnFr *z, const mclBnFr *x, const mclBnFr *y);
void mclBnFr_neg(mclBnFr *y, const mclBnFr *x);
mclSize mclBnFr_serialize(void *buf, mclSize maxBufSize, const mclBnFr *x);
int mclBnFr_setByCSPRNG(mclBnFr *x);
int mclBnFr_setHashOf(mclBnFr *x, const void *buf, mclSize bufSize);
void mclBnFr_setInt(mclBnFr *y, mclInt x);
void mclBnFr_setInt32(mclBnFr *y, int x);
void mclBnFr_sub(mclBnFr *z, const mclBnFr *x, const mclBnFr *y);
typedef struct {} mclBnG1;
void mclBnG1_add(mclBnG1 *z, const mclBnG1 *x, const mclBnG1 *y);
void mclBnG1_clear(mclBnG1 *x);
void mclBnG1_dbl(mclBnG1 *y, const mclBnG1 *x);
mclSize mclBnG1_deserialize(mclBnG1 *x, const void *buf, mclSize bufSize);
mclSize mclBnG1_getStr(char *buf, mclSize maxBufSize, const mclBnG1 *x, int ioMode);
int mclBnG1_hashAndMapTo(mclBnG1 *x, const void *buf, mclSize bufSize);
int mclBnG1_isEqual(const mclBnG1 *x, const mclBnG1 *y);
int mclBnG1_isZero(const mclBnG1 *x);
void mclBnG1_mul(mclBnG1 *z, const mclBnG1 *x, const mclBnFr *y);
void mclBnG1_mulCT(mclBnG1 *z, const mclBnG1 *x, const mclBnFr *y);
void mclBnG1_neg(mclBnG1 *y, const mclBnG1 *x);
mclSize mclBnG1_serialize(void *buf, mclSize maxBufSize, const mclBnG1 *x);
int mclBnG1_setStr(mclBnG1 *x, const char *buf, mclSize bufSize, int ioMode);
void mclBnG1_sub(mclBnG1 *z, const mclBnG1 *x, const mclBnG1 *y);
typedef struct {} mclBnG2;
void mclBnG2_add(mclBnG2 *z, const mclBnG2 *x, const mclBnG2 *y);
void mclBnG2_clear(mclBnG2 *x);
void mclBnG2_dbl(mclBnG2 *y, const mclBnG2 *x);
mclSize mclBnG2_deserialize(mclBnG2 *x, const void *buf, mclSize bufSize);
mclSize mclBnG2_getStr(char *buf, mclSize maxBufSize, const mclBnG2 *x, int ioMode);
int mclBnG2_hashAndMapTo(mclBnG2 *x, const void *buf, mclSize bufSize);
int mclBnG2_isEqual(const mclBnG2 *x, const mclBnG2 *y);
int mclBnG2_isZero(const mclBnG2 *x);
void mclBnG2_mul(mclBnG2 *z, const mclBnG2 *x, const mclBnFr *y);
void mclBnG2_neg(mclBnG2 *y, const mclBnG2 *x);
mclSize mclBnG2_serialize(void *buf, mclSize maxBufSize, const mclBnG2 *x);
int mclBnG2_setStr(mclBnG2 *x, const char *buf, mclSize bufSize, int ioMode);
void mclBnG2_sub(mclBnG2 *z, const mclBnG2 *x, const mclBnG2 *y);
typedef struct {} mclBnGT;
void mclBnGT_add(mclBnGT *z, const mclBnGT *x, const mclBnGT *y);
void mclBnGT_clear(mclBnGT *x);
mclSize mclBnGT_deserialize(mclBnGT *x, const void *buf, mclSize bufSize);
void mclBnGT_div(mclBnGT *z, const mclBnGT *x, const mclBnGT *y);
mclSize mclBnGT_getStr(char *buf, mclSize maxBufSize, const mclBnGT *x, int ioMode);
void mclBnGT_inv(mclBnGT *y, const mclBnGT *x);
int mclBnGT_isEqual(const mclBnGT *x, const mclBnGT *y);
int mclBnGT_isOne(const mclBnGT *x);
int mclBnGT_isZero(const mclBnGT *x);
void mclBnGT_mul(mclBnGT *z, const mclBnGT *x, const mclBnGT *y);
void mclBnGT_neg(mclBnGT *y, const mclBnGT *x);
void mclBnGT_pow(mclBnGT *z, const mclBnGT *x, const mclBnFr *y);
mclSize mclBnGT_serialize(void *buf, mclSize maxBufSize, const mclBnGT *x);
void mclBnGT_setInt(mclBnGT *y, mclInt x);
int mclBnGT_setStr(mclBnGT *x, const char *buf, mclSize bufSize, int ioMode);
void mclBnGT_sub(mclBnGT *z, const mclBnGT *x, const mclBnGT *y);
int mclBn_FrEvaluatePolynomial(mclBnFr *out, const mclBnFr *cVec, mclSize cSize, const mclBnFr *x);
int mclBn_FrLagrangeInterpolation(mclBnFr *out, const mclBnFr *xVec, const mclBnFr *yVec, mclSize k);
int mclBn_G1EvaluatePolynomial(mclBnG1 *out, const mclBnG1 *cVec, mclSize cSize, const mclBnFr *x);
int mclBn_G1LagrangeInterpolation(mclBnG1 *out, const mclBnFr *xVec, const mclBnG1 *yVec, mclSize k);
int mclBn_G2EvaluatePolynomial(mclBnG2 *out, const mclBnG2 *cVec, mclSize cSize, const mclBnFr *x);
int mclBn_G2LagrangeInterpolation(mclBnG2 *out, const mclBnFr *xVec, const mclBnG2 *yVec, mclSize k);
void mclBn_finalExp(mclBnGT *y, const mclBnGT *x);
mclSize mclBn_getCurveOrder(char *buf, mclSize maxBufSize);
mclSize mclBn_getFieldOrder(char *buf, mclSize maxBufSize);
int mclBn_getOpUnitSize(void);
int mclBn_getUint64NumToPrecompute(void);
void mclBn_millerLoop(mclBnGT *z, const mclBnG1 *x, const mclBnG2 *y);
void mclBn_pairing(mclBnGT *z, const mclBnG1 *x, const mclBnG2 *y);
void mclBn_precomputeG2(uint64_t *Qbuf, const mclBnG2 *Q);
void mclBn_precomputedMillerLoop(mclBnGT *f, const mclBnG1 *P, const uint64_t *Qbuf);
void mclBn_precomputedMillerLoop2(mclBnGT *f, const mclBnG1 *P1, const uint64_t *Q1buf, const mclBnG1 *P2, const uint64_t *Q2buf);
