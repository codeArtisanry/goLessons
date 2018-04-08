package main

import (
	"context"
	"log"

	cdp "github.com/knq/chromedp"
)

func main() {
	var err error

	// create context
	ctxt, cancel := context.WithCancel(context.Background())
	defer cancel()

	// create chrome instance
	c, err := cdp.New(ctxt)
	if err != nil {
		log.Fatal(err)
	}

	// run task list
	// var site, res string
	err = c.Run(ctxt, scrapeSite())
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

	log.Println("Program has successfully run.")
}

func scrapeSite() cdp.Tasks {
	return cdp.Tasks{
		cdp.Navigate("https://bloc-chat-8c229.firebaseapp.com/"),
		// cdp.Sleep(1 * time.Second),
		cdp.SendKeys(`input`, "Mr. Scrape", cdp.ByQuery),
		// cdp.Sleep(1 * time.Second),
		cdp.Click(`button.btn.btn-primary`, cdp.ByQuery),
		cdp.Click(`a.room-link.ng-binding`, cdp.ByQuery),
		// cdp.Sleep(1 * time.Second),
		cdp.SendKeys(`input`, "Mr. Scrape was here", cdp.ByQuery),
		// cdp.Sleep(1 * time.Second),
	}
}
