package main

import (
	"flag"
	"fmt"
	"github.com/prometheus/common/expfmt"
	"io"
	//"io/ioutil"
	"log"
	"net/http"
	"time"
)

func get_metrics(url string) io.Reader {
	timeout := time.Duration(60 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, newReqErr := http.NewRequest("GET", url, nil)
	if newReqErr != nil {
		log.Fatal("http.NewRequest error:", newReqErr)
	}

	res, getErr := client.Do(req)
	defer res.Body.Close()
	if getErr != nil {
		log.Fatal("client.Do error:", getErr)
	}

	//body, readErr := ioutil.ReadAll(res.Body)
	//if readErr != nil {
	//	log.Fatal("ioutil.ReadAll error:", readErr)
	//}
	//fmt.Println(string(body))

	return res.Body
}

func retriveAddr() string {
	proto := flag.String("proto", "http", "specify protocol used for scrape")
	host := flag.String("host", "localhost", "specify the host to scrape")
	port := flag.String("port", "9100", "specify the port to scrape")
	flag.Parse()
	addr := fmt.Sprintf("%s://%s:%s/metrics", *proto, *host, *port)
	return addr
}

func main() {
	addr := retriveAddr()
	rawMetrics := get_metrics(addr)
	parser := expfmt.TextParser{}
	metrics, err := parser.TextToMetricFamilies(rawMetrics)
	if err != nil {
		log.Fatal("parser.TextToMetricFamilies error:", err)
	}
	fmt.Println(metrics)
}
