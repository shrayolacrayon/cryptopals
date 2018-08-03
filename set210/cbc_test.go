package set210

import (
	"fmt"
	"testing"
)

func TestCBC(t *testing.T) {
	toEncrypt := []byte("hello I am going to test this but I want to see \n what it looks like when there are more than 2 blocks")
	key := []byte("YELLOW SUBMARINE")
	encrypted, err := CBCEncrypt(toEncrypt, key)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	decrypted, err := CBCDecrypt(encrypted, key)
	if err != nil {
		t.Errorf("unexpected error %s", err)
	}
	if string(toEncrypt) != string(decrypted[:len(toEncrypt)]) {
		t.Errorf("expected decrypted string to be the same")
	}
}
