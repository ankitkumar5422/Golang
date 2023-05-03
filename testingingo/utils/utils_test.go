package utils

import(
	"testing"

)

func TestEncryptDecrypt(t *testing.T){
	originalText := "Text massages"

	encryptedText, err := Encrypt(originalText)
	if err != nil {
		t.Errorf("Error Encrypting Text %v", err)
	}

	decryptedText , err := Decrypt(encryptedText)
	if err != nil {
		t.Errorf("Error Decrypting Text %v", err)
	}

	if decryptedText != originalText {
		t.Errorf("Expected Decrypted text is be %q but got %q", originalText, decryptedText)
	}



}