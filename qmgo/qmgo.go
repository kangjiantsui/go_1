package main

import (
	"context"
	"encoding/json"
	"github.com/qiniu/qmgo"
	rawOptions "github.com/qiniu/qmgo/options"
	"go.mongodb.org/mongo-driver/bson"
	"go_01/pb"
	"io/ioutil"
	"os"
	"path"
	"strconv"
)

const (
	uri1      = "mongodb://134.175.186.173"
	database1 = "basketball_develop"
	username  = "basketball"
	password  = "37fsmuo2oh"
	jsonFile  = "ranks1.json"

	//uri1 = "mongodb://localhost"
	//database1="basketball"
	//jsonFile="ranks1.json"

)

func main() {
	ctx := context.Background()

	qOpts := rawOptions.ClientOptions{}

	maxPoolSize := uint64(200)
	minPoolSize := uint64(1)
	mConfig := &qmgo.Config{
		Uri:         uri1,
		Database:    database1,
		MaxPoolSize: &maxPoolSize,
		MinPoolSize: &minPoolSize,
		Auth: &qmgo.Credential{
			Username: username,
			Password: password,
		},
	}

	client, err := qmgo.NewClient(ctx, mConfig, qOpts)
	if err != nil {
		panic(err)
	}

	database := client.Database(database1)

	c := context.Background()

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	file, err := os.Open(path.Join(wd, "ranks1.json"))
	if err != nil {
		panic(err)
	}
	all, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}

	rank := &ScoreMap{}
	err = json.Unmarshal(all, rank)
	if err != nil {
		panic(err)
	}

	users := map[uint64]map[string]int32{}
	for _, e := range rank.Ranks {
		for k, v := range e.Ranks {
			uid, err := strconv.Atoi(k)
			if err != nil {
				panic(err)
			}
			if _, ok := users[uint64(uid)]; !ok {
				users[uint64(uid)] = map[string]int32{}
			}
			users[uint64(uid)][e.BoardName] = int32(v*-1 + 1)
		}
	}
	if err != nil {
		panic(err)
	}

	for k, v := range users {
		var ret2 *BaseBoard
		err = database.Collection("user").Find(c, bson.M{"_id": 100375}).Select(bson.M{
			"leaderboard.valmap": 1,
		}).One(&ret2)

		if _, ok := v["season"]; ok {
			ret2.LeaderboardBoard.ValMap[pb.RANK_TYPE_CUP_RANK][pb.TIME_RANGE_SEASON_TIME] = v["season"]
		}
		if _, ok := v["week"]; ok {
			ret2.LeaderboardBoard.ValMap[pb.RANK_TYPE_CUP_RANK][pb.TIME_RANGE_WEEK_TIME] = v["week"]
		}
		err = database.Collection("user").UpdateOne(ctx, bson.M{"_id": k}, bson.M{"$set": bson.M{"leaderboard.valmap": ret2.LeaderboardBoard.ValMap}})
		if err != nil {
			panic(err)
		}
	}

}
