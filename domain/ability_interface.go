package domain

type Ability interface {
	String() string
	Choice() []Action
}

type Action interface {
	Do()
}
