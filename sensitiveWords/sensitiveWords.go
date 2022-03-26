package main

import (
	"fmt"
	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
	"github.com/go-redis/redis/v7"
	"time"
)

const (
	SyncCd           = time.Hour                   // 敏感词库同步CD
	ThesaurusKey     = "basketball:sensitiveWords" // redis中敏感词的key
	MaxOnlineSeconds = 86400 * 3
)

var (
	client *redis.Client
)

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     "192.168.199.92:6379",
		Password: "", // no password set
		DB:       3,  // use default DB
	})
}

var (
	thesaurus    = make([]string, 0) // 基础敏感词配置
	Thesaurus    []*ThesaurusDefine  // excel表中配置的敏感词
	memStore     *store.MemoryStore
	filterManage *filter.DirtyManager
	SyncTime     int64 // 最后同步敏感词库时间戳
)

type ThesaurusDefine struct {

	//mask:[1,2]
	Index int32 `mask:"[1 2]"`

	Content string `mask:"[1 2]"`
}

func main() {
	ret, err := checkSensitiveWords(" ")
	if err != nil {
		return
	}
	fmt.Println(ret)
}

func checkSensitiveWords(str string) (bool, error) {
	err := initThesaurus()
	if err != nil {
		return false, err
	}
	result, err := filterManage.Filter().Filter(str, '*', '@')
	if err != nil {
		return false, err
	}
	if result != nil {
		return false, nil
	}

	// 第三方检测
	/*	pass, err := checkViaApi(str)
		if err != nil {
			return false, err
		}
		if !pass {
			return false, nil
		}*/
	return true, nil
}

// 初始化敏感词库
func initThesaurus() error {
	var err error
	if client.SCard(ThesaurusKey).Val() == 0 {
		for _, e := range Thesaurus {
			thesaurus = append(thesaurus, e.Content)
		}
		client.SAdd(ThesaurusKey, thesaurus)
		memStore, err = store.NewMemoryStore(store.MemoryConfig{DataSource: thesaurus})
		filterManage = filter.NewDirtyManager(memStore)
		SyncTime = time.Now().Unix()
	}
	if len(thesaurus) == 0 {
		thesaurus = client.SMembers(ThesaurusKey).Val()
		memStore, err = store.NewMemoryStore(store.MemoryConfig{DataSource: thesaurus})
		filterManage = filter.NewDirtyManager(memStore)
		SyncTime = time.Now().Unix()
	}

	// 同步内存词库和redis词库
	if time.Duration(time.Now().Unix()-SyncTime) > SyncCd && client.SCard(ThesaurusKey).Val() != int64(len(thesaurus)) {
		thesaurus = client.SMembers(ThesaurusKey).Val()
		memStore, err = store.NewMemoryStore(store.MemoryConfig{DataSource: thesaurus})
		filterManage = filter.NewDirtyManager(memStore)
		SyncTime = time.Now().Unix()
	}
	return err
}
