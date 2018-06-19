package main

import (
	"net/http"
	"os/exec"
	"runtime"
	"time"
)

var err error

//go:generate goversioninfo -icon=icon.ico
func main() {
	// dir, _ := homedir.Dir()

	// config, err := ioutil.ReadFile(dir + "config")
	// if err != nil {
	// 	log.Fatal(err)
	// }

	var c = &http.Client{}
	c.Timeout = 500 * time.Millisecond
	resp, _ := c.Head("https://baidu.com")
	// if err != nil {
	// 	fmt.Println("err:", err)
	// }
	if resp != nil && resp.StatusCode == 200 {
		// fmt.Println("外网")
		openBrowser("https://lzkp.lybb.gov.cn")
	} else {
		// fmt.Println("内网")
		openBrowser("http://15.14.12.150")
	}
	// if resp.StatusCode == 200 {
	// 	err := openBrowser("https:fenwickelliott.io")
	// 	if err != nil {
	// 		fmt.Println(err)
	// 	}
	// }
	// fmt.Println(resp)
}

// Start tries to open the URL in a browser.
func openBrowser(url string) error {
	var args []string
	var cmd *exec.Cmd
	switch runtime.GOOS {
	case "darwin":
		args = []string{"open"}
	case "windows":
		args = []string{"cmd", "/c", "start"}
		cmd = exec.Command(args[0], append(args[1:], url)...)
	default:
		args = []string{"xdg-open"}

	}
	cmd = exec.Command(args[0], append(args[1:], url)...)
	hideWindow(cmd)
	return cmd.Start()
}
