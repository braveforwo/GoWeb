package server

import "GoWeb/domain"

type RegisterService interface {
	register(user domain.User) bool
}
