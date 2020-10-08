package src

type ResponseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Username string `json:"username"`
	FullName string `json:"fullName"`
	Avatar   string `json:"avatar"`
}

type ResponseUserLinks struct {
	Status    int    `json:"status"`
	Username  string `json:"username"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vkontakte"`
	Facebook  string `json:"facebook"`
	Avatar    string `json:"avatar"`
}

type ResponseError struct {
	OriginalError error      `json:"-"`
	Status        int        `json:"status"`
	Messages      []Messages `json:"messages"`
}

type Messages struct {
	ErrorName string `json:"errorName"`
	Message   string `json:"message"`
}

func (c ResponseError) Error() string {
	return c.OriginalError.Error()
}

func (response *ResponseUser) WriteResponse(user User) {
	response.Status = 200
	response.Email = user.Email
	response.Username = user.Username
	response.FullName = user.FullName
	response.Avatar = user.Avatar
}

func (response *ResponseUserLinks) WriteResponse(username string, links UserLinks, avatar string) {
	response.Status = 200
	response.Username = username
	response.Telegram = links.Telegram
	response.Instagram = links.Instagram
	response.Github = links.Github
	response.Bitbucket = links.Bitbucket
	response.Vk = links.Vk
	response.Facebook = links.Facebook
	response.Avatar = avatar
}
