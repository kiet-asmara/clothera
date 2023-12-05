package handler

import "fmt"

var (
	ErrorAlreadyExists  = fmt.Errorf("record already exists")
	ErrorDuplicateEntry = fmt.Errorf("duplicate entry")
	ErrorRecordNotFound = fmt.Errorf("record not found")
)
