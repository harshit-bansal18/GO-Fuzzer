package main

import (
	"fmt"
	"bufio"
	"net/http"
	"log"
	"sync"
	"regexp"
	"os"

	"github.com/fatih/color"

)

type pgStatus struct {
	URL string
	status int
}
func (p pgStatus) toString() string {
	return fmt.Sprintf("[%d] %s\n", p.status, p.URL)
}

func checkUrl(baseUrl string, urlChan chan string, urlOk *[]pgStatus, WG *sync.WaitGroup){
	defer WG.Done()
	path := <- urlChan
	re := regexp.MustCompile("^/")
	path = re.ReplaceAllString(path, "")
	fullUrl := fmt.Sprintf("%s%s", baseUrl, path)
	resp, err := http.Get(fullUrl)
	if err != nil {
		fmt.Printf("[Error] %s\n%s\n", fullUrl, err)
		return
	}
	if resp.StatusCode >= 200 && resp.StatusCode < 400 {
		*urlOk = append(*urlOk, pgStatus{fullUrl, resp.StatusCode})
		color.Green(fmt.Sprintf("[%d] %s\n", resp.StatusCode, fullUrl))
	} else {
		color.Red(fmt.Sprintf("[%d] %s\n", resp.StatusCode, fullUrl))
	}
}

func main(){
	fmt.Printf("Welcome to Go-Fuzzer!!\n\n")
	WG := new(sync.WaitGroup)
	params := ParseOptions(os.Args)
	urlChan := make(chan string, params.nThreads)
	urlOk := [] pgStatus{}
	f, err := os.Open(params.wl)
	if(err != nil){
		log.Fatalln("Can't open wordlist file!")

	}
	defer f.Close()
	scanner := bufio.NewScanner(f)
	for scanner.Scan() {
		urlChan <- scanner.Text()
		WG.Add(1)
		go checkUrl(params.URL, urlChan, &urlOk, WG)
	}

	WG.Wait()
	color.Blue("\nFound URLs: (%d)\n", len(urlOk))
	for _, u := range urlOk {
		fmt.Printf("%s\n", u.toString())
	}
}

