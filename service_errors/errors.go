package service_errors

type ServiceError struct {
	ServiceErrorHeader string `json:"header"`
	ServiceError  string `json:"error"`
	ServicePayload interface{} `json:"payload"`
}

func NewServiceError(err error, payload interface{}) *ServiceError {
	serviceError := ServiceError{
		ServiceError:  err.Error(),
		ServicePayload: payload,
	}
	return &serviceError
}