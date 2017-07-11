package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"io"
	"strings"
)

func Document(contest, problem string) (*goquery.Document, error) {
	url := fmt.Sprintf("http://codeforces.com/contest/%s/problem/%s", contest, problem)
	return goquery.NewDocument(url)
}

func Examples(w io.Writer, contest, problem, className string) error {
	doc, err := Document(contest, problem)
	if err != nil {
		return err
	}

	doc.Find(className).Each(func(i int, s *goquery.Selection) {
		html, _ := s.Find("pre").Html()
		fmt.Fprintln(w, strings.Replace(html, "<br/>", "\n", 2))
	})

	return nil
}
