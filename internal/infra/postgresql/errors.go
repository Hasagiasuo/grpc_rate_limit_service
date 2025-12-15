package postgresql

import "errors"

var (
	ErrCannotCreateServiceTable = errors.New("cannot create services table")
	ErrCannotAddNewService = errors.New("cannot add new service into services table")
	ErrServiceNotFound = errors.New("service not found")
	ErrCannotUpdateServiceBalance = errors.New("cannot update service balance")
	ErrCannotDeleteService = errors.New("cannot delete service")
)