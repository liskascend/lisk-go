package crypto

import (
	"bytes"
	"encoding/hex"
	"testing"
)

var (
	defaultBytes                = []byte{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	default8BytesReversed       = []byte{7, 6, 5, 4, 3, 2, 1, 0}
	defaultBytesBigNumberString = "18591708106338011145"

	unpaddedData       = []byte{0, 1, 2, 3, 4}
	paddedData         = []byte{0, 1, 2, 3, 4, 2, 2}
	invalidPaddedData  = []byte{0, 1, 2, 3, 4, 2, 20}
	invalidPaddedData2 = []byte{0, 1, 2, 3, 4, 2, 3}
)

func TestGetSHA256Hash(t *testing.T) {
	if val := GetSHA256Hash(defaultPassphrase); hex.EncodeToString(val[:]) != defaultPassphraseHash {
		t.Errorf("GetSHA256Hash(%v)=%v; want %v", defaultPassphrase, hex.EncodeToString(val[:]), defaultPassphraseHash)
	}
}

func TestGetFirstEightBytesReversed(t *testing.T) {
	if val := GetFirstEightBytesReversed(defaultBytes); !bytes.Equal(val, default8BytesReversed) {
		t.Errorf("GetFirstEightBytesReversed(%v)=%v; want %v", defaultBytes, val, default8BytesReversed)
	}

	if val := GetFirstEightBytesReversed(nil); val != nil {
		t.Errorf("GetFirstEightBytesReversed(%v)=%v; want %v", nil, val, nil)
	}
}

func TestGetBigNumberStringFromBytes(t *testing.T) {
	if val := GetBigNumberStringFromBytes(defaultBytes); val != defaultBytesBigNumberString {
		t.Errorf("GetBigNumberStringFromBytes(%v)=%v; want %v", defaultBytes, val, defaultBytesBigNumberString)
	}
}
