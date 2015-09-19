package search

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"text/template"
)

var engines = map[string]string{
	"bing":   "http://global.bing.com/search?q={{.search}}&setmkt=en-us&setlang=en-us",
	"google": "https://google.com/search?q={{.search}}",
	"ddg":    "https://duckduckgo.com/search?q={{.search}}",
}

func makeRedirectQueryString(engine string, query string) string {
	tmpl, err := template.New("qstring").Parse(engines[engine])
	if err != nil {
		fmt.Println(err)
		return engines["bing"]
	}
	var out bytes.Buffer
	if err := tmpl.Execute(&out, map[string]string{"search": query}); err != nil {
		fmt.Println(err)
		return engines["bing"]
	}
	return out.String()
}

// Search determines country of origin for search and sets the search engine
// to Bing global in English when inside China or Google if using VPN
func Search(c *gin.Context) {
	searchString := c.Query("q")

	resp, err := http.Get("http://ipinfo.io/country")
	if err != nil {
		c.Redirect(301, makeRedirectQueryString("bing", ""))
	}
	defer resp.Body.Close()

	countryB, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		c.Redirect(301, makeRedirectQueryString("bing", ""))
	}
	country := strings.Trim(string(countryB), "\n\r ")

	engine := "google"
	switch country {
	case "CN":
		engine = "bing"
	}
	log.Printf("Found country %s, using %s", string(country), engine)

	c.Redirect(301, makeRedirectQueryString(engine, searchString))
}
