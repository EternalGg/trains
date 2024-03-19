package BrowserHistory

type BrowserHistory struct {
	UrlCache []string
	Now      int
	HomePage string
}

func Constructor(homepage string) BrowserHistory {
	browser := new(BrowserHistory)
	browser.UrlCache = make([]string, 1)
	browser.UrlCache[0] = homepage
	browser.Now = 1
	return *browser
}

func (this *BrowserHistory) Visit(url string) {
	if url == this.UrlCache[this.Now-1] {
		return
	} else {
		this.UrlCache = this.UrlCache[:this.Now]
		this.UrlCache = append(this.UrlCache, url)
		this.Now++
	}
}

func (this *BrowserHistory) Back(steps int) string {
	if steps-this.Now >= 0 {
		this.Now = 1
		return this.HomePage
	} else {
		this.Now = this.Now - steps
		return this.UrlCache[this.Now-1]
	}
}

func (this *BrowserHistory) Forward(steps int) string {
	if steps+this.Now > len(this.UrlCache) {
		this.Now = len(this.UrlCache)
		return this.UrlCache[len(this.UrlCache)-1]
	} else {
		this.Now += steps
		return this.UrlCache[this.Now-1]
	}
}
