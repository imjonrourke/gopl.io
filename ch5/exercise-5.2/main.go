package main

import (
  "fmt"
  "golang.org/x/net/html"
  "os"
)

func main() {
  doc, err := html.Parse(os.Stdin)
  elementMap := make(map[string]int)
  if err != nil {
    fmt.Fprintf(os.Stderr, "HTML Parsing error #{err}\n")
  }
  fmt.Printf("result of map %v\n", parseHtml(elementMap, doc))
}

func parseHtml(m map[string]int, n *html.Node) map[string]int {
  if n.Type == html.ElementNode {
    m[n.Data]++
  }
  if n.FirstChild != nil {
    parseHtml(m, n.FirstChild)
  }
  if n.NextSibling != nil {
    parseHtml(m, n.NextSibling)
  }
  return m
}
