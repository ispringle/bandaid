package main

import (
	"flag"
	"fmt"
	dto "github.com/prometheus/client_model/go"
	"github.com/prometheus/common/expfmt"
	"io"
	//"io/ioutil"
	"net/http"
	"time"
)

func get_metrics(url string) (io.Reader, error) {
	timeout := time.Duration(60 * time.Second)
	client := &http.Client{
		Timeout: timeout,
	}

	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	res, getErr := client.Do(req)
	if getErr != nil {
		return nil, getErr
	}
	defer res.Body.Close()

	//body, readErr := ioutil.ReadAll(res.Body)
	//if readErr != nil {
	//	return nil, readErr
	//}
	//fmt.Println(body)

	return res.Body, nil
}

func parseMetrics(raw io.Reader) (expfmt.Decoder, error) {
	return expfmt.NewDecoder(raw, "text"), nil
}

func parseMetricsAlso(raw io.Reader) (map[string]*dto.MetricFamily, error) {
	var parser expfmt.TextParser
	metrics, metricErr := parser.TextToMetricFamilies(raw)
	if metricErr != nil {
		return nil, metricErr
	}
	return metrics, nil
}

func retriveAddr() (string, error) {
	proto := flag.String("proto", "http", "specify protocol used for scrape")
	host := flag.String("host", "localhost", "specify the host to scrape")
	port := flag.String("port", "9100", "specify the port to scrape")
	flag.Parse()
	addr := fmt.Sprintf("%s://%s:%s", *proto, *host, *port)
	return addr, nil
}

func main() {
	addr, flagErr := retriveAddr()
	if flagErr != nil {
		fmt.Println(flagErr)
	}

	rawMetrics, err := get_metrics(addr)
	if err != nil {
		fmt.Println(err)
	}

	metrics := dto.MetricFamily{}
	parser, parseErr := parseMetrics(rawMetrics)
	if parseErr != nil {
		fmt.Println(parseErr)
	}

	decodeErr := parser.Decode(&metrics)
	if decodeErr != nil {
		fmt.Println(decodeErr)
	}
	fmt.Println(metrics)

	textMetrics, parseErr := parseMetricsAlso(rawMetrics)
	if parseErr != nil {
		fmt.Println(parseErr)
	}
	fmt.Println(textMetrics)
}
