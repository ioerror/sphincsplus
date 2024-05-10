package sphincsplus

import (
	"github.com/stretchr/testify/require"
	"testing"
)

func TestSchemeName(t *testing.T) {
	t.Logf("SPHINCS+ SignatureScheme: %s", SignatureScheme)
	t.Logf("SPHINCS+ scheme from Name(): %s", Name())
	require.Equal(t, SignatureScheme, Name())
}
