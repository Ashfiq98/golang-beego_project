package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
    "bytes"
	beego "github.com/beego/beego/v2/server/web"
)

// CatController handles requests related to cats.
type CatController struct {
	beego.Controller
}

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

type Breed struct {
	ID               string `json:"id"`
	Name             string `json:"name"`
	Temperament      string `json:"temperament"`
	Origin           string `json:"origin"`
	Description      string `json:"description"`
	LifeSpan         string `json:"life_span"`
	WikipediaURL     string `json:"wikipedia_url"`
	ReferenceImageID string `json:"reference_image_id"`
}

var (
	breeds []Breed
	once   sync.Once
)

// ShowCat fetches and displays a random cat image using a Go channel.
func (c *CatController) ShowCat() {
	apiKey := "live_8Vq87uY7jXkcqmqwhODWVdzEp9iUzbog1G0hxJgh6gphgTP9sjK23Pbnir5Xl5JY" // Replace with your actual API key
	url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?api_key=%s", apiKey)

	responseChannel := make(chan []CatResponse)
	errorChannel := make(chan error)

	go func() {
		resp, err := http.Get(url)
		if err != nil {
			errorChannel <- err
			return
		}
		defer resp.Body.Close()

		var catResponse []CatResponse
		if err := json.NewDecoder(resp.Body).Decode(&catResponse); err != nil {
			errorChannel <- err
			return
		}

		responseChannel <- catResponse
	}()

	select {
	case catResponse := <-responseChannel:
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
	case err := <-errorChannel:
		c.Ctx.WriteString("Error fetching or parsing cat image response")
		log.Println("Error:", err)
		return
	}

	c.TplName = "index.tpl"
}

// FetchBreeds fetches and stores cat breeds from the API.
// FetchBreeds fetches and stores cat breeds from the API.
func fetchBreeds(ch chan<- []Breed, errCh chan<- error) {
	resp, err := http.Get("https://api.thecatapi.com/v1/breeds")
	if err != nil {
		errCh <- fmt.Errorf("failed to fetch breeds: %v", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		errCh <- fmt.Errorf("failed to read response body: %v", err)
		return
	}

	var fetchedBreeds []Breed
	if err := json.Unmarshal(body, &fetchedBreeds); err != nil {
		errCh <- fmt.Errorf("failed to unmarshal breeds: %v", err)
		return
	}

	ch <- fetchedBreeds
}

// FetchImagesByBreed fetches images for a given breed using its ID.
func fetchImagesByBreed(breedID string, ch chan<- []CatResponse, errCh chan<- error) {
	apiUrl := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?breed_ids=%s&limit=5", breedID)
	resp, err := http.Get(apiUrl)
	if err != nil {
		errCh <- fmt.Errorf("failed to fetch images for breed %s: %v", breedID, err)
		return
	}
	defer resp.Body.Close()

	var catResponse []CatResponse
	if err := json.NewDecoder(resp.Body).Decode(&catResponse); err != nil {
		errCh <- fmt.Errorf("failed to decode images for breed %s: %v", breedID, err)
		return
	}

	ch <- catResponse
}

// FetchAndStoreBreeds fetches breeds once and caches them in memory.
func fetchAndStoreBreeds() {
	ch := make(chan []Breed)
	errCh := make(chan error)

	go fetchBreeds(ch, errCh)

	select {
	case data := <-ch:
		breeds = data
		log.Println("Breeds fetched and stored successfully!")
	case err := <-errCh:
		log.Fatalf("Error fetching breeds: %v", err)
	}
}

// GetBreedsHandler handles requests to fetch and display breeds.
func (c *CatController) GetBreedsHandler() {
	once.Do(fetchAndStoreBreeds)
	c.Data["json"] = breeds
	c.ServeJSON()
}

// GetBreedImagesHandler handles fetching images for each breed dynamically.
func (c *CatController) GetBreedImagesHandler() {
	once.Do(fetchAndStoreBreeds) // Ensure breeds are fetched only once

	// Create channels for breed images and errors
	imageCh := make(chan []CatResponse)
	errCh := make(chan error)

	// Fetch images for each breed concurrently
	for _, breed := range breeds {
		go fetchImagesByBreed(breed.ID, imageCh, errCh)
	}

	// Collect the results
	var allImages []CatResponse
	for i := 0; i < len(breeds); i++ {
		select {
		case images := <-imageCh:
			allImages = append(allImages, images...)
		case err := <-errCh:
			log.Println("Error fetching images:", err)
		}
	}

	// Send the collected images to the view
	c.Data["json"] = allImages
	c.ServeJSON()
}
func (c *CatController) VoteOnImage() {
    // Parse the incoming vote data from the request body
    var vote struct {
        ImageID string `json:"image_id"`
        SubID   string `json:"sub_id"`
        Value   int    `json:"value"`
    }

    err := json.Unmarshal(c.Ctx.Input.RequestBody, &vote)
    if err != nil {
        c.Ctx.WriteString("Invalid request")
        return
    }

    // Construct the request body to send to the Cat API
    voteRequest := map[string]interface{}{
        "image_id": vote.ImageID,
        "sub_id":   vote.SubID,
        "value":    vote.Value,
    }

    // Send the POST request to the Cat API
    client := &http.Client{}
    reqBody, err := json.Marshal(voteRequest)
    if err != nil {
        c.Ctx.WriteString("Error marshalling vote data")
        return
    }

    req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/votes", bytes.NewBuffer(reqBody))
    if err != nil {
        c.Ctx.WriteString("Error creating vote request")
        return
    }
	apiKey := "live_8Vq87uY7jXkcqmqwhODWVdzEp9iUzbog1G0hxJgh6gphgTP9sjK23Pbnir5Xl5JY"

    req.Header.Set("Content-Type", "application/json")
    req.Header.Set("x-api-key", apiKey) // Replace with your API key

    resp, err := client.Do(req)
    if err != nil {
        c.Ctx.WriteString("Error sending vote request")
        return
    }
    defer resp.Body.Close()

    // Handle the response and provide feedback to the user
    if resp.StatusCode == http.StatusOK {
        c.Ctx.WriteString("Vote submitted successfully!")
    } else {
        c.Ctx.WriteString("Failed to submit vote")
    }
}
