package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/tencentyun/scf-go-lib/cloudfunction"
	"github.com/tencentyun/scf-go-lib/events"

	"github.com/fghwett/leetcodedaily/notify"
	"github.com/fghwett/leetcodedaily/task"
)

type DefineEvent struct {
	Notify struct {
		ServerChan struct {
			SecretKey string `json:"secretKey"`
		} `json:"serverChan"`
	} `json:"notify"`
}

func hello(ctx context.Context, event events.TimerEvent) {
	param := &DefineEvent{}
	err := json.Unmarshal([]byte(event.Message), param)
	if err != nil {
		fmt.Printf("解析定时事件附加参数错误 %s\n", err)
		return
	}

	t := task.New()
	t.Do()

	if err := notify.Send(param.Notify.ServerChan.SecretKey, "LeeCode每日一题", t.GetResult()); err != nil {
		fmt.Printf("通知发送失败 %s\n", err)
	}
}

func main() {
	cloudfunction.Start(hello)
}
