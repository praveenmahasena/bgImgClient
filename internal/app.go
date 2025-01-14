package internal

import (
	"github.com/praveenmahasena/bgclient/internal/api"
	"github.com/praveenmahasena/bgclient/internal/img"
)

func Run() error {
	d := api.New("tcp", ":42069")
	b, err := d.Dial()
	if err != nil {
		return err
	}
	fileErr:=img.WriteImg(b)
	if fileErr!=nil{
		return fileErr
	}
	return img.Set()
}
