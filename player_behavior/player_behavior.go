package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"path"
	"regexp"
	"strconv"
	"strings"
	"time"
	"unsafe"
)

// 无内存拷贝的字符串转 bytes
func str2bytes(s string) []byte {
	x := (*[2]uintptr)(unsafe.Pointer(&s))
	b := [3]uintptr{x[0], x[1], x[1]}
	return *(*[]byte)(unsafe.Pointer(&b))
}

func main() {
	f5()
}

func f5() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	file, err := os.Open(path.Join(wd, "logic_1-baseinfo.log"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uidMap := map[int][]*cmdAndTime{}
	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		if strings.Contains(line, `cmd:`) && strings.Contains(line, `seq`) && strings.Contains(line, `gate`) {
			uid, _ := strconv.Atoi(line[strings.Index(line, `uid`)+4 : strings.Index(line, `uid`)+10])
			fmt.Print(line, "\t")
			fmt.Print(uid, "\t")
			if _, ok := uidMap[uid]; !ok {
				uidMap[uid] = []*cmdAndTime{}
			}
			var cmd int
			cmd, err := strconv.Atoi(line[strings.Index(line, `cmd`)+4 : strings.Index(line, `cmd`)+9])
			if err != nil {
				cmd, err = strconv.Atoi(line[strings.Index(line, `cmd`)+4 : strings.Index(line, `cmd`)+8])
			}
			fmt.Print(cmd, "\t")
			year, err := strconv.Atoi(line[strings.Index(line, `t`)+4 : strings.Index(line, `t`)+8])
			if err != nil {
				return
			}
			month, err := strconv.Atoi(line[strings.Index(line, `t`)+9 : strings.Index(line, `t`)+11])
			if err != nil {
				return
			}
			day, err := strconv.Atoi(line[strings.Index(line, `t`)+12 : strings.Index(line, `t`)+14])
			if err != nil {
				return
			}
			hour, err := strconv.Atoi(line[strings.Index(line, `t`)+15 : strings.Index(line, `t`)+17])
			if err != nil {
				return
			}
			minute, err := strconv.Atoi(line[strings.Index(line, `t`)+18 : strings.Index(line, `t`)+20])
			if err != nil {
				return
			}
			second, err := strconv.Atoi(line[strings.Index(line, `t`)+21 : strings.Index(line, `t`)+23])
			if err != nil {
				return
			}
			fmt.Print("\t")
			fmt.Print(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local))
			fmt.Print(cmd, "\t")

			fmt.Println()
			fmt.Print(strings.Replace(line[strings.Index(line, `req`)+4 : strings.Index(line, `rsp`)-1],`\`,``,-1))
			fmt.Println()

			t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
			uidMap[uid] = append(uidMap[uid], &cmdAndTime{Cmd: cmd, Time: &t})
		}
	}
	fmt.Println(uidMap)
}

func f4() {
	s1 := "abc"
	fmt.Println(s1[2:3])
}

func f3() {
	s1 := `cmd:10001 seq:10 uid:106835"}`
	fmt.Println(strings.Index(s1, `cmd:`))
	fmt.Println(s1[strings.Index(s1, `seq`)+4:])
}

type cmdAndTime struct {
	Cmd  int
	Time *time.Time
}

func f2() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	file, err := os.Open(path.Join(wd, "logic_1-baseinfo.log"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	uidMap := map[int][]*cmdAndTime{}

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}
		if strings.Contains(line, `cmd:`) && strings.Contains(line, `seq`) && !strings.Contains(line, `gate`) {
			fmt.Print(line)
			//fmt.Print(line[strings.Index(line, `cmd`)+4 : strings.Index(line, `cmd`)+9])
			uid, _ := strconv.Atoi(line[strings.Index(line, `uid`)+4 : strings.Index(line, `uid`)+10])
			fmt.Print(uid, "\t")
			if _, ok := uidMap[uid]; !ok {
				uidMap[uid] = []*cmdAndTime{}
			}

			var cmd int
			cmd, err := strconv.Atoi(line[strings.Index(line, `cmd`)+4 : strings.Index(line, `cmd`)+9])
			if err != nil {
				cmd, err = strconv.Atoi(line[strings.Index(line, `cmd`)+4 : strings.Index(line, `cmd`)+8])
			}
			year, err := strconv.Atoi(line[strings.Index(line, `t`)+4 : strings.Index(line, `t`)+8])
			if err != nil {
				return
			}
			month, err := strconv.Atoi(line[strings.Index(line, `t`)+9 : strings.Index(line, `t`)+11])
			if err != nil {
				return
			}
			day, err := strconv.Atoi(line[strings.Index(line, `t`)+12 : strings.Index(line, `t`)+14])
			if err != nil {
				return
			}
			hour, err := strconv.Atoi(line[strings.Index(line, `t`)+15 : strings.Index(line, `t`)+17])
			if err != nil {
				return
			}
			minute, err := strconv.Atoi(line[strings.Index(line, `t`)+18 : strings.Index(line, `t`)+20])
			if err != nil {
				return
			}
			second, err := strconv.Atoi(line[strings.Index(line, `t`)+21 : strings.Index(line, `t`)+23])
			if err != nil {
				return
			}
			fmt.Print("\t")
			fmt.Print(time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local))
			fmt.Println()
			t := time.Date(year, time.Month(month), day, hour, minute, second, 0, time.Local)
			uidMap[uid] = append(uidMap[uid], &cmdAndTime{Cmd: cmd, Time: &t})
		}
	}
	//fmt.Println(uidMap)
}

func f1() {
	wd, err := os.Getwd()
	if err != nil {
		return
	}
	file, err := os.Open(path.Join(wd, "logic_1-baseinfo.log"))
	if err != nil {
		panic(err)
	}
	defer file.Close()

	stat, err := file.Stat()
	if err != nil {
		panic(err)
	}
	var size = stat.Size()
	fmt.Println("file size=", size)
	reg1 := regexp.MustCompile(`uid.......`)
	if reg1 == nil {
		panic("regexp err")
	}

	buf := bufio.NewReader(file)
	for {
		line, err := buf.ReadString('\n')
		line = strings.TrimSpace(line)
		strings.Replace(line, "\\", "", -1)
		if err != nil {
			if err == io.EOF {
				fmt.Println("File read ok!")
				break
			} else {
				fmt.Println("Read file error!", err)
				return
			}
		}

		fmt.Println(reg1.FindAllStringSubmatch(strings.Replace(strings.Replace(strings.Replace(strings.Replace(strings.Replace(line, `\`, ``, -1), `"'`, ``, -1), `{`, ``, -1), `"`, ``, -1), `$numberLong`, ``, -1), -1))
	}
}
