package main

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"

	"github.com/bramvdbogaerde/go-scp"
	"golang.org/x/crypto/ssh"

	_ "github.com/go-sql-driver/mysql" // mysql driver
)

//go:generate goversioninfo -icon=icon.ico
var config *ssh.ClientConfig

func init() {
	config = Genconfig()
}
func Genconfig() *ssh.ClientConfig {
	var User string = "root"
	var password string = "HR2018!!"
	// var connectKey string = "/home/ubuntu/.ssh/id_rsa"
	var connectKey string
	var permBytes []byte

	config := &ssh.ClientConfig{}
	if connectKey != "" {
		// Read PublicKey

		buffer, err := ioutil.ReadFile(connectKey)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:%s%n", err)
			return nil
		}
		if len(permBytes) != 0 {
			buffer = permBytes
		}
		key, err := ssh.ParsePrivateKey(buffer)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error:%s%n", err)
			return nil
		}

		// Create ssh client config for KeyAuth
		config = &ssh.ClientConfig{
			User: User,
			Auth: []ssh.AuthMethod{
				ssh.PublicKeys(key)},
			Timeout:         60 * time.Second,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	} else {
		// Create ssh client config for PasswordAuth
		config = &ssh.ClientConfig{
			User: User,
			Auth: []ssh.AuthMethod{
				ssh.Password(password)},
			Timeout:         60 * time.Second,
			HostKeyCallback: ssh.InsecureIgnoreHostKey(),
		}
	}
	return config
}

func DialSSH(hostport, cmdStr string) int {

	// New Connext create
	conn, err := ssh.Dial("tcp", hostport, config)
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot connect %v: %v \n", hostport, err)
		return 1
	}

	// New Session
	session, err := conn.NewSession()
	if err != nil {
		fmt.Fprintf(os.Stderr, "cannot open new session: %v \n", err)
		return 1
	}

	// go func() {
	// 	time.Sleep(2419200 * time.Second)
	// 	conn.Close()
	// }()

	session.Stdout = os.Stdout
	session.Stderr = os.Stderr

	err = session.Run(cmdStr)
	session.Close()
	conn.Close()
	return 0
}

func SCP(hostport, src, dst string) (err error) {

	// Create a new SCP client
	client := scp.NewClient(hostport, config)

	// Connect to the remote server
	err = client.Connect()
	if err != nil {
		fmt.Println("Couldn't establisch a connection to the remote server ", err)
		return err
	}

	// Open a file
	f, _ := os.Open(src)

	// Close session after the file has been copied
	defer client.Session.Close()

	// Close the file after it has been copied
	defer f.Close()

	// Finaly, copy the file over
	// Usage: CopyFile(fileReader, remotePath, permission)

	client.CopyFile(f, dst, "0644")
	return nil
}

func createDeploy() {

	var data string = `#! /bin/sh
 echo "wo shi shui "`

	ioutil.WriteFile("deploy.sh", []byte(data), 0755)
}

func createDir(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	fmt.Println("[createDir] Connect to server:", nginx_host)
	DialSSH(nginx_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)
	DialSSH(tomcat_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)
	DialSSH(lzkpbi_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)
}
func deployFrontend(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	// 1. check frontend
	if Exists("lotus.tar.gz") {
		fmt.Println("[deployFrontend] Connect to server:", nginx_host)
		err := SCP(nginx_host, "lotus.tar.gz", "/docker/update/lotus.tar.gz")
		if err != nil {
			return err
		}
		DialSSH(nginx_host, `rm -fr /docker/rollback/lotus; mv /docker/bianban/lotus /docker/rollback/;tar -zxf /docker/update/lotus.tar.gz -C /docker/bianban/`)
		os.Mkdir("del", 0755)
		os.Rename("lotus.tar.gz", "del/lotus.tar.gz")
	}
	return nil
}

func deployTomcat(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	// 1. check frontend
	if Exists("ROOT.war") {
		fmt.Println("[deployTomcat] Connect to server:", nginx_host)

		// 1. backup old
		DialSSH(nginx_host, `rm -f /docker/rollback/ROOT.war; mv /docker/bianban/lzkpv4/ROOT.war /docker/rollback/ROOT.war `)

		err := SCP(nginx_host, "./ROOT.war", "/docker/bianban/lzkpv4/ROOT.war")
		if err != nil {
			fmt.Println(err)
			return err
		}
		DialSSH(nginx_host, `sh ~/lzkpv4/deploy.sh tomcat;`)
		DialSSH(tomcat_host, `systemctl restart tomcat8`)
		os.Mkdir("del", 0755)
		os.Rename("ROOT.war", "del/ROOT.war")
	}
	return nil
}
func deployBackend(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	// 1. check frontend
	if Exists("backendv4.jar") {
		fmt.Println("[deployBackend] Connect to server:", nginx_host)

		DialSSH(nginx_host, "rm -f /docker/rollback/backendv4.jar; mv /docker/bianban/backendv4/backendv4.jar /docker/rollback/backendv4.jar")
		err := SCP(nginx_host, "backendv4.jar", "/docker/bianban/backendv4/backendv4.jar")
		if err != nil {
			return err
		}
		DialSSH(nginx_host, `sh ~/lzkpv4/deploy.sh backend`)
		os.Mkdir("del", 0755)
		os.Rename("backendv4.jar", "del/backendv4.jar")
	}
	return nil
}

func deployBIfront(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	if Exists("bi.tar.gz") {
		fmt.Println("[deployBIfront] Connect to server:", nginx_host)

		err := SCP(nginx_host, "bi.tar.gz", "/docker/update/bi.tar.gz")
		if err != nil {
			return err
		}
		DialSSH(nginx_host, `rm -fr /docker/rollback/bi; mv /docker/bianban/bi /docker/rollback/;tar -zxf /docker/update/bi.tar.gz -C /docker/bianban/`)
		os.Mkdir("del", 0755)
		os.Rename("bi.tar.gz", "del/bi.tar.gz")
	}
	return nil
}

func deployLzkpbi(wg *sync.WaitGroup) error {
	wg.Add(1)
	defer wg.Done()

	var cmdlzkpbi = `
sed -i -e "s|.dbServerName=.*|.dbServerName=15.14.12.152:3306|g" \
	-e "s|.dbName=.*|.dbName=lzkp_bi|g" \
	-e "s|.dbUser=.*|.dbUser=lzkp|g" \
	-e "s|.dbPws=.*|.dbPws=yqhtfjzm|g" \
	-e "s|redis.host=.*|redis.host=15.14.12.154|g" \
	-e "s|redis.password=.*|redis.password=hangruan2018|g" \
	-e "s|redis.database=.*|redis.database=0|g" \
	-e "s|fileupload.PrivaPath=.*|fileupload.PrivaPath=lzkpbi/@1\!now\!yyyyMMdd@|g" /docker/tomcat/webapps/lzkpbi/WEB-INF/classes/utility/lzkpbi.properties
`
	// 1. check frontend
	if Exists("lzkpbi.war") {
		fmt.Println("[deployLzkpbi] Connect to server:", lzkpbi_host)

		fmt.Println("Deploy lzkpbi !")
		// 1. backup old
		DialSSH(lzkpbi_host, `rm -f /docker/rollback/lzkpbi.war; mv /docker/update/lzkpbi.war /docker/rollback/lzkpbi.war `)

		err := SCP(lzkpbi_host, "lzkpbi.war", "/docker/update/lzkpbi.war")
		if err != nil {
			fmt.Println(err)
			return err
		}
		DialSSH(lzkpbi_host, `rm -fr /docker/tomcat/webapps/lzkpbi/;unzip /docker/update/lzkpbi.war -d /docker/tomcat/webapps/lzkpbi/;`+cmdlzkpbi)
		DialSSH(lzkpbi_host, `systemctl restart tomcat8`)
		os.Mkdir("del", 0755)
		os.Rename("lzkpbi.war", "del/lzkpbi.war")
	}
	return nil
}

func deployetl(wg *sync.WaitGroup) {
	wg.Add(1)
	defer wg.Done()

	if Exists("etl.zip") {
		fmt.Println("[deployetl] Connect to server:", lzkpbi_host)

		SCP(lzkpbi_host, "etl.zip", "/docker/bianban/etl.zip")

		DialSSH(lzkpbi_host, `rm -fr /root/etl; unzip /docker/bianban/etl.zip -d /root/;`)
		os.Mkdir("del", 0755)
		os.Rename("etl.zip", "del/etl.zip")
	}
}

// 判断所给路径文件/文件夹是否存在
func Exists(path string) bool {
	_, err := os.Stat(path) //os.Stat获取文件信息
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

// Load for loading data for testing
func MysqlLoad(file, dbname string) {
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		fmt.Errorf("Fail to connect database. %s", err.Error())
	}
	defer db.Close()
	db.Ping()

	_, err = db.Exec("USE " + dbname)
	if err != nil {
		fmt.Printf("[db.EXEC] switch databases %s -> %s : \n", dbname, err.Error())
		return
	}

	content, err := ioutil.ReadFile(file)
	if err != nil {
		fmt.Errorf("readfile error")
	}

	defer func() {
		if err := recover(); err != nil {
			fmt.Println("[EXEC RECOVER] error ", err)
		}
	}()

	queries := strings.Split(string(content), ";")
	// fmt.Println(queries)
	for _, query := range queries {
		if strings.TrimSpace(query) != "" {

			// fmt.Println(query)
			_, err := db.Exec(query)
			if err != nil {
				fmt.Errorf("[db.EXEC] %s -> %s : ", query, err.Error())
			}
			// a, _ := res.RowsAffected()
			// fmt.Println(a)
		}
	}
	fmt.Println("import ", file, " finish !")
}

var name = map[string]bool{
	"mengshan":  true,
	"mengyin":   true,
	"pingyi":    true,
	"shizhi":    true,
	"tancheng":  true,
	"yinan":     true,
	"yishui":    true,
	"feixian":   true,
	"gaoxinqu":  true,
	"hedong":    true,
	"jingkaiqu": true,
	"junan":     true,
	"luozhuang": true,
	"lanling":   true,
	"lanshan":   true,
	"lingang":   true,
	"linshu":    true,
}

func main() {

	var wg sync.WaitGroup
	go createDir(&wg)
	// lotus.tar.gz
	go deployFrontend(&wg)
	// ROOT.war
	go deployTomcat(&wg)
	// backendv4.jar
	go deployBackend(&wg)
	// bi.tar.gz
	go deployBIfront(&wg)
	// lzkpbi.war
	go deployLzkpbi(&wg)
	// etl.zip
	go deployetl(&wg)

	files, _ := filepath.Glob("*.sql")
	for _, f := range files {
		db := strings.Split(f, "_")
		if name[db[0]] {
			fmt.Println(f)
			MysqlLoad(f, db[0])
		}
	}
	wg.Wait()
}

var nginx_host = "15.14.12.150:22"
var tomcat_host = "15.14.12.151:22"
var lzkpbi_host = "15.14.12.153:22"

// var nginx_host = "118.190.117.250:3009"
// var tomcat_host = "118.190.117.250:3009"
// var lzkpbi_host = "118.190.117.250:3009"

// var dsn = `root:000000@tcp(192.168.5.100:3306)/?parseTime=true&loc=Local`

var dsn = `lzkp:yqhtfjzm@tcp(15.14.12.152:3306)/?parseTime=true&loc=Local`
