import "time"

type Person struct {
	ID      int       `meddler:"id,pk"`
	Name    string    `meddler:"name"`
	Age     int       //自动转换为 `meddler:"Age"`
	salary  int       //不导出 meddler 无法看到它。 将被忽略。
	Created time.Time `meddler:"created,localtime"`
	Closed  time.Time `meddler:",localtimez"`
}