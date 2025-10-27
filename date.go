package securitytxt

import "time"

func (s *SecurityTxt) HasExpiration() bool {
	return s.Expiration() != nil
}

func (s *SecurityTxt) Expiration() *time.Time {
	// Returns the actual expiration date
	return s.expires
}

func (s *SecurityTxt) Expired() bool {
	// Returns true if the document is expired
	if !s.HasExpiration() {
		// No expiration date means "not expired"
		return false
	}
	return time.Now().After(*s.expires)
}

func (s *SecurityTxt) DaysUntilExpiration() int {
	if !s.HasExpiration() {
		return 0
	}
	return int(time.Until(*s.Expiration()).Hours() / 24)
}
