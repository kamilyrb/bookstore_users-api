package users

import "github.com/kamilyrb/bookstore_users-api/utils/errors"

type User struct {
	Id          int64  `json:"id"`
	FirstName   string `json:"first_name"`
	LastName    string `json:"last_name"`
	Email       string `json:"email"`
	DateCreated string `json:"date_created"`
}

func (user *User) Validate() *errors.RestErr {
	if user.Email == "" {
		return errors.NewBadRequestError("Invalid email address")
	}
	return nil
}
