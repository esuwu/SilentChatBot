package usecase

import (
	"encoding/json"
	"fmt"
	"github.com/esuwu/SilentChatBot/models"
	"github.com/esuwu/SilentChatBot/tools/errors"
	"github.com/esuwu/SilentChatBot/tools/messages"
	"github.com/esuwu/SilentChatBot/tools/urls"
	"log"
	"net/http"
	"strconv"
)

func (u *useCase) MainHandler(body *models.MessageT) (*messages.MessageString, error) {
	response := messages.MessageString{}
	if body == nil {
		return nil, errors.ReceivedEmptyMessage
	}
	if !u.repository.IsUserExist(body.From.Id) {
		user := models.User{UserID: body.From.Id, UserChatID: body.Chat.Id, NickName: body.From.Username}
		u.repository.CreateUser(&user)
	}
	if !u.repository.IsUserBusy(body.From.Id){
		switch body.Text {
		case "/start" :
			response.Message = messages.Start
		case "/help" :
			response.Message = messages.Help
		case "/joinChat":
			response.Message = messages.Join
			freeGroups, err := u.repository.GetFreeGroups()
			if err != nil {
				fmt.Print(err.Error())
			}
			if len(freeGroups) != 0 {
				group := freeGroups[0]
				anotherUser, err := u.repository.GetAnotherUser(body.From.Id, group.GroupID)
				if err != nil {
					fmt.Print(err.Error())
				}
				err = u.repository.MakeUserBusy(body.From.Id)
				if err != nil {
					fmt.Print(err.Error())
				}
				err = u.repository.AddNextUserToGroup(body.From.Id, group)
				if err != nil {
					fmt.Print(err.Error())
				}
				err = u.NotifyAnotherFirstUser(anotherUser.UserID)
				if err != nil {
					fmt.Print(err.Error())
				}
				response.Message = messages.JoinFinally
				break
			}
			_, err = u.repository.CreateGroupAndAddUser(body.From.Id)
			if err != nil {
				fmt.Print(err.Error())
			}
			err = u.repository.MakeUserBusy(body.From.Id)
			if err != nil {
				fmt.Print(err.Error())
			}



		case "/leaveChat":
			response.Message = messages.LeaveError
		default:
			response.Message = messages.Default
		}
	} else {
		switch body.Text {
		case "/leaveChat":
			err := u.repository.LeaveChat(body.From.Id)
			if err != nil {
				log.Print(err.Error())
			}
			response.Message = messages.Leave

		default:
			if u.repository.IsUserInBusyChat(body.From.Id) {
				groupID := u.repository.GetGroupID(body.From.Id)
				anotherUser, err := u.repository.GetAnotherUser(body.From.Id, groupID)
				if err != nil {
					fmt.Print(err.Error())
				}
				u.SendMessage(anotherUser.UserChatID, body.Text)
				response.Message = messages.Sent
			} else {
				response.Message = messages.YourFriendIsNotComingYet
			}

		}
	}



	return &response, nil
}

func (u *useCase) NotifyAnotherFirstUser(chatID int) error{
	u.SendMessage(chatID, messages.JoinFinally)
	return nil
}

func (u *useCase) SendMessage(chatId int, text string) (models.SendMessageResponseT, error) {
	url := urls.BaseTelegramURL + "/bot" + u.token + "/" + urls.SendMessage
	url = url + "?chat_id=" + strconv.Itoa(chatId) + "&text=" + text
	response := getResponse(url)

	sendMessage := models.SendMessageResponseT{}
	err := json.Unmarshal(response, &sendMessage)
	if err != nil {
		return sendMessage, err
	}

	return sendMessage, nil
}

func getResponse(url string) []byte {
	response := make([]byte, 0)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println(err)

		return response
	}
	defer resp.Body.Close()
	for true {
		bs := make([]byte, 1024)
		n, err := resp.Body.Read(bs)
		response = append(response, bs[:n]...)

		if n == 0 || err != nil{
			break
		}
	}
	return response
}

