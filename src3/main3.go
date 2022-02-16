package main

import (
	"context"
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"fmt"
	filter "github.com/antlinker/go-dirtyfilter"
	"github.com/antlinker/go-dirtyfilter/store"
	"github.com/go-redis/redis/v7"
	"github.com/satori/go.uuid"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	gourl "net/url"
	"os"
	"reflect"
	"strings"
	"time"
	"unsafe"
)

func main() {
	f33()
}

func f33() {
	var t1 time.Time
	fmt.Println(time.Now().After(t1))
}

func f32() {
	t1 := time.Date(1, 1, 1, 1, 1, 1, 1, time.Local)
	t2 := time.Date(1, 1, 1, 1, 1, 1, 1, time.Local)
	fmt.Println(t1 == t2)
}

func f31() {
	fmt.Println(time.Now().AddDate(0, 0, int(time.Monday-time.Now().Weekday())).Day())
}

func f30() {
	fmt.Println(int(time.Now().Month()), time.Now().Year(), time.Now().Day())
}

func f29() {
	fmt.Println(time.Now().Add(time.Hour * -5).String())
}

func f28() {
	fmt.Println(time.Now().String()[0:len("yyyy-MM-DD")])
}

func f27() {
	var a1 []int32
	//a1 = append(a1, 1)
	fmt.Println(a1)
}

func f26() {
	fmt.Printf("补偿玩家:%d,前球员%d,补偿%d,后球员%d", 12, []int{1, 2, 3}, []int{5, 6}, []int{1, 2, 3, 5, 6})
}

func f25() {
	dir, err := os.Getwd()
	if err != nil {
		return
	}
	fmt.Println(dir)
	//fileName := "C:\Users\SAGI-PC-00015\Desktop\temp\事件分析_全量数据_20210708_20210722.csv"
	fileName1 := dir + "\\事件分析_全量数据_20210708_20210722.csv"
	fs, err := os.Open(fileName1)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	count := 0
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		count++
		if count == 1 {
			continue
		}
		fmt.Println(row)
		if count > 20 {
			break
		}
	}
}

type TimePorint struct {
	StartTime int64 `bson:"startTime"` //开始时间
	EndTime   int64 `bson:"endTime"`   //结束时间
}

func f24() {
	var (
		client     *mongo.Client
		err        error
		db         *mongo.Database
		collection *mongo.Collection
	)
	//1.建立连接
	if client, err = mongo.Connect(context.TODO(), options.Client().ApplyURI("mongodb://localhost:27017").SetConnectTimeout(5*time.Second)); err != nil {
		fmt.Print(err)
		return
	}
	//2.选择数据库 my_db
	db = client.Database("basketball")

	//3.选择表 my_collection
	collection = db.Collection("my_collection")
	collection = collection

	tp := &TimePorint{StartTime: 1, EndTime: 2}
	if _, err := collection.InsertOne(context.TODO(), tp); err != nil {
		fmt.Print(err)
		return
	}
}

func f23() {
	clientOptions := options.Client().ApplyURI("mongodb://localhost:27017")
	// 连接到MongoDB
	mgoCli, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	// 检查连接
	err = mgoCli.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal(err)
	}

}

func f22() {
	//file, err := excelize.OpenFile("C:\Users\SAGI-PC-00015\Desktop\temp\事件分析_全量数据_20210708_20210722.csv")
	//准备读取文件
	fileName := "C:\\Users\\SAGI-PC-00015\\Desktop\temp\\事件分析_全量数据_20210708_20210722.csv"
	fs, err := os.Open(fileName)
	if err != nil {
		log.Fatalf("can not open the file, err is %+v", err)
	}
	defer fs.Close()
	r := csv.NewReader(fs)
	//针对大文件，一行一行的读取文件
	count := 0
	for {
		row, err := r.Read()
		if err != nil && err != io.EOF {
			log.Fatalf("can not read, err is %+v", err)
		}
		if err == io.EOF {
			break
		}
		fmt.Println(row)
		count++
		if count > 20 {
			break
		}
	}
}

func f21() {
	a1 := []int{1, 2, 3}
	a2 := []int{1, 2}
	for _, e := range a1 {
		for _, f := range a2 {
			if e == f {

			}
		}
	}
}

func f20() {
	s := "1,2,3;4,5,6"
	for _, e := range strings.Split(s, ";") {
		fmt.Println(e)
		for _, f := range strings.Split(e, ",") {
			fmt.Println(f)
		}
	}
}

func f19() {
	t1, err := time.ParseInLocation("2006-01-02", "2021-07-22", time.Local)
	if err != nil {
		return
	}
	t2, err := time.ParseInLocation("2006-01-02", "2021-07-21", time.Local)
	if err != nil {
		return
	}
	fmt.Println(t1.Sub(t2).Hours() / 24)
	fmt.Println(t1.YearDay())
	fmt.Println(t1.Add(time.Hour * -1).YearDay())
}

func f18() {
	m1 := map[int][]int{}
	m1[1] = append(m1[1], 1)
	fmt.Println(m1)
}

func f17() {
	a1 := []int{1, 2, 3}
	fmt.Println(json.Marshal(a1))
}

func f16() {
	fmt.Println(time.Now().Hour())

}

func f15() {
	m1 := map[int][]int{1: {1, 2}}
	for _, e := range m1[1] {
		fmt.Println(e)
	}
}

func f14() {
	m1 := map[int]int{}
	m1[1]++
	fmt.Println(m1)
}

type S1 struct {
}

type S2 struct {
}

func f13() {
	a1 := make([]interface{}, 0)
	a1 = append(a1, &S1{})
	a2, ok := a1[0].(*S1)
	fmt.Println(ok)
	fmt.Println(a2)
}

func f12() {
	timeFormat := "2006-01-02"
	t1 := time.Now().In(time.Local)
	fmt.Println(t1.Unix())
	fmt.Println(t1.Unix() / 86400)
	t2, _ := time.ParseInLocation(timeFormat, "2021-07-13", time.Local)
	fmt.Println(t2.Unix())
	fmt.Println(t2.Unix() / 86400)
	t3, _ := time.ParseInLocation(timeFormat, "2021-07-13", time.UTC)
	fmt.Println(t3.Unix() - t2.Unix())
}

type T1 struct {
	Name bool
}

func f11() {
	t1 := &T1{}
	fmt.Println(t1.Name)
}

func f10() {
	array := []int{1, 2, 3, 4}
	for i, e := range array {
		if e == 2 {
			array[i] = 3
		}
	}
	fmt.Println(array)
}

func f9() {
	for i := 0; i < 10; i++ {
		fmt.Println(uuid.NewV4().String())
	}
}

func f8() {
	fmt.Println("err:runtime error: invalid memory address or nil pointer dereference stack:goroutine 343 [running]:\nruntime/debug.Stack(0xc000db4e40, 0x13af140, 0x2087310)\n\t/usr/loca    l/go/src/runtime/debug/stack.go:24 +0x9f\ngameserver/game/pkg/components/handler.(*CompHandlerGRPC).HandleHttpRequest.func1(0xc000db5850, 0xc000db56bc, 0xc036b53e3abfd957, 0x5152a7abe3, 0x20acca0, 0xc000d1def0, 0x19398, 0xc000d5    a7d0, 0xc000db5770)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/components/handler/handler.go:128 +0x3c7\npanic(0x13af140, 0x2087310)\n\t/usr/local/go/src/runtime/panic.go:965 +0x1b9\ngameserver/game/pkg/gamesubmap.    (*BugFix).Init(0xc000f4cc90, 0x19398)\n\t<autogenerated>:1 +0x32\ngameserver/game/pkg/submap.(*SubMap).initSubDataArr(0xc000f60460, 0x184a718, 0xc000d72b40, 0xc000f7e180, 0x1, 0x1, 0x1)\n\t/mnt/d/workspace/go/game-server2/game-s    erver/game/pkg/submap/submap.go:262 +0x2cb\ngameserver/game/pkg/submap.(*SubMap).GetData(0xc000f60460, 0x184a718, 0xc000d72b40, 0xc000db5410, 0x1, 0x1, 0x12745c4)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/submap/s    ubmap.go:150 +0x225\ngameserver/game/pkg/submap.(*SubMap).GetOne(...)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/submap/submap.go:108\ngameserver/game/pkg/service/bugfix/box_unsaved_bug.(*DataManager).ReimburseActo    rAndClothes(0xc00052a2a0, 0x184a718, 0xc000d72b40, 0x19398, 0xc000ec79c0, 0x0)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/service/bugfix/box_unsaved_bug/data_manager.go:40 +0x13b\ngameserver/game/pkg/service/main_i    nterface.(*DataManager).GetMainInterfaceMsg(0xc00063c340, 0x184a718, 0xc000d72b40, 0x19398, 0xc000d4de00, 0x1842900, 0x7f6fd15d4c60)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/service/main_interface/data_manager.go    :102 +0x53b\ngameserver/game/pkg/service/main_interface.GetMainInterfaceMsg(...)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/service/main_interface/mainInterface.go:9\ngameserver/game/pkg/api.(*GameService).GetMainI    nterfaceMsg(0x20de2e8, 0x184a718, 0xc000d72b40, 0xc0000778e0, 0x20de2e8, 0x0, 0xc000095658)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/api/mainInterface.go:14 +0xb3\ngameserver/common/pb._Game_GetMainInterfaceMsg_H    andler(0x1566120, 0x20de2e8, 0x184a718, 0xc000d72b40, 0xc000e02588, 0x0, 0x184a718, 0xc000d72b40, 0x0, 0x0)\n\t/mnt/d/workspace/go/game-server2/game-server/common/pb/service_game_grpc.pb.go:2127 +0x217\ngameserver/game/pkg/compo    nents/handler.(*CompHandlerGRPC).HandleHttpRequest(0xc000a33060, 0x184a718, 0xc000d72930, 0xc000d1def0, 0x0, 0x0, 0x0)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/components/handler/handler.go:142 +0x5b0\ngameserver    /game/pkg/cluster/logic.(*ModuleLogic).HandleHttpRequest(0xc000b1eb40, 0x184a718, 0xc000d72930, 0xc000d1def0, 0xc000b1eb40, 0x184a718, 0xc000d72930)\n\t/mnt/d/workspace/go/game-server2/game-server/game/pkg/cluster/logic/module.g    o:120 +0x52\ngameserver/common/pb._Rpc_HandleHttpRequest_Handler.func1(0x184a718, 0xc000d72930, 0x14baca0, 0xc000d1def0, 0x184a718, 0xc000d72930, 0x185b820, 0xc000d2b680)\n\t/mnt/d/workspace/go/game-server2/game-server/common/pb    /cluster_grpc.pb.go:504 +0x89\ngithub.com/opentracing-contrib/go-grpc.OpenTracingServerInterceptor.func1(0x184a718, 0xc000d72930, 0x14baca0, 0xc000d1def0, 0xc000077800, 0xc000e024f8, 0x0, 0x0, 0x0, 0x0)\n\t/mnt/d/workspace/go/ga    me-server2/game-server/vendor/github.com/opentracing-contrib/go-grpc/server.go:57 +0x302\ngameserver/common/pb._Rpc_HandleHttpRequest_Handler(0x15562e0, 0xc000b1eb40, 0x184a718, 0xc000d728a0, 0xc000d45140, 0xc000a330a0, 0x184a71    8, 0xc000d728a0, 0xc000e06000, 0xd0)\n\t/mnt/d/workspace/go/game-server2/game-server/common/pb/cluster_grpc.pb.go:506 +0x150\ngoogle.golang.org/grpc.(*Server).processUnaryRPC(0xc0006c8e00, 0x1856c58, 0xc0004ea480, 0xc000b33300,     0xc000a6aa20, 0x209aad0, 0x0, 0x0, 0x0)\n\t/mnt/d/workspace/go/game-server2/game-server/vendor/google.golang.org/grpc/server.go:1217 +0x52b\ngoogle.golang.org/grpc.(*Server).handleStream(0xc0006c8e00, 0x1856c58, 0xc0004ea480, 0x    c000b33300, 0x0)\n\t/mnt/d/workspace/go/game-server2/game-server/vendor/google.golang.org/grpc/server.go:1540 +0xd0c\ngoogle.golang.org/grpc.(*Server).serveStreams.func1.2(0xc0005bf0d0, 0xc0006c8e00, 0x1856c58, 0xc0004ea480, 0xc    000b33300)\n\t/mnt/d/workspace/go/game-server2/game-server/vendor/google.golang.org/grpc/server.go:878 +0xab\ncreated by google.golang.org/grpc.(*Server).serveStreams.func1\n\t/mnt/d/workspace/go/game-server2/game-server/vendor/    google.golang.org/grpc/server.go:876 +0x1fd\n")
}

type ac struct {
	code string `json:"code"`
	wid  string
	ic   int
}

type acs struct {
	ac []ac
}

func slice(s string) (b []byte) {
	pbytes := (*reflect.SliceHeader)(unsafe.Pointer(&b))
	pstring := (*reflect.StringHeader)(unsafe.Pointer(&s))
	pbytes.Data = pstring.Data
	pbytes.Len = pstring.Len
	pbytes.Cap = pstring.Len
	return
}

func f7() {
	s := "  [{\"003W1XZMUM\", \"wid\": \"\", \"ic\": 0}, {\"003W20HZZ6\", \"wid\": \"\", \"ic\": 0}, {\"003W21Q513\", \"wid\": \"\", \"ic\": 0}, {\"003W231S2Z\", \"wid\": \"\", \"ic\": 0}, {\"003W25HA4N\", \"wid\": \"\", \"ic\": 0}, {\"003W26CZ21\", \"wid\": \"\", \"ic\": 0}, {\"003W27L07Z\", \"wid\": \"\", \"ic\": 0}, {\"003W297RTV\", \"wid\": \"\", \"ic\": 0}, {\"003W2AK6AZ\", \"wid\": \"\", \"ic\": 0}, {\"003W2BBHCQ\", \"wid\": \"\", \"ic\": 0}, {\"003W2CRB0E\", \"wid\": \"\", \"ic\": 0}, {\"003W2DBVAF\", \"wid\": \"\", \"ic\": 0}, {\"003W2FNAR7\", \"wid\": \"\", \"ic\": 0}, {\"003W2G6VTC\", \"wid\": \"\", \"ic\": 0}, {\"003W2HXMXG\", \"wid\": \"\", \"ic\": 0}, {\"003W2YQ3ZP\", \"wid\": \"\", \"ic\": 0}, {\"003W2JWHRC\", \"wid\": \"\", \"ic\": 0}, {\"003W2LQEXD\", \"wid\": \"\", \"ic\": 0}, {\"003W2NBPAB\", \"wid\": \"\", \"ic\": 0}, {\"003W2ZG48D\", \"wid\": \"\", \"ic\": 0}, {\"003W2ZURAR\", \"wid\": \"\", \"ic\": 0}, {\"003W2QSM0Y\", \"wid\": \"\", \"ic\": 0}, {\"003W2SNB08\", \"wid\": \"\", \"ic\": 0}, {\"003W2TN4Q5\", \"wid\": \"\", \"ic\": 0}, {\"003W2UCSVU\", \"wid\": \"\", \"ic\": 0}, {\"003W2W8HFJ\", \"wid\": \"\", \"ic\": 0}, {\"003W2WPTL2\", \"wid\": \"\", \"ic\": 0}, {\"003W30AF0K\", \"wid\": \"\", \"ic\": 0}, {\"003W326T1T\", \"wid\": \"\", \"ic\": 0}, {\"003W32ZAGB\", \"wid\": \"\", \"ic\": 0}, {\"003W34H7Y7\", \"wid\": \"\", \"ic\": 0}, {\"003W35L6S7\", \"wid\": \"\", \"ic\": 0}, {\"003W36D0H7\", \"wid\": \"\", \"ic\": 0}, {\"003W37JRVR\", \"wid\": \"\", \"ic\": 0}, {\"003W39BG8X\", \"wid\": \"\", \"ic\": 0}, {\"003W3ATZ32\", \"wid\": \"\", \"ic\": 0}, {\"003W3CK6ME\", \"wid\": \"\", \"ic\": 0}, {\"003W3DE7NS\", \"wid\": \"\", \"ic\": 0}, {\"003W3EZJK8\", \"wid\": \"\", \"ic\": 0}, {\"003W3FQGTE\", \"wid\": \"\", \"ic\": 0}, {\"003W3H89AZ\", \"wid\": \"\", \"ic\": 0}, {\"003W3YF8N2\", \"wid\": \"\", \"ic\": 0}, {\"003W3KB7TV\", \"wid\": \"\", \"ic\": 0}, {\"003W3LDT68\", \"wid\": \"\", \"ic\": 0}, {\"003W3LRD30\", \"wid\": \"\", \"ic\": 0}, {\"003W3NWH8B\", \"wid\": \"\", \"ic\": 0}, {\"003W3ZP4HY\", \"wid\": \"\", \"ic\": 0}, {\"003W3PNA6Q\", \"wid\": \"\", \"ic\": 0}, {\"003W3R1Y68\", \"wid\": \"\", \"ic\": 0}, {\"003W3T15DX\", \"wid\": \"\", \"ic\": 0}, {\"003W3U71NA\", \"wid\": \"\", \"ic\": 0}, {\"003W3UQEN9\", \"wid\": \"\", \"ic\": 0}, {\"003W3WKGAL\", \"wid\": \"\", \"ic\": 0}, {\"003W3XBPU8\", \"wid\": \"\", \"ic\": 0}, {\"003W40NMUF\", \"wid\": \"\", \"ic\": 0}, {\"003W41U006\", \"wid\": \"\", \"ic\": 0}, {\"003W43G2EU\", \"wid\": \"\", \"ic\": 0}, {\"003W44YV5E\", \"wid\": \"\", \"ic\": 0}, {\"003W45RKCX\", \"wid\": \"\", \"ic\": 0}, {\"003W47BA3N\", \"wid\": \"\", \"ic\": 0}, {\"003W48UCZB\", \"wid\": \"\", \"ic\": 0}, {\"003W4A7DFY\", \"wid\": \"\", \"ic\": 0}, {\"003W4BEGN1\", \"wid\": \"\", \"ic\": 0}, {\"003W4CVKHS\", \"wid\": \"\", \"ic\": 0}, {\"003W4E5U3B\", \"wid\": \"\", \"ic\": 0}, {\"003W4F9UTQ\", \"wid\": \"\", \"ic\": 0}, {\"003W4GP707\", \"wid\": \"\", \"ic\": 0}, {\"003W4Y3YYP\", \"wid\": \"\", \"ic\": 0}, {\"003W4J9Z0R\", \"wid\": \"\", \"ic\": 0}, {\"003W4JQ44Q\", \"wid\": \"\", \"ic\": 0}, {\"003W4L2Y3R\", \"wid\": \"\", \"ic\": 0}, {\"003W4MF3MZ\", \"wid\": \"\", \"ic\": 0}, {\"003W4Z2J9F\", \"wid\": \"\", \"ic\": 0}, {\"003W4PW5EE\", \"wid\": \"\", \"ic\": 0}, {\"003W4QKUMC\", \"wid\": \"\", \"ic\": 0}, {\"003W4RUD61\", \"wid\": \"\", \"ic\": 0}, {\"003W4T214U\", \"wid\": \"\", \"ic\": 0}, {\"003W4UA2Z0\", \"wid\": \"\", \"ic\": 0}, {\"003W4W0GKJ\", \"wid\": \"\", \"ic\": 0}, {\"003W4X3F68\", \"wid\": \"\", \"ic\": 0}, {\"003W50HLLG\", \"wid\": \"\", \"ic\": 0}, {\"003W51CY1L\", \"wid\": \"\", \"ic\": 0}, {\"003W52T8SY\", \"wid\": \"\", \"ic\": 0}, {\"003W54QJ4R\", \"wid\": \"\", \"ic\": 0}, {\"003W55ULDN\", \"wid\": \"\", \"ic\": 0}, {\"003W56C7VF\", \"wid\": \"\", \"ic\": 0}, {\"003W58SVRT\", \"wid\": \"\", \"ic\": 0}, {\"003W59CVLR\", \"wid\": \"\", \"ic\": 0}, {\"003W5A9GEZ\", \"wid\": \"\", \"ic\": 0}, {\"003W5CEAH1\", \"wid\": \"\", \"ic\": 0}, {\"003W5CRUGE\", \"wid\": \"\", \"ic\": 0}, {\"003W5E1SQJ\", \"wid\": \"\", \"ic\": 0}, {\"003W5G6B9F\", \"wid\": \"\", \"ic\": 0}, {\"003W5GTZFV\", \"wid\": \"\", \"ic\": 0}, {\"003W5YVF6V\", \"wid\": \"\", \"ic\": 0}, {\"003W5KCYHB\", \"wid\": \"\", \"ic\": 0}, {\"003W5KMWGK\", \"wid\": \"\", \"ic\": 0}, {\"003W5LZ8NH\", \"wid\": \"\", \"ic\": 0}, {\"003W5Z66G5\", \"wid\": \"\", \"ic\": 0}, {\"003W5ZCLPC\", \"wid\": \"\", \"ic\": 0}, {\"003W5QCDZH\", \"wid\": \"\", \"ic\": 0}, {\"003W5R7FQB\", \"wid\": \"\", \"ic\": 0}, {\"003W5SN1XQ\", \"wid\": \"\", \"ic\": 0}, {\"003W5U0D42\", \"wid\": \"\", \"ic\": 0}, {\"003W5UQC4M\", \"wid\": \"\", \"ic\": 0}, {\"003W5WRZ3T\", \"wid\": \"\", \"ic\": 0}, {\"003W5XPGF4\", \"wid\": \"\", \"ic\": 0}, {\"003W61MW2X\", \"wid\": \"\", \"ic\": 0}, {\"003W630P81\", \"wid\": \"\", \"ic\": 0}, {\"003W63GT90\", \"wid\": \"\", \"ic\": 0}, {\"003W650BTK\", \"wid\": \"\", \"ic\": 0}, {\"003W665MT6\", \"wid\": \"\", \"ic\": 0}, {\"003W66WNQF\", \"wid\": \"\", \"ic\": 0}, {\"003W691KTF\", \"wid\": \"\", \"ic\": 0}, {\"003W6A1RD5\", \"wid\": \"\", \"ic\": 0}, {\"003W6B7YKX\", \"wid\": \"\", \"ic\": 0}, {\"003W6D735V\", \"wid\": \"\", \"ic\": 0}, {\"003W6DMALR\", \"wid\": \"\", \"ic\": 0}, {\"003W6FL0Z9\", \"wid\": \"\", \"ic\": 0}, {\"003W6FWS5M\", \"wid\": \"\", \"ic\": 0}, {\"003W6Y4M21\", \"wid\": \"\", \"ic\": 0}, {\"003W6J9SUD\", \"wid\": \"\", \"ic\": 0}, {\"003W6KEJ3R\", \"wid\": \"\", \"ic\": 0}, {\"003W6L4X0K\", \"wid\": \"\", \"ic\": 0}, {\"003W6N91JB\", \"wid\": \"\", \"ic\": 0}, {\"003W6ZPFX5\", \"wid\": \"\", \"ic\": 0}, {\"003W6Q5CLG\", \"wid\": \"\", \"ic\": 0}, {\"003W6R9B17\", \"wid\": \"\", \"ic\": 0}, {\"003W6SY9PK\", \"wid\": \"\", \"ic\": 0}, {\"003W6T2EL1\", \"wid\": \"\", \"ic\": 0}, {\"003W6V73BQ\", \"wid\": \"\", \"ic\": 0}, {\"003W6VE5D5\", \"wid\": \"\", \"ic\": 0}, {\"003W6XMWFL\", \"wid\": \"\", \"ic\": 0}, {\"003W70NKZ0\", \"wid\": \"\", \"ic\": 0}, {\"003W724JT6\", \"wid\": \"\", \"ic\": 0}, {\"003W738XXK\", \"wid\": \"\", \"ic\": 0}, {\"003W73T9YM\", \"wid\": \"\", \"ic\": 0}, {\"003W751BFM\", \"wid\": \"\", \"ic\": 0}, {\"003W771R5T\", \"wid\": \"\", \"ic\": 0}, {\"003W77MWDP\", \"wid\": \"\", \"ic\": 0}, {\"003W7A5531\", \"wid\": \"\", \"ic\": 0}, {\"003W7AXYK5\", \"wid\": \"\", \"ic\": 0}, {\"003W7CNT7P\", \"wid\": \"\", \"ic\": 0}, {\"003W7DQVCG\", \"wid\": \"\", \"ic\": 0}, {\"003W7EEKXS\", \"wid\": \"\", \"ic\": 0}, {\"003W7G4KW8\", \"wid\": \"\", \"ic\": 0}, {\"003W7HR3TG\", \"wid\": \"\", \"ic\": 0}, {\"003W7Y8JLS\", \"wid\": \"\", \"ic\": 0}, {\"003W7JCNXB\", \"wid\": \"\", \"ic\": 0}, {\"003W7LMGTT\", \"wid\": \"\", \"ic\": 0}, {\"003W7MQ9W0\", \"wid\": \"\", \"ic\": 0}, {\"003W7ND2E3\", \"wid\": \"\", \"ic\": 0}, {\"003W7PACE9\", \"wid\": \"\", \"ic\": 0}, {\"003W7Q31G0\", \"wid\": \"\", \"ic\": 0}, {\"003W7RPRU2\", \"wid\": \"\", \"ic\": 0}, {\"003W7SLSTN\", \"wid\": \"\", \"ic\": 0}, {\"003W7UYS6N\", \"wid\": \"\", \"ic\": 0}, {\"003W7VV61A\", \"wid\": \"\", \"ic\": 0}, {\"003W7VXRW1\", \"wid\": \"\", \"ic\": 0}, {\"003W8096A8\", \"wid\": \"\", \"ic\": 0}, {\"003W812PKX\", \"wid\": \"\", \"ic\": 0}, {\"003W81SJ9X\", \"wid\": \"\", \"ic\": 0}, {\"003W83MEDD\", \"wid\": \"\", \"ic\": 0}, {\"003W85YM6R\", \"wid\": \"\", \"ic\": 0}, {\"003W86Q52C\", \"wid\": \"\", \"ic\": 0}, {\"003W870AQW\", \"wid\": \"\", \"ic\": 0}, {\"003W892SHX\", \"wid\": \"\", \"ic\": 0}, {\"003W89WUDT\", \"wid\": \"\", \"ic\": 0}, {\"003W8BFARN\", \"wid\": \"\", \"ic\": 0}, {\"003W8D5LF9\", \"wid\": \"\", \"ic\": 0}, {\"003W8DDL83\", \"wid\": \"\", \"ic\": 0}, {\"003W8ERMVV\", \"wid\": \"\", \"ic\": 0}, {\"003W8G99FF\", \"wid\": \"\", \"ic\": 0}, {\"003W8Y9W9Q\", \"wid\": \"\", \"ic\": 0}, {\"003W8YUUDY\", \"wid\": \"\", \"ic\": 0}, {\"003W8JR476\", \"wid\": \"\", \"ic\": 0}, {\"003W8M7N3U\", \"wid\": \"\", \"ic\": 0}, {\"003W8MS79H\", \"wid\": \"\", \"ic\": 0}, {\"003W8NQ1XG\", \"wid\": \"\", \"ic\": 0}, {\"003W8PPDEA\", \"wid\": \"\", \"ic\": 0}, {\"003W8R6YRB\", \"wid\": \"\", \"ic\": 0}, {\"003W8SH3RT\", \"wid\": \"\", \"ic\": 0}, {\"003W8T3MTP\", \"wid\": \"\", \"ic\": 0}, {\"003W8UGXW8\", \"wid\": \"\", \"ic\": 0}, {\"003W8VKTD3\", \"wid\": \"\", \"ic\": 0}, {\"003W8XN028\", \"wid\": \"\", \"ic\": 0}, {\"003W90C2TL\", \"wid\": \"\", \"ic\": 0}, {\"003W91WYC7\", \"wid\": \"\", \"ic\": 0}, {\"003W937927\", \"wid\": \"\", \"ic\": 0}, {\"003W949Z4N\", \"wid\": \"\", \"ic\": 0}, {\"003W956Y4B\", \"wid\": \"\", \"ic\": 0}, {\"003W97GZCS\", \"wid\": \"\", \"ic\": 0}, {\"003W97MV1G\", \"wid\": \"\", \"ic\": 0}, {\"003W9966HW\", \"wid\": \"\", \"ic\": 0}, {\"003W9AJXB4\", \"wid\": \"\", \"ic\": 0}, {\"003W9BUQC1\", \"wid\": \"\", \"ic\": 0}, {\"003W9D6BSK\", \"wid\": \"\", \"ic\": 0}, {\"003W9EE01P\", \"wid\": \"\", \"ic\": 0}, {\"003W9GYLU7\", \"wid\": \"\", \"ic\": 0}, {\"003W9HFAY9\", \"wid\": \"\", \"ic\": 0}, {\"003W9Y4ZAV\", \"wid\": \"\", \"ic\": 0}, {\"003W9KDA99\", \"wid\": \"\", \"ic\": 0}, {\"003W9KS3GV\", \"wid\": \"\", \"ic\": 0}, {\"003W9MPR55\", \"wid\": \"\", \"ic\": 0}, {\"003W9N6132\", \"wid\": \"\", \"ic\": 0}, {\"003W9PH5N6\", \"wid\": \"\", \"ic\": 0}, {\"003W9QDE30\", \"wid\": \"\", \"ic\": 0}, {\"003W9S2N40\", \"wid\": \"\", \"ic\": 0}, {\"003W9T61VV\", \"wid\": \"\", \"ic\": 0}, {\"003W9TM97T\", \"wid\": \"\", \"ic\": 0}, {\"003W9VV42T\", \"wid\": \"\", \"ic\": 0}, {\"003W9X1UCU\", \"wid\": \"\", \"ic\": 0}, {\"003W9XUSJJ\", \"wid\": \"\", \"ic\": 0}, {\"003WA1M73S\", \"wid\": \"\", \"ic\": 0}, {\"003WA2PY57\", \"wid\": \"\", \"ic\": 0}, {\"003WA3G33H\", \"wid\": \"\", \"ic\": 0}, {\"003WA4VJPU\", \"wid\": \"\", \"ic\": 0}, {\"003WA6T42Y\", \"wid\": \"\", \"ic\": 0}, {\"003WA868WX\", \"wid\": \"\", \"ic\": 0}, {\"003WA8UFFJ\", \"wid\": \"\", \"ic\": 0}, {\"003WA9UVG3\", \"wid\": \"\", \"ic\": 0}, {\"003WAC2YYF\", \"wid\": \"\", \"ic\": 0}, {\"003WACNBKZ\", \"wid\": \"\", \"ic\": 0}, {\"003WADF6G2\", \"wid\": \"\", \"ic\": 0}, {\"003WAF2L6Y\", \"wid\": \"\", \"ic\": 0}, {\"003WAH616J\", \"wid\": \"\", \"ic\": 0}, {\"003WAY6VFB\", \"wid\": \"\", \"ic\": 0}, {\"003WAJ74TA\", \"wid\": \"\", \"ic\": 0}, {\"003WAKML7K\", \"wid\": \"\", \"ic\": 0}, {\"003WAM6QMX\", \"wid\": \"\", \"ic\": 0}, {\"003WAMWNKW\", \"wid\": \"\", \"ic\": 0}, {\"003WAZDYKU\", \"wid\": \"\", \"ic\": 0}, {\"003WAPYW99\", \"wid\": \"\", \"ic\": 0}, {\"003WAQVKGX\", \"wid\": \"\", \"ic\": 0}, {\"003WASZNGL\", \"wid\": \"\", \"ic\": 0}, {\"003WATF33B\", \"wid\": \"\", \"ic\": 0}, {\"003WAV0TP4\", \"wid\": \"\", \"ic\": 0}, {\"003WAVXHP7\", \"wid\": \"\", \"ic\": 0}, {\"003WAWZLD2\", \"wid\": \"\", \"ic\": 0}, {\"003WB08SAU\", \"wid\": \"\", \"ic\": 0}, {\"003WB2DDGJ\", \"wid\": \"\", \"ic\": 0}, {\"003WB2WT5B\", \"wid\": \"\", \"ic\": 0}, {\"003WB4A539\", \"wid\": \"\", \"ic\": 0}, {\"003WB5K1T6\", \"wid\": \"\", \"ic\": 0}, {\"003WB6FR1U\", \"wid\": \"\", \"ic\": 0}, {\"003WB8QHXW\", \"wid\": \"\", \"ic\": 0}, {\"003WB8V0HA\", \"wid\": \"\", \"ic\": 0}, {\"003WBAYDPB\", \"wid\": \"\", \"ic\": 0}, {\"003WBBKT3U\", \"wid\": \"\", \"ic\": 0}, {\"003WBDT1Z9\", \"wid\": \"\", \"ic\": 0}, {\"003WBECJHF\", \"wid\": \"\", \"ic\": 0}, {\"003WBGK51J\", \"wid\": \"\", \"ic\": 0}, {\"003WBHR1P4\", \"wid\": \"\", \"ic\": 0}, {\"003WBY3ZAW\", \"wid\": \"\", \"ic\": 0}, {\"003WBJJ56R\", \"wid\": \"\", \"ic\": 0}, {\"003WBL4RD9\", \"wid\": \"\", \"ic\": 0}, {\"003WBLWTAW\", \"wid\": \"\", \"ic\": 0}, {\"003WBZ29V5\", \"wid\": \"\", \"ic\": 0}, {\"003WBP76MH\", \"wid\": \"\", \"ic\": 0}, {\"003WBPV9U4\", \"wid\": \"\", \"ic\": 0}, {\"003WBR6L27\", \"wid\": \"\", \"ic\": 0}, {\"003WBSJWER\", \"wid\": \"\", \"ic\": 0}, {\"003WBUDPFK\", \"wid\": \"\", \"ic\": 0}, {\"003WBUZRNG\", \"wid\": \"\", \"ic\": 0}, {\"003WBX2VHG\", \"wid\": \"\", \"ic\": 0}, {\"003WBXJU0V\", \"wid\": \"\", \"ic\": 0}, {\"003WC0K2SZ\", \"wid\": \"\", \"ic\": 0}, {\"003WC32JV3\", \"wid\": \"\", \"ic\": 0}, {\"003WC4A0DW\", \"wid\": \"\", \"ic\": 0}, {\"003WC5ME13\", \"wid\": \"\", \"ic\": 0}, {\"003WC5PVAJ\", \"wid\": \"\", \"ic\": 0}, {\"003WC7UQ95\", \"wid\": \"\", \"ic\": 0}, {\"003WC8BYAF\", \"wid\": \"\", \"ic\": 0}, {\"003WCA2ZKT\", \"wid\": \"\", \"ic\": 0}, {\"003WCBFBC4\", \"wid\": \"\", \"ic\": 0}, {\"003WCCJCZP\", \"wid\": \"\", \"ic\": 0}, {\"003WCDFL4J\", \"wid\": \"\", \"ic\": 0}, {\"003WCFFYPD\", \"wid\": \"\", \"ic\": 0}, {\"003WCGAFPZ\", \"wid\": \"\", \"ic\": 0}, {\"003WCYB33K\", \"wid\": \"\", \"ic\": 0}, {\"003WCJQ9X2\", \"wid\": \"\", \"ic\": 0}, {\"003WCL0U7P\", \"wid\": \"\", \"ic\": 0}, {\"003WCLDMAS\", \"wid\": \"\", \"ic\": 0}, {\"003WCNAZAL\", \"wid\": \"\", \"ic\": 0}, {\"003WCZLVJQ\", \"wid\": \"\", \"ic\": 0}, {\"003WCPY6JQ\", \"wid\": \"\", \"ic\": 0}, {\"003WCQMX5U\", \"wid\": \"\", \"ic\": 0}, {\"003WCS4EW0\", \"wid\": \"\", \"ic\": 0}, {\"003WCU2G5N\", \"wid\": \"\", \"ic\": 0}, {\"003WCUYUJP\", \"wid\": \"\", \"ic\": 0}, {\"003WCVNL38\", \"wid\": \"\", \"ic\": 0}, {\"003WCWZF9D\", \"wid\": \"\", \"ic\": 0}, {\"003WD0WFBS\", \"wid\": \"\", \"ic\": 0}, {\"003WD26LC9\", \"wid\": \"\", \"ic\": 0}, {\"003WD3BWBT\", \"wid\": \"\", \"ic\": 0}, {\"003WD42Y6E\", \"wid\": \"\", \"ic\": 0}, {\"003WD5V7G7\", \"wid\": \"\", \"ic\": 0}, {\"003WD6T7HQ\", \"wid\": \"\", \"ic\": 0}, {\"003WD8RV02\", \"wid\": \"\", \"ic\": 0}, {\"003WD93WRZ\", \"wid\": \"\", \"ic\": 0}]"
	acs1 := &acs{}
	json.Unmarshal(slice(s), acs1)
	fmt.Println(acs1)
}

func f6() {
	//s := "  [{\"003W1XZMUM\", \"wid\": \"\", \"ic\": 0}, {\"003W20HZZ6\", \"wid\": \"\", \"ic\": 0}, {\"003W21Q513\", \"wid\": \"\", \"ic\": 0}, {\"003W231S2Z\", \"wid\": \"\", \"ic\": 0}, {\"003W25HA4N\", \"wid\": \"\", \"ic\": 0}, {\"003W26CZ21\", \"wid\": \"\", \"ic\": 0}, {\"003W27L07Z\", \"wid\": \"\", \"ic\": 0}, {\"003W297RTV\", \"wid\": \"\", \"ic\": 0}, {\"003W2AK6AZ\", \"wid\": \"\", \"ic\": 0}, {\"003W2BBHCQ\", \"wid\": \"\", \"ic\": 0}, {\"003W2CRB0E\", \"wid\": \"\", \"ic\": 0}, {\"003W2DBVAF\", \"wid\": \"\", \"ic\": 0}, {\"003W2FNAR7\", \"wid\": \"\", \"ic\": 0}, {\"003W2G6VTC\", \"wid\": \"\", \"ic\": 0}, {\"003W2HXMXG\", \"wid\": \"\", \"ic\": 0}, {\"003W2YQ3ZP\", \"wid\": \"\", \"ic\": 0}, {\"003W2JWHRC\", \"wid\": \"\", \"ic\": 0}, {\"003W2LQEXD\", \"wid\": \"\", \"ic\": 0}, {\"003W2NBPAB\", \"wid\": \"\", \"ic\": 0}, {\"003W2ZG48D\", \"wid\": \"\", \"ic\": 0}, {\"003W2ZURAR\", \"wid\": \"\", \"ic\": 0}, {\"003W2QSM0Y\", \"wid\": \"\", \"ic\": 0}, {\"003W2SNB08\", \"wid\": \"\", \"ic\": 0}, {\"003W2TN4Q5\", \"wid\": \"\", \"ic\": 0}, {\"003W2UCSVU\", \"wid\": \"\", \"ic\": 0}, {\"003W2W8HFJ\", \"wid\": \"\", \"ic\": 0}, {\"003W2WPTL2\", \"wid\": \"\", \"ic\": 0}, {\"003W30AF0K\", \"wid\": \"\", \"ic\": 0}, {\"003W326T1T\", \"wid\": \"\", \"ic\": 0}, {\"003W32ZAGB\", \"wid\": \"\", \"ic\": 0}, {\"003W34H7Y7\", \"wid\": \"\", \"ic\": 0}, {\"003W35L6S7\", \"wid\": \"\", \"ic\": 0}, {\"003W36D0H7\", \"wid\": \"\", \"ic\": 0}, {\"003W37JRVR\", \"wid\": \"\", \"ic\": 0}, {\"003W39BG8X\", \"wid\": \"\", \"ic\": 0}, {\"003W3ATZ32\", \"wid\": \"\", \"ic\": 0}, {\"003W3CK6ME\", \"wid\": \"\", \"ic\": 0}, {\"003W3DE7NS\", \"wid\": \"\", \"ic\": 0}, {\"003W3EZJK8\", \"wid\": \"\", \"ic\": 0}, {\"003W3FQGTE\", \"wid\": \"\", \"ic\": 0}, {\"003W3H89AZ\", \"wid\": \"\", \"ic\": 0}, {\"003W3YF8N2\", \"wid\": \"\", \"ic\": 0}, {\"003W3KB7TV\", \"wid\": \"\", \"ic\": 0}, {\"003W3LDT68\", \"wid\": \"\", \"ic\": 0}, {\"003W3LRD30\", \"wid\": \"\", \"ic\": 0}, {\"003W3NWH8B\", \"wid\": \"\", \"ic\": 0}, {\"003W3ZP4HY\", \"wid\": \"\", \"ic\": 0}, {\"003W3PNA6Q\", \"wid\": \"\", \"ic\": 0}, {\"003W3R1Y68\", \"wid\": \"\", \"ic\": 0}, {\"003W3T15DX\", \"wid\": \"\", \"ic\": 0}, {\"003W3U71NA\", \"wid\": \"\", \"ic\": 0}, {\"003W3UQEN9\", \"wid\": \"\", \"ic\": 0}, {\"003W3WKGAL\", \"wid\": \"\", \"ic\": 0}, {\"003W3XBPU8\", \"wid\": \"\", \"ic\": 0}, {\"003W40NMUF\", \"wid\": \"\", \"ic\": 0}, {\"003W41U006\", \"wid\": \"\", \"ic\": 0}, {\"003W43G2EU\", \"wid\": \"\", \"ic\": 0}, {\"003W44YV5E\", \"wid\": \"\", \"ic\": 0}, {\"003W45RKCX\", \"wid\": \"\", \"ic\": 0}, {\"003W47BA3N\", \"wid\": \"\", \"ic\": 0}, {\"003W48UCZB\", \"wid\": \"\", \"ic\": 0}, {\"003W4A7DFY\", \"wid\": \"\", \"ic\": 0}, {\"003W4BEGN1\", \"wid\": \"\", \"ic\": 0}, {\"003W4CVKHS\", \"wid\": \"\", \"ic\": 0}, {\"003W4E5U3B\", \"wid\": \"\", \"ic\": 0}, {\"003W4F9UTQ\", \"wid\": \"\", \"ic\": 0}, {\"003W4GP707\", \"wid\": \"\", \"ic\": 0}, {\"003W4Y3YYP\", \"wid\": \"\", \"ic\": 0}, {\"003W4J9Z0R\", \"wid\": \"\", \"ic\": 0}, {\"003W4JQ44Q\", \"wid\": \"\", \"ic\": 0}, {\"003W4L2Y3R\", \"wid\": \"\", \"ic\": 0}, {\"003W4MF3MZ\", \"wid\": \"\", \"ic\": 0}, {\"003W4Z2J9F\", \"wid\": \"\", \"ic\": 0}, {\"003W4PW5EE\", \"wid\": \"\", \"ic\": 0}, {\"003W4QKUMC\", \"wid\": \"\", \"ic\": 0}, {\"003W4RUD61\", \"wid\": \"\", \"ic\": 0}, {\"003W4T214U\", \"wid\": \"\", \"ic\": 0}, {\"003W4UA2Z0\", \"wid\": \"\", \"ic\": 0}, {\"003W4W0GKJ\", \"wid\": \"\", \"ic\": 0}, {\"003W4X3F68\", \"wid\": \"\", \"ic\": 0}, {\"003W50HLLG\", \"wid\": \"\", \"ic\": 0}, {\"003W51CY1L\", \"wid\": \"\", \"ic\": 0}, {\"003W52T8SY\", \"wid\": \"\", \"ic\": 0}, {\"003W54QJ4R\", \"wid\": \"\", \"ic\": 0}, {\"003W55ULDN\", \"wid\": \"\", \"ic\": 0}, {\"003W56C7VF\", \"wid\": \"\", \"ic\": 0}, {\"003W58SVRT\", \"wid\": \"\", \"ic\": 0}, {\"003W59CVLR\", \"wid\": \"\", \"ic\": 0}, {\"003W5A9GEZ\", \"wid\": \"\", \"ic\": 0}, {\"003W5CEAH1\", \"wid\": \"\", \"ic\": 0}, {\"003W5CRUGE\", \"wid\": \"\", \"ic\": 0}, {\"003W5E1SQJ\", \"wid\": \"\", \"ic\": 0}, {\"003W5G6B9F\", \"wid\": \"\", \"ic\": 0}, {\"003W5GTZFV\", \"wid\": \"\", \"ic\": 0}, {\"003W5YVF6V\", \"wid\": \"\", \"ic\": 0}, {\"003W5KCYHB\", \"wid\": \"\", \"ic\": 0}, {\"003W5KMWGK\", \"wid\": \"\", \"ic\": 0}, {\"003W5LZ8NH\", \"wid\": \"\", \"ic\": 0}, {\"003W5Z66G5\", \"wid\": \"\", \"ic\": 0}, {\"003W5ZCLPC\", \"wid\": \"\", \"ic\": 0}, {\"003W5QCDZH\", \"wid\": \"\", \"ic\": 0}, {\"003W5R7FQB\", \"wid\": \"\", \"ic\": 0}, {\"003W5SN1XQ\", \"wid\": \"\", \"ic\": 0}, {\"003W5U0D42\", \"wid\": \"\", \"ic\": 0}, {\"003W5UQC4M\", \"wid\": \"\", \"ic\": 0}, {\"003W5WRZ3T\", \"wid\": \"\", \"ic\": 0}, {\"003W5XPGF4\", \"wid\": \"\", \"ic\": 0}, {\"003W61MW2X\", \"wid\": \"\", \"ic\": 0}, {\"003W630P81\", \"wid\": \"\", \"ic\": 0}, {\"003W63GT90\", \"wid\": \"\", \"ic\": 0}, {\"003W650BTK\", \"wid\": \"\", \"ic\": 0}, {\"003W665MT6\", \"wid\": \"\", \"ic\": 0}, {\"003W66WNQF\", \"wid\": \"\", \"ic\": 0}, {\"003W691KTF\", \"wid\": \"\", \"ic\": 0}, {\"003W6A1RD5\", \"wid\": \"\", \"ic\": 0}, {\"003W6B7YKX\", \"wid\": \"\", \"ic\": 0}, {\"003W6D735V\", \"wid\": \"\", \"ic\": 0}, {\"003W6DMALR\", \"wid\": \"\", \"ic\": 0}, {\"003W6FL0Z9\", \"wid\": \"\", \"ic\": 0}, {\"003W6FWS5M\", \"wid\": \"\", \"ic\": 0}, {\"003W6Y4M21\", \"wid\": \"\", \"ic\": 0}, {\"003W6J9SUD\", \"wid\": \"\", \"ic\": 0}, {\"003W6KEJ3R\", \"wid\": \"\", \"ic\": 0}, {\"003W6L4X0K\", \"wid\": \"\", \"ic\": 0}, {\"003W6N91JB\", \"wid\": \"\", \"ic\": 0}, {\"003W6ZPFX5\", \"wid\": \"\", \"ic\": 0}, {\"003W6Q5CLG\", \"wid\": \"\", \"ic\": 0}, {\"003W6R9B17\", \"wid\": \"\", \"ic\": 0}, {\"003W6SY9PK\", \"wid\": \"\", \"ic\": 0}, {\"003W6T2EL1\", \"wid\": \"\", \"ic\": 0}, {\"003W6V73BQ\", \"wid\": \"\", \"ic\": 0}, {\"003W6VE5D5\", \"wid\": \"\", \"ic\": 0}, {\"003W6XMWFL\", \"wid\": \"\", \"ic\": 0}, {\"003W70NKZ0\", \"wid\": \"\", \"ic\": 0}, {\"003W724JT6\", \"wid\": \"\", \"ic\": 0}, {\"003W738XXK\", \"wid\": \"\", \"ic\": 0}, {\"003W73T9YM\", \"wid\": \"\", \"ic\": 0}, {\"003W751BFM\", \"wid\": \"\", \"ic\": 0}, {\"003W771R5T\", \"wid\": \"\", \"ic\": 0}, {\"003W77MWDP\", \"wid\": \"\", \"ic\": 0}, {\"003W7A5531\", \"wid\": \"\", \"ic\": 0}, {\"003W7AXYK5\", \"wid\": \"\", \"ic\": 0}, {\"003W7CNT7P\", \"wid\": \"\", \"ic\": 0}, {\"003W7DQVCG\", \"wid\": \"\", \"ic\": 0}, {\"003W7EEKXS\", \"wid\": \"\", \"ic\": 0}, {\"003W7G4KW8\", \"wid\": \"\", \"ic\": 0}, {\"003W7HR3TG\", \"wid\": \"\", \"ic\": 0}, {\"003W7Y8JLS\", \"wid\": \"\", \"ic\": 0}, {\"003W7JCNXB\", \"wid\": \"\", \"ic\": 0}, {\"003W7LMGTT\", \"wid\": \"\", \"ic\": 0}, {\"003W7MQ9W0\", \"wid\": \"\", \"ic\": 0}, {\"003W7ND2E3\", \"wid\": \"\", \"ic\": 0}, {\"003W7PACE9\", \"wid\": \"\", \"ic\": 0}, {\"003W7Q31G0\", \"wid\": \"\", \"ic\": 0}, {\"003W7RPRU2\", \"wid\": \"\", \"ic\": 0}, {\"003W7SLSTN\", \"wid\": \"\", \"ic\": 0}, {\"003W7UYS6N\", \"wid\": \"\", \"ic\": 0}, {\"003W7VV61A\", \"wid\": \"\", \"ic\": 0}, {\"003W7VXRW1\", \"wid\": \"\", \"ic\": 0}, {\"003W8096A8\", \"wid\": \"\", \"ic\": 0}, {\"003W812PKX\", \"wid\": \"\", \"ic\": 0}, {\"003W81SJ9X\", \"wid\": \"\", \"ic\": 0}, {\"003W83MEDD\", \"wid\": \"\", \"ic\": 0}, {\"003W85YM6R\", \"wid\": \"\", \"ic\": 0}, {\"003W86Q52C\", \"wid\": \"\", \"ic\": 0}, {\"003W870AQW\", \"wid\": \"\", \"ic\": 0}, {\"003W892SHX\", \"wid\": \"\", \"ic\": 0}, {\"003W89WUDT\", \"wid\": \"\", \"ic\": 0}, {\"003W8BFARN\", \"wid\": \"\", \"ic\": 0}, {\"003W8D5LF9\", \"wid\": \"\", \"ic\": 0}, {\"003W8DDL83\", \"wid\": \"\", \"ic\": 0}, {\"003W8ERMVV\", \"wid\": \"\", \"ic\": 0}, {\"003W8G99FF\", \"wid\": \"\", \"ic\": 0}, {\"003W8Y9W9Q\", \"wid\": \"\", \"ic\": 0}, {\"003W8YUUDY\", \"wid\": \"\", \"ic\": 0}, {\"003W8JR476\", \"wid\": \"\", \"ic\": 0}, {\"003W8M7N3U\", \"wid\": \"\", \"ic\": 0}, {\"003W8MS79H\", \"wid\": \"\", \"ic\": 0}, {\"003W8NQ1XG\", \"wid\": \"\", \"ic\": 0}, {\"003W8PPDEA\", \"wid\": \"\", \"ic\": 0}, {\"003W8R6YRB\", \"wid\": \"\", \"ic\": 0}, {\"003W8SH3RT\", \"wid\": \"\", \"ic\": 0}, {\"003W8T3MTP\", \"wid\": \"\", \"ic\": 0}, {\"003W8UGXW8\", \"wid\": \"\", \"ic\": 0}, {\"003W8VKTD3\", \"wid\": \"\", \"ic\": 0}, {\"003W8XN028\", \"wid\": \"\", \"ic\": 0}, {\"003W90C2TL\", \"wid\": \"\", \"ic\": 0}, {\"003W91WYC7\", \"wid\": \"\", \"ic\": 0}, {\"003W937927\", \"wid\": \"\", \"ic\": 0}, {\"003W949Z4N\", \"wid\": \"\", \"ic\": 0}, {\"003W956Y4B\", \"wid\": \"\", \"ic\": 0}, {\"003W97GZCS\", \"wid\": \"\", \"ic\": 0}, {\"003W97MV1G\", \"wid\": \"\", \"ic\": 0}, {\"003W9966HW\", \"wid\": \"\", \"ic\": 0}, {\"003W9AJXB4\", \"wid\": \"\", \"ic\": 0}, {\"003W9BUQC1\", \"wid\": \"\", \"ic\": 0}, {\"003W9D6BSK\", \"wid\": \"\", \"ic\": 0}, {\"003W9EE01P\", \"wid\": \"\", \"ic\": 0}, {\"003W9GYLU7\", \"wid\": \"\", \"ic\": 0}, {\"003W9HFAY9\", \"wid\": \"\", \"ic\": 0}, {\"003W9Y4ZAV\", \"wid\": \"\", \"ic\": 0}, {\"003W9KDA99\", \"wid\": \"\", \"ic\": 0}, {\"003W9KS3GV\", \"wid\": \"\", \"ic\": 0}, {\"003W9MPR55\", \"wid\": \"\", \"ic\": 0}, {\"003W9N6132\", \"wid\": \"\", \"ic\": 0}, {\"003W9PH5N6\", \"wid\": \"\", \"ic\": 0}, {\"003W9QDE30\", \"wid\": \"\", \"ic\": 0}, {\"003W9S2N40\", \"wid\": \"\", \"ic\": 0}, {\"003W9T61VV\", \"wid\": \"\", \"ic\": 0}, {\"003W9TM97T\", \"wid\": \"\", \"ic\": 0}, {\"003W9VV42T\", \"wid\": \"\", \"ic\": 0}, {\"003W9X1UCU\", \"wid\": \"\", \"ic\": 0}, {\"003W9XUSJJ\", \"wid\": \"\", \"ic\": 0}, {\"003WA1M73S\", \"wid\": \"\", \"ic\": 0}, {\"003WA2PY57\", \"wid\": \"\", \"ic\": 0}, {\"003WA3G33H\", \"wid\": \"\", \"ic\": 0}, {\"003WA4VJPU\", \"wid\": \"\", \"ic\": 0}, {\"003WA6T42Y\", \"wid\": \"\", \"ic\": 0}, {\"003WA868WX\", \"wid\": \"\", \"ic\": 0}, {\"003WA8UFFJ\", \"wid\": \"\", \"ic\": 0}, {\"003WA9UVG3\", \"wid\": \"\", \"ic\": 0}, {\"003WAC2YYF\", \"wid\": \"\", \"ic\": 0}, {\"003WACNBKZ\", \"wid\": \"\", \"ic\": 0}, {\"003WADF6G2\", \"wid\": \"\", \"ic\": 0}, {\"003WAF2L6Y\", \"wid\": \"\", \"ic\": 0}, {\"003WAH616J\", \"wid\": \"\", \"ic\": 0}, {\"003WAY6VFB\", \"wid\": \"\", \"ic\": 0}, {\"003WAJ74TA\", \"wid\": \"\", \"ic\": 0}, {\"003WAKML7K\", \"wid\": \"\", \"ic\": 0}, {\"003WAM6QMX\", \"wid\": \"\", \"ic\": 0}, {\"003WAMWNKW\", \"wid\": \"\", \"ic\": 0}, {\"003WAZDYKU\", \"wid\": \"\", \"ic\": 0}, {\"003WAPYW99\", \"wid\": \"\", \"ic\": 0}, {\"003WAQVKGX\", \"wid\": \"\", \"ic\": 0}, {\"003WASZNGL\", \"wid\": \"\", \"ic\": 0}, {\"003WATF33B\", \"wid\": \"\", \"ic\": 0}, {\"003WAV0TP4\", \"wid\": \"\", \"ic\": 0}, {\"003WAVXHP7\", \"wid\": \"\", \"ic\": 0}, {\"003WAWZLD2\", \"wid\": \"\", \"ic\": 0}, {\"003WB08SAU\", \"wid\": \"\", \"ic\": 0}, {\"003WB2DDGJ\", \"wid\": \"\", \"ic\": 0}, {\"003WB2WT5B\", \"wid\": \"\", \"ic\": 0}, {\"003WB4A539\", \"wid\": \"\", \"ic\": 0}, {\"003WB5K1T6\", \"wid\": \"\", \"ic\": 0}, {\"003WB6FR1U\", \"wid\": \"\", \"ic\": 0}, {\"003WB8QHXW\", \"wid\": \"\", \"ic\": 0}, {\"003WB8V0HA\", \"wid\": \"\", \"ic\": 0}, {\"003WBAYDPB\", \"wid\": \"\", \"ic\": 0}, {\"003WBBKT3U\", \"wid\": \"\", \"ic\": 0}, {\"003WBDT1Z9\", \"wid\": \"\", \"ic\": 0}, {\"003WBECJHF\", \"wid\": \"\", \"ic\": 0}, {\"003WBGK51J\", \"wid\": \"\", \"ic\": 0}, {\"003WBHR1P4\", \"wid\": \"\", \"ic\": 0}, {\"003WBY3ZAW\", \"wid\": \"\", \"ic\": 0}, {\"003WBJJ56R\", \"wid\": \"\", \"ic\": 0}, {\"003WBL4RD9\", \"wid\": \"\", \"ic\": 0}, {\"003WBLWTAW\", \"wid\": \"\", \"ic\": 0}, {\"003WBZ29V5\", \"wid\": \"\", \"ic\": 0}, {\"003WBP76MH\", \"wid\": \"\", \"ic\": 0}, {\"003WBPV9U4\", \"wid\": \"\", \"ic\": 0}, {\"003WBR6L27\", \"wid\": \"\", \"ic\": 0}, {\"003WBSJWER\", \"wid\": \"\", \"ic\": 0}, {\"003WBUDPFK\", \"wid\": \"\", \"ic\": 0}, {\"003WBUZRNG\", \"wid\": \"\", \"ic\": 0}, {\"003WBX2VHG\", \"wid\": \"\", \"ic\": 0}, {\"003WBXJU0V\", \"wid\": \"\", \"ic\": 0}, {\"003WC0K2SZ\", \"wid\": \"\", \"ic\": 0}, {\"003WC32JV3\", \"wid\": \"\", \"ic\": 0}, {\"003WC4A0DW\", \"wid\": \"\", \"ic\": 0}, {\"003WC5ME13\", \"wid\": \"\", \"ic\": 0}, {\"003WC5PVAJ\", \"wid\": \"\", \"ic\": 0}, {\"003WC7UQ95\", \"wid\": \"\", \"ic\": 0}, {\"003WC8BYAF\", \"wid\": \"\", \"ic\": 0}, {\"003WCA2ZKT\", \"wid\": \"\", \"ic\": 0}, {\"003WCBFBC4\", \"wid\": \"\", \"ic\": 0}, {\"003WCCJCZP\", \"wid\": \"\", \"ic\": 0}, {\"003WCDFL4J\", \"wid\": \"\", \"ic\": 0}, {\"003WCFFYPD\", \"wid\": \"\", \"ic\": 0}, {\"003WCGAFPZ\", \"wid\": \"\", \"ic\": 0}, {\"003WCYB33K\", \"wid\": \"\", \"ic\": 0}, {\"003WCJQ9X2\", \"wid\": \"\", \"ic\": 0}, {\"003WCL0U7P\", \"wid\": \"\", \"ic\": 0}, {\"003WCLDMAS\", \"wid\": \"\", \"ic\": 0}, {\"003WCNAZAL\", \"wid\": \"\", \"ic\": 0}, {\"003WCZLVJQ\", \"wid\": \"\", \"ic\": 0}, {\"003WCPY6JQ\", \"wid\": \"\", \"ic\": 0}, {\"003WCQMX5U\", \"wid\": \"\", \"ic\": 0}, {\"003WCS4EW0\", \"wid\": \"\", \"ic\": 0}, {\"003WCU2G5N\", \"wid\": \"\", \"ic\": 0}, {\"003WCUYUJP\", \"wid\": \"\", \"ic\": 0}, {\"003WCVNL38\", \"wid\": \"\", \"ic\": 0}, {\"003WCWZF9D\", \"wid\": \"\", \"ic\": 0}, {\"003WD0WFBS\", \"wid\": \"\", \"ic\": 0}, {\"003WD26LC9\", \"wid\": \"\", \"ic\": 0}, {\"003WD3BWBT\", \"wid\": \"\", \"ic\": 0}, {\"003WD42Y6E\", \"wid\": \"\", \"ic\": 0}, {\"003WD5V7G7\", \"wid\": \"\", \"ic\": 0}, {\"003WD6T7HQ\", \"wid\": \"\", \"ic\": 0}, {\"003WD8RV02\", \"wid\": \"\", \"ic\": 0}, {\"003WD93WRZ\", \"wid\": \"\", \"ic\": 0}]"
}

func Contain(obj interface{}, target interface{}) bool {
	targetValue := reflect.ValueOf(target)
	switch reflect.TypeOf(target).Kind() {
	case reflect.Slice, reflect.Array:
		for i := 0; i < targetValue.Len(); i++ {
			if targetValue.Index(i).Interface() == obj {
				return true
			}
		}
	case reflect.Map:
		if targetValue.MapIndex(reflect.ValueOf(obj)).IsValid() {
			return true
		}
	}
	return false
}
func f5() {
	obj := int32(1)
	target := []int{1, 2, 3}
	fmt.Println(Contain(obj, target))
}

func f3() {
	client := redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
	client.SAdd("basketball:sensitiveWords", 1)
	client.SAdd("basketball:sensitiveWords", 2)
	client.SAdd("basketball:sensitiveWords", 3)
	c := client.SMembers("basketball:sensitiveWords")
	for _, e := range c.Val() {
		fmt.Println(e)
	}

}

type s3 struct {
	Sug string `json:"sug"`
}

type s2 struct {
	Labels []s3 `json:"labels"`
}

type s1 struct {
	Showapi_res_body s2
}

func f4() {
	fmt.Println("error\",\"t\":\"2021-06-18T18:13:51.899+0800\",\"msg\":\"user load from db err:mongo: no documents in result\",\"service\":\"app\",\"service\":\"sub\",\"trace_id\":\"3ee4d5c7058b6928\",\"span_id\":\"5ac413456761f17a\",\"uid\":63795,\"stack\":\"goroutine 1023 [running]:\nruntime/debug.     Stack(0xc0002d6d00, 0x153f8f5, 0x1d)\n\t/home/liupan/.g/go/src/runtime/debug/stack.go:24 +0x9f\ngameserver/game/pkg/submap.(*SubMap).loadFromDB(0xc00152b7e0, 0x17e2ab8, 0xc001507fb0, 0xc001821320, 0x1, 0x1, 0x1)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/su     bmap/submap.go:297 +0x347\ngameserver/game/pkg/submap.(*SubMap).GetData(0xc00152b7e0, 0x17e2ab8, 0xc001507fb0, 0xc001a45438, 0x1, 0x1, 0x18)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/submap/submap.go:132 +0x3a5\ngameserver/game/pkg/submap.(*SubMap).GetOne(     ...)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/submap/submap.go:103\ngameserver/game/pkg/service/gloryroad.(*DataManager).GetGloryroadInfo(0xc00052ab90, 0x17e2ab8, 0xc001507fb0, 0xf933, 0x14365c0, 0x1, 0xc00137dec0)\n\t/mnt/d/work/SagiGames/Code/game-serve     r/game/pkg/service/gloryroad/gloryroad.go:126 +0xed\ngameserver/game/pkg/service/gloryroad.(*DataManager).GetCurCup(0xc00052ab90, 0x17e2ab8, 0xc001507fb0, 0xf933, 0xc0005e9720)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/service/gloryroad/data_manager.go:38      +0x49\ngameserver/game/pkg/service/player.GetPlayerPersonal(0x17e2ab8, 0xc001507fb0, 0x18cf9, 0x0, 0x17dc400, 0x7f8576118330)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/service/player/data_manager.go:121 +0x391\ngameserver/game/pkg/api.(*GameService).Player     Personal(0x203e808, 0x17e2ab8, 0xc001507fb0, 0xc0009a4000, 0x203e808, 0x0, 0xc0010d1638)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/api/player.go:39 +0x66\ngameserver/common/pb._Game_PlayerPersonal_Handler(0x15042a0, 0x203e808, 0x17e2ab8, 0xc001507fb0, 0xc0     00a84ba0, 0x0, 0x17e2ab8, 0xc001507fb0, 0x0, 0x0)\n\t/mnt/d/work/SagiGames/Code/game-server/common/pb/service_game_grpc.pb.go:1363 +0x217\ngameserver/game/pkg/components/handler.(*CompHandlerGRPC).HandleRequest(0xc000775ce0, 0x17e2ab8, 0xc001507f50, 0xc000bc8540, 0     x0, 0x0, 0x0)\n\t/mnt/d/work/SagiGames/Code/game-server/game/pkg/components/handler/handler.go:221 +0x5f3\ngameserver/game/pkg/cluster/logic.(*ModuleLogic).HandleRequest(0xc00075ec80, 0x17e2ab8, 0xc001507f50, 0xc000bc8540, 0xc00075ec80, 0x17e2ab8, 0xc001507f50)\n\t     /mnt/d/work/SagiGames/Code/game-server/game/pkg/cluster/logic/module.go:111 +0x52\ngameserver/common/pb._Rpc_HandleRequest_Handler.func1(0x17e2ab8, 0xc001507f50, 0x149d3a0, 0xc000bc8540, 0x17e2ab8, 0xc001507f50, 0x17f35e0, 0xc000695680)\n\t/mnt/d/work/SagiGames/Cod     e/game-server/common/pb/cluster_grpc.pb.go:324 +0x89\ngithub.com/opentracing-contrib/go-grpc.OpenTracingServerInterceptor.func1(0x17e2ab8, 0xc001507f50, 0x149d3a0, 0xc000bc8540, 0xc00124a8c0, 0xc000a84b58, 0x0, 0x0, 0x0, 0x0)\n\t/mnt/d/work/SagiGames/Code/game-serv     er/vendor/github.com/opentracing-contrib/go-grpc/server.go:57 +0x302\ngameserver/common/pb._Rpc_HandleRequest_Handler(0x14ec000, 0xc00075ec80, 0x17e2ab8, 0xc001507ef0, 0xc000d5f0e0, 0xc000775d20, 0x17e2ab8, 0xc001507ef0, 0xc0010c41e0, 0x2a)\n\t/mnt/d/work/SagiGames     /Code/game-server/common/pb/cluster_grpc.pb.go:326 +0x150\ngoogle.golang.org/grpc.(*Server).processUnaryRPC(0xc00054d880, 0x17eeb38, 0xc000c09200, 0xc00097a400, 0xc0005e49f0, 0x1ff9b20, 0x0, 0x0, 0x0)\n\t/mnt/d/work/SagiGames/Code/game-server/vendor/google.golang.o     rg/grpc/server.go:1217 +0x52b\ngoogle.golang.org/grpc.(*Server).handleStream(0xc00054d880, 0x17eeb38, 0xc000c09200, 0xc00097a400, 0x0)\n\t/mnt/d/work/SagiGames/Code/game-server/vendor/google.golang.org/grpc/server.go:1540 +0xd0c\ngoogle.golang.org/grpc.(*Server).se     rveStreams.func1.2(0xc0008f9290, 0xc00054d880, 0x17eeb38, 0xc000c09200, 0xc00097a400)\n\t/mnt/d/work/SagiGames/Code/game-server/vendor/google.golang.org/grpc/server.go:878 +0xab\ncreated by google.golang.org/grpc.(*Server).serveStreams.func1\n\t/mnt/d/work/SagiGame     s/Code/game-server/vendor/google.golang.org/grpc/server.go:876 +0x1fd\n")
}

func f2() {
	// 云市场分配的密钥Id
	secretId := "AKID9dkMDi0l7oiwjjaiFmlI2s47y39c8qZb2U54"
	// 云市场分配的密钥Key
	secretKey := "JGk1CAofa3EY9MJq4X92kTm6EmXTsr716l2TEVRm"
	source := "market"

	// 签名
	auth, datetime, _ := calcAuthorization(source, secretId, secretKey)

	// 请求方法
	method := "GET"
	// 请求头
	headers := map[string]string{"X-Source": source, "X-Date": datetime, "Authorization": auth}

	// 查询参数
	queryParams := make(map[string]string)
	queryParams["content"] = "习他爸爸"
	// body参数
	bodyParams := make(map[string]string)

	// url参数拼接
	url := "https://service-owyxx71r-1255468759.ap-shanghai.apigateway.myqcloud.com/release/sensitiveWords"
	if len(queryParams) > 0 {
		url = fmt.Sprintf("%s?%s", url, urlencode(queryParams))
	}

	bodyMethods := map[string]bool{"POST": true, "PUT": true, "PATCH": true}
	var body io.Reader = nil
	if bodyMethods[method] {
		body = strings.NewReader(urlencode(bodyParams))
		headers["Content-Type"] = "application/x-www-form-urlencoded"
	}

	client := &http.Client{
		Timeout: 5 * time.Second,
	}
	request, err := http.NewRequest(method, url, body)
	if err != nil {
		panic(err)
	}
	for k, v := range headers {
		request.Header.Set(k, v)
	}
	response, err := client.Do(request)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	bodyBytes, err := ioutil.ReadAll(response.Body)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(bodyBytes))

	ss1 := &s1{}
	json.Unmarshal(bodyBytes, ss1)
	fmt.Println(ss1)
}

func calcAuthorization(source string, secretId string, secretKey string) (auth string, datetime string, err error) {
	timeLocation, _ := time.LoadLocation("Etc/GMT")
	datetime = time.Now().In(timeLocation).Format("Mon, 02 Jan 2006 15:04:05 GMT")
	signStr := fmt.Sprintf("x-date: %s\nx-source: %s", datetime, source)

	// hmac-sha1
	mac := hmac.New(sha1.New, []byte(secretKey))
	mac.Write([]byte(signStr))
	sign := base64.StdEncoding.EncodeToString(mac.Sum(nil))

	auth = fmt.Sprintf("hmac id=\"%s\", algorithm=\"hmac-sha1\", headers=\"x-date x-source\", signature=\"%s\"",
		secretId, sign)

	return auth, datetime, nil
}

func urlencode(params map[string]string) string {
	var p = gourl.Values{}
	for k, v := range params {
		p.Add(k, v)
	}
	return p.Encode()
}

func f1() {
	filterText := `我是需要过滤的33内容，内容为：**，需要过22滤。。。`
	memStore, err := store.NewMemoryStore(store.MemoryConfig{
		DataSource: []string{"文件", "11"},
	})
	if err != nil {
		panic(err)
	}
	err = memStore.Write("22")
	if err != nil {
		return
	}
	filterManage := filter.NewDirtyManager(memStore)
	result, err := filterManage.Filter().Filter(filterText, '*', '@')
	if err != nil {
		panic(err)
	}
	fmt.Println(result)
	if len(result) != 0 {
		fmt.Println("有敏感词")
	}
}
