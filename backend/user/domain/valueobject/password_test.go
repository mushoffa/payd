package valueobject

import (
	"golang.org/x/crypto/bcrypt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Password_String(t *testing.T) {
	// Given
	input := "input password"

	// When
	result := Password(input).String()

	// Then
	assert.Equal(t, result, input)
}

func Test_Password_EmptyString(t *testing.T) {
	// Given
	input := ""

	// When
	result, err := Password(input).Hash()

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, err, ErrEmptyString)
	assert.Empty(t, result)
}

func Test_Password_Hash(t *testing.T) {
	// Given
	input := "adsfgh"

	// When
	result, err := Password(input).Hash()

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func Test_Password_Hash_EmptyString(t *testing.T) {
	// Given
	input := "a"

	// When
	result, err := Password(input).Hash()

	// Then
	assert.Nil(t, err)
	assert.NotNil(t, result)
	assert.NotEmpty(t, result)
}

func Test_Password_IsHashed_True(t *testing.T) {
	// Given
	plaintext := "adsfagadf"
	encrypted, err := Password(plaintext).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	result := Password(encrypted).IsHashed()

	// Then
	assert.NotNil(t, result)
	assert.True(t, result)
}

func Test_Password_IsHashed_False(t *testing.T) {
	// Given
	input := "adsfagadf"

	// When
	result := Password(input).IsHashed()

	// Then
	assert.NotNil(t, result)
	assert.False(t, result)
}

func Test_Password_Compare_EmptyString(t *testing.T) {
	// Given
	password := Password("12345")

	// When
	result, err := password.Compare("")

	// Then
	assert.Nil(t, err)
	assert.False(t, result)
}

func Test_Password_Compare_Both_Plaintext_Equal(t *testing.T) {
	// Given
	plaintext_a := "12345"
	plaintext_b := "12345"
	password := Password(plaintext_a)

	// When
	result, err := password.Compare(plaintext_b)

	// Then
	assert.Nil(t, err)
	assert.True(t, result)
}

func Test_Password_Compare_Both_Plaintext_NotEqual(t *testing.T) {
	// Given
	plaintext_a := "12345"
	plaintext_b := "123456"
	password := Password(plaintext_a)

	// When
	result, err := password.Compare(plaintext_b)

	// Then
	assert.Nil(t, err)
	assert.False(t, result)
}

func Test_Password_Compare_Plaintext_and_Hashed_Equal(t *testing.T) {
	// Given
	plaintext := "12345"
	hashed, err := Password(plaintext).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	password := Password(plaintext)
	result, err := password.Compare(hashed)

	// Then
	assert.Nil(t, err)
	assert.True(t, result)
}

func Test_Password_Compare_Plaintext_and_Hashed_NotEqual(t *testing.T) {
	// Given
	plaintext_a := "12345"
	plaintext_b := "123456"
	hashed, err := Password(plaintext_b).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	result, err := Password(plaintext_a).Compare(hashed)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, err, bcrypt.ErrMismatchedHashAndPassword)
	assert.False(t, result)
}

func Test_Password_Compare_Hashed_and_Plaintext_Equal(t *testing.T) {
	// Given
	plaintext := "12345"
	hashed, err := Password(plaintext).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	result, err := Password(plaintext).Compare(hashed)

	// Then
	assert.Nil(t, err)
	assert.True(t, result)
}

func Test_Password_Compare_Hashed_and_Plaintext_NotEqual(t *testing.T) {
	// Given
	plaintext_a := "12345"
	plaintext_b := "123456"
	hashed, err := Password(plaintext_a).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	result, err := Password(hashed).Compare(plaintext_b)

	// Then
	assert.NotNil(t, err)
	assert.Equal(t, err, bcrypt.ErrMismatchedHashAndPassword)
	assert.False(t, result)
}

func Test_Password_Compare_Both_Hashed_Equal(t *testing.T) {
	// Given
	plaintext := "12345"
	hashed, err := Password(plaintext).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	result, err := Password(hashed).Compare(hashed)

	// Then
	assert.Nil(t, err)
	assert.True(t, result)
}

func Test_Password_Compare_Both_Hashed_NotEqual(t *testing.T) {
	// Given
	plaintext_a := "12345"
	plaintext_b := "1234567"

	hashed_a, err := Password(plaintext_a).Hash()
	if err != nil {
		t.Fatal(err)
	}

	hashed_b, err := Password(plaintext_b).Hash()
	if err != nil {
		t.Fatal(err)
	}

	// When
	result, err := Password(hashed_a).Compare(hashed_b)

	// Then
	assert.Nil(t, err)
	assert.False(t, result)
}
