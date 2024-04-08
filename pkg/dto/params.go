package dto

type ResponseParams struct {
	StatusCode int
	Message    string
	Data       any
}

type FilterParams struct {
	Page   int
	Limit  int
	Offset int
	Search string
}