package bench

import (
	"database/sql"
	"fmt"
	"testing"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

/*type mysql struct {
	dbo sql.DB
}*/

func TestSetup(t *testing.T) {
	// mysql.RegisterDial()
	// start := time.Now()
	db, e := sql.Open("mysql", "bench:bench@tcp(192.168.5.100:3306)/bench?charset=utf8&multiStatements=true")
	err(e)
	const sqlstr = `
	CREATE TABLE IF NOT EXISTS USER (
		Uid  TEXT PRIMARY KEY,
	  Username  TEXT,
	  Age    TEXT
	);
`
	fmt.Println(db.Exec(`
	CREATE TABLE IF NOT EXISTS USERs (
		Uid  int(10) PRIMARY KEY
	);-- --aaaa
	-- aaa
	CREATE TABLE IF NOT EXISTS retross (
		Id        int(10) PRIMARY KEY
	  );
	`))
	if _, err := db.Exec(sqlstr); err != nil {
		// err(err)
		// fmt.Println(err)
	}
	db.Exec("use bench;")
	/*
	 * empty your database before insert
	 */

	/*
	 * uncomment below methods for insert, select, update for each cycle
	 */

	/*
	 * send no of rows
	 */
	//u := insert(db,1000000)
	//u := update(db,"test")
	//u := del(db)

	//log.Printf("\t%d\t%s\n",u,time.Since(start))

	// o := find(db)

	// log.Printf("\t%d\t%s\n", o, time.Since(start))
}

func insert(db *sql.DB, n int) int64 {
	s, e := db.Prepare("insert userinfo set username=?,departname=?,created=?")
	err(e)

	var u int64
	for i := 0; i < n; i++ {
		r, e := s.Exec("reoxey", "Web", time.Now())
		err(e)
		i, e := r.LastInsertId()
		err(e)

		u = i
	}
	return u
}

func update(db *sql.DB, d string) int64 {
	s, e := db.Prepare("update userinfo set username=?,departname=? where uid <=1")
	err(e)

	r, e := s.Exec(d, d)
	err(e)

	u, e := r.RowsAffected()
	err(e)

	return u
}

func find(db *sql.DB) int64 {
	s, e := db.Query("select * from userinfo where 1")
	err(e)

	var u int64
	u = 0
	for s.Next() {
		var uid int
		var usn string
		var dpt string
		var time string
		s.Scan(&uid, &usn, &dpt, &time)
		fmt.Print(uid)
		fmt.Print(" ", usn)
		fmt.Print(" ", dpt)
		fmt.Println(" ", time)
		u++
	}
	return u
}

func del(db *sql.DB) int64 {
	s, e := db.Prepare("delete from userinfo")
	err(e)
	r, e := s.Exec()
	err(e)

	u, e := r.RowsAffected()
	err(e)

	return u
}

func err(e error) {
	if e != nil {
		panic(e)
	}
}
