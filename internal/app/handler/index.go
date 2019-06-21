package handler

import (
	"context"
	"go-micro-starter/api/proto/app"
	"go-micro-starter/internal/app/service"
)

type AppHandler struct {
}

func (a *AppHandler) StartPing(ctx context.Context, req *app.Ping, res *app.Pong) error {
	r := service.Test(req.Content)
	res.PingContent = req.Content
	res.Content = r
	return nil
}
