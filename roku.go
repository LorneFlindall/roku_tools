package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"

	"github.com/parnurzeal/gorequest"
)

func main() {

	ipAddress := flag.String("ip", "ipaddress", "i.e. 10.1.13.221")
	appID := flag.String("launch", "app_id", "-launch=app_id found running -query=true i.e.dev, 2213 ")
	query := flag.String("query", "false", "-query=true to query channels on device")
	homePress := flag.String("home", "false", "=true to go home ")
	flag.Parse()

	for i, arg := range os.Args {
		// print index and value
		fmt.Println("item", i, "is", arg)
	}
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
	if *appID == "app_id" {
		fmt.Println("-appID target not set")
	} else {
		launchApp(ipAddress, appID)
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
	titleList := []string{"title1", "title2", "title3"}
	for _, title := range titleList {
		resp, body, errs := request.Post(url).
			Set("X-Custom-Header", "myvalue").
			Send(`{"title":"` + title + `"}`).
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

func home(ipAddress *string) {
	ipValue := *ipAddress
	url := "http://" + ipValue + ":8060/keypress/home"
	fmt.Println("URL:>", url)
	request := gorequest.New()
	titleList := []string{"title1", "title2", "title3"}
	for _, title := range titleList {
		resp, body, errs := request.Post(url).
			Set("X-Custom-Header", "myvalue").
			Send(`{"title":"` + title + `"}`).
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
