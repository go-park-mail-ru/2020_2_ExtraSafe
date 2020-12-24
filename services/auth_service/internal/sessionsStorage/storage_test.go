package sessionsStorage

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"github.com/tarantool/go-tarantool"
	"testing"
)

func TestStorage_CreateUserSession(t *testing.T) {
	tConn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{ User: "guest" })

	if err != nil {
		fmt.Println("Connection refused")
		return
	}
	defer tConn.Close()

	userID := int64(1)
	SID := "lalalalalalalalala"
	sessionStorage := NewStorage(tConn)

	err = sessionStorage.CreateUserSession(userID, SID)
	assert.Equal(t, nil, err)
}

func TestStorage_CreateUserSessionFail(t *testing.T) {
	tConn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{ User: "guest" })

	if err != nil {
		fmt.Println("Connection refused")
		return
	}
	defer tConn.Close()


	userID := int64(1)
	SID := "lalalalalalalalala"
	sessionStorage := NewStorage(tConn)

	err = sessionStorage.CreateUserSession(userID, SID)
	assert.Error(t, err)
}

func TestStorage_CheckUserSession(t *testing.T) {
	tConn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{ User: "guest" })

	if err != nil {
		fmt.Println("Connection refused")
		return
	}
	defer tConn.Close()


	userID := int64(1)
	SID := "lalalalalalalalala"
	sessionStorage := NewStorage(tConn)

	ID, err := sessionStorage.CheckUserSession(SID)
	assert.Equal(t, userID, ID)
}

func TestStorage_CheckUserSessionFail(t *testing.T) {
	tConn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{ User: "guest" })

	if err != nil {
		fmt.Println("Connection refused")
		return
	}
	defer tConn.Close()

	//userID := int64(1)
	SID := "lalalala"
	sessionStorage := NewStorage(tConn)

	_, err = sessionStorage.CheckUserSession(SID)
	assert.Error(t, err)
	//assert.Equal(t, userID, ID)
}

func TestStorage_DeleteUserSession(t *testing.T) {
	tConn, err := tarantool.Connect("tarantool:3301", tarantool.Opts{ User: "guest" })

	if err != nil {
		fmt.Println("Connection refused")
		return
	}
	defer tConn.Close()

	SID := "lalalalalalalalala"
	sessionStorage := NewStorage(tConn)

	err = sessionStorage.DeleteUserSession(SID)
	assert.Equal(t, nil, err)
}
