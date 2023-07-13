package gatari

import "github.com/go-resty/resty/v2"

const (
	baseUri = "https://api.gatari.pw"
)

type GatariClient struct {
	Client *resty.Client
}
