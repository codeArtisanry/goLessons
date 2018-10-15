package zlog

import (
	"os"
	"runtime"
	"time"

	"github.com/rs/zerolog"
)

/*
zerolog allows for logging at the following levels (from highest to lowest):
panic (zerolog.PanicLevel, 5)
fatal (zerolog.FatalLevel, 4)
error (zerolog.ErrorLevel, 3)
warn (zerolog.WarnLevel, 2)
info (zerolog.InfoLevel, 1)
debug (zerolog.DebugLevel, 0)
log.Fatal().
		Err(err).
		Str("c.Service", c.Service).
		Msgf("Cannot start %s", c.Service)
log.Info().
    Str("foo", "bar").
    Int("n", 123).
    Msg("hello world")
*/

var (
	Log        zerolog.Logger
	ConsoleLog zerolog.Logger
)

type Conf struct {
	Dir            string
	Service        string
	WriteToFile    bool
	WriteToConsole bool
}

func Run(c *Conf) (err error) {

	if c.WriteToConsole {
		ConsoleLog = zerolog.New(zerolog.ConsoleWriter{
			Out:     os.Stdout,
			NoColor: false,
		}).With().
			Timestamp().
			Str("service", c.Service).
			Logger()
	}

	if !c.WriteToFile {
		return
	}

	//----FILE WRITER
	t := time.Now()
	day := 86400000000000
	now := (day - (t.Hour() * 3600000000000)) - ((3600000000000 - t.Minute()*60000000000) - (60000000000 - t.Second()*1000000000))

	switch string(c.Dir[len(c.Dir)-1]) {
	case `\`:
	case `/`:
	default:
		switch runtime.GOOS {
		case "windows":
			c.Dir = c.Dir + `\`
		default:
			c.Dir = c.Dir + "/"
		}
	}
	// 目录不存在，创建
	if _, err = os.Stat(c.Dir); err != nil {
		err = os.MkdirAll(c.Dir, 0755)
		if err != nil {
			panic(err)
		}
	}

	for {

		file := &os.File{}
		file, err = os.Create(c.Dir + c.Service + "_" + t.Format("2006-01-02_15-04-05") + ".log")

		if err == nil {
			Log = zerolog.New(file).With().
				Timestamp().
				Str("service", c.Service).
				Logger()

			Log.Info()

		} else {
			panic(err)
		}

		Log.Info().Msg("ZLOG STARTING")

		time.Sleep(time.Duration(day - now))

		file.Close()
		t = time.Now()
		now = (day - (t.Hour() * 3600000000000)) - ((3600000000000 - t.Minute()*60000000000) - (60000000000 - t.Second()*1000000000))
	}

}
