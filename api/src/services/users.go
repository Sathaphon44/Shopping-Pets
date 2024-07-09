package services

import (
	"shoppets-api/domain/entities"
	"shoppets-api/domain/repository"
	"shoppets-api/src/utils"

	"golang.org/x/crypto/bcrypt"
)

type usersRepo struct {
	RepoUsers repository.IUsersRepo
}

type IUsersServices interface {
	Show_users_all() (interface{}, error)
	RegisterService(username, email, password string) (interface{}, error)
	LoginServices(email, password string) (interface{}, error)
	VerifyTokenService(token string) (*entities.ResponseUsers, error)
	DeleteAccountService(token string) (interface{}, error)
}

func NewUsersServices(repo repository.IUsersRepo) IUsersServices {
	return &usersRepo{
		RepoUsers: repo,
	}
}

func (c usersRepo) Show_users_all() (interface{}, error) {
	data, err := c.RepoUsers.FindUsers()
	if err != nil {
		return nil, err
	}
	return data, nil
}

func (c usersRepo) LoginServices(email, password string) (interface{}, error) {
	// Fetch the hashed password for the given email from your database
	result, err := c.RepoUsers.FindOneUser(email)
	if err != nil {
		return "ไม่พบบัญชีของท่าน", err
	}

	// Compare the provided password with the stored hashed password
	if err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(password)); err != nil {
		return "รหัสผ่านของท่านไม่ถูกต้อง", err
	}

	// If the password is correct, generate a token
	token, err := utils.CreateToken(email)
	if err != nil {
		return "เกิดข้อผิดพลาด", err
	}

	return token, nil
}

func (c usersRepo) VerifyTokenService(token string) (*entities.ResponseUsers, error) {
	email, err := utils.VerifyToken(token)
	if err != nil {
		return nil, err
	}
	if email != "" {
		if result, err := c.RepoUsers.FindOneUser(email); err != nil {
			return nil, err
		} else {
			return result, nil
		}

	} else {
		return nil, nil
	}
}

func (c usersRepo) RegisterService(username, email, password string) (interface{}, error) {
	pHash, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		return nil, err
	} else {
		account_amount, err := c.RepoUsers.CountUsers(email)
		if err != nil {
			return nil, err
		} else if account_amount == int64(0) {
			status_insert, err := c.RepoUsers.InsertOneUser(username, email, pHash)
			if err != nil {
				return nil, err
			}
			return status_insert, nil
		}
		return account_amount, nil
	}
}

func (c usersRepo) DeleteAccountService(token string) (interface{}, error) {
	status_token, err := utils.VerifyToken(token)
	if err != nil {
		return nil, err
	} else if status_token != "" {
		status_delete, err := c.RepoUsers.DeleteOneUser(status_token)
		if err != nil {
			return "เกิดข้อผิดพลาดไม่สามารถดำเนินการลบบัญชีได้", err
		}
		return status_delete, nil

	}
	return nil, nil
}
