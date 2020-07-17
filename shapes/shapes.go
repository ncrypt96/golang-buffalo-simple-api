package shapes

//SuccessResponse is the response sent on success
type SuccessResponse struct {
	Data *Data
}

//Data is sent when the result is success
type Data struct {
	Message string
}

//ErrorResponse is the response sent on error
type ErrorResponse struct {
	Error *Error
}

//Error is sent when the result is error
type Error struct {
	Code    int
	Message string
}
