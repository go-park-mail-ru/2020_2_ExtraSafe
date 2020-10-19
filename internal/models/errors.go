package models

type ServeError struct {
	OriginalError error    `json:"-"`
	Codes         []string `json:"codes"`
}

func (c ServeError) Error() string {
	return c.OriginalError.Error()
}
