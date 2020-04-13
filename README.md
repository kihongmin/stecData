#used Library & how to set
- goquery
$ go get github.com/PuerkitoBio/goquery

- chromedp
$ go get -u github.com/chromedp/chromedp

1. Crawling informations about url, title ... etc
    - Chromedp
    (1) Run new Context
        ex) ctx, cancle := chromedp.NewContext(context.Background())    <- create context
            err := chromedp.Run(ctx,chromedp.Navigate(baseURL),)        <- go to URL

    (2) Get __node__ which have what we want
        ex) err := chromedp.Run(ctx,                                    <- always run context 
                    chromedp.Nodes([selector], &[saving var],)          <- get nodes having selector address
                    
    (3) Set how many pages we crawl
        >Because it varys from site to site, you have to search it by using (2)
        
    (4) Get information which I want
        1) Get Nodes by using (2)  
        
        2) If it is attribute value
            To get information from Nodes, we have to do it
            
            for index, row := range Nodes{                              <- saved Nodes by &[saving var]
                repository[index] = row.AttributeValue([attribute name])<- what we want!
            }
            
            
        3) If it is text bounded by tag
            To get information from Nodes, we have to do it
            
            for index, row := range Nodes{
                repository[index] = __row.Children[0].NodeValue__       <- what we want!
            }
            
            
2. Get urls from crawled data and get bodytext of the url
    -Chromedp
        it is same with way of 1.
    -http and goquery
    (1) use get method by http
        ex) res, err := http.Get(url)
        
    (2) set quit by using defer
        ex) defer res.Body.Close()
        
    (3) Read it by goquery
        ex) doc, err = goquery.NewDocumentFromReader(res.Body)
        
    (4) using Find method to get text data
        ex) doc.Find("[selector]").Each(func(in int, tablehtml *goquery.Selection){
            repository = __tablehtml.Text()__ <- what we want!
        }


# Data Description
{
1. url:
    -ex) https://gitlab.com/geekermeter/data/-/edit/master/README.md
2. title
    -ex) [FIFA ONLINE 4] Feature Game Client Engineer
3. origin
    -ex) EA Korea
4. start_date
    -ex) 3.11 (need to be same, but not yet)
5. newbie
    -ex) 신입, 경력
6. content

}

# 크롤링 대상 홈페이지

	- coupang
	- kakao
	- naver
	- netmarble
	- nexon
	- programmers
	- rocketpunch
-----------------------------

# 진행상황

|사이트|URL|Title|origin|start_date|newbie|content|auto|
|:--:|:--:|:--:|:--:|:--:|:--:|:--:|:--:|
|coupang|O|X|X|X|X|X|X|
|kakao|O|X|X|X|X|X|X|
|naver|O|O|O|O|X|O|O|
|netmarble|O|O|O|O|O|O|O|
|nexon|O|O|O|O|O|O|O|
|programmers|O|O|O|O|O|O|O|
|rocketpunch|O|O|O|O|O|O|O|
