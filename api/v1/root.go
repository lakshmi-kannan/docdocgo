package v1

import (
	"fmt"
	"net/http"
	"os"

	"github.com/mailgun/scroll"
	"github.com/mailgun/vulcand/plugin"
)

func Run(registry *plugin.Registry) error {
	options, err := ParseCommandLine()
	if err != nil {
		return fmt.Println("Failed parsing command line arguments: %s.", err)
	}
	service := NewService(options, registry)
	if err := service.Start(); err != nil {
		fmt.Println("Failed to start service: %v.", err)
		return fmt.Println("Service start failure: %s.", err)
	} else {
		fmt.Println("Service exited gracefully.")
	}
	return nil
}

type Service struct {
	options  Options
	registry *plugin.Registry
	apiApp   *scroll.App
	errorC   chan error
	sigC     chan os.Signal
}

func NewService(options Options, registry *plugin.Registry) *Service {
	return &Service{
		registry: registry,
		options:  options,
		errorC:   make(chan error),
		// Channel receiving signals has to be non blocking, otherwise the service can miss a signal.
		sigC: make(chan os.Signal, 1024),
	}
}

func (s *Service) Start() error {
	fmt.Println("Service started.")
	return nil
}
