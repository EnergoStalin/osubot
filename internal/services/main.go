package services

import (
	"os"
	"time"

	"github.com/EnergoStalin/osubot/pkg/gatari"
	"github.com/EnergoStalin/osubot/pkg/store"
	"github.com/EnergoStalin/pixiv"
	"github.com/go-resty/resty/v2"
)

var (
	GatariClient = gatari.GatariClient{
		Client: resty.New(),
	}
	KVStore     = store.KvFsStore{Root: os.Getenv("KVFS_STORE_ROOT")}
	PixivApp    = pixiv.NewApp()
	PixivAcc, _ = pixiv.LoadAuth(os.Getenv("PIXIV_TOKEN"), os.Getenv("PIXIV_REFRESH_TOKEN"), time.Now().Add(3600))
)
