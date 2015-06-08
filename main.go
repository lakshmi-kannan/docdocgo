package main

import (
	"fmt"
	"os"
	"runtime"

	"github.com/lakshmi-kannan/docdocgo/api/v1"
	"github.com/mailgun/vulcand/plugin/registry"
)

func main() {
	runtime.GOMAXPROCS(runtime.NumCPU())
	if err := v1.Run(registry.GetRegistry()); err != nil {
		fmt.Printf("Service exited with error: %s\n", err)
		os.Exit(255)
	} else {
		fmt.Println("Service exited gracefully")
	}
}
