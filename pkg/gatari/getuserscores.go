package gatari

import (
	"context"
	"encoding/json"

	"github.com/EnergoStalin/osubot/pkg/util"
)

type GetUserRecentScoresConfig struct {
	Mode     int
	F        int
	PpFilter int
	L        int
	Id       int
	Special  int
}

func (g *GatariClient) GetUserRecentScores(cfg *GetUserRecentScoresConfig, ctx context.Context) (r *UserRecents, err error) {
	res, err := g.Client.R().SetQueryParams(util.SToMap(cfg)).SetContext(ctx).Get(baseUri + "/user/scores/recent")
	if err != nil {
		return
	}

	p := new(UserRecents)
	err = json.Unmarshal(res.Body(), p)
	if err != nil {
		return
	}

	r = p

	return
}

type UserRecents struct {
	Code   int64   `json:"code"`
	Count  int64   `json:"count"`
	Scores []Score `json:"scores"`
}

type Score struct {
	Accuracy      float64 `json:"accuracy"`
	Beatmap       Beatmap `json:"beatmap"`
	CommentsCount int64   `json:"comments_count"`
	Completed     int64   `json:"completed"`
	Count100      int64   `json:"count_100"`
	Count300      int64   `json:"count_300"`
	Count50       int64   `json:"count_50"`
	CountGekis    int64   `json:"count_gekis"`
	CountKatu     int64   `json:"count_katu"`
	CountMiss     int64   `json:"count_miss"`
	FullCombo     bool    `json:"full_combo"`
	ID            int64   `json:"id"`
	Isfav         bool    `json:"isfav"`
	MaxCombo      int64   `json:"max_combo"`
	Mods          int64   `json:"mods"`
	PlayMode      int64   `json:"play_mode"`
	Pp            float64 `json:"pp"`
	Ranking       string  `json:"ranking"`
	Score         int64   `json:"score"`
	Time          int64   `json:"time"`
	Verified      bool    `json:"verified"`
	Views         int64   `json:"views"`
}

type Beatmap struct {
	Ar                 float64 `json:"ar"`
	Artist             string  `json:"artist"`
	BeatmapID          int64   `json:"beatmap_id"`
	BeatmapMd5         string  `json:"beatmap_md5"`
	BeatmapsetID       int64   `json:"beatmapset_id"`
	BPM                int64   `json:"bpm"`
	Creator            string  `json:"creator"`
	Difficulty         float64 `json:"difficulty"`
	Fc                 int64   `json:"fc"`
	HitLength          int64   `json:"hit_length"`
	Od                 float64 `json:"od"`
	Ranked             int64   `json:"ranked"`
	RankedStatusFrozen int64   `json:"ranked_status_frozen"`
	SongName           string  `json:"song_name"`
	Title              string  `json:"title"`
	Version            string  `json:"version"`
}
