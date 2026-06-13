package core

import (
	"log"

	"github.com/go-playground/validator/v10"
)

var (
	Validator  *validator.Validate
	specialMap map[rune]bool = make(map[rune]bool)
)

func initVariables() {
	list := "¡¿ªºçÇñÑ!@#$%^&*(),.?:{}|<>_+-=[]';/^~`\\\""

	for _, char := range list {
		specialMap[char] = true
	}
}

func securePasswordValidator(fl validator.FieldLevel) bool {
	var flags int = 0

	for _, char := range fl.Field().String() {
		if specialMap[char] {
			flags |= 1
		} else if char >= '0' && char <= '9' {
			flags |= 2
		} else if char >= 'A' && char <= 'Z' {
			flags |= 4
		} else if char >= 'a' && char <= 'z' {
			flags |= 8
		}

		if flags == 15 {
			return true
		}
	}

	return false
}

func RegisterCustomValidators() error {
	initVariables()

	Validator = validator.New()

	err := Validator.RegisterValidation("securepassword", securePasswordValidator)

	if err != nil {
		log.Printf("Err register password validator: %v", err)
		return err
	}

	return nil
}
