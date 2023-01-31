package user_test

import (
	"testing"

	. "github.com/zalgonoise/cloaki/user"
	"github.com/zalgonoise/x/errors"
)

func TestValidateUsername(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Run("Simple", func(t *testing.T) {
			input := "user"
			err := ValidateUsername(input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Success", func(t *testing.T) {
		t.Run("Special", func(t *testing.T) {
			input := "user_one"
			err := ValidateUsername(input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Fail", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			input := ""
			err := ValidateUsername(input)
			if !errors.Is(ErrEmptyUsername, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrEmptyUsername, err)
				return
			}
		})
		t.Run("TooShort", func(t *testing.T) {
			input := "x"
			err := ValidateUsername(input)
			if !errors.Is(ErrShortUsername, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrShortUsername, err)
				return
			}
		})
		t.Run("TooLong", func(t *testing.T) {
			// 26 chars
			input := "useruseruseruseruseruserus"
			err := ValidateUsername(input)
			if !errors.Is(ErrLongUsername, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrLongUsername, err)
				return
			}
		})
		t.Run("InvalidSpecial", func(t *testing.T) {
			input := "user one"
			err := ValidateUsername(input)
			if !errors.Is(ErrInvalidUsername, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrInvalidUsername, err)
				return
			}
		})
	})
}

func TestValidateName(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Run("Simple", func(t *testing.T) {
			input := "User"
			err := ValidateName(input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Success", func(t *testing.T) {
		t.Run("Special", func(t *testing.T) {
			input := "User One"
			err := ValidateName(input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Fail", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			input := ""
			err := ValidateName(input)
			if !errors.Is(ErrEmptyName, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrEmptyName, err)
				return
			}
		})
		t.Run("TooShort", func(t *testing.T) {
			input := "U"
			err := ValidateName(input)
			if !errors.Is(ErrShortName, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrShortName, err)
				return
			}
		})
		t.Run("TooLong", func(t *testing.T) {
			// 26 chars
			input := "Useruseruseruseruseruserus"
			err := ValidateName(input)
			if !errors.Is(ErrLongName, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrLongName, err)
				return
			}
		})
		t.Run("InvalidSpecial", func(t *testing.T) {
			input := "User_One"
			err := ValidateName(input)
			if !errors.Is(ErrInvalidName, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrInvalidName, err)
				return
			}
		})
	})
}

func BenchmarkValidatePassword(b *testing.B) {
	input := "L0ng_4nD-C0mP!3X^P@$$W0RD+!?L0ng_4nD-C0mP!3X^P@$$W0RD+!?"
	var err error

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		err = ValidatePassword(input)
	}
	_ = err
}

func TestValidatePassword(t *testing.T) {
	t.Run("Success", func(t *testing.T) {
		t.Run("Simple", func(t *testing.T) {
			input := "SecretPassword"
			err := ValidatePassword(input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Success", func(t *testing.T) {
		t.Run("Special", func(t *testing.T) {
			input := "Secret0!["
			err := ValidatePassword(input)
			if err != nil {
				t.Errorf("unexpected error: %v", err)
				return
			}
		})
	})

	t.Run("Fail", func(t *testing.T) {
		t.Run("Empty", func(t *testing.T) {
			input := ""
			err := ValidatePassword(input)
			if !errors.Is(ErrEmptyPassword, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrEmptyPassword, err)
				return
			}
		})
		t.Run("TooShort", func(t *testing.T) {
			input := "x"
			err := ValidatePassword(input)
			if !errors.Is(ErrShortPassword, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrShortPassword, err)
				return
			}
		})
		t.Run("TooLong", func(t *testing.T) {
			// 301 chars
			input := "1111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111111"
			err := ValidatePassword(input)
			if !errors.Is(ErrLongPassword, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrLongPassword, err)
				return
			}
		})
		t.Run("InvalidSpecial", func(t *testing.T) {
			input := "Special Password"
			err := ValidatePassword(input)
			if !errors.Is(ErrInvalidPassword, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrInvalidPassword, err)
				return
			}
		})

		t.Run("InvalidRepeat", func(t *testing.T) {
			input := "Special0000"
			err := ValidatePassword(input)
			if !errors.Is(ErrInvalidPassword, err) {
				t.Errorf("unexpected error: wanted %v ; got %v", ErrInvalidPassword, err)
				return
			}
		})
	})
}
