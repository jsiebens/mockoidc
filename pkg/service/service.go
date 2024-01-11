package service

import (
	"context"
	"github.com/bufbuild/connect-go"
	"github.com/jsiebens/mockoidc"
	v1 "github.com/jsiebens/mockoidc/pkg/gen/mockoidc/v1"
	"github.com/jsiebens/mockoidc/pkg/gen/mockoidc/v1/mockoidcv1connect"
)

func New(queue *mockoidc.UserQueue) mockoidcv1connect.MockOIDCServiceHandler {
	return &service{queue: queue}
}

type service struct {
	queue *mockoidc.UserQueue
}

func (s *service) PushUser(ctx context.Context, c *connect.Request[v1.PushUserRequest]) (*connect.Response[v1.PushUserResponse], error) {
	user := mockoidc.DefaultUser()
	user.Subject = c.Msg.Subject
	user.Email = c.Msg.Email
	user.PreferredUsername = c.Msg.PreferredUsername

	s.queue.Push(user)

	return connect.NewResponse(&v1.PushUserResponse{}), nil
}
