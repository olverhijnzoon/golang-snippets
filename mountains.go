package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math"
	"sort"
)

type Location struct {
	Latitude  float64
	Longitude float64
}

type Mountain struct {
	Name     string
	Height   float64
	Location Location
}

func (m *Mountain) Distance(lat, lon float64) float64 {
	const R = 6371 // Earth's radius in kilometers
	dLat := (lat - m.Location.Latitude) * math.Pi / 180
	dLon := (lon - m.Location.Longitude) * math.Pi / 180
	a := math.Sin(dLat/2)*math.Sin(dLat/2) + math.Cos(lat*math.Pi/180)*math.Cos(m.Location.Latitude*math.Pi/180)*math.Sin(dLon/2)*math.Sin(dLon/2)
	c := 2 * math.Atan2(math.Sqrt(a), math.Sqrt(1-a))
	d := R * c
	return d
}

func FindClosestMountain(mountains []Mountain, lat, lon float64) Mountain {
	closest := mountains[0]
	closestDistance := closest.Distance(lat, lon)
	for _, m := range mountains[1:] {
		d := m.Distance(lat, lon)
		if d < closestDistance {
			closest = m
			closestDistance = d
		}
	}
	return closest
}

func writeToJSONFile(filepath string, mountains []Mountain) error {
	jsonData, err := json.MarshalIndent(mountains, "", " ")
	if err != nil {
		return err
	}

	err = ioutil.WriteFile(filepath, jsonData, 0644)
	if err != nil {
		return err
	}
	return nil
}

func readFromJSONFile(filepath string) ([]Mountain, error) {
	jsonData, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	var mountains []Mountain
	err = json.Unmarshal(jsonData, &mountains)
	if err != nil {
		return nil, err
	}

	return mountains, nil
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Mountains")

	/*
		This Golang snippet plays with representation of mountain data and information. It defines a struct named Mountain, which encapsulates the name, height, and location of a mountain.

		The snippet then initializes an array of Mountain structs with information of some of the highest mountains in the world, including Mount Everest, K2, Kangchenjunga and more. The array is then sorted by height, allowing us to easily identify the three tallest mountains. The snippet also includes a method Distance which utilizes the Haversine formula to calculate the distance between a mountain and a given location.

		Additionally, the snippet includes a function FindClosestMountain that finds the closest mountain to a given location by iterating through all the mountains and calculating their distance. To make the data more persistent and easy to share, the snippet also includes a function writeToJSONFile that writes the mountain data to a JSON file and a function readFromJSONFile that reads the mountain data from a JSON file.

		Finally, the snippet ends by printing the data read from the json file, making it easy for the user to access and utilize the information.
	*/

	mountains := []Mountain{
		{Name: "Mount Everest", Height: 8848, Location: Location{Latitude: 27.988056, Longitude: 86.925278}},
		{Name: "K2", Height: 8611, Location: Location{Latitude: 35.8825, Longitude: 76.513333}},
		{Name: "Kangchenjunga", Height: 8586, Location: Location{Latitude: 27.702508, Longitude: 88.147576}},
		{Name: "Lhotse", Height: 8516, Location: Location{Latitude: 27.960833, Longitude: 86.923611}},
		{Name: "Makalu", Height: 8485, Location: Location{Latitude: 27.888333, Longitude: 87.088889}},
		{Name: "Cho Oyu", Height: 8188, Location: Location{Latitude: 28.093611, Longitude: 86.660833}},
		{Name: "Dhaulagiri", Height: 8167, Location: Location{Latitude: 28.694722, Longitude: 83.493056}},
		{Name: "Manaslu", Height: 8163, Location: Location{Latitude: 28.539722, Longitude: 84.562222}},
		{Name: "Nanga Parbat", Height: 8126, Location: Location{Latitude: 35.2375, Longitude: 74.589167}},
	}

	// Sort the mountains by height
	sort.Slice(mountains, func(i, j int) bool {
		return mountains[i].Height > mountains[j].Height
	})

	fmt.Println("Tallest mountains:")
	for _, m := range mountains[:3] {
		fmt.Printf("%s (%.0fm)\n", m.Name, m.Height)
	}

	// Find closest mountain to a location
	lat, lon := 28.3949, 84.1240
	closest := FindClosestMountain(mountains, lat, lon)
	fmt.Printf("\nClosest mountain to %.4f, %.4f is %s (%.0fm)\n", lat, lon, closest.Name, closest.Height)

	// write the data to json file
	err := writeToJSONFile("mountains.json", mountains)
	if err != nil {
		fmt.Println(err)
	}

	// read the data from json file
	data, err := readFromJSONFile("mountains.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("\nData read from json file:", data)
}
