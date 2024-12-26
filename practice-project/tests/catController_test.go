package tests

import (
	"bytes"
	"encoding/json"
	"io"
    "io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/beego/beego/v2/server/web/context"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"

	"practice-project/controllers"
)

// MockHTTPClient is our mock HTTP client
type MockHTTPClient struct {
	mock.Mock
}

func (m *MockHTTPClient) Do(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// mockHTTPTransport implements http.RoundTripper interface
type mockHTTPTransport struct {
	mock.Mock
}

func (m *mockHTTPTransport) RoundTrip(req *http.Request) (*http.Response, error) {
	args := m.Called(req)
	return args.Get(0).(*http.Response), args.Error(1)
}

// setupTestContext creates a new context for testing
func setupTestContext(method, path string, body io.Reader) (*controllers.CatController, *httptest.ResponseRecorder) {
	// Create a new test recorder
	w := httptest.NewRecorder()

	// Create a new request
	req := httptest.NewRequest(method, path, body)
	if method == "POST" {
		req.Header.Set("Content-Type", "application/json")
	}

	// Create a new context
	ctx := context.NewContext()
	ctx.Reset(w, req)

	// Initialize the input
	ctx.Input.Context = ctx

	// Create and initialize the controller
	controller := &controllers.CatController{}
	controller.Init(ctx, "CatController", "TestMethod", nil)

	return controller, w
}

// Mock function for fetchImagesByBreed
func MockFetchImagesByBreed(breedID string, imageCh chan []controllers.CatResponse, errCh chan error) {
	mockImages := []controllers.CatResponse{
		{URL: "http://example.com/cat1.jpg", ID: breedID},
		{URL: "http://example.com/cat2.jpg", ID: breedID},
	}
	imageCh <- mockImages
	close(imageCh)
	close(errCh)
}

// TestShowCat tests the ShowCat handler
func TestShowCat(t *testing.T) {
	// Setup test context
	controller, _ := setupTestContext("GET", "/cat", nil)

	// Create mock response
	mockJSON := `[{"url":"http://example.com/cat.jpg","id":"test123","width":500,"height":400}]`
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	// Create and configure mock transport
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.Anything).Return(mockResponse, nil)

	// Replace default transport
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call the handler
	controller.ShowCat()

	// Assert response
	assert.Equal(t, "index.tpl", controller.TplName)
	assert.Equal(t, "http://example.com/cat.jpg", controller.Data["CatImage"])
}

//getcat data



// TestVoteUp tests the VoteUp handler
func TestVoteUp(t *testing.T) {
	// Setup test context
	controller, w := setupTestContext("POST", "/vote/up?image_id=test123", nil)

	// Mock getCurrentVote response
	mockVoteJSON := `[{"image_id":"test123","value":0}]`
	mockVoteResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockVoteJSON)),
	}

	// Mock vote submission response
	mockSubmitJSON := `{"message":"success"}`
	mockSubmitResponse := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(bytes.NewBufferString(mockSubmitJSON)),
	}

	// Configure mock transport
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.MatchedBy(func(req *http.Request) bool {
		return req.Method == "GET"
	})).Return(mockVoteResponse, nil)
	mockTransport.On("RoundTrip", mock.MatchedBy(func(req *http.Request) bool {
		return req.Method == "POST"
	})).Return(mockSubmitResponse, nil)

	// Replace transport
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call handler
	controller.VoteUp()

	// Decode response
	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)

	// Assert response
	assert.Equal(t, "Vote toggled successfully", response["message"])
}
func TestVoteDown(t *testing.T) {
	// Setup test context
	controller, w := setupTestContext("POST", "/vote/down?image_id=test123", nil)

	// Mock getCurrentVote response
	mockVoteJSON := `[{"image_id":"test123","value":1}]` // Assuming the current vote is 1 (upvote)
	mockVoteResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockVoteJSON)),
	}

	// Mock vote submission response
	mockSubmitJSON := `{"message":"success"}`
	mockSubmitResponse := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(bytes.NewBufferString(mockSubmitJSON)),
	}

	// Configure mock transport
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.MatchedBy(func(req *http.Request) bool {
		return req.Method == "GET"
	})).Return(mockVoteResponse, nil)
	mockTransport.On("RoundTrip", mock.MatchedBy(func(req *http.Request) bool {
		return req.Method == "POST"
	})).Return(mockSubmitResponse, nil)

	// Replace transport
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call handler
	controller.VoteDown()

	// Decode response
	var response map[string]string
	json.NewDecoder(w.Body).Decode(&response)

	// Assert response
	assert.Equal(t, "Vote toggled successfully", response["message"])
}
func TestVoteHistory_Success(t *testing.T) {
    // Mock API response data
    mockResponse := []controllers.VoteHistoryResponse{
        {ID: 1, ImageID: "abc123", Value: 1, SubID: "test123"},
        {ID: 2, ImageID: "xyz789", Value: 0, SubID: "test123"},
    }

    // Create a mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        assert.Equal(t, "live_8Vq87uY7jXkcqmqwhODWVdzEp9iUzbog1G0hxJgh6gphgTP9sjK23Pbnir5Xl5JY", r.Header.Get("x-api-key"))
        assert.Contains(t, r.URL.String(), "sub_id=test123")
        w.WriteHeader(http.StatusOK)
        json.NewEncoder(w).Encode(mockResponse)
    }))
    defer server.Close()

    // Replace the API URL with the mock server's URL
    originalURL := controllers.APIBaseURL // Assume `APIBaseURL` is a variable holding the base URL in your app
    controllers.APIBaseURL = server.URL
    defer func() { controllers.APIBaseURL = originalURL }()

    // Setup the test
    controller, w := setupTestContext("GET", "/vote-history", nil) // Replace `setupTestContext` with your helper function for tests

    // Call the VoteHistory method
    controller.VoteHistory()

    // Decode the response
    var response []controllers.VoteHistoryResponse
    err := json.NewDecoder(w.Body).Decode(&response)
    assert.NoError(t, err)
    assert.Equal(t, mockResponse, response)
}

func TestVoteHistory_Failure(t *testing.T) {
    // Create a mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusInternalServerError)
    }))
    defer server.Close()

    // Replace the API URL with the mock server's URL
    originalURL := controllers.APIBaseURL
    controllers.APIBaseURL = server.URL
    defer func() { controllers.APIBaseURL = originalURL }()

    // Setup the test
    controller, w := setupTestContext("GET", "/vote-history", nil)

    // Call the VoteHistory method
    controller.VoteHistory()

    // Decode the response
    var response map[string]string
    err := json.NewDecoder(w.Body).Decode(&response)
    assert.NoError(t, err)
    assert.Equal(t, "Failed to fetch vote history", response["error"])
}

func TestVoteHistory_DecodeError(t *testing.T) {
    // Create a mock server
    server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
        w.WriteHeader(http.StatusOK)
        w.Write([]byte(`invalid json`)) // Send invalid JSON
    }))
    defer server.Close()

    // Replace the API URL with the mock server's URL
    originalURL := controllers.APIBaseURL
    controllers.APIBaseURL = server.URL
    defer func() { controllers.APIBaseURL = originalURL }()

    // Setup the test
    controller, w := setupTestContext("GET", "/vote-history", nil)

    // Call the VoteHistory method
    controller.VoteHistory()

    // Decode the response
    var response map[string]string
    err := json.NewDecoder(w.Body).Decode(&response)
    assert.NoError(t, err)
    assert.Equal(t, "Failed to decode response", response["error"])
}
// TestCreateFavourite tests the CreateFavourite handler
func TestCreateFavourite(t *testing.T) {
	// Create test request body
	requestBody := map[string]string{
		"image_id": "test123",
		"sub_id":   "user123",
	}
	bodyBytes, _ := json.Marshal(requestBody)

	// Setup test context
	controller, w := setupTestContext("POST", "/favourites", bytes.NewBuffer(bodyBytes))

	// Mock API response
	mockJSON := `{"id": 1, "message": "SUCCESS"}`
	mockResponse := &http.Response{
		StatusCode: http.StatusCreated,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	// Configure mock transport
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.Anything).Return(mockResponse, nil)

	// Replace transport
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call handler
	controller.CreateFavourite()

	// Assert response code
	assert.Equal(t, http.StatusOK, w.Code)

	// Decode and verify response
	var response map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)
	assert.Equal(t, float64(1), response["id"])
	assert.Equal(t, "SUCCESS", response["message"])
}

// TestGetFavourites tests the GetFavourites handler
func TestGetFavourites(t *testing.T) {
	// Setup test context
	controller, w := setupTestContext("GET", "/favourites?sub_id=test123", nil)

	// Mock response from The Cat API
	mockJSON := `[{"id": 1, "image_id": "test123", "sub_id": "test123"}]`
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       io.NopCloser(bytes.NewBufferString(mockJSON)),
	}

	// Configure mock transport
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.Anything).Return(mockResponse, nil)

	// Replace default transport
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call the method
	controller.GetFavourites()

	// Decode response
	var response []map[string]interface{}
	json.NewDecoder(w.Body).Decode(&response)

	// Assertions
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Len(t, response, 1)
	assert.Equal(t, float64(1), response[0]["id"])
	assert.Equal(t, "test123", response[0]["image_id"])
}

// Mock function to simulate fetching breeds from the API
func MockFetchBreeds(ch chan<- []controllers.Breed, errCh chan<- error) {
	mockBreeds := []controllers.Breed{
		{
			ID:          "abys",
			Name:        "Abyssinian",
			Temperament: "Active, Energetic, Independent",
			Origin:      "Egypt",
			Description: "Test description",
		},
	}
	ch <- mockBreeds
	close(ch)
	close(errCh)
}


// TestFetchBreeds tests the fetchBreeds function
func TestFetchBreeds(t *testing.T) {
	ch := make(chan []controllers.Breed)
	errCh := make(chan error)

	// Simulate fetchBreeds
	go MockFetchBreeds(ch, errCh)

	select {
	case breeds := <-ch:
		assert.NotEmpty(t, breeds)
		assert.Equal(t, "abys", breeds[0].ID)
		assert.Equal(t, "Abyssinian", breeds[0].Name)
	case err := <-errCh:
		t.Fatal(err)
	}
}

// TestFetchImagesByBreed tests the fetchImagesByBreed function
func TestFetchImagesByBreed(t *testing.T) {
	ch := make(chan []controllers.CatResponse)
	errCh := make(chan error)

	// Simulate fetchImagesByBreed
	go MockFetchImagesByBreed("abys", ch, errCh)

	select {
	case images := <-ch:
		assert.NotEmpty(t, images)
		assert.Equal(t, "http://example.com/cat1.jpg", images[0].URL)
		assert.Equal(t, "abys", images[0].ID)
	case err := <-errCh:
		t.Fatal(err)
	}
}

// TestFetchAndStoreBreeds tests the fetchAndStoreBreeds function
func TestFetchAndStoreBreeds(t *testing.T) {
	// Setup channels
	ch := make(chan []controllers.Breed)
	errCh := make(chan error)

	// Mock the fetchBreeds function with the channel
	go MockFetchBreeds(ch, errCh)

	// Simulate storing breeds
	// go controllers.FetchAndStoreBreeds()

	// Assert the channel has received data
	select {
	case breeds := <-ch:
		assert.NotEmpty(t, breeds)
		assert.Equal(t, "abys", breeds[0].ID)
		assert.Equal(t, "Abyssinian", breeds[0].Name)
	case err := <-errCh:
		t.Fatal(err)
	}
}

// TestGetBreedsHandler tests the GetBreedsHandler method
func TestGetBreedsHandler(t *testing.T) {
	// Setup the request and response recorder
	controller, w := setupTestContext("GET", "/breeds", nil)

	// Mock the fetchBreeds function
	mockBreeds := []controllers.Breed{
		{
			ID:          "abys",
			Name:        "Abyssinian",
			Temperament: "Active, Energetic, Independent",
			Origin:      "Egypt",
			Description: "Test description",
		},
	}
	mockJSON, _ := json.Marshal(mockBreeds)
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(mockJSON)),
	}

	// Replace transport with mock response
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.Anything).Return(mockResponse, nil)

	// Replace default transport with mock
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call the handler
	controller.GetBreedsHandler()

	// Decode response and assert
	var response []controllers.Breed
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, mockBreeds, response)
}

// TestGetBreedImagesHandler tests the GetBreedImagesHandler method
func TestGetBreedImagesHandler(t *testing.T) {
	// Setup the request and response recorder
	controller, w := setupTestContext("GET", "/breed/images", nil)

	// Mock the image fetch function
	mockImages := []controllers.CatResponse{
		{URL: "http://example.com/cat1.jpg", ID: "abys"},
		{URL: "http://example.com/cat2.jpg", ID: "abys"},
	}
	mockJSON, _ := json.Marshal(mockImages)
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(mockJSON)),
	}

	// Replace transport with mock response
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.Anything).Return(mockResponse, nil)

	// Replace default transport with mock
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call the handler
	controller.GetBreedImagesHandler()

	// Decode response and assert
	var response []controllers.CatResponse
	json.NewDecoder(w.Body).Decode(&response)

	assert.Equal(t, mockImages, response)
}

// TestFetchImagesByBreedHandler tests the FetchImagesByBreedHandler method
func TestFetchImagesByBreedHandler(t *testing.T) {
	// Setup the request and response recorder
	controller, w := setupTestContext("GET", "/breed/images/abys", nil)

	// Mock the image fetch response data
	mockImages := []controllers.CatResponse{
		{URL: "http://example.com/cat1.jpg", ID: "abys"},
		{URL: "http://example.com/cat2.jpg", ID: "abys"},
	}
	mockJSON, _ := json.Marshal(mockImages)
	mockResponse := &http.Response{
		StatusCode: http.StatusOK,
		Body:       ioutil.NopCloser(bytes.NewBuffer(mockJSON)),
	}

	// Create a mock transport
	mockTransport := &mockHTTPTransport{}
	mockTransport.On("RoundTrip", mock.Anything).Return(mockResponse, nil)

	// Replace default transport with the mock
	origTransport := http.DefaultTransport
	http.DefaultTransport = mockTransport
	defer func() { http.DefaultTransport = origTransport }()

	// Call the handler (this will trigger the HTTP call inside the handler)
	controller.FetchImagesByBreedHandler()

	// Decode response and assert
	var response []controllers.CatResponse
	err := json.NewDecoder(w.Body).Decode(&response)
	if err != nil {
		t.Fatal(err)
	}

	// Assert that the response matches the mocked data
	assert.Equal(t, mockImages, response)

	// Ensure that the mock transport's RoundTrip method was called
	mockTransport.AssertExpectations(t)
}
