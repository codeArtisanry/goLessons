package bench

type user struct {
	ID       int    `xorm:"ID" db:"ID"`
	Username string `xorm:"Username" db:"Username"`
	Password string `xorm:"Password" db:"Password"`
}

var dsn = "bench:bench@118.190.83.129/bench?charset=utf8"
