package util

import (
	"fmt"
	"os"
	"os/exec"
)

// CleanScreen Funcion que limpia la pantalla
func CleanScreen() {
	c := exec.Command("clear")
	c.Stdout = os.Stdout
	c.Run()
}

func HolaFatake() {
	fmt.Println("[+] Modulo FatakeGo cargado...")
}
