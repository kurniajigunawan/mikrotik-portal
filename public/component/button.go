package component

import "github.com/kurniajigunawan/mikrotik-portal/public/style"

type ButtonSolid struct {
	Text            string
	Type            string
	BackgroundColor style.Color
}

type ButtonLink struct {
	Text      string
	LinkURL   string
	TextColor style.Color
	Icon      string
}

func (b ButtonLink) HasIcon() bool {
	return b.Icon != ""
}
