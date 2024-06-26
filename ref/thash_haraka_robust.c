//go:build cgo && ((linux && amd64) || (darwin && amd64) || (darwin && arm64) || (windows && amd64)) && ((sphincs_haraka_128f || sphincs_haraka_192f || sphincs_haraka_256f) || (!sphincs_haraka_128f && !sphincs_haraka_128s && !sphincs_haraka_192f && !sphincs_haraka_192s && !sphincs_haraka_256f && !sphincs_haraka_256s && !sphincs_sha2_128f && !sphincs_sha2_128s && !sphincs_sha2_192f && !sphincs_sha2_192s && !sphincs_sha2_256f && !sphincs_sha2_256s && !sphincs_shake_128f && !sphincs_shake_128s && !sphincs_shake_192f && !sphincs_shake_192s && !sphincs_shake_256f && !sphincs_shake_256s))

#include <stdint.h>
#include <string.h>

#include "thash.h"
#include "address.h"
#include "params.h"
#include "utils.h"

#include "haraka.h"

/**
 * Takes an array of inblocks concatenated arrays of SPX_N bytes.
 */
void thash(unsigned char *out, const unsigned char *in, unsigned int inblocks,
           const spx_ctx *ctx, uint32_t addr[8])
{
    SPX_VLA(uint8_t, buf, SPX_ADDR_BYTES + inblocks*SPX_N);
    SPX_VLA(uint8_t, bitmask, inblocks*SPX_N);
    unsigned char outbuf[32];
    unsigned char buf_tmp[64];
    unsigned int i;

    if (inblocks == 1) {
        /* F function */
        /* Since SPX_N may be smaller than 32, we need a temporary buffer. */
        memset(buf_tmp, 0, 64);
        memcpy(buf_tmp, addr, 32);

        haraka256(outbuf, buf_tmp, ctx);
        for (i = 0; i < inblocks * SPX_N; i++) {
            buf_tmp[SPX_ADDR_BYTES + i] = in[i] ^ outbuf[i];
        }
        haraka512(outbuf, buf_tmp, ctx);
        memcpy(out, outbuf, SPX_N);
    } else {
        /* All other tweakable hashes*/
        memcpy(buf, addr, 32);
        haraka_S(bitmask, inblocks * SPX_N, buf, SPX_ADDR_BYTES, ctx);

        for (i = 0; i < inblocks * SPX_N; i++) {
            buf[SPX_ADDR_BYTES + i] = in[i] ^ bitmask[i];
        }

        haraka_S(out, SPX_N, buf, SPX_ADDR_BYTES + inblocks*SPX_N, ctx);
    }
}
