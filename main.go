package main

import (
    "fmt"
    "net/http"
    "net/url"
    "strconv"
    "strings"
)

func sendPost1() {
    data := make(url.Values)
    data["name"] = []string{"rnben"}

    res, err := http.PostForm("http://127.0.0.1/tpost", data)
    if err != nil {
        fmt.Println(err.Error())
        return
    }
    defer res.Body.Close()
    fmt.Println("post send success")
}

func sendPost2(post_data string) {
	apiUrl := "http://127.0.0.1:1111/"
	resource := "/v1/push/packing"

	u, _ := url.ParseRequestURI(apiUrl)
	u.Path = resource
	urlStr := u.String() // "http://127.0.0.1/tpost"

	client := &http.Client{}
	r, _ := http.NewRequest("POST", urlStr, strings.NewReader(post_data)) // URL-encoded payload
	r.Header.Add("Content-Type", "application/json")
	r.Header.Add("accept", "application/json")
	//r.Header.Add("Content-Length", strconv.Itoa(len(data.Encode())))
	resp, err := client.Do(r)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	defer resp.Body.Close()
	fmt.Println(post_data)

}

func main() {
	var i int
	for i = 600000; i < 700000; i++ {
		outband := `{ "nodeInfo" : {"nodeId" : "【LG仓库】", "nodeRole" : "【LG仓库】", "nodeActorId": "", "nodeActorRole" : ""},
		"typ" : 1, "containerId" : "CK201911180008",
		"subIdList" : [ "9319932002169d45d78ea6e4b389745e", "3282779051477666208cc4e78517da32" ],
		"packingInfo" : "厂家装箱，等待出库`+ strconv.Itoa(i) +`", "time" : "2019-11-18 09:30:45"}`;
		sendPost2(outband)
	}

}
