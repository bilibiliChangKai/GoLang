package main

import (
	"fmt"
	//"os"
	"os/exec"
)

func main() {
	cmd := exec.Command("lp", "-dlp1")
	out, err := cmd.CombinedOutput()
	if err != nil {
		//fmt.Println(err)
	}
	fmt.Println(string(out))
}
