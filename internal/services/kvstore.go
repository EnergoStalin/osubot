package services

import (
	"os"

	"github.com/EnergoStalin/osubot/pkg/store"
)

var (
	KVStore store.KvFsStore = store.KvFsStore{Root: os.Getenv("KVFS_STORE_ROOT")}
)
