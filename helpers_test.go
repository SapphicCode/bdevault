package bdevault

import "testing"

func TestGetRecoveryKey(t *testing.T) {
	key, err := GetRecoveryKey("C:")
	if err != nil {
		t.Fatal(err)
	}
	if key == "" {
		t.Fatalf("Key was empty")
	}
	t.Log(key)
}

func TestGetRecoveryKeys(t *testing.T) {
	keyMap, err := GetRecoveryKeys()
	if err != nil {
		t.Error(err)
	}
	t.Log(keyMap)
}
