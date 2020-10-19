package server

import "GoWeb/domain"

type LoginService interface {
	Login(user *domain.User) error
}
