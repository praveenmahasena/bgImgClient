package img

import (
	"fmt"
	"os"
	"os/exec"
)

func WriteImg(f []byte)error{
	err:=os.WriteFile("./bgImg.jpg",f,0666)
	if err!=nil{
		return fmt.Errorf("error during file write io \n %v",err)
	}
	return nil
}

func Set()error{
	err:=exec.Command("feh","--bg-max","./bgImg.jpg").Run()
	//err:=exec.Command("feh","--bg-center","--auto-zoom","./bgImg.jpg","-b","-X").Run() // original one
	if err!=nil{
		return fmt.Errorf("error during setting up bg Img \n %v",err)
	}
	return nil
}
