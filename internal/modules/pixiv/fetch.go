package pixiv

import (
	"encoding/json"
	"fmt"

	"github.com/EnergoStalin/osubot/internal/services"
	"github.com/everpcpc/pixiv"
)

func fetchAll(app *pixiv.AppPixivAPI, id uint64) ([]pixiv.Illust, error) {
	var (
		err     error
		next    int            = 0
		illusts                = []pixiv.Illust{}
		s       []pixiv.Illust = []pixiv.Illust{}
	)

	PIXIV := services.KVStore.Get("PIXIV" + fmt.Sprint(id))
	if PIXIV == "" {
		for {
			s, next, err = app.UserBookmarksIllust(id, "private", next, "")
			illusts = append(illusts, s...)
			if next == 0 {
				break
			}
		}
		text, _ := json.Marshal(illusts)
		services.KVStore.Put("PIXIV"+fmt.Sprint(id), string(text))
	} else {
		json.Unmarshal([]byte(PIXIV), &illusts)
	}

	return illusts, err
}
