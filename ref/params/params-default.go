//go:build cgo && (!sphincs_haraka_128f && !sphincs_haraka_128s && !sphincs_haraka_192f && !sphincs_haraka_192s && !sphincs_haraka_256f && !sphincs_haraka_256s && !sphincs_sha2_128f && !sphincs_sha2_128s && !sphincs_sha2_192f && !sphincs_sha2_192s && !sphincs_sha2_256f && !sphincs_sha2_256s && !sphincs_shake_128f && !sphincs_shake_128s && !sphincs_shake_192f && !sphincs_shake_192s && !sphincs_shake_256f && !sphincs_shake_256s) && (!haraka_aesni && !sha2_avx2 && !shake_a64 && !shake_avx2 && !ref)
package params

var BuildTagSignatureName = "sphincs_haraka_128f"
