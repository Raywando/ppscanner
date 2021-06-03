package main

import (
        "context"
        "regexp"
        "os"
        "bufio"
        "sync"
        "fmt"
        "strings"
        "github.com/chromedp/chromedp"
)

func getProperty(query string) string {
        re := regexp.MustCompile(`\d`)
        property := re.ReplaceAllString(query, "")

        property = strings.ReplaceAll(property, "[]", "")
        property = strings.ReplaceAll(property, "[", ".")
        property = strings.ReplaceAll(property, "]", "")

        return property
}

func main() {
	var payload string = "?__proto__[5]=Raywando&__proto__[ALLOWED_ATTR][1]=Raywando&__proto__[Config][SiteOptimization][recommendationApiURL]=Raywando&__proto__[attrs][0][value]=Raywando&__proto__[crossDomain]=Raywando&__proto__[dataType]=Raywando&__proto__[delegateTarget]=Raywando&__proto__[div][1]=Raywando&__proto__[div][2]=Raywando&__proto__[handleObj]=Raywando&__proto__[is]=Raywando&__proto__[jquery]=Raywando&__proto__[name]=Raywando&__proto__[onerror][]=Raywando&__proto__[script][1]=Raywando&__proto__[script][2]=Raywando&__proto__[src]=Raywando&__proto__[src][]=Raywando&__proto__[template][innerHTML]=Raywando&__proto__[template][nodeType]=Raywando&__proto__[url]=Raywando&__proto__[whiteList][img][1]=Raywando&__proto__[xxx]=Raywando&__proto__.array=Raywando&__proto__.test=Raywando&__proto__[4]=Raywando&__proto__[ALLOWED_ATTR][0]=Raywando&__proto__[BOOMR]=Raywando&__proto__[CLOSURE_BASE_PATH]=Raywando&__proto__[Config][SiteOptimization][enabled]=Raywando&__proto__[attrs][0][name]=Raywando&__proto__[attrs][src]=Raywando&__proto__[context]=Raywando&__proto__[data]=Raywando&__proto__[div][0]=Raywando&__proto__[div][intro]=Raywando&__proto__[documentMode]=Raywando&__proto__[hif][]=Raywando&__proto__[innerHTML]=Raywando&__proto__[innerText]=Raywando&__proto__[onerror]=Raywando&__proto__[onload]=Raywando&__proto__[preventDefault]=Raywando&__proto__[props][]=Raywando&__proto__[script][0]=Raywando&__proto__[sourceURL]=Raywando&__proto__[src]=Raywando&__proto__[src][]=Raywando&__proto__[srcdoc][]=Raywando&__proto__[tagName]=Raywando&__proto__[template]=Raywando&__proto__[test]=Raywando&__proto__[test]=Raywando&__proto__[url]=Raywando&__proto__[url][]=Raywando&__proto__[whiteList][img][0]=Raywando&__proto__[xxx]=Raywando"

        sc := bufio.NewScanner(os.Stdin)

        var wg sync.WaitGroup
        urls := make(chan string)

	for i := 0; i < 10; i++ {
		wg.Add(1)
                go func() {
			for url := range urls {

				fmt.Printf("Testing: %s\n", url)

				fullUrl := url+payload

				//fmt.Printf("Input url: %s\n", url)
				//fmt.Printf("Full URL: %s\n", fullUrl)

				r, _ := regexp.Compile("__proto__[^=]+")
				queries := r.FindAllString(fullUrl, -1)
				ctx, cancel := chromedp.NewContext(context.Background())

				// run task list
				var res []byte
				chromedp.Run(ctx,
					chromedp.Navigate(fullUrl),
					chromedp.ActionFunc(func(ctx context.Context) error {
					for _, query := range queries {
						property := getProperty(query)

						chromedp.Evaluate(property, &res).Do(ctx)

						if strings.Contains(string(res), "Raywando"){
							fmt.Printf("POLLUTED - %s?%v=Raywando\n", url, query)
						}

						/*if err2 != nil {
							fmt.Printf("error in ActionFunc: %s\n", err2)
						} else {
							fmt.Printf("Property %s outputs: %v\n", property, string(res))
						}*/
					}
					return nil
					}),
				)
				cancel()

				/*
				if err != nil {
					log.Fatal(err)
				}*/
			}
			wg.Done()
		}()
	}

	for sc.Scan() {
                url := sc.Text()
                urls <- url
        }

        close(urls)
        wg.Wait()
	fmt.Printf("Done!\n")
}
