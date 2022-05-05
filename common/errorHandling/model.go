package errorHandling

type BaseError struct {
	error   `json:"-"`
	Message string                 `json:"message"`
	Status  int                    `json:"status"`
	Fields  []ValidationFieldError `json:"fields,omitempty"`
}

type ValidationFieldError struct {
	Message string `json:"message"`
	Field   string `json:"field"`
}
