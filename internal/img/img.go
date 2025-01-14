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
	err:=exec.Command("feh","--bg-scale","./bgImg.jpg").Run()
	if err!=nil{
		return fmt.Errorf("error during setting up bg Img \n %v",err)
	}
	return nil
}
