package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Println("Sistema operativo y  arquitectura", runtime.GOOS, runtime.GOARCH)
}
