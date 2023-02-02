package main

import (
	"encoding/base64"
	"testing"
)

func TestAesEncryptDecrypt(t *testing.T) {
	origin := []byte("github.com")
	encrypt, err := AesEncrypt(origin)
	if err != nil {
		t.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(encrypt)
	origin, err = base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		t.Fatal(err)
	}
	decrypt, err := AesDecrypt(origin)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("base64:%s origin:%s decrypt:%s\n", base64Str, origin, decrypt)
}

func TestAesDecrypt(t *testing.T) {
	base64Str := "N4BDspfO+T0qsPnk1PJOtA=="
	origin, err := base64.StdEncoding.DecodeString(base64Str)
	if err != nil {
		t.Fatal(err)
	}
	t.Log("origin:", origin)
	decrypt, err := AesDecrypt(origin)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("base64:%s origin:%s decrypt:%s\n", base64Str, origin, decrypt)
}

func TestAesEncrypt(t *testing.T) {
	origin := []byte("github.com")
	encrypt, err := AesEncrypt(origin)
	if err != nil {
		t.Fatal(err)
	}
	base64Str := base64.StdEncoding.EncodeToString(encrypt)
	t.Logf("origin:%s encrypt:%s base64:%s \n", origin, encrypt, base64Str)
}
