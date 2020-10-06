package main

type responseUser struct {
	Status   int    `json:"status"`
	Email    string `json:"email"`
	Nickname string `json:"nickname"`
	FullName string `json:"fullname"`
}

type responseUserLinks struct {
	Status    int    `json:"status"`
	Nickname  string `json:"nickname"`
	Telegram  string `json:"telegram"`
	Instagram string `json:"instagram"`
	Github    string `json:"github"`
	Bitbucket string `json:"bitbucket"`
	Vk        string `json:"vk"`
	Facebook  string `json:"facebook"`
}

type responseError struct {
	Status  int    `json:"status"`
	Message string `json:"message"`
}

func (response *responseError) WriteResponse(message string) {
	response.Status = 500
	response.Message = message
}

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
