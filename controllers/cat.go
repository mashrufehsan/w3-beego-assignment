package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

type CatController struct {
	beego.Controller
}

func (c *CatController) URLMapping() {
	c.Mapping("GetACat", c.GetACat)
	c.Mapping("CreateAFavourite", c.CreateAFavourite)
	c.Mapping("GetFavourites", c.GetFavourites)
	c.Mapping("DeleteAFavourite", c.DeleteAFavourite)
}

func (c *CatController) GetACat() {
	apiURL := "https://api.thecatapi.com/v1/images/search"

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go FetchData(apiURL, resultChan, errChan)

	// Set a timeout for the API request
	select {
	case result := <-resultChan:
		if images, ok := result.([]interface{}); ok && len(images) > 0 {
			if imageMap, ok := images[0].(map[string]interface{}); ok {
				id, idOk := imageMap["id"].(string)
				url, urlOk := imageMap["url"].(string)
				if idOk && urlOk {
					c.Data["json"] = map[string]string{
						"id":  id,
						"url": url,
					}
				} else {
					c.Ctx.Output.SetStatus(http.StatusInternalServerError)
					c.Data["json"] = map[string]string{"error": "Failed to parse id or url"}
				}
			} else {
				c.Ctx.Output.SetStatus(http.StatusInternalServerError)
				c.Data["json"] = map[string]string{"error": "Unexpected data structure"}
			}
		} else {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Data["json"] = map[string]string{"error": "No image data found"}
		}
		c.ServeJSON()

	case err := <-errChan:
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()

	case <-time.After(10 * time.Second):
		// Timeout after 10 seconds
		c.Ctx.Output.SetStatus(http.StatusRequestTimeout)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
		c.ServeJSON()
	}
}

func (c *CatController) CreateAFavourite() {
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to read request body"}
		c.ServeJSON()
		return
	}

	// Unmarshal the JSON body
	var requestBody map[string]string
	if err := json.Unmarshal(body, &requestBody); err != nil {
		fmt.Println("Error unmarshalling request body:", err)
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	apiKey := beego.AppConfig.String("CatAPIKey")
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?api_key=%s", apiKey)

	// Marshal the request body to send to the external API
	body, err = json.Marshal(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error marshaling request body"}
		c.ServeJSON()
		return
	}

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go PostData(apiURL, body, resultChan, errChan)

	// Handle results and errors with a timeout
	select {
	case result := <-resultChan:
		// Use type switch to handle different result types
		switch res := result.(type) {
		case map[string]interface{}:
			c.Data["json"] = res
		case string:
			c.Data["json"] = map[string]string{"message": res}
		default:
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Data["json"] = map[string]string{"error": "Unexpected response format"}
		}
		c.ServeJSON()
	case <-errChan:
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Failed to make POST request"}
		c.ServeJSON()
	case <-time.After(10 * time.Second):
		c.Ctx.Output.SetStatus(http.StatusRequestTimeout)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
		c.ServeJSON()
	}
}

func (c *CatController) GetFavourites() {
	apiKey := beego.AppConfig.String("CatAPIKey")
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/favourites?sub_id=mashruf&order=DESC&api_key=%s", apiKey)

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go FetchData(apiURL, resultChan, errChan)

	// Set a timeout for the API request
	select {
	case result := <-resultChan:
		if favourites, ok := result.([]interface{}); ok {
			// Process the favourites data as needed
			c.Data["json"] = favourites
		} else {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Data["json"] = map[string]string{"error": "Unexpected data structure"}
		}
		c.ServeJSON()

	case err := <-errChan:
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": err.Error()}
		c.ServeJSON()

	case <-time.After(10 * time.Second):
		// Timeout after 10 seconds
		c.Ctx.Output.SetStatus(http.StatusRequestTimeout)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
		c.ServeJSON()
	}
}

func (c *CatController) DeleteAFavourite() {
	// Extract the favouriteId from the URL parameter
	favouriteID := c.Ctx.Input.Param(":favourite_id")

	// Read the API key from the Beego configuration
	apiKey := beego.AppConfig.String("CatAPIKey")

	// Construct the API URL
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/favourites/%s?api_key=%s", favouriteID, apiKey)

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the DELETE request in a separate goroutine
	go DeleteData(apiURL, resultChan, errChan)

	// Set a timeout for the API request
	select {
	case result := <-resultChan:
		// Handle successful response
		c.Data["json"] = result
		c.ServeJSON()

	case <-errChan:
		// Handle error
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Failed to fetch data from TheCatAPI"}
		c.ServeJSON()

	case <-time.After(10 * time.Second):
		// Timeout after 10 seconds
		c.Ctx.Output.SetStatus(http.StatusRequestTimeout)
		c.Data["json"] = map[string]string{"error": "Request timed out"}
		c.ServeJSON()
	}
}
