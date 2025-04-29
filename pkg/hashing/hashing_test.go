package hashing

import "testing"

func TestHashAndVerify(t *testing.T) {
	password := "password"

	options := HashOptions{
		Value:   password,
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}

	hash, err := Hash(options)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}

	match, err := VerifyPassword(password, hash)
	if err != nil {
		t.Fatalf("Verifying failed: %v", err)
	}
	if !match {
		t.Fatalf("Incorrect password")
	}
}

func TestIncorrectHashAndVerify(t *testing.T) {
	password := "password"
	wrongPassword := "wrongPassword"

	options := HashOptions{
		Value:   password,
		Time:    1,
		Memory:  64 * 1024,
		Threads: 4,
		KeyLen:  32,
	}

	hash, err := Hash(options)
	if err != nil {
		t.Fatalf("Hashing failed: %v", err)
	}

	match, err := VerifyPassword(wrongPassword, hash)
	if err != nil {
		t.Fatalf("Verifying failed: %v", err)
	}
	if match {
		t.Fatalf("Incorrect password mathes")
	}
}

func TestUniqueSalts(t *testing.T) {
	options := HashOptions{
		Value:   "samePassword",
		Time:    1,
		Memory:  65536,
		Threads: 4,
		KeyLen:  32,
	}

	hash1, err := Hash(options)
	if err != nil {
		t.Fatal(err)
	}

	hash2, err := Hash(options)
	if err != nil {
		t.Fatal(err)
	}

	if hash1 == hash2 {
		t.Error("Hashes should be different due to different salts")
	}
}
