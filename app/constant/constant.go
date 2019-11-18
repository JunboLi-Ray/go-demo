package constant

import "errors"

/*
	错误提示
*/
var ParamError = errors.New("Param Error")

var ShutdownPasswordError = errors.New("Shutdown Password Error")

var JsonFail = errors.New("Json Fail")
