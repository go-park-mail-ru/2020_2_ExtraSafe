package tests

import (
	"2020_2_ExtraSafe/src"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestWriteResponse(t *testing.T) {
	response := new(src.ResponseUser)

	user := src.User{
		ID:       12,
		Email:    "sas@mail.ru",
		Username: "some",
		Password: "1234",
		FullName: "Petr",
		Links:    nil,
		Avatar:   "default",
	}

	response.WriteResponse(user)

	expectResponse := src.ResponseUser{
		Status:   200,
		Email:    "sas@mail.ru",
		Username: "some",
		FullName: "Petr",
		Avatar:   "default",
	}

	assert.Equal(t, *response, expectResponse)
}


func TestWriteResponseLinks(t *testing.T) {
	response := new(src.ResponseUserLinks)

	userLinks := src.UserLinks{
		Telegram:  "@telegram",
		Instagram: "@instagram",
		Github:    "github/bab",
		Bitbucket: "bitbucket/ket",
		Vk:        "vk.com",
		Facebook:  "facebook",
	}

	response.WriteResponse("some", userLinks, "default")

	expectResponse := src.ResponseUserLinks{
		Status:    200,
		Username:  "some",
		Telegram:  "@telegram",
		Instagram: "@instagram",
		Github:    "github/bab",
		Bitbucket: "bitbucket/ket",
		Vk:        "vk.com",
		Facebook:  "facebook",
		Avatar:    "default",
	}

	assert.Equal(t, *response, expectResponse)
}

