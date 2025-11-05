package component

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/kurniajigunawan/mikrotik-portal/public/style"
)

type ToastType string

const (
	ToastTypeSuccess ToastType = "success"
	ToastTypeInfo    ToastType = "info"
	ToastTypeWarning ToastType = "warning"
	ToastTypeError   ToastType = "error"
)

type Toast struct {
	id string
	// generate id component
	Text string
	Type ToastType
}

func (e *Toast) GetID() string {
	if e.id != "" {
		return e.id
	}
	e.id = fmt.Sprintf("toast-%s", uuid.New().String())
	return e.id
}

func (e Toast) GetColor() style.Color {
	switch e.Type {
	case ToastTypeSuccess:
		return style.ColorGreen
	case ToastTypeInfo:
		return style.ColorBlue
	case ToastTypeWarning:
		return style.ColorYellow
	case ToastTypeError:
		return style.ColorRed
	}
	return style.ColorStone
}
