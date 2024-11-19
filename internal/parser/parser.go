package parser

import (
	"errors"
	"fmt"
	"io"
	"net/url"
	"slices"
	"strings"

	"golang.org/x/net/html"
)

func (p *Parser) FiltersUrls(raw_url *string) (bool, error) {
	// parse url extension if exist

	u, err := url.Parse(*raw_url)
	if err != nil {
		return false, err
	}
	// filter on extensions we look for
	// paths := strings.Split(u.Path, ".")
	index := slices.IndexFunc(p.Config.FilteredFileExtensions, func(s string) bool { return strings.Contains(u.Path, s) })
	if index == -1 {
		return false, errors.New("not found")
	} else {
		return true, nil
	}
}

// TODO parse all ressources defined in config
func (p *Parser) GetHrefs(n *html.Node, urls *[]string) error {
	if n.Attr == nil {
		return errors.New("no hrefs")
	}
	for _, a := range n.Attr {
		if a.Key == "href" {
			if strings.Contains(a.Val, "css") {
				fmt.Println("CSS FOUND")
			}
			valid, err := p.FiltersUrls(&a.Val)
			if err != nil {
				continue
			}
			if valid {
				fmt.Println(a.Val)
				*urls = append(*urls, a.Val)
			}
		}

	}
	// Traverse child nodes
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		err := p.GetHrefs(c, urls)
		if err != nil {
			continue
		}
	}
	return nil
}

func (p *Parser) GetAllUrls(body io.ReadCloser, urls *[]string) (err error) {
	doc, err := html.Parse(body)
	if err != nil {
		return err
	}
	var get_urls func(*html.Node)
	get_urls = func(n *html.Node) {
		if n.Type == html.ElementNode {
			err := p.GetHrefs(n, urls)
			if err != nil { //nolint:all
			}
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			get_urls(c)
		}
	}
	get_urls(doc)
	return err

}
