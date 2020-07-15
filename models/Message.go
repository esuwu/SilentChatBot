package models

type WebHookReqBody struct {
	Message MessageT `json:"message"`
}

type SendMessageResponseT struct {
	Ok bool `json:"ok"`
	Result MessageT `json:"result"`
}

type MessageT struct {
	MessageId int `json:"message_id"`
	From FromT `json:"from"`
	Chat ChatT `json:"chat"`
	Date int `json:"date"`
	Text string `json:"text"`
}


type FromT struct {
	Id int `json:"id"`
	IsBot bool `json:"is_bot"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Language string `json:"language_code"`
}

type ChatT struct {
	Id int `json:"id"`
	FirstName string `json:"first_name"`
	LastName string `json:"last_name"`
	Username string `json:"username"`
	Type string `json:"type"`
}