package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	beego "github.com/beego/beego/v2/server/web"
)

// CatController handles requests to show cat images.
type CatController struct {
	beego.Controller
}

// Struct for holding the API response
type CatResponse struct {
	URL    string `json:"url"`
	ID     string `json:"id"`
	Width  int    `json:"width"`
	Height int    `json:"height"`
	Breeds []struct {
		Name        string `json:"name"`
		Origin      string `json:"origin"`
		Description string `json:"description"`
		URL         string `json:"wikipedia_url"`
	} `json:"breeds"`
}

// ShowCat fetches and displays a random cat image.
func (c *CatController) ShowCat() {
	apiKey := "live_8Vq87uY7jXkcqmqwhODWVdzEp9iUzbog1G0hxJgh6gphgTP9sjK23Pbnir5Xl5JY" // Replace with your actual API key
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?api_key=%s", apiKey)
	// url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=10&breed_ids=beng&api_key=%s", apiKey)

	// Make the HTTP request to fetch the cat image
	resp, err := http.Get(url)
	if err != nil {
		c.Ctx.WriteString("Error fetching cat image")
		fmt.Println("Error fetching cat image:", err)  // Print error
		return
	}
	defer resp.Body.Close()

	// Check if we received a valid response
	fmt.Println("HTTP Status Code:", resp.StatusCode)  // Print status code

	// Parse the response JSON
	var catResponse []CatResponse
	if err := json.NewDecoder(resp.Body).Decode(&catResponse); err != nil {
		c.Ctx.WriteString("Error parsing cat image response")
		fmt.Println("Error decoding JSON:", err)  // Print JSON decoding error
		return
	}

	// Debug print the response data
	// fmt.Println("Cat API Response:", catResponse)  // Print the entire response to check the structure

	// Pass the cat image URL and breed info to the view
	if len(catResponse) > 0 {
		c.Data["CatImage"] = catResponse[0].URL
		c.Data["CatID"] = catResponse[0].ID
		c.Data["CatWidth"] = catResponse[0].Width
		c.Data["CatHeight"] = catResponse[0].Height
		if len(catResponse[0].Breeds) > 0 {
			c.Data["CatBreedName"] = catResponse[0].Breeds[0].Name
			c.Data["CatBreedOrigin"] = catResponse[0].Breeds[0].Origin
			c.Data["CatBreedDescription"] = catResponse[0].Breeds[0].Description
			c.Data["CatBreedURL"] = catResponse[0].Breeds[0].URL
		} else {
			c.Data["CatBreedName"] = "Unknown"
		}
	} else {
		c.Data["CatImage"] = "No image found"
	}

	// Render the HTML template
	c.TplName = "index.tpl"
}
