package models

import (
	"errors"
	"html"
	"log"
	"strings"

	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

type User struct {
	ID          uint32 `gorm:"primary_key;auto_increment" json:"id"`
	NamaLengkap string `gorm:"size:255;not null;unique" json:"namalengkap"`
	Username    string `gorm:"size:100;not null;unique" json:"username"`
	Password    string `gorm:"size:100;not null;" json:"password"`
	Foto        string `gorm:"size:100" json:"foto"`
}

func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
}

func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

func (u *User) BeforeSave() error {
	hashedPassword, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) Prepare() {
	u.ID = 0
	u.NamaLengkap = html.EscapeString(strings.TrimSpace(u.NamaLengkap))
	u.Username = html.EscapeString(strings.TrimSpace(u.Username))
	u.Foto = html.EscapeString(strings.TrimSpace(u.Username))
}

func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.NamaLengkap == "" {
			return errors.New("Required Nama Lengkap")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Username == "" {
			return errors.New("Required Username")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Username == "" {
			return errors.New("Required Username")
		}
		return nil

	default:
		if u.NamaLengkap == "" {
			return errors.New("Required Nama Lengkap")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Username == "" {
			return errors.New("Required Username")
		}
		return nil
	}
}

func (u *User) SaveUser(db *gorm.DB) (*User, error) {

	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError(err) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash the password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumns(
		map[string]interface{}{
			"password":    u.Password,
			"namalengkap": u.NamaLengkap,
			"username":    u.Username,
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	// This is the display the updated user
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {

	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})

	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}
