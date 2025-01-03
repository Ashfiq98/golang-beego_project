package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"sync"
	"time"
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

// CatCache holds the image data in memory with a lock to ensure thread-safety
var CatCache struct {
    sync.Mutex
    ImageData    *CatResponse
    LastFetched  time.Time
}
type FavouriteRequest struct {
    ImageID string `json:"image_id"`
    SubID   string `json:"sub_id"`
}
type VoteRequest struct {
    ImageID string `json:"image_id"`
    SubID   string `json:"sub_id"`
    Value   int    `json:"value"` // 1 for upvote, 0 for downvote
}

type VoteResponse struct {
    ID      int    `json:"id"`
    Message string `json:"message"`
}
type VoteHistoryResponse struct {
    ID        int       `json:"id"`
    ImageID   string    `json:"image_id"`
    SubID     string    `json:"sub_id"`
    CreatedAt time.Time `json:"created_at"`
    Value     int       `json:"value"`
    Image     struct {
        ID  string `json:"id"`
        URL string `json:"url"`
        } `json:"image"`
    }
    
var requestData struct {
        ImageID string `json:"image_id"`
        SubID   string `json:"sub_id"`
    }
    
var APIBaseURL = "https://api.thecatapi.com/v1" // Default base URL

// ShowCat fetches generate cats data with load 👇
func (c *CatController) ShowCat() {
    apiKey, err := beego.AppConfig.String("api_key")
    url := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?api_key=%s", apiKey)

    // Fetch data synchronously to ensure it completes before sending a response
    resp, err := http.Get(url)
    if err != nil {
        log.Println("Error fetching cat data:", err)
        c.Ctx.WriteString("Error fetching cat data")
        return
    }
    defer resp.Body.Close()

    var catResponse []CatResponse
    if err := json.NewDecoder(resp.Body).Decode(&catResponse); err != nil {
        log.Println("Error decoding cat data:", err)
        c.Ctx.WriteString("Error decoding cat data")
        return
    }

    // Store the fetched data in-memory (in CatCache)
    CatCache.Lock()
    if len(catResponse) > 0 {
        CatCache.ImageData = &catResponse[0] // Store the first image in the cache
        CatCache.LastFetched = time.Now()
    }
    CatCache.Unlock()

    // Now we can send the image URL to the template
    if len(catResponse) > 0 {
        c.Data["CatImage"] = catResponse[0].URL
    } else {
        c.Data["CatImage"] = "No image found"
    }

    // Render the template
    c.TplName = "index.tpl"
}

// API to get the stored data (for frontend to fetch)
func (c *CatController) GetCatData() {
    // Ensure data is available in CatCache
    CatCache.Lock()
    if CatCache.ImageData != nil {
        c.Data["json"] = []CatResponse{*CatCache.ImageData}
    } else {
        c.Data["json"] = []string{"No cat data available"}
    }
    CatCache.Unlock()

    // Serve the JSON data
    c.ServeJSON()
}
// ShowCat fetches generate cats data with load 👆


// VOTING CODE : 👇

// VoteUp handles upvoting a cat image.
const subID = "test123" // Use a fixed sub_id for testing

func (c *CatController) VoteUp() {
    apiKey, err := beego.AppConfig.String("api_key")
    imageID := c.GetString("image_id")
    if imageID == "" {
        c.Data["json"] = map[string]string{"error": "Image ID is required"}
        c.ServeJSON()
        return
    }

    // Fetch the current vote state for the given image ID
    currentVote, err := getCurrentVote(imageID, subID)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get current vote"}
        c.ServeJSON()
        return
    }

    // Determine the new vote value (toggle between 1 and 0)
    var newVoteValue int
    if currentVote == 1 {
        newVoteValue = 0 // If it's already upvoted, remove the vote (set to 0)
    } else {
        newVoteValue = 1 // Otherwise, upvote (set to 1)
    }

    // Create a new vote request with the toggled value
    vote := VoteRequest{
        ImageID: imageID,
        SubID:   subID,
        Value:   newVoteValue,
    }

    // Send the updated vote
    err = sendVote(apiKey, vote)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]string{"message": "Vote toggled successfully"}
    c.ServeJSON()
}

func (c *CatController) VoteDown() {
    imageID := c.GetString("image_id")
    apiKey, err := beego.AppConfig.String("api_key")
    if imageID == "" {
        c.Data["json"] = map[string]string{"error": "Image ID is required"}
        c.ServeJSON()
        return
    }

    // Fetch the current vote state for the given image ID
    currentVote, err := getCurrentVote(imageID, subID)
    if err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to get current vote"}
        c.ServeJSON()
        return
    }

    // Determine the new vote value (toggle between -1 and 0)
    var newVoteValue int
    if currentVote == -1 {
        newVoteValue = 0 // If it's already downvoted, remove the vote (set to 0)
    } else {
        newVoteValue = -1 // Otherwise, downvote (set to -1)
    }

    // Create a new vote request with the toggled value
    vote := VoteRequest{
        ImageID: imageID,
        SubID:   subID,
        Value:   newVoteValue,
    }

    // Send the updated vote
    err = sendVote(apiKey, vote)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }

    c.Data["json"] = map[string]string{"message": "Vote toggled successfully"}
    c.ServeJSON()
}

// Helper function to get the current vote state for an image
func getCurrentVote(imageID, subID string) (int, error) {
    apiKey, err := beego.AppConfig.String("api_key")   
    url := fmt.Sprintf("https://api.thecatapi.com/v1/votes?sub_id=%s", subID)

    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        return 0, fmt.Errorf("failed to fetch vote history: %v", err)
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        return 0, fmt.Errorf("failed to fetch vote history: %v", resp.Status)
    }

    var votes []VoteHistoryResponse
    if err := json.NewDecoder(resp.Body).Decode(&votes); err != nil {
        return 0, fmt.Errorf("failed to decode vote history response: %v", err)
    }

    // Check if there's a vote for this imageID
    for _, vote := range votes {
        if vote.ImageID == imageID {
            return vote.Value, nil
        }
    }

    // If no previous vote is found, return 0 (no vote)
    return 0, nil
}


func (c *CatController) VoteHistory() {
    apiKey, err := beego.AppConfig.String("api_key")
    url := fmt.Sprintf("https://api.thecatapi.com/v1/votes?sub_id=%s", subID) // Use fixed sub_id
    
    req, _ := http.NewRequest("GET", url, nil)
    req.Header.Set("x-api-key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.Data["json"] = map[string]string{"error": err.Error()}
        c.ServeJSON()
        return
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        c.Data["json"] = map[string]string{"error": "Failed to fetch vote history"}
        c.ServeJSON()
        return
    }

    var votes []VoteHistoryResponse
    if err := json.NewDecoder(resp.Body).Decode(&votes); err != nil {
        c.Data["json"] = map[string]string{"error": "Failed to decode response"}
        c.ServeJSON()
        return
    }

    c.Data["json"] = votes
    c.ServeJSON()
}
// VOTING CODE : 👆

// Helper to send a vote
func sendVote(apiKey string, vote VoteRequest) error {
	body, err := json.Marshal(vote)
	if err != nil {
		return fmt.Errorf("failed to marshal vote request: %v", err)
	}

	req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/votes", bytes.NewBuffer(body))
	if err != nil {
		return fmt.Errorf("failed to create vote request: %v", err)
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("x-api-key", apiKey)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send vote request: %v", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected response status: %s", resp.Status)
	}

	return nil
}

//BREEDS CODE : 👇

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

// FetchImagesByBreedHandler handles requests to fetch images for a specific breed.
func (c *CatController) FetchImagesByBreedHandler() {
    once.Do(fetchAndStoreBreeds) // Ensure breeds are fetched only once

    // Get breedID from the URL parameter
    breedID := c.Ctx.Input.Param(":breedID")
    if breedID == "" {
        c.Data["json"] = map[string]string{"error": "Breed ID is required"}
        c.ServeJSON()
        return
    }

    // Create channels for fetching images
    imageCh := make(chan []CatResponse)
    errCh := make(chan error)

    // Fetch images for the specific breed
    go fetchImagesByBreed(breedID, imageCh, errCh)

    // Wait for the result
    select {
    case images := <-imageCh:
        c.Data["json"] = images
    case err := <-errCh:
        c.Data["json"] = map[string]string{"error": err.Error()}
    }

    c.ServeJSON()
}
//BREEDS CODE 👆


/// FAVOURITES CODE 👇


func (c *CatController) CreateFavourite() {
    apiKey, err := beego.AppConfig.String("api_key")    
    // Read the raw body first for debugging
    body, err := ioutil.ReadAll(c.Ctx.Request.Body)
    if err != nil {
        fmt.Println("Error reading body:", err)
        c.CustomAbort(400, "Error reading request body")
        return
    }
    
    // Log the raw request body
    // fmt.Println("Raw request body:", string(body))
    
    // Parse the JSON
    
    if err := json.Unmarshal(body, &requestData); err != nil {
        fmt.Println("Error parsing JSON:", err)
        c.CustomAbort(400, "Invalid JSON format")
        return
    }
    
    fmt.Printf("Parsed data - ImageID: %s, SubID: %s\n", requestData.ImageID, requestData.SubID)
    
    // Continue with your existing code...
    requestBody := map[string]string{
        "image_id": requestData.ImageID,
        "sub_id":   requestData.SubID,
    }
    
    jsonPayload, err := json.Marshal(requestBody)
    if err != nil {
        fmt.Println("Error marshaling request:", err)
        c.CustomAbort(500, "Failed to prepare request payload")
        return
    }
    
    req, err := http.NewRequest("POST", "https://api.thecatapi.com/v1/favourites", bytes.NewBuffer(jsonPayload))
    if err != nil {
        c.CustomAbort(500, "Failed to create request")
        return
    }
    
    req.Header.Set("x-api-key", apiKey)
    req.Header.Set("Content-Type", "application/json")
    
    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        fmt.Println("Error making request:", err)
        c.CustomAbort(500, "Failed to make API request")
        return
    }
    defer resp.Body.Close()
    
    respBody, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        fmt.Println("Error reading response:", err)
        c.CustomAbort(500, "Failed to read response")
        return
    }
    
    // fmt.Println("API Response:", string(respBody))
    
    if resp.StatusCode != http.StatusOK && resp.StatusCode != http.StatusCreated {
        c.CustomAbort(resp.StatusCode, string(respBody))
        return
    }
    
    c.Data["json"] = json.RawMessage(respBody)
    c.ServeJSON()
}

func (c *CatController) GetFavourites() {
    apiKey, err := beego.AppConfig.String("api_key")
    subID := c.GetString("sub_id")

    url := "https://api.thecatapi.com/v1/favourites"
    if subID != "" {
        url += "?sub_id=" + subID
    }

    req, err := http.NewRequest("GET", url, nil)
    if err != nil {
        c.CustomAbort(500, "Failed to create request")
    }
    req.Header.Set("x-api-key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.CustomAbort(500, "Failed to make API request")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        c.CustomAbort(resp.StatusCode, string(body))
    }

    var favourites []map[string]interface{}
    if err := json.NewDecoder(resp.Body).Decode(&favourites); err != nil {
        c.CustomAbort(500, "Failed to parse response")
    }

    c.Data["json"] = favourites
    c.ServeJSON()
}

func (c *CatController) DeleteFavourite() {
    apiKey, err := beego.AppConfig.String("api_key")
    favouriteID := c.Ctx.Input.Param(":favouriteId")

    if favouriteID == "" {
        c.CustomAbort(400, "favouriteId is required")
    }

    url := "https://api.thecatapi.com/v1/favourites/" + favouriteID
    req, err := http.NewRequest("DELETE", url, nil)
    if err != nil {
        c.CustomAbort(500, "Failed to create request")
    }
    req.Header.Set("x-api-key", apiKey)

    client := &http.Client{}
    resp, err := client.Do(req)
    if err != nil {
        c.CustomAbort(500, "Failed to make API request")
    }
    defer resp.Body.Close()

    if resp.StatusCode != http.StatusOK {
        body, _ := ioutil.ReadAll(resp.Body)
        c.CustomAbort(resp.StatusCode, string(body))
    }

    c.Data["json"] = map[string]string{"message": "Favourite deleted successfully"}
    c.ServeJSON()
}
/// FAVOURITES CODE 👆