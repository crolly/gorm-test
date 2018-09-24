package models

import (
	"encoding/json"
	"time"

	"github.com/gobuffalo/validate"
	"github.com/gobuffalo/validate/validators"
	"github.com/jinzhu/gorm"
	"github.com/pkg/errors"
	"github.com/satori/go.uuid"
)

type User struct {
	ID        uuid.UUID `gorm:"primary_key;type:char(36);"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time
	Name      string `json:"name"`
}

// String is not required by pop and may be deleted
func (u User) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

// Users is not required by pop and may be deleted
type Users []User

// String is not required by pop and may be deleted
func (u Users) String() string {
	ju, _ := json.Marshal(u)
	return string(ju)
}

func (u *User) BeforeCreate(tx *gorm.DB) error {
	var err error
	u.ID, err = uuid.NewV4()
	return err
}

// Validate gets run every time you call a "pop.Validate*" (pop.ValidateAndSave, pop.ValidateAndCreate, pop.ValidateAndUpdate) method.
// This method is not required and may be deleted.
func (u *User) Validate(tx *gorm.DB) (*validate.Errors, error) {
	return validate.Validate(
		&validators.StringIsPresent{Field: u.Name, Name: "Name"},
	), nil
}

// ValidateCreate gets run every time you call "pop.ValidateAndCreate" method.
// This method is not required and may be deleted.
func (u *User) ValidateCreate(tx *gorm.DB) (*validate.Errors, error) {
	verrs, err := u.Validate(tx)
	if err != nil {
		return verrs, errors.WithStack(err)
	}
	e := tx.Create(u)
	return verrs, e.Error
}

// ValidateUpdate gets run every time you call "pop.ValidateAndUpdate" method.
// This method is not required and may be deleted.
func (u *User) ValidateUpdate(tx *gorm.DB) (*validate.Errors, error) {
	verrs, err := u.Validate(tx)
	if err != nil {
		return verrs, errors.WithStack(err)
	}
	e := tx.Model(u).Updates(u)
	return verrs, e.Error
}
