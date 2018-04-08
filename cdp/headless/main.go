package main

import (
	"context"
	"io/ioutil"
	"log"

	cdp "github.com/knq/chromedp"
	cdptypes "github.com/knq/chromedp/cdp"
	"github.com/knq/chromedp/client"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome
	c, err := cdp.New(ctxt, cdp.WithTargets(client.New().WatchPageTargets(ctxt)), cdp.WithLog(log.Printf))
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	var site, res string
	err = c.Run(ctxt, googleSearch("site:brank.as", "Easy Money Management", &site, &res))
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("saved screenshot of #testimonials from search result listing `%s` (%s)", res, site)
}

func googleSearch(q, text string, site, res *string) cdp.Tasks {
	var buf []byte
	// sel := fmt.Sprintf(`//a[text()[contains(., '%s')]]`, text)
	return cdp.Tasks{
		cdp.Navigate(`https://www.baidu.com`),
		// cdp.Sleep(2 * time.Second),
		//*[@id="kw"]#kw
		// cdp.WaitVisible(`#kw`, cdp.ByID),
		// cdp.SendKeys(`#kw`, q+"\n", cdp.ByID),
		// cdp.WaitVisible(`.c-title-en > a:nth-child(1)`, cdp.ByID),

		// cdp.Text(sel, res),
		// cdp.Click(sel),
		// cdp.Sleep(2 * time.Second),
		// cdp.WaitVisible(`#footer`, cdp.ByQuery),
		// cdp.WaitNotVisible(`div.v-middle > div.la-ball-clip-rotate`, cdp.ByQuery),
		// cdp.Location(site),
		cdp.Screenshot(`#lg`, &buf, cdp.ByID),
		cdp.ActionFunc(func(context.Context, cdptypes.Handler) error {
			return ioutil.WriteFile("testimonials.png", buf, 0644)
		}),
	}
}
