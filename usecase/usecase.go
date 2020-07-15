package usecase

import (
	"github.com/esuwu/SilentChatBot/models"
	"github.com/esuwu/SilentChatBot/repository"
	"github.com/esuwu/SilentChatBot/tools/messages"
)

type UseCase interface {
	MainHandler(body *models.MessageT) (*messages.MessageString, error)
	SendMessage(chatId int, text string) (models.SendMessageResponseT, error)
	NotifyAnotherFirstUser(chatID int) error

}

type useCase struct {
	repository repository.Repository
	token string
}

func NewUseCase(repo repository.Repository, myToken string) UseCase {
	return &useCase{
		repository: repo,
		token : myToken,
	}
}