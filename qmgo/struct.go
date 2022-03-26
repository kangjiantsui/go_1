package main

import (
	"go_01/pb"
	"time"
)

type ScoreBoard struct {
	BoardName string
	Ranks     map[string]float64
}

type ScoreMap struct {
	Ranks []*ScoreBoard
}

type BaseBoard struct {
	LeaderboardBoard `bson:"leaderboard"`
}

type LeaderboardBoard struct {
	Season         int32                                    // 记录赛季
	Week           int32                                    // 记录周
	AreaCode       pb.AREA_CODE                             // 玩家地区代号
	ValMap         map[pb.RANK_TYPE]map[pb.TIME_RANGE]int32 // 增量map
	SwitchAreaTime *time.Time                               // 换区时间记录
}
