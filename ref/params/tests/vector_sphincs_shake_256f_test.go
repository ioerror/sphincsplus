//go:build cgo && ((linux && amd64) || (darwin && amd64)) && SPECIFIC_BUILD_TAG

package tess

import (
	"encoding/hex"
	"testing"

	"github.com/stretchr/testify/require"
  sphincsplus "github.com/katzenpost/sphincsplus/ref"
)

func TestSignVerifyVectorSpecifics(t *testing.T) {

	privKey := new(PrivateKey)
	privKeyBytes, err := hex.DecodeString("")
	require.NoError(t, err)
	err = privKey.FromBytes(privKeyBytes)
	require.NoError(t, err)

	pubKey := new(PublicKey)
	pubKeyBytes, err := hex.DecodeString("")
	require.NoError(t, err)
	err = pubKey.FromBytes(pubKeyBytes)

	signature, err := hex.DecodeString("")
	require.NoError(t, err)

	t.Logf("SPHINCS+ scheme parameters: %s", SchemeName())
	t.Logf("priv key %x", privKey.Bytes())
	t.Logf("pub key %x", pubKey.Bytes())
	t.Logf("sig %x", signature)
	message := []byte("i am a message")
	require.True(t, pubKey.Verify(signature, message))

}
