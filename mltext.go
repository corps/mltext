// Copyright 2013 Zachary Collins
// Implements a very simple approximate xml / html -> plain text conversion
package mltext

import (
	"code.google.com/p/go.net/html"
	"code.google.com/p/go.net/html/atom"
	"io"
	"regexp"
	"strings"
)

// Reads in the html given in the input reader, then attempts to convert it
// to some plain text.  err is given from code.google.com/p/go.net/html.Parse
func ToText(htmlReader io.Reader) (text string, err error) {
	doc, err := html.Parse(htmlReader)
	if err != nil {
		return text, err
	}

	return textOfNode(doc, false), nil
}

func textOfNode(node *html.Node, isBoxBoundary bool) (text string) {
	if node.Type == html.ElementNode {
		switch node.DataAtom {
		case atom.P, atom.Div, atom.Li:
			if !isBoxBoundary {
				text += "\n"
				isBoxBoundary = true
			}
		case atom.Br:
			text += "\n"
		}
	} else if node.Type == html.TextNode {
		if len(strings.TrimSpace(node.Data)) > 0 {
			text += htmlNormalizedText(node.Data)
		}
	}

	for child := node.FirstChild; child != nil; child = child.NextSibling {
		childText := textOfNode(child, isBoxBoundary)
		text += childText
		if strings.HasSuffix(childText, "\n") {
			isBoxBoundary = true
		} else if len(strings.TrimSpace(childText)) > 0 {
			isBoxBoundary = false
		}
	}

	return text
}

var (
	doubledWhitespace = regexp.MustCompile(`\s\s+`)
)

func htmlNormalizedText(text string) (normalized string) {
	return doubledWhitespace.ReplaceAllString(text, " ")
}
