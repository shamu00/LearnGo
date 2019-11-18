package links

import (
	"fmt"
	"net/http"

	"golang.org/x/net/html"
)

func Extract(url string) ([]string, error) {
	rsp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if rsp.StatusCode != http.StatusOK {
		rsp.Body.Close()
		return nil, fmt.Errorf("getting %s: %s", url, rsp.Status)
	}

	doc, err := html.Parse(rsp.Body)
	rsp.Body.Close()
	if err != nil {
		return nil, fmt.Errorf("parsing %s as HTML, %v", url, err)
	}

	var links []string
	visitNode := func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "a" {
			for _, a := range n.Attr {
				if a.Key != "href" {
					continue
				}
				link, err := rsp.Request.URL.Parse(a.Val)
				if err != nil {
					continue
				}
				links = append(links, link.String())
			}
		}
	}
	forEachNode(doc, visitNode, nil)

	return links, nil
}
