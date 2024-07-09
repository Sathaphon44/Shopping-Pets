package utils

func ExtractTokenFromHeader(header string) string {
	// Check if the header starts with "Bearer "
	if len(header) > 7 && header[:7] == "Bearer " {
		token := header[7:]
		if token == "" {
			return "Invalid or missing token"
		}
		return token
	}
	return ""
}
