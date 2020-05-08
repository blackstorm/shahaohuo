package service

import "errors"

var ResourcesExist = errors.New("resources exist")
var ResourcesNotFound = errors.New("resources not found")
var ServiceError = errors.New("service error")
var UsernameOrPasswordError = errors.New("username or password error")

func IsResourcesExist(err error) bool {
	return ResourcesExist == err
}

func IsResourcesNotFound(err error) bool {
	return ResourcesNotFound == err
}

func IsServiceError(err error) bool {
	return ServiceError == err
}

func IsUsernameOrPasswordError(err error) bool {
	return UsernameOrPasswordError == err
}