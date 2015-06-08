package service

import (
	"time"
)

type Options struct {
	ApiPort      int
	ApiInterface string

	PidPath string
	Port    int

	Interface string
	CertPath  string

	EtcdNodes       listOptions
	EtcdKey         string
	EtcdCaFile      string
	EtcdCertFile    string
	EtcdKeyFile     string
	EtcdConsistency string

	Log string
	// LogSeverity severity

	ServerReadTimeout    time.Duration
	ServerWriteTimeout   time.Duration
	ServerMaxHeaderBytes int

	EndpointDialTimeout time.Duration
	EndpointReadTimeout time.Duration

	SealKey string

	StatsdAddr   string
	StatsdPrefix string
}
