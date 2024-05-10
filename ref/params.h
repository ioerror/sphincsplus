#ifndef PARAMS
#define PARAMS = sphincs-haraka-128f
#endif

#define sphincs_haraka_128f 0x01
#define sphincs_haraka_128s 0x02
#define sphincs_haraka_192f 0x03
#define sphincs_haraka_192s 0x04
#define sphincs_haraka_256f 0x05
#define sphincs_haraka_256s 0x06
#define sphincs_haraka_set (sphincs_haraka_128f | sphincs_haraka_128s | \
        sphincs_haraka_192f | sphincs_haraka_192s | sphincs_haraka_256f | \
        sphincs_haraka_256s)

#define HARAKA_PARAM(b) ((sphincs_haraka_set & b) == b)

#define sphincs_sha2_128f 0x07
#define sphincs_sha2_128s 0x08
#define sphincs_sha2_192f 0x09
#define sphincs_sha2_192s 0x10
#define sphincs_sha2_256f 0x11
#define sphincs_sha2_256s 0x12
#define sphincs_sha2_set (sphincs_sha2_128f | sphincs_sha2_128s | \
        sphincs_sha2_192f | sphincs_sha2_192s | sphincs_sha2_256f | \
        sphincs_sha2_256s)

#define SHA2_PARAM(c) ((sphincs_sha2_set & c) == c)

#define sphincs_shake_128f 0x13
#define sphincs_shake_128s 0x14
#define sphincs_shake_192f 0x15
#define sphincs_shake_192s 0x16
#define sphincs_shake_256f 0x17
#define sphincs_shake_256f_s sphincs-shake-256f
#define sphincs_shake_256s 0x18
#define sphincs_shake_set (sphincs_shake_128f | sphincs_shake_128s | \
        sphincs_shake_192f | sphincs_shake_192s | sphincs_shake_256f | \
        sphincs_shake_256s)

#define SHAKE_PARAM(d) ((sphincs_shake_set & d) == d)

#define VALID_PARAM(e) (((sphincs_haraka_set | sphincs_sha2_set | sphincs_shake_set) & e) == e)
#define MATCH_PARAM(f) ((PARAMS & f) == f)

#if defined(PARAMS) && VALID_PARAM(PARAMS)
#if HARAKA_PARAM(PARAMS)
  #if MATCH_PARAM(sphincs_haraka_128f)
    #define PARAMS_H sphincs-haraka-128f
  #elif MATCH_PARAM(sphincs_haraka_128s)
    #define PARAMS_H sphincs-haraka-128s
  #elif MATCH_PARAM(sphincs_haraka_192f) 
    #define PARAMS_H sphincs-haraka-192f
  #elif MATCH_PARAM(sphincs_haraka_192s)
    #define PARAMS_H sphincs-haraka-192s
  #elif MATCH_PARAM(sphincs_haraka_256f)
    #define PARAMS_H sphincs-haraka-256f
  #elif MATCH_PARAM(sphincs_haraka_256s)
    #define PARAMS_H sphincs-haraka-256s
  #endif
#elif SHA2_PARAM(PARAMS)
  #if MATCH_PARAM(sphincs_sha2_128f)
    #define PARAMS_H sphincs-sha2-128f
  #elif MATCH_PARAM(sphincs_sha2_128s)
    #define PARAMS_H sphincs-sha2-128s
  #elif MATCH_PARAM(sphincs_sha2_192f) 
    #define PARAMS_H sphincs-sha2-192f
  #elif MATCH_PARAM(sphincs_sha2_192s)
    #define PARAMS_H sphincs-sha2-192s
  #elif MATCH_PARAM(sphincs_sha2_256f)
    #define PARAMS_H sphincs-sha2-256f
  #elif MATCH_PARAM(sphincs_sha2_256s)
    #define PARAMS_H sphincs-sha2-256s
  #endif
#elif SHAKE_PARAM(PARAMS)
  #if MATCH_PARAM(sphincs_shake_128f)
    #define PARAMS_H sphincs-shake-128f
  #elif MATCH_PARAM(sphincs_shake_128s)
    #define PARAMS_H sphincs-shake-128s
  #elif MATCH_PARAM(sphincs_shake_192f) 
    #define PARAMS_H sphincs-shake-192f
  #elif MATCH_PARAM(sphincs_shake_192s)
    #define PARAMS_H sphincs-shake-192s
  #elif MATCH_PARAM(sphincs_shake_256f)
    #define PARAMS_H sphincs-shake-256f
  #elif MATCH_PARAM(sphincs_shake_256s)
    #define PARAMS_H sphincs-shake-256s
  #endif
#endif
#endif

#define str(s) #s
#define xstr(s) str(s)

#if defined(PARAMS_H)
#include xstr(params/params-PARAMS_H.h)
#else
#error "Set PARAMS to a valid parameter"
#endif
