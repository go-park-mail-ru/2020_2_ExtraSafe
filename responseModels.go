package main

type responseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Username string `json:"nickname"`
	FullName string `json:"fullname"`
	Avatar string `json:"avatar"`
}

type responseUserLinks struct {
	Status    int    `json:"status"`
	Username  string `json:"username"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
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
	response.Username = user.Username
	response.FullName = user.FullName
}

func (response *responseUserLinks) WriteResponse(username string, links UserLinks) {
	response.Status = 200
	response.Username = username
	response.Telegram = links.Telegram
	response.Instagram = links.Instagram
	response.Github = links.Github
	response.Bitbucket = links.Bitbucket
	response.Vk = links.Vk
	response.Facebook = links.Facebook
}
