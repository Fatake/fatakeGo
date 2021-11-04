package util

import (
	"fmt"
	"os"
	"os/exec"
	"runtime"
)

var clear map[string]func()

func init() {
	clear = make(map[string]func())
	clear["linux"] = func() {
		cmd := exec.Command("clear")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
	clear["windows"] = func() {
		cmd := exec.Command("cmd", "/c", "cls")
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}

// CleanScreen Funcion que limpia la pantalla
func CleanScreen() {
	value, ok := clear[runtime.GOOS]
	if ok {
		value()
	} else {
		panic("[!] No puedo limpiar la pantalla en: " + runtime.GOOS)
	}
}

func HolaFatake() {
	fmt.Println("[+] Modulo FatakeGo cargado...")
}
