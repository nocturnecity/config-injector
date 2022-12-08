package main

import (
	"flag"
	"fmt"
	"github.com/PuerkitoBio/goquery"
	json "github.com/json-iterator/go"
	yaml "gopkg.in/yaml.v2"
	"os"
	"strings"
)

var (
	version = "dev"
	commit  = "none"
	date    = "unknown"
)

func yaml2json(yamlString string) (string, error) {
	var data interface{}
	err := yaml.Unmarshal([]byte(yamlString), &data)
	if err != nil {
		return "", err
	}
	//print data structure
	jsonString, err := json.Marshal(data)
	if err != nil {
		return "", err
	}
	return string(jsonString), nil
}

func insertJSON(htmlString string, jsonString string, selector string, value string) (string, error) {
	doc, err := goquery.NewDocumentFromReader(strings.NewReader(htmlString))
	if err != nil {
		return "", err
	}
	scriptBlock := doc.Find(fmt.Sprintf("script[%s='%s']", selector, value))
	scriptBlock.SetHtml(jsonString)
	indexString, err := doc.Html()
	if err != nil {
		return "", err
	}
	return indexString, nil
}

func main() {
	helpFlag := flag.Bool("help", false, "print help and exit")
	versionFlag := flag.Bool("version", false, "print version and exit")
	yamlFileFlag := flag.String("yaml", "", "path to yaml file")
	htmlFileFlag := flag.String("html", "", "path to html file")

	configAttrFlag := flag.String("config-attr", "data-gig-selector", "attribute name of config block")
	configAttrValueFlag := flag.String("config-attr-value", "config", "attribute value of config block")

	flag.Parse()

	if *helpFlag {
		flag.PrintDefaults()
		os.Exit(0)
	}

	if *versionFlag {
		fmt.Println(version)
		fmt.Println(commit)
		fmt.Println(date)
		os.Exit(0)
	}

	if *yamlFileFlag == "" {
		fmt.Println("yaml file not set")
		os.Exit(1)
	}

	if *htmlFileFlag == "" {
		fmt.Println("html file not set")
		os.Exit(1)
	}

	//read yaml file
	yamlFile, err := os.ReadFile(*yamlFileFlag)
	if err != nil {
		fmt.Println("yaml file read error:", err)
		os.Exit(1)
	}
	jsonString, err := yaml2json(string(yamlFile))
	if err != nil {
		fmt.Println("yaml to json error:", err)
		os.Exit(1)
	}

	htmlString, err := os.ReadFile(*htmlFileFlag)
	if err != nil {
		fmt.Println("html file read error:", err)
		os.Exit(1)
	}

	indexString, err := insertJSON(string(htmlString), jsonString, *configAttrFlag, *configAttrValueFlag)
	if err != nil {
		fmt.Println("insert json error:", err)
		os.Exit(1)
	}
	fmt.Println(indexString)

}
