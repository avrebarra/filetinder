package sdk

import (
	"fmt"

	"github.com/imroc/req"
	"github.com/shrotavre/filetinder/internal/config"
	"github.com/shrotavre/filetinder/internal/filetinder"
)

var basicHeader req.Header = req.Header{"Accept": "application/json"}

// FileTinderSDK is entity to communicate with FileTinder server
type FileTinderSDK struct {
	BaseURI string
}

// New create new sdk instance
func New() *FileTinderSDK {
	sdk := FileTinderSDK{
		BaseURI: fmt.Sprintf("http://localhost:%d", config.DefaultPort),
	}

	return &sdk
}

// NewTarget post new target to active server
func (sdk *FileTinderSDK) NewTarget(p NewTargetParams) error {
	targetURL := sdk.BaseURI + "/api/targets"
	payload := req.Param{
		"url": p.URL,
	}

	_, err := req.Post(targetURL, basicHeader, req.BodyJSON(payload))
	if err != nil {
		return err
	}

	return nil
}

// ListTarget list all target existed in active server
func (sdk *FileTinderSDK) ListTarget() (targets []*filetinder.Target, err error) {
	url := fmt.Sprintf("http://localhost:%d/api/targets", config.DefaultPort)
	r, err := req.Get(url, req.Header{"Accept": "application/json"})
	if err != nil {
		return nil, err
	}

	var ts []*filetinder.Target
	r.ToJSON(&ts)

	return ts, nil
}

// DelTarget delete meant target
func (sdk *FileTinderSDK) DelTarget(id int64) error {
	url := fmt.Sprintf("http://localhost:%d/api/targets", config.DefaultPort)
	_, err := req.Delete(url, req.Header{"Accept": "application/json"})
	if err != nil {
		return err
	}

	return nil
}

// KillServer kill currently running server via HTTP call
func (sdk *FileTinderSDK) KillServer() error {
	url := fmt.Sprintf("http://localhost:%d/api/funcs/stop-server", config.DefaultPort)
	_, err := req.Post(url)
	if err != nil {
		return err
	}

	return nil
}
