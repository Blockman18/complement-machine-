package main

import (
	"fmt"
	"time"

	"github.com/pkg/browser"
)

func main() {
	fmt.Print("you are amazing!")
	time.Sleep(5 * time.Second)
	fmt.Print("\nyou are going to have a nice life!")
	time.Sleep(5 * time.Second)
	fmt.Print("\nim seeing that you are very smart!")
	time.Sleep(5 * time.Second)
	fmt.Print("\nyou are going to be great!")
	time.Sleep(5 * time.Second)
	browser.OpenURL("https://youtu.be/IM6oVTDxiHg")
	time.Sleep(5 * time.Second)

}
