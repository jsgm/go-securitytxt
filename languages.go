package securitytxt

import iso6391 "github.com/emvi/iso-639-1"

func (s *SecurityTxt) Language() iso6391.Language {
	// Returns the first appearing language.
	// Note that the order in which they appear is not an indication of priority; the
	// listed languages are intended to have equal priority.
	langs := s.Languages()
	return langs[0]
}

func (s *SecurityTxt) Languages() []iso6391.Language {
	if len(s.preferredLanguages) == 0 {
		// If this field is absent, security researchers may assume that English
		// is the language to be used (as per Section 4.5 of [RFC2277]).
		return []iso6391.Language{iso6391.FromCode("en")}
	}

	langs := make([]iso6391.Language, len(s.preferredLanguages))
	copy(langs, s.preferredLanguages)
	return langs
}

func (s *SecurityTxt) ContainsLanguage(code string) bool {
	for _, language := range s.Languages() {
		if language.Code == code {
			return true
		}
	}
	return false
}
