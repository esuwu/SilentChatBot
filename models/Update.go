package models

type UpdateT struct {
	Ok bool `json:"ok"`
	Result []ResultT `json:"result"`
}

type ResultT struct {
	UpdateId int `json:"update_id"`
	Message MessageT `json:"message"`
}


