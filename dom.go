package ketchup

import (
	"regexp"
	"strings"
)

var xmlTag = regexp.MustCompile(`(\<.+?\>)|(\<//?\w+\>\\?)`)
var clTag = regexp.MustCompile(`\<\/\w+\>`)
var selfClosingTag = regexp.MustCompile(`\<.+\/\>`)
var tagContent = regexp.MustCompile(`(.+?)\<\/`)
var tagName = regexp.MustCompile(`(\<\w+)`)
var attr = regexp.MustCompile(`\w+=".+?"`)

func extractAttributes(tag string) []*Attribute {
	rawAttrArray := attr.FindAllString(tag, -1)
	elementAttrs := []*Attribute{}

	for i := 0; i < len(rawAttrArray); i++ {
		attrStringSlice := strings.Split(rawAttrArray[i], "=")
		attr := &Attribute{
			Name:  attrStringSlice[0],
			Value: strings.Trim(attrStringSlice[1], "\""),
		}

		elementAttrs = append(elementAttrs, attr)
	}

	return elementAttrs
}

func isVoidElement(tagName string) bool {
	var isVoid bool
	switch tagName {
	case "area",
		"base",
		"br",
		"col",
		"command",
		"embed",
		"hr",
		"img",
		"input",
		"keygen",
		"link",
		"meta",
		"param",
		"source",
		"track",
		"wbr":
		isVoid = true
	default:
		isVoid = false
	}

	return isVoid
}

func ParseDocument(document string) *NodeDOM {
	DOM_Tree := &NodeDOM{
		Element:  "root",
		Content:  "THDWB",
		Children: []*NodeDOM{},
		Style:    nil,
		Parent:   nil,
	}

	lastNode := DOM_Tree
	parseDocument := xmlTag.MatchString(document)
	document = strings.ReplaceAll(document, "\n", "")

	for parseDocument == true {
		var currentNode *NodeDOM

		currentTag := xmlTag.FindString(document)
		currentTagIndex := xmlTag.FindStringIndex(document)

		if string(currentTag[1]) == "!" {
			document = strings.Replace(document, currentTag, "", 1)
		} else {
			if clTag.MatchString(currentTag) {
				contentStringMatch := tagContent.FindStringSubmatch(document)
				contentString := ""

				if len(contentStringMatch) > 1 {
					contentString = contentStringMatch[1]
				}

				if clTag.MatchString(contentString) {
					lastNode.Content = ""
				} else {
					lastNode.Content = strings.TrimSpace(contentString)
				}

				lastNode = lastNode.Parent
			} else {
				currentTagName := strings.Trim(tagName.FindString(currentTag), "<")
				extractedAttributes := extractAttributes(currentTag)
				elementStylesheet := GetElementStylesheet(currentTagName, extractedAttributes)

				currentNode = &NodeDOM{
					Element:    currentTagName,
					Content:    "",
					Children:   []*NodeDOM{},
					Attributes: extractedAttributes,
					Style:      elementStylesheet,
					Parent:     lastNode,
				}

				lastNode.Children = append(lastNode.Children, currentNode)

				if !isVoidElement(currentTagName) {
					lastNode = currentNode
				}
			}

			document = document[currentTagIndex[1]:len(document)]
		}

		if !xmlTag.MatchString(document) {
			parseDocument = false
		}
	}

	if len(DOM_Tree.Children) > 0 {
		return DOM_Tree.Children[0]
	}

	return DOM_Tree
}
