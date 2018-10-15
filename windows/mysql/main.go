package main

import (
	"bytes"
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"time"

	"github.com/rinetd/go-learning/windows/mysql/asset"
	"github.com/rs/zerolog"
)

var (
	// dbhost = flag.String("host", "192.168.5.100", "数据库地址")
	dbhost     = flag.String("h", "15.14.12.152", "数据库地址")
	dbport     = flag.String("P", "3306", "端口")
	dbuser     = flag.String("u", "lzkp", "数据库用户名")
	dbpassword = flag.String("p", "yqhtfjzm", "端口")
	force      = flag.Bool("f", false, "强制执行")
	// bindata    = flag.Bool("bindata", false, "是否执行内部文件")
)

// logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
var Log zerolog.Logger

func main() {
	file, err := os.OpenFile("log_"+time.Now().Format("2006-01-02")+".log", os.O_CREATE|os.O_APPEND|os.O_RDWR, 0755)
	// file, err := os.Create("log_" + time.Now().Format("2006-01-02") + ".log")
	if err != nil {
		fmt.Println("create file err")
	}
	defer file.Close()
	Log = zerolog.New(file).With().Timestamp().Str("", "").Logger()
	Log = zerolog.New(zerolog.ConsoleWriter{
		// Out: os.Stdout,
		Out:     file,
		NoColor: true,
	}).With().Timestamp().Str("", "").Logger()

	// Log.Info().Msg("aa")
	Log.Info().Msg("----------------------------")

	flag.Parse()
	InitDB()
	fmt.Println("3. 正在执行的文件:")
	ParseGlob("**.sql")
	ParseGlob("**/**.sql")
	// if *bindata {
	// 	RunAssets(Asset_files)
	// }
	fmt.Println("4. 执行完成!\n按任意键退出！")
	// ioutil.WriteFile("执行成功", []byte(time.Now().String()), 0755)
	Log.Debug().Msg("执行完成！")
	fmt.Scanln()
}

func ParseGlob(filename string) (s []string) {

	// files, _ := filepath.Glob("*.sql")
	files, _ := filepath.Glob(filename)
	for _, f := range files {
		fmt.Println(f)
	}
	println()
	if *force {
		fmt.Println("文件加载完毕！是否继续执行?")
		fmt.Scanln()
	}
	return RunFiles(files)
}

func RunFiles(files []string) (s []string) {
	if files == nil {
		return nil
	}
	for _, f := range files {
		fmt.Println(f)
		cc := ParseFile(f)
		base := filepath.Base(f)
		dbs := strings.Split(base, "_")

		var run = func(sdb string) (s []string) {
			fmt.Print("           Run on db :", sdb)
			Log.Debug().Msg("Run on db : " + sdb)
			if *force {
				fmt.Println("是否继续执行?")
				fmt.Scanln()
			}
			s = append(s, sdb)
			for _, c := range cc {

				// Log.Info().Msg(c)

				if _, err := mysqlMap[sdb].Exec(c); err != nil {
					fmt.Println(err)
					Log.Error().Msg(err.Error())
					Log.Error().Msg(c)
				}

			}
			fmt.Println(" 导入完成")
			return s
		}
		for _, fild := range dbs {
			if fild == "all" {
				dbs = dbnames
				break
			}
			if fild == "bi" {
				// dbs = []string{"bi"}
				s = run("bi")
			}
		}

		for _, db := range dbs {
			for _, sdb := range dbnames {
				if sdb == db {
					s = run(db)
				}
			}
		}
	}
	return s
}

func RunAssets(files []string) (s []string) {
	if files == nil {
		return nil
	}
	for _, f := range files {
		fmt.Println(f)
		bs, err := asset.Asset(f) // 根据地址获取对应内容
		if err != nil {
			fmt.Println("asset:", err)
			return
		}
		br := bytes.NewReader(bs)
		cc := Parsesql(br)

		dbs := strings.Split(f, "_")
		for _, fild := range dbs {
			if fild == "all" {
				dbs = dbnames
				break
			}
		}

		for _, db := range dbs {
			for _, sdb := range dbnames {
				if sdb == db {

					fmt.Println("Run on db:", sdb)
					s = append(s, sdb)
					for _, c := range cc {
						if _, err := mysqlMap[db].Exec(c); err != nil {
							fmt.Println(err)
						}
					}
				}
			}
		}
	}
	return s
}
