package main

import (
	"context"
	"fmt"

	"github.com/tencentyun/scf-go-lib/cloudfunction"

	"leetcodedaily/notify"
	"leetcodedaily/task"
)

type DefineEvent struct {
	Notify struct {
		ServerChan struct {
			SecretKey string `json:"secretKey"`
		} `json:"serverChan"`
	} `json:"notify"`
}

func hello(ctx context.Context, event DefineEvent) {

	t := task.New()
	t.Do()

	if err := notify.Send(event.Notify.ServerChan.SecretKey, "LeeCode每日一题", t.GetResult()); err != nil {
		fmt.Printf("通知发送失败 %s\n", err)
	}
}

func main() {
	cloudfunction.Start(hello)
}
