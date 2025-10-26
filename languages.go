package securitytxt

import iso6391 "github.com/emvi/iso-639-1"

func (s *SecurityTxt) PreferredLanguage() iso6391.Language {
	langs := s.PreferredLanguages()
	return langs[0]
}

func (s *SecurityTxt) PreferredLanguages() []iso6391.Language {
	if len(s.preferredLanguages) == 0 {
		// If this field is absent, security researchers may assume that English
		// is the language to be used (as per Section 4.5 of [RFC2277]).
		return []iso6391.Language{iso6391.FromCode("en")}
	}

	langs := make([]iso6391.Language, len(s.preferredLanguages))
	copy(langs, s.preferredLanguages)
	return langs
}

func (s *SecurityTxt) IsPreferredLanguage(code string) bool {
	return iso6391.ValidCode(code) && code == s.PreferredLanguage().Code
}
