package auth

import (
	"context"
	"sync"
	"fmt"
	"github.com/renosyah/CustomAuth/Golang/model"
	"io"
)

type Auth_Server struct{
	Broadcast     chan CallbackData
	ClientStreams map[string]chan CallbackData
	streamsMtx    sync.RWMutex
}

func (a *Auth_Server) Authlogin(ctx context.Context, in *UserData) (*UserData,error){

	user := &model.User{
		Email:in.Email,
		Username: in.Username,
		Password:in.Password,
	}

	result,err := user.LoginUser(ctx,dbPool)
	if err != nil {
		fmt.Println(err.Error())
	}

	return &UserData{
		Id:result.Id,
		Name:result.Name,
		Email:result.Email,
		Username:result.Username,
		Password:result.Password,
	},nil
}

func (a *Auth_Server) WaitCallback(stream AuthService_WaitCallbackServer) error {
	tkn, _ := a.extractToken(stream.Context())

	go a.sendBroadcasts(stream,tkn)

	for{
		_,err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
	}
}