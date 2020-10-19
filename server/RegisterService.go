package server

import "GoWeb/domain"

type RegisterService interface {
	Register(user *domain.User) error
}
