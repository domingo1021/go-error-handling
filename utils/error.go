package utils

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"net/http"
)

type CustomizedError struct {
	Code    int
	Message string
	Err     error
}

func (ce *CustomizedError) Unwrap() error {
	return ce.Err
}

func (ce *CustomizedError) Error() string {
	return fmt.Sprintf("Code: %d, Message: %s, Error: %v", ce.Code, ce.Message, ce.Err)
}

type ApiError struct {
	RequestId  string
	StatusCode int
	CustomizedError
}

func (ae *ApiError) Error() string {
	return fmt.Sprintf("RequestId: %s, StatusCode: %d, Code: %d, Message: %s \n",
		ae.RequestId, ae.StatusCode, ae.Code, ae.Message)
}

func NewApiError(requestId string, statusCode int, code int, message string, err error) *ApiError {
	return &ApiError{
		RequestId:  requestId,
		StatusCode: statusCode,
		CustomizedError: CustomizedError{
			Code:    code,
			Message: message,
			Err:     err,
		},
	}
}

type DbError struct {
	CustomizedError
}

func NewDbNotFoundError() *DbError {
	return &DbError{
		CustomizedError: CustomizedError{
			Code:    ResourceNotFoundCode,
			Message: "Resource not found.",
			Err:     errors.New("resource not found"),
		},
	}
}

const (
	BadRequestCode       = 400000
	ResourceNotFoundCode = 404000
)

func HandleError(err error) {
	if err == nil {
		return
	}

	var de *DbError
	var ae *ApiError
	if errors.As(err, &de) {
		if de.Code == ResourceNotFoundCode {
			fmt.Printf(
				"Api error from db: %v\n",
				NewApiError(
					uuid.New().String(),
					http.StatusNotFound,
					ResourceNotFoundCode,
					"Resource not found",
					errors.New("resource not found"),
				),
			)
		}
	} else if errors.As(err, &ae) {
		fmt.Printf("Api error: %v\n", err)
	}
}
