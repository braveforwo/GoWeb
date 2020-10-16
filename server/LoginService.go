package server

import "GoWeb/domain"

type LoginService interface {
	login(user domain.User)
}
