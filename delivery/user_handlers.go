package delivery

import (
	"encoding/json"
	"fmt"
	"github.com/esuwu/SilentChatBot/models"
	"github.com/valyala/fasthttp"
)



func (handlers *Handlers) MainHandler(ctx *fasthttp.RequestCtx) {
	body := models.WebHookReqBody{}
	err := json.Unmarshal(ctx.PostBody(), &body)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Print(body.Message.Text)
	response, err := handlers.usecase.MainHandler(&body.Message)
	if err != nil {
		handlers.usecase.SendMessage(body.Message.Chat.Id, "Моя внутренняя критическая ошибка...")
	}
	handlers.usecase.SendMessage(body.Message.Chat.Id, response.Message)
	fmt.Println("reply sent")
}




