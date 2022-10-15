// this is frap written in go
package main

import (
	"encoding/json"
	"io/ioutil"
	"log"

	"github.com/gocolly/colly"
)

// Rapport is the datastructure of the information we want to show
type Rapport struct {
	Namn  string `json:"Namn"`
	Datum string `json:"Datum"`
	Art   string `json:"Art"`
	Langd string `json:"Langd"`
	Plats string `json:"Plats"`
	Metod string `json:"Metod"`
}

func main() {
	url := "https://kagealven.com/fangstrapporter-aktuella/"
	rap := make([]Rapport, 0)
	agent := "Mozilla/5.0 (Windows NT 6.1) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/41.0.2228.0 Safari/537.36"
	c := colly.NewCollector(
		colly.AllowedDomains("kagealven.com"),
		colly.UserAgent(agent),
	)
	c.OnHTML("table", func(e *colly.HTMLElement) {
		nyRapport := Rapport{}
		e.ForEach("tr", func(_ int, row *colly.HTMLElement) {
			row.ForEach("td", func(_ int, data *colly.HTMLElement) {
				switch data.Index {
				case 0:
					nyRapport.Namn = data.Text
				case 1:
					nyRapport.Datum = data.Text
				case 2:
					nyRapport.Art = data.Text
				case 8:
					nyRapport.Langd = data.Text
				case 10:
					nyRapport.Plats = data.Text
				case 7:
					nyRapport.Metod = data.Text
				}
			})
			rap = append(rap, nyRapport)
		})
	})

	c.Visit(url)
	/* enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", " ")
	enc.Encode(rap)
	fmt.Println("Done!") */
	writeJSON(rap)
}

func writeJSON(data []Rapport) {
	file, err := json.MarshalIndent(data, "", " ")
	if err != nil {
		log.Fatalf("unable to create JSON file")
		return
	}
	_ = ioutil.WriteFile("rapporter.json", file, 0644)
}
