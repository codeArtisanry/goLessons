package main

import (
	"context"
	"log"
	"time"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/runner"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt,
		// cdp.WithTargets(
		// 	client.New().WatchPageTargets(ctxt),
		// ),
		cdp.WithRunnerOptions(
			// runner.Headless(runner.DefaultChromePath, 9222),
			// runner.Flag("headless", true),
			// runner.UserAgent("baidu"),
			// runner.Proxy(`socks5://118.190.83.129:1080`),
			// runner.Flag("disable-gpu", true),
			// runner.Flag("proxy-server", "socks5://118.190.83.129:1080"),
			// runner.Flag("no-first-run", true),
			// runner.Flag("no-default-browser-check", true),
			// runner.Flag("hide-scrollbars", "true"),
			// runner.Flag("disable-web-security", true),
			// runner.Flag("window-size", "800,420"),
			runner.Flag("start-maximized", true),
		),
		cdp.WithLog(log.Printf),
	)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var val1, val2, val3, val4 string
	err = c.Run(ctxt, sendkeys(&val1, &val2, &val3, &val4))
	if err != nil {
		log.Fatal(err)
	}

	// shutdown chrome
	err = c.Shutdown(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// wait for chrome to finish
	err = c.Wait()
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("#output value: %s", html)

	log.Printf("#input1 value: %s", val1)
	log.Printf("#textarea1 value: %s", val2)
	log.Printf("#input2 value: %s", val3)
	log.Printf("#select1 value: %s", val4)
	for a := range ids {

		log.Println("Nodes", a)
	}
}

var urlstr = "chinaebr.com"

var html string
var ids []*cdptypes.Node

func sendkeys(val1, val2, val3, val4 *string) cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate("http://myip.ipip.net"),
		cdp.OuterHTML("/html/body/pre", &html, cdp.BySearch),
		cdp.Nodes("pre", &ids, cdp.ByQuery),
		cdp.Sleep(1 * time.Second),
		// cdp.Sleep(300 * time.Millisecond),
		// cdp.Title(&urlstr),

		// cdp.Navigate("file:" + os.Getenv("GOPATH") + "/src/github.com/knq/chromedp/testdata/visible.html"),
		// cdp.WaitVisible(`#input1`, cdp.ByID),
		// cdp.WaitVisible(`#textarea1`, cdp.ByID),
		// cdp.SendKeys(`#textarea1`, kb.End+"\b\b\n\naoeu\n\ntest1\n\nblah2\n\n\t\t\t\b\bother box!\t\ntest4", cdp.ByID),
		// cdp.Value(`#input1`, val1, cdp.ByID),
		// cdp.Value(`#textarea1`, val2, cdp.ByID),
		// cdp.SetValue(`#input2`, "test3", cdp.ByID),
		// cdp.Value(`#input2`, val3, cdp.ByID),
		// cdp.SendKeys(`#select1`, kb.ArrowDown+kb.ArrowDown, cdp.ByID),
		// cdp.Value(`#select1`, val4, cdp.ByID),

		// cdp.Sleep(100 * time.Second),
	}
}
