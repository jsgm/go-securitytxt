package securitytxt

import "fmt"

type Encryption struct {
	Source   EncryptionSource
	RawValue string
	Value    string
	Priority int
}

type EncryptionSource string

const (
	EncryptionSourceHTTPS       EncryptionSource = "https"       // "https://example.com/pgp-key.txt"
	EncryptionSourceDNS         EncryptionSource = "dns"         // "dns:5d2d37...example.com?type=OPENPGPKEY"
	EncryptionSourceOpenPGP4FPR EncryptionSource = "openpgp4fpr" // "openpgp4fpr:5f2de552..."
	EncryptionSourceUnknown     EncryptionSource = "unknown"     //
)

// TODO: Verify that the message is correctly signed

func (e *Encryption) FetchKey() (string, error) {
	switch e.Source {
	case EncryptionSourceHTTPS:
		//return downloadHTTPSKey(e.RawValue)
	case EncryptionSourceDNS:
		//return fetchDNSOpenPGPKey(e.RawValue)
	case EncryptionSourceOpenPGP4FPR:
		//return downloadKeyFromFingerprint(e.RawValue)
	default:
		return "", fmt.Errorf("unknown encryption source: %s", e.RawValue)
	}

	return "", nil
}
