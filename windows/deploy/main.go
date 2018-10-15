package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"sync"
)

//go:generate goversioninfo -icon=icon.ico

func createDeploy() {

	var data string = `#! /bin/sh
 echo "wo shi shui "`

	ioutil.WriteFile("deploy.sh", []byte(data), 0755)
}

func deployFrontend(wg *sync.WaitGroup) error {
	defer wg.Done()

	// 1. check frontend
	if Exists("lotus.tar.gz") {
		DialSSH(nginx_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)

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
	defer wg.Done()
	if Exists("lzkpv4.war") {
		DialSSH(tomcat_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)

		fmt.Println("[deployTomcat] Connect to server:", nginx_host)

		// 1. backup old
		DialSSH(nginx_host, `rm -f /docker/rollback/ROOT.war; mv /docker/bianban/lzkpv4/ROOT.war /docker/rollback/ROOT.war `)

		err := SCP(nginx_host, "./lzkpv4.war", "/docker/bianban/lzkpv4/ROOT.war")
		if err != nil {
			fmt.Println(err)
			return err
		}
		DialSSH(nginx_host, `sh ~/lzkpv4/deploy.sh tomcat;`)
		DialSSH(tomcat_host, `systemctl restart tomcat8`)
		os.Mkdir("del", 0755)
		os.Rename("lzkpv4.war", "del/lzkpv4.war")
	}
	// 1. check frontend
	if Exists("ROOT.war") {
		DialSSH(tomcat_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)

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
func deployLogplan(wg *sync.WaitGroup) error {

	defer wg.Done()
	var cmd = `sed -i -e "/logPlan.enable/s/^##//" \
	-e "s|.dbServerName=.*|.dbServerName=15.14.12.152:3306|g" \
	-e "s|sysDbConfig.dbName=.*|sysDbConfig.dbName=shizhi|g" \
	-e "s|dbConfigs.DbConfig\[0\].dbName=.*|dbConfigs.DbConfig[0].dbName=shizhi|g" \
	-e "s|.dbUser=.*|.dbUser=lzkp|g" \
	-e "s|.dbPws=.*|.dbPws=yqhtfjzm|g" \
	-e "s|redis.host=.*|redis.host=15.14.12.154|g" \
	-e "s|redis.password=.*|redis.password=hangruan2018|g" \
	-e "s|redis.database=.*|redis.database=0|g" \
	-e "s|fileupload.PrivaPath=.*|fileupload.PrivaPath=logplan/@1\!now\!yyyyMMdd@|g" /docker/tomcat/webapps/logplan/WEB-INF/classes/utility/hrlzkp.properties
`
	if Exists("logplan.war") {
		fmt.Println("[logplan] Connect to server:", lzkpbi_host)

		DialSSH(lzkpbi_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)
		err := SCP(lzkpbi_host, "logplan.war", "/docker/update/logplan.war")
		if err != nil {
			fmt.Println(err)
			return err
		}
		DialSSH(lzkpbi_host, `rm -fr /docker/tomcat/webapps/logplan/;unzip /docker/update/logplan.war -d /docker/tomcat/webapps/logplan/;`)
		DialSSH(lzkpbi_host, cmd)
		DialSSH(lzkpbi_host, `systemctl restart tomcat8`)
		os.Mkdir("del", 0755)
		os.Rename("logplan.war", "del/logplan.war")

	}
	return nil
}
func deployBIfront(wg *sync.WaitGroup) error {
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

	// sed -i 's/15.14.12/192.168.5' /docker/tomcat/webapps/lzkpbi/WEB-INF/classes/utility/lzkpbi.properties
	// 1. check frontend
	if Exists("lzkpbi.war") {
		fmt.Println("[deployLzkpbi] Connect to server:", lzkpbi_host)

		fmt.Println("Deploy lzkpbi !")
		DialSSH(lzkpbi_host, `mkdir -p /docker/update/ /docker/tomcat/webapps/ /docker/bianban/lzkpv4/ /docker/bianban/backendv4/ /docker/rollback/`)

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
	return err == nil || os.IsExist(err)
}

func main() {
	fmt.Println("start deploy!")
	var wg sync.WaitGroup
	wg.Add(7)
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
	//
	go deployLogplan(&wg)
	// etl.zip
	go deployetl(&wg)
	// DBExecAll()

	wg.Wait()
	fmt.Println("Finish deploy!")
	fmt.Scanln()

}
