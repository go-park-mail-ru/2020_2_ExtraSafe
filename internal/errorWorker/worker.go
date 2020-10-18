package errorWorker

type ResponseError struct {
	OriginalError error    `json:"-"`
	Status        int      `json:"status"`
	Codes         []string `json:"codes"`
}

func (c ResponseError) Error() string {
	return c.OriginalError.Error()
}
