package sphincsplus

import (
  "github.com/katzenpost/sphincsplus/shake-avx2/keccak4x"
)

func Keccak4xEnabled() bool {
  return keccak4x.Keccak4x_enabled()
}
