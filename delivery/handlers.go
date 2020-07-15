package delivery

import useCase "github.com/esuwu/SilentChatBot/usecase"


type Handlers struct {
	usecase useCase.UseCase
}

func NewHandlers(ucases useCase.UseCase) *Handlers {
	return &Handlers{
		usecase: ucases,
	}
}