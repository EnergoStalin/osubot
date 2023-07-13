package services

import (
	"os"

	"github.com/EnergoStalin/osubot/pkg/gatari"
	"github.com/EnergoStalin/osubot/pkg/store"
	"github.com/go-resty/resty/v2"
)

var (
	GatariClient = gatari.GatariClient{
		Client: resty.New(),
	}
	KVStore = store.KvFsStore{Root: os.Getenv("KVFS_STORE_ROOT")}
)
