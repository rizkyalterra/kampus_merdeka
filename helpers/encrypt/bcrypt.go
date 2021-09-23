package encrypt

import "golang.org/x/crypto/bcrypt"

func Hash(password string) (string, error) {
	result, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.MinCost)
	if err != nil {
		return "", err
	}
	return string(result), nil
}

func Addition(a, b int) int {
	result := a + b
	if result < 0 {
		result = 0
	}
	return result
}
