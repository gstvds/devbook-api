package models

import (
	"api/src/providers/hash_provider"
	"encoding/json"
	"errors"
	"github.com/go-playground/locales/pt_BR"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	br "github.com/go-playground/validator/v10/translations/pt_BR"
	"strings"
	"time"
)

type User struct {
	Id        uint64    `gorm:"autoIncrement;primaryKey" json:"id,omitempty"`
	Name      string    `gorm:"type:varchar(50);not null" json:"name,omitempty" validate:"required"`
	Username  string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"username,omitempty" validate:"required"`
	Email     string    `gorm:"type:varchar(50);not null;uniqueIndex" json:"email,omitempty" validate:"required,email"`
	Password  string    `gorm:"type:varchar(100);not null" json:"password,omitempty" validate:"required"`
	CreatedAt time.Time `gorm:"autoCreateTime" json:"crated_at,omitempty"`
}

func translator(validate *validator.Validate) ut.Translator {
	ptBR := pt_BR.New()
	universalTranslator := ut.New(ptBR, ptBR)
	translation, _ := universalTranslator.GetTranslator("pt_BR")
	br.RegisterDefaultTranslations(validate, translation)

	return translation
}

func (user *User) Validate(stage string) error {
	var validate = validator.New()
	translation := translator(validate)

	if err := user.format(stage); err != nil {
		return err
	}

	if stage == "register" {
		if err := validate.Struct(user); err != nil {
			errs := err.(validator.ValidationErrors)
			translatedErrs := errs.Translate(translation)

			output, _ := json.Marshal(translatedErrs)
			return errors.New(string(output))
		}
	} else {
		if err := validate.StructExcept(user, "Password"); err != nil {
			errs := err.(validator.ValidationErrors)
			translatedErrs := errs.Translate(translation)

			output, _ := json.Marshal(translatedErrs)
			return errors.New(string(output))
		}
	}

	return nil
}

func (user *User) format(stage string) error {
	user.Username = strings.TrimSpace(user.Username)
	user.Name = strings.TrimSpace(user.Name)
	user.Email = strings.TrimSpace(user.Email)

	if stage == "register" {
		hashedPassword, err := hash_provider.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(hashedPassword)
	}

	return nil
}
