package errors

import "errors"

var LoginBusy = errors.New("login is busy")
var UserIsAbsent = errors.New("user is absent")

var UserUnauthorized = errors.New("user is unauthorized")
var ListTooShort = errors.New("list argument must be larger")
var AccessError = errors.New("access error")

var NoneRowsAffected = errors.New("no rows affected")
