package jsonplaceholder

import (
	"encoding/json"
)

func (p *Resource) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Album) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Comment) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Photo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Post) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Todo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *User) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}
