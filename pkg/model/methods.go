package model

import (
	"encoding/json"
	"fmt"
)

func (p *Album) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Album) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d ===>>> Title: %s", p.Id, p.UserId, p.Title)
}

func (p *Album) IsZeroId() bool {
	return p.Id == 0
}

func (p *Album) SetId(n int) {
	p.Id = n
}

func (p *Comment) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Comment) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> Post Id: %d ===>>> Name: %s", p.Id, p.PostId, p.Name)
}

func (p *Comment) IsZeroId() bool {
	return p.Id == 0
}

func (p *Comment) SetId(n int) {
	p.Id = n
}

func (p *Photo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Photo) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> Album Id: %d ===>>> Title: %s", p.Id, p.AlbumId, p.Title)
}

func (p *Photo) IsZeroId() bool {
	return p.Id == 0
}

func (p *Photo) SetId(n int) {
	p.Id = n
}

func (p *Post) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Post) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d ===>>> Title: %s", p.Id, p.UserId, p.Title)
}

func (p *Post) IsZeroId() bool {
	return p.Id == 0
}

func (p *Post) SetId(n int) {
	p.Id = n
}

func (p *Todo) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *Todo) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> User Id: %d ===>>> Title: %s", p.Id, p.UserId, p.Title)
}

func (p *Todo) IsZeroId() bool {
	return p.Id == 0
}

func (p *Todo) SetId(n int) {
	p.Id = n
}

func (p *User) Show() []byte {
	bodyResult, _ := json.Marshal(p)
	return bodyResult
}

func (p *User) Echo() string {
	return fmt.Sprintf("Id: %d ===>>> Name: %s ===>>> Email: %s", p.Id, p.Name, p.Email)
}

func (p *User) IsZeroId() bool {
	return p.Id == 0
}

func (p *User) SetId(n int) {
	p.Id = n
}
