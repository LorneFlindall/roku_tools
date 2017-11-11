package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
	"os"
	"strings"
	"time"

	"github.com/parnurzeal/gorequest"
)

func main() {

	ipAddress := flag.String("ip", "ipaddress", "-ip=10.1.13.221")
	appID := flag.String("launch", "", "-launch=app_id found running -query=true i.e.dev, 2213 ")
	sendText := flag.String("send", "false", "-send=text string to enter ")
	query := flag.String("query", "false", "-query=true to query channels on device")
	homePress := flag.String("home", "false", "-home=true to go home ")
	keyPress := flag.String("key", "false", "-key=KeyName i.e.Rev,Fwd,Play,Select,Left,Right,Down,Up,Back,InstantReplay,Info,Backspace,Search,Enter")

	flag.Parse()

	for i, arg := range os.Args {
		// print index and value
		fmt.Println("item", i, "is", arg)
	}
	if *ipAddress == "ipaddress" {
		fmt.Println("-ip=ipaddress target not set")
	} else {

		if *query == "false" {
			fmt.Println("-query target not set")
		} else {
			queryApps(ipAddress)
		}
		if *homePress == "false" {
			fmt.Println("-homePress target not set")
		} else {
			home(ipAddress)
		}
		if *appID == "" {
			fmt.Println("-appID= target not set, using default")
		} else {
			launchApp(ipAddress, appID)
		}
		if *sendText == "false" {
			fmt.Println("-sendText=text target not set")
		} else {
			send(ipAddress, sendText)
		}
		if *keyPress == "false" {
			fmt.Println("-keyPress= target not set")
		} else {
			key(ipAddress, keyPress)
		}
	}
}

func queryApps(ipAddress *string) {
	ipValue := *ipAddress
	url := "http://" + ipValue + ":8060/query/apps"
	fmt.Println("URL:>", url)

	resp, errs := http.Get("http://" + ipValue + ":8060/query/apps")

	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	if resp.StatusCode != http.StatusOK {
		fmt.Println("Status error: ", resp.StatusCode)
	}

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Read body:", err)
	}

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Body Data:", string(data))

}

func launchApp(ipAddress *string, appID *string) {
	ipValue := *ipAddress
	appIDVal := *appID
	url := "http://" + string(ipValue) + ":8060/launch/" + string(appIDVal)
	fmt.Println("URL:>", url)
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Set("X-Custom-Header", "myvalue").
		End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", body)

}

func home(ipAddress *string) {
	ipValue := *ipAddress
	url := "http://" + ipValue + ":8060/keypress/home"
	fmt.Println("URL:>", url)
	request := gorequest.New()
	resp, body, errs := request.Post(url).
		Set("X-Custom-Header", "myvalue").
		End()
	if errs != nil {
		fmt.Println(errs)
		os.Exit(1)
	}
	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	fmt.Println("response Body:", body)

}

func send(ipAddress *string, sendText *string) {
	ipValue := *ipAddress
	stringValue := *sendText

	a := []rune(stringValue)
	for i, r := range a {
		fmt.Printf("i%d r %c\n", i, r)
		t := url.QueryEscape(string(r))
		url := ("http://" + ipValue + ":8060/keypress/Lit_" + string(t))
		fmt.Println("URL:>", url)
		request := gorequest.New()

		resp, body, errs := request.Post(url).
			Set("X-Custom-Header", "myvalue").
			End()
		if errs != nil {
			fmt.Println(errs)
			os.Exit(1)
		}
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", body)
	}
}

func key(ipAddress *string, keyName *string) {
	ipValue := *ipAddress
	stringValue := *keyName
	fmt.Println("keypress:>", string(stringValue))

	// Split on comma.
	result := strings.Split(*keyName, ",")

	// Display all elements.
	for i := range result {
		fmt.Println("results", result[i])
		url := ("http://" + ipValue + ":8060/keypress/" + result[i])
		fmt.Println("URL:>", url)
		request := gorequest.New()

		resp, body, errs := request.Post(url).
			Set("X-Custom-Header", "myvalue").
			End()
		if errs != nil {
			fmt.Println(errs)
			os.Exit(1)
		}
		fmt.Println("response Status:", resp.Status)
		fmt.Println("response Headers:", resp.Header)
		fmt.Println("response Body:", body)
		time.Sleep(time.Second * 1)
	}

}
