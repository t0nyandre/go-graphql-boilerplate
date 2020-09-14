package user

import (
	"time"

	"github.com/alexedwards/argon2id"
	"github.com/t0nyandre/go-graphql-boilerplate/internal/logger"
	"gorm.io/gorm"
)

// User model
type User struct {
	ID        string    `json:"_id,omitempty" gorm:"type:varchar(25);unique;not null"`
	Username  string    `json:"username,omitempty" gorm:"type:varchar(80);unique;not null"`
	Password  string    `json:"password,omitmepty" gorm:"not null"`
	Email     string    `json:"email,omitempty" gorm:"unique;not null"`
	CreatedAt time.Time `json:"created_at,omitempty"`
	UpdatedAt time.Time `json:"updated_at,omitempty"`
}

func (u *User) hashPassword() string {
	hash, err := argon2id.CreateHash(u.Password, &argon2id.Params{
		Iterations:  3,
		Memory:      4096,
		Parallelism: 2,
		SaltLength:  16,
		KeyLength:   32,
	})
	if err != nil {
		logger := logger.NewLogger()
		logger.Panicw("Could not hash password for user: %s", u.Username)
	}

	return hash
}

func (u *User) verifyPassword(password string) bool {
	match, err := argon2id.ComparePasswordAndHash(password, u.Password)
	if err != nil {
		return false
	}

	return match
}

// BeforeCreate hook sets the created date and hashes the users password
func (u *User) BeforeCreate(tx *gorm.DB) error {
	t := time.Now()
	tx.Statement.SetColumn("CreatedAt", t.Format(time.RFC822Z))
	tx.Statement.SetColumn("UpdatedAt", t.Format(time.RFC822Z))
	tx.Statement.SetColumn("Password", u.hashPassword())
	return nil
}

// BeforeUpdate sets the updated date
func (u *User) BeforeUpdate(tx *gorm.DB) error {
	t := time.Now()
	tx.Statement.SetColumn("UpdatedAt", t.Format(time.RFC822Z))
	return nil
}
