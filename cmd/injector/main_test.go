package main

import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"strings"
	"testing"
)

// Long HTML String fixture
var htmlString = `<!DOCTYPE html>
<html lang="en">
<head>
    <script type="application/json" data-gig-selector="config">
        {}
    </script>
</head>
<body>
</body>
</html>`

// Long YAML String fixture
var yamlString = `---
  apiurl: https://api.com
  main_site_url: https://site.com`

func TestInjection(t *testing.T) {
	var selector = "data-gig-selector"
	var value = "config"
	jsonString, err := yaml2json(yamlString)
	if err != nil {
		t.Errorf("failed to convert yaml to json: %v", err)
	}
	htmlString, err := insertJSON(htmlString, jsonString, selector, value)
	if err != nil {
		t.Errorf("failed to insert json into html: %v", err)
	}
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		t.Errorf("failed to parse html: %v", err)
	}
	scriptBlock := doc.Find(fmt.Sprintf("script[%s='%s']", selector, value))
	yamlString := scriptBlock.Text()
	if yamlString != yamlString {
		t.Errorf("yamlString does not match")
	}
}
