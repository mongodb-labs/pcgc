package pcgc

import "time"

// HTTPClientTimeouts allows users of this code to tweak all necessary timeouts
type HTTPClientTimeouts struct {
	DialTimeout           time.Duration
	ExpectContinueTimeout time.Duration
	IdleConnectionTimeout time.Duration
	ResponseHeaderTimeout time.Duration
	TLSHandshakeTimeout   time.Duration
	// GlobalTimeout the maximum allowed duration to complete a single HTTP request and response
	GlobalTimeout time.Duration
}

// InitTimeouts initializes the timeouts struct using default values
func InitTimeouts() *HTTPClientTimeouts {
	return &HTTPClientTimeouts{
		DialTimeout:           10 * time.Second,
		ExpectContinueTimeout: 2 * time.Second,
		IdleConnectionTimeout: 10 * time.Second,
		ResponseHeaderTimeout: 10 * time.Second,
		TLSHandshakeTimeout:   10 * time.Second,
		GlobalTimeout:         30 * time.Second,
	}
}
