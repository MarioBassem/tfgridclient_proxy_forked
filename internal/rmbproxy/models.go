package rmbproxy

import (
	"bytes"
	"net/http"

	"github.com/threefoldtech/substrate-client"
)

// MessageIdentifier to get the specific result
type MessageIdentifier struct {
	Retqueue string `json:"retqueue" example:"xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx"`
}

// App is the main app objects
type App struct {
	resolver TwinExplorerResolver
}

// Flags for the App cmd command
type Flags struct {
	Debug        string
	Substrate    string
	Address      string
	Domain       string
	TLSEmail     string
	CA           string
	CertCacheDir string
}

// TwinExplorerResolver is Substrate resolver
type TwinExplorerResolver struct {
	client *substrate.Substrate
}

type twinClient struct {
	dstIP string
}

// TwinClient interface
type TwinClient interface {
	SubmitMessage(msg bytes.Buffer) (*http.Response, error)
	GetResult(msgID MessageIdentifier) (*http.Response, error)
}
