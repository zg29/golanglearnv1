package main

import (
    "encoding/csv"
    "log"
    "os"
    "fmt"

    "github.com/gocolly/colly"
)

// initialize a data structure to keep the scraped data
type Weather struct {
    city, temp, condition string
}
/*
This main function defines a colly collector, scrapes a weather website to find basic weather data 
(temp, condition, city) then puts that data into a csv
@param: none
@return: none
*/
func main() {
    // define collector
    c := colly.NewCollector(
        colly.AllowedDomains("www.wunderground.com"),
    )

    var weathers []Weather
    weather := Weather{}
    // scrape data using the HTML
    c.OnHTML(".current-temp", func(e *colly.HTMLElement) {

        weather.temp = e.ChildText(".wu-value")
        fmt.Println(weather.temp)


    })
    c.OnHTML(".current-temp", func(e *colly.HTMLElement) {

        weather.temp = e.ChildText(".wu-value")
        fmt.Println(weather.temp)


    })
    c.OnHTML(".condition-icon", func(a *colly.HTMLElement) {

        weather.condition = a.ChildText("p")
        fmt.Println(weather.condition)

    })
    // output the scraped data to a CSV
    c.OnScraped(func(r *colly.Response) {
        weathers = append(weathers,weather)
        file, err := os.Create("weather.csv")
        if err != nil {
            log.Fatalln("Failed to create output CSV file", err)
        }
        defer file.Close()

        writer := csv.NewWriter(file)

        headers := []string{
            "Temp",
            "Condition",
            "city",
        }
        writer.Write(headers)

        for _, weather := range weathers {
            record := []string{
                weather.temp,
                weather.condition,
            }

            writer.Write(record)
        }
        defer writer.Flush()
    })
    // visit specific URL
    c.Visit("https://www.wunderground.com/weather/us/il/metamora/KILMETAM36?utm_source=HomeCard&utm_content=Button&cm_ven=HomeCardButton")

}
