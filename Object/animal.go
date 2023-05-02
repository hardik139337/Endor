package object

import "reflect"

type Animal struct {
	Name    string `json:"name"`
	ID      string `json:"id"`
	Type    string `json:"type"`
	OwnerID string `json:"owner_id"`
}

func (p *Animal) GetKind() string {
	return reflect.TypeOf(p).String()
}
func (p *Animal) GetID() string {
	return p.ID
}
func (p *Animal) GetName() string {
	return p.Name
}
func (p *Animal) SetID(s string) {
	p.ID = s
}
func (p *Animal) SetName(s string) {
	p.Name = s
}
