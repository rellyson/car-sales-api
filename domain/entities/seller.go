package entities

import (
	"errors"

	"golang.org/x/crypto/bcrypt"
)

var (
	ErrPasswordDoesNotMeetMinReqs error = errors.New("password does not meet the minimum requirements")
)

type Seller struct {
	BaseEntity
	FullName string `json:"full_name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (s *Seller) IsPasswordValid(p string) bool {
	err := bcrypt.CompareHashAndPassword(
		[]byte(s.Password), []byte(p))

	return err == nil
}

func (s *Seller) HashPassword(p string) error {
	if len(p) < 8 {
		return ErrPasswordDoesNotMeetMinReqs
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(p), bcrypt.MinCost)
	if err != nil {
		return err
	}

	s.Password = string(hash)

	return nil
}
