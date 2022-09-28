package utils

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/k0kubun/pp/v3"
)

type Type int64

const (
	PrintTypeUnknown Type = iota  // 0
	PrintTypeNormal
	PrintTypeWarning
)

func Print(message interface{}, printType Type)  {
	switch printType {
	case PrintTypeWarning:
		color.Yellow("WARNING:")
	default:

	}
	pp.Println(message)
	fmt.Println("----------------------")
}
