package object

type Object interface {
	GetKind() string
	GetID() string
	SetID(string)
	SetName(string)
	GetName() string
}
