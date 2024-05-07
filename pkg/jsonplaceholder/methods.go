package jsonplaceholder

import (
	"encoding/json"
	"fmt"
)

func (p *Resource) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Resource) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d", p.Id, p.UserId)
}

func (p *Album) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Album) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d ===>>> Title: %s", p.Id, p.UserId, p.Title)
}

func (p *Comment) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Comment) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> Post Id: %d ===>>> Name: %s", p.Id, p.PostId, p.Name)
}

func (p *Photo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Photo) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> Album Id: %d ===>>> Title: %s", p.Id, p.AlbumId, p.Title)
}

func (p *Post) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Post) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d ===>>> Title: %s", p.Id, p.UserId, p.Title)
}

func (p *Todo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Todo) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d ===>>> Title: %s", p.Id, p.UserId, p.Title)
}

func (p *User) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *User) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> Name: %s ===>>> Email: %s", p.Id, p.Name, p.Email)
}
