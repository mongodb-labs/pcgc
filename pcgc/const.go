package pcgc

import "time"

// RequestTimeout the maximum duration for completing HTTP requests
const RequestTimeout = 10 * time.Second

// ResponseHeaderTimeout the maximum duration for receiving the response header
const ResponseHeaderTimeout = 10 * time.Second

// HTTPRequestResponseTimeout the maximum allowed duration to complete any HTTP requests and responses
const HTTPRequestResponseTimeout = 30 * time.Second

// ContentTypeJSON defines the JSON content type
const ContentTypeJSON = "application/json; charset=UTF-8"
