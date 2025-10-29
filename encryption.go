package securitytxt

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

// TO-DO:
// -----BEGIN PGP SIGNED MESSAGE-----
// Hash: SHA256
// -----BEGIN PGP SIGNATURE-----
// -----END PGP SIGNATURE-----
