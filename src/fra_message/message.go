package message

import (
	"fmt"
	"github.com/kataras/iris"
	"io/ioutil"
	"net/http"
)

type Message struct {
	PostType    string `json:"post_type" validate:"required"`
	MessageType string `json:"message_type" validate:"required"`
	SubType     string `json:"sub_type" validate:"required"`
	MessageId   int64  `json:"message_id" validate:"required"`
	UserId      int64  `json:"user_id" validate:"required"`
	Message     string `json:"message" validate:"required"`
	RawMessage  string `json:"raw_message" validate:"required"`
	Font        int64  `json:"font" validate:"required"`
}

func InitHttpListener() {
	app := iris.Default()

	app.Post("/getMessage", func(ctx iris.Context) {
		var message Message
		if err := ctx.ReadJSON(&message); err != nil {
			fmt.Println(err)
		}

		fmt.Println(message.Message)
		fmt.Println(message.UserId)
		fmt.Println(message.MessageType)
		fmt.Println(message)
	})
	// listen and serve on http://0.0.0.0:8080.
	app.Run(iris.Addr(":8089"))
}

func SendPrivateMessage(userId string, message string) {
	resp, err := http.Get("http://127.0.0.1:5700/send_private_msg?user_id=" + userId + "&message=" + message)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}

func SendGroupMessage(groupId string, message string) {
	resp, err := http.Get("http://127.0.0.1:5700/send_group_msg?group_id=" + groupId + "&message=" + message)

	if err != nil {
		fmt.Println(err)
		return
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(body))
}
