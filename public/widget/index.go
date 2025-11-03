package widget

import (
	component "github.com/kurniajigunawan/mikrotik-portal/public/component"
)

type Form struct {
	ActionURL    string
	Method       string
	Input        []Input
	SubmitButton component.ButtonSolid
}

type Input struct {
	ID    string
	Label string
	Name  string
	Type  string
}

type Heading struct {
	Title    string
	Subtitle string
}

type MenuItem struct {
	LinkURL   string
	Title     string
	Subtitle  string
	Icon      string
	IconColor string
}
