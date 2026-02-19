package analyzer

import (
	"unicode"
)

// isLowercase checks if the first letter is lowercase
func isLowercase(s string) bool {
	if len(s) == 0 {
		return true
	}
	firstRune := []rune(s)[0]
	return unicode.IsLower(firstRune)
}

// containsNonEnglish checks if string contains non-English characters
func containsNonEnglish(s string) bool {
	for _, r := range s {
		// Skip common punctuation and whitespace
		if unicode.IsSpace(r) || unicode.IsPunct(r) || unicode.IsDigit(r) {
			continue
		}
		
		// Check if it's not a basic Latin letter
		if !unicode.Is(unicode.Latin, r) || r > 127 {
			return true
		}
	}
	return false
}

// containsSpecialChars checks for special characters and emoji
func containsSpecialChars(s string) bool {
	// Define allowed characters: letters, digits, spaces, basic punctuation
	for _, r := range s {
		// Allow basic punctuation that's normal in logs
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			continue
		}
		
		// Allow common punctuation: . , : - _ /
		if r == '.' || r == ',' || r == ':' || r == '-' || r == '_' || r == '/' {
			continue
		}
		
		// Everything else is considered special (!, ?, emoji, etc.)
		return true
	}
	return false
}

// containsSensitiveData checks for sensitive data keywords
func containsSensitiveData(s string) bool {
	sensitiveKeywords := []string{
		"password", "pwd", "passwd",
		"token", "jwt", "bearer",
		"api_key", "apikey", "api-key",
		"secret", "private_key", "private-key",
		"credit_card", "card_number",
		"ssn", "social_security",
	}
	
	// Convert to lowercase for case-insensitive comparison
	lower := toLower(s)
	
	for _, keyword := range sensitiveKeywords {
		if contains(lower, keyword) {
			return true
		}
	}
	
	return false
}

// toLower converts string to lowercase
func toLower(s string) string {
	runes := []rune(s)
	for i, r := range runes {
		runes[i] = unicode.ToLower(r)
	}
	return string(runes)
}

// contains checks if s contains substr
func contains(s, substr string) bool {
	return len(s) >= len(substr) && indexOfSubstring(s, substr) >= 0
}

// indexOfSubstring finds the index of substr in s
func indexOfSubstring(s, substr string) int {
	if len(substr) == 0 {
		return 0
	}
	if len(s) < len(substr) {
		return -1
	}
	
	for i := 0; i <= len(s)-len(substr); i++ {
		if s[i:i+len(substr)] == substr {
			return i
		}
	}
	return -1
}

// removeSpecialChars removes special characters and emoji from string
func removeSpecialChars(s string) string {
	var result []rune
	
	for _, r := range s {
		// Keep letters, digits, spaces
		if unicode.IsLetter(r) || unicode.IsDigit(r) || unicode.IsSpace(r) {
			result = append(result, r)
			continue
		}
		
		// Keep common punctuation: . , : - _ /
		if r == '.' || r == ',' || r == ':' || r == '-' || r == '_' || r == '/' {
			result = append(result, r)
			continue
		}
		
		// Skip everything else (special chars, emoji)
	}
	
	// Trim trailing spaces that might be left after removing chars
	str := string(result)
	// Simple trim - remove trailing spaces
	for len(str) > 0 && str[len(str)-1] == ' ' {
		str = str[:len(str)-1]
	}
	
	return str
}
