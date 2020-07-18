package responsecodes

/*
general
*/

var ErrResponseDatabase = ErrorResponse{
	Error: &Error{500, "Database exception"},
}

/*
 u/add
*/

var SuccessResponseOnAdd = SuccessResponse{
	Data: &Data{
		"The provided name has been successfully added to the database",
	},
}

var ErrResponseOnAddMissing = ErrorResponse{
	Error: &Error{422, "Either Name or Quote is missing"},
}

var ErrResponseOnAddDatabase = ErrorResponse{
	Error: &Error{500, "There was a problem adding to the database"},
}

/*
u/get
*/

var ErrResponseOnGetInvalidParameter = ErrorResponse{
	Error: &Error{422, "Please check the parameters and try again"},
}

var ErrResponseOnGetNotExist = ErrorResponse{
	Error: &Error{404, "The user does not exist"},
}

func SuccessResponseOnGet(quote string) SuccessResponseGet {
	return SuccessResponseGet{
		Data: &DataGet{quote},
	}
}
