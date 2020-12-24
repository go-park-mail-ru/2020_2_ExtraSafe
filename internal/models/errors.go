package models

var ServerError = "500"

type ServeError struct {
	OriginalError error    `json:"-"`
	Codes         []string `json:"codes"`
	Descriptions  []string `json:"descriptions"`
	MethodName    string   `json:"method_name"`
}

func (c ServeError) Error() string {
	return c.OriginalError.Error()
}

type MultiErrors struct {
	Codes         []string `json:"codes"`
	Descriptions  []string `json:"descriptions"`
}