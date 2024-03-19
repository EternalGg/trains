package bug

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gocolly/colly/extensions"
	"io"
	"net/http"
	"os"
	"path"
	"regexp"
)

import (
	"github.com/PuerkitoBio/goquery"
)

var t = "(http|https|ftp)://((((25[0-5])|(2[0-4]\\d)|(1\\d{2})|([1-9]?\\d)\\.){3}((25[0-5])|(2[0-4]\\d)|(1\\d{2})|([1-9]?\\d)))|(([\\w-]+\\.)+(net|com|org|gov|edu|mil|info|travel|pro|museum|biz|[a-z]{2})))(/[\\w\\-~#]+)*(/[\\w-]+\\.[\\w]{2,4})?([\\?=&%_]?[\\w-]+)*"

func Bug() {
	//url := "https://bad.news/search/q-%E8%A3%B8%E8%88%9E/page-2"
	url := "www.baidu.com"
	dom, _ := goquery.NewDocument(url)
	str, _ := dom.Html()
	fmt.Println(str)
	//dom.Find(".video").Each(func(i int, selection *goquery.Selection) {
	//	href, _ := selection.Attr("Poster")
	//	text := selection.Text()
	//	fmt.Println(i, url+href, text)
	//})
}

func Colly() {
	c := colly.NewCollector()

	extensions.RandomUserAgent(c)
	//extensions.Referrer(c)
	c.SetProxy("http://127.0.0.1:7890/")
	// 遍历所有的a标签，如果有href属性就回调

	// 在请求之前打印“Visiting xxx”
	c.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())

	})

	c.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		// Print link
		b, _ := regexp.MatchString(t, link)
		if b && len(link) > 45 {
			fmt.Println(link)

		}
	})
	c.Visit("https://bad.news/search/q-%E8%B7%B3%E8%88%9E/type-/page-25")

}

func query(str string) bool {
	if str == "https" {
		return true
	} else {
		return false
	}
}

func DownloadFile(filepath string, url string) error {

	// Get the data
	resp, err := http.Get(url)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}

func DownloadFile2(url string) error {

	// Get the data

	resp, err := http.Get(url)

	if err != nil {
		return err
	}
	defer resp.Body.Close()

	// Create the file
	filepath := path.Base(resp.Request.URL.String())
	out, err := os.Create(filepath)
	if err != nil {
		return err
	}
	defer out.Close()

	// Write the body to file
	_, err = io.Copy(out, resp.Body)
	return err
}
