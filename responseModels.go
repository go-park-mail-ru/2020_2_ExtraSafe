package main

type responseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Nickname string `json:"username"`
	FullName string `json:"fullName"`
}

type responseUserLinks struct {
	Status    int    `json:"status"`
	Nickname  string `json:"username"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vk"`
	Facebook  string `json:"facebook"`
}

type responseError struct {
	OriginalError error      `json:"-"`
	Status        int        `json:"status"`
	Messages      []Messages `json:"messages"`
}

type Messages struct {
	ErrorName string `json:"errorName"`
	Message   string `json:"message"`
}

func (c responseError) Error() string {
	return c.OriginalError.Error()
}

/*func (c *responseError) WriteResponse(message string) {
	c.Status = 500
	c.Message = message
}*/

func (response *responseUser) WriteResponse(user User) {
	response.Status = 200
	response.Email = user.Email
	response.Nickname = user.Nickname
	response.FullName = user.FullName
}

func (response *responseUserLinks) WriteResponse(nickname string, links UserLinks) {
	response.Status = 200
	response.Nickname = nickname
	response.Telegram = links.Telegram
	response.Instagram = links.Instagram
	response.Github = links.Github
	response.Bitbucket = links.Bitbucket
	response.Vk = links.Bitbucket
	response.Facebook = links.Facebook
}
