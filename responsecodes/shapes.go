package responsecodes

/*
General responses
*/

type SuccessResponse struct {
	Data *Data
}

type Data struct {
	Message string 
}

type ErrorResponse struct {
	Error *Error
}

type Error struct {
	Code int
	Message string
}

/*
u/get
*/

type SuccessResponseGet struct {
	Data *DataGet
}

type DataGet struct {
	Quote string
}
