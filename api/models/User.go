package models

import (
	"errors"
	"html"
	"log"
	"strings"
	"time"

	"github.com/badoux/checkmail"
	"github.com/jinzhu/gorm"
	"golang.org/x/crypto/bcrypt"
)

// User struct to store data about userss
type User struct {
	ID        uint32    `gorm:"primary_key;auto_increment" json:"id"`
	Nickname  string    `gorm:"size:255;not null" json:"nickname"`
	Email     string    `gorm:"size:100;not null" json:"email"`
	Password  string    `gorm:"size:100;not null" json:"password"`
	Role      string    `gorm:"size:100;not null" json:"role"`
	CreatedAt time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAT time.Time `gorm:"default:CURRENT_TIMESTAMP" json:"updated_at"`
}

// Hash func to encrypt
func Hash(password string) ([]byte, error) {
	return bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

}

// VerifyPassword func to compare two passwords
func VerifyPassword(hashedPassword, password string) error {
	return bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(password))
}

// BeforeSave func to prepare password
func (u *User) BeforeSave() error {
	hashedPassowrd, err := Hash(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassowrd)
	return nil
}

// Prepare func to fill model with dummy data
func (u *User) Prepare() {
	u.ID = 0
	u.Nickname = html.EscapeString((strings.TrimSpace(u.Nickname)))
	u.Email = html.EscapeString((strings.TrimSpace(u.Email)))
	u.Role = html.EscapeString((strings.TrimSpace(u.Role)))
	u.CreatedAt = time.Now()
	u.UpdatedAT = time.Now()
}

// Validate func to check for errors
func (u *User) Validate(action string) error {
	switch strings.ToLower(action) {
	case "update":
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required email")
		}
		if u.Role == "" {
			return errors.New("Required role")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	case "login":
		if u.Password == "" {
			return errors.New("Required Passowrd")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	default:
		if u.Nickname == "" {
			return errors.New("Required Nickname")
		}
		if u.Password == "" {
			return errors.New("Required Password")
		}
		if u.Email == "" {
			return errors.New("Required Email")
		}
		if u.Role == "" {
			return errors.New("Required role")
		}
		if err := checkmail.ValidateFormat(u.Email); err != nil {
			return errors.New("Invalid Email")
		}
		return nil
	}
}

// SaveUser func to insert data into database
func (u *User) SaveUser(db *gorm.DB) (*User, error) {
	var err error
	err = db.Debug().Create(&u).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// FindAllUsers func to get all user in database
func (u *User) FindAllUsers(db *gorm.DB) (*[]User, error) {
	var err error
	users := []User{}
	err = db.Debug().Model(&User{}).Limit(100).Find(&users).Error
	if err != nil {
		return &[]User{}, err
	}
	return &users, err
}

// FindUserByID func to get user by id
func (u *User) FindUserByID(db *gorm.DB, uid uint32) (*User, error) {
	var err error
	err = db.Debug().Model(User{}).Where("id = ? ", uid).Take(&u).Error
	if err != nil {
		return &User{}, err
	}
	if gorm.IsRecordNotFoundError((err)) {
		return &User{}, errors.New("User Not Found")
	}
	return u, err
}

// UpdateAUser func to update data about user
func (u *User) UpdateAUser(db *gorm.DB, uid uint32) (*User, error) {

	// To hash password
	err := u.BeforeSave()
	if err != nil {
		log.Fatal(err)
	}
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).UpdateColumn(
		map[string]interface{}{
			"password":  u.Password,
			"nickname":  u.Nickname,
			"email":     u.Email,
			"role":      u.Role,
			"update_at": time.Now(),
		},
	)
	if db.Error != nil {
		return &User{}, db.Error
	}
	err = db.Debug().Model(&User{}).Where("id = ?", uid).Error
	if err != nil {
		return &User{}, err
	}
	return u, nil
}

// DeleteAUser func to remove user
func (u *User) DeleteAUser(db *gorm.DB, uid uint32) (int64, error) {
	db = db.Debug().Model(&User{}).Where("id = ?", uid).Take(&User{}).Delete(&User{})
	if db.Error != nil {
		return 0, db.Error
	}
	return db.RowsAffected, nil
}