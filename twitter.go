package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Twitter")

	/*
		Example usage of the Twitter API version 2.
	*/

	// Set up HTTP client and request to Twitter API
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://api.twitter.com/2/tweets/counts/recent?query=lang%3Aen%20%22%24TSLA%22&start_time=2022-12-08T00:00:00.000Z&end_time=2022-12-13T00:00:00.000Z&granularity=day&search_count.fields=tweet_count", nil)

	// Ask for bearer token input and display it - supersafe
	var token string
	fmt.Print("Enter bearer token: ")
	fmt.Scan(&token)

	// Add necessary authorization header with bearer token to request
	req.Header.Add("Authorization", "Bearer "+token)
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println("Error sending HTTP request:", err)
		return
	}

	// Read response body
	body, _ := ioutil.ReadAll(resp.Body)
	// Unmarshal JSON response into a map
	var result map[string]interface{}
	json.Unmarshal(body, &result)
	// Access the "data" field in the response
	data := result["data"].([]interface{})
	// Iterate over the array of data objects and print their fields
	for _, datum := range data {
		m := datum.(map[string]interface{})
		fmt.Println("end:", m["end"])
		fmt.Println("start:", m["start"])
		fmt.Println("tweet_count:", m["tweet_count"])
	}
	fmt.Println(result)

	// Extract the "tweet_count" field from each data object
	var counts []int
	for _, datum := range data {
		m := datum.(map[string]interface{})
		counts = append(counts, int(m["tweet_count"].(float64)))
	}

	// Create a new plot
	p := plot.New()
	// Set up plot title, axis labels, and legend
	p.Title.Text = "Twitter Trend Counts"
	p.X.Label.Text = "Days"
	p.Y.Label.Text = "Tweet Counts"
	// Create a line plot with the count data points
	l, err := plotter.NewLine(plotter.XYs{
		{1.0, float64(counts[0])},
		{2.0, float64(counts[1])},
		{3.0, float64(counts[2])},
		{4.0, float64(counts[3])},
		{5.0, float64(counts[4])},
	})
	if err != nil {
		panic(err)
	}
	// Add the plot to the plot
	p.Add(l)
	// Save the plot to an image file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "trend_counts.png"); err != nil {
		panic(err)
	}
}
