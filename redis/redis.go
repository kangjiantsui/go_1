package main

import (
	"fmt"
	"github.com/go-redis/redis/v7"
	"os"
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
}

func main() {
	f14()
}

func f14() {
	fileName := "敏感词"
	dstFile, err := os.Create(fileName)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer dstFile.Close()
	s := "hello world"
	dstFile.WriteString(s + "\n")
}

func f13() {
	result, _, err := client.SScan("basketball:sensitiveWords", 0, "*", -1).Result()
	if err != nil {
		return
	}
	for _, e := range result {
		fmt.Println(e)
	}
}

func f12() {
	result, err := client.ZRange("basketball:sensitiveWords", 0, -1).Result()
	if err != nil {
		panic(err)
	}
	for _, e := range result {
		fmt.Println(e)
	}
}

func f11() {
	client = redis.NewClient(&redis.Options{
		Addr:     "172.16.0.49:6379",
		Password: "B8ahtB8XTl16VssF",
		DB:       9,
	})

	client.ZCard("basketball:sensitiveWords")
}

func f10() {
	result, err := client.ZRangeWithScores("ztest", 0, 1).Result()
	if err != nil {
		return
	}
	for _, e := range result {
		fmt.Println(e.Score)
		fmt.Println(e.Member)
	}
}

func f9() {
	fmt.Println(client.ZScore("ztest", "kj").Result())
}

func f8() {
	p := client.Pipeline()
	p.ZAdd("ztest", &redis.Z{Member: "kj", Score: 8.5})
	p.ZCard("ztest")
	p.ZRank("ztest", "kj")
	cmders, err := p.Exec()
	if err != nil {
		return
	}
	for _, e := range cmders {
		intCmd := e.(*redis.IntCmd)
		fmt.Println(intCmd.Val())
	}

}

func f7() {
	client.ZAdd("ztest", &redis.Z{Score: 123456789123456789.54324234234, Member: "a"})
}

func f6() {
	client.ZAdd("ztest", &redis.Z{Score: 1, Member: "a"})
	client.ZAdd("ztest", &redis.Z{Score: 2, Member: "b"})
	client.ZAdd("ztest", &redis.Z{Score: 6, Member: "c"})
	client.ZAdd("ztest", &redis.Z{Score: 5, Member: "d"})
	client.ZAdd("ztest", &redis.Z{Score: 7, Member: "e"})
	client.ZAdd("ztest", &redis.Z{Score: 8, Member: "f"})
	client.ZAdd("ztest", &redis.Z{Score: 9, Member: "z"})

	fmt.Println(client.ZRange("ztest", 0, -1))
}

func f5() {
	client.ZAdd("ztest", &redis.Z{Score: 1, Member: "a"})
	client.ZAdd("ztest", &redis.Z{Score: 2, Member: "b"})
	client.ZAdd("ztest", &redis.Z{Score: 6, Member: "c"})
	client.ZAdd("ztest", &redis.Z{Score: 5, Member: "d"})
	client.ZAdd("ztest", &redis.Z{Score: 7, Member: "e"})
	client.ZAdd("ztest", &redis.Z{Score: 8, Member: "f"})
	client.ZAdd("ztest", &redis.Z{Score: 9, Member: "z"})

	fmt.Println(client.ZRank("ztest", "a").Val())
}

func f4() {
	client.ZAdd("ztest", &redis.Z{Score: 1, Member: "a"})
	client.ZAdd("ztest", &redis.Z{Score: 2, Member: "b"})
	client.ZAdd("ztest", &redis.Z{Score: 6, Member: "c"})
	client.ZAdd("ztest", &redis.Z{Score: 5, Member: "d"})
	client.ZAdd("ztest", &redis.Z{Score: 7, Member: "e"})
	client.ZAdd("ztest", &redis.Z{Score: 8, Member: "f"})
	client.ZAdd("ztest", &redis.Z{Score: 9, Member: "z"})
	fmt.Println(client.ZRange("ztest", 0, -1))

}

func f3() {
	keys := []string{}
	p := client.Pipeline()
	p.ZAdd("ztest", &redis.Z{Score: 1, Member: "a"})
	p.ZAdd("ztest", &redis.Z{Score: 2, Member: "b"})
	p.ZAdd("ztest", &redis.Z{Score: 6, Member: "c"})
	p.ZCard("ztest")
	p.ZCard("ztest")
	p.ZCard("ztest")
	keys = append(keys, "1")
	keys = append(keys, "1")
	keys = append(keys, "1")
	cmders, err := p.Exec()
	if err != nil {
		return
	}

	for i, e := range cmders {
		if len(keys) > i {
			continue
		}
		fmt.Println(e.(*redis.IntCmd).Val())
	}

}

func f2() {
	client.Del("ztest")
	client.ZAdd("ztest", &redis.Z{Score: 1, Member: "a"})
	client.ZAdd("ztest", &redis.Z{Score: 2, Member: "b"})
	client.ZAdd("ztest", &redis.Z{Score: 6, Member: "c"})
	client.ZAdd("ztest", &redis.Z{Score: 5, Member: "d"})
	client.ZAdd("ztest", &redis.Z{Score: 7, Member: "e"})
	client.ZAdd("ztest", &redis.Z{Score: 8, Member: "f"})
	client.ZAdd("ztest", &redis.Z{Score: 9, Member: "z"})

	client.ZRemRangeByRank("ztest", 0, client.ZCard("ztest").Val()-(3+1))

	fmt.Println(client.ZRevRange("ztest", 0, -1).Val())

}

func f1() {
	client.ZAdd("ztest", &redis.Z{Score: 1, Member: "a"})
	client.ZAdd("ztest", &redis.Z{Score: 2, Member: "b"})
	client.ZAdd("ztest", &redis.Z{Score: 5, Member: "c"})
	client.ZAdd("ztest", &redis.Z{Score: 6, Member: "d"})
	client.ZAdd("ztest", &redis.Z{Score: 7, Member: "e"})
	client.ZAdd("ztest", &redis.Z{Score: 8, Member: "f"})
}
