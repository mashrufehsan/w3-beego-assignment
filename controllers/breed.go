package controllers

import (
	"fmt"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

type BreedController struct {
	beego.Controller
}

func (c *BreedController) URLMapping() {
	c.Mapping("GetBreeds", c.GetBreeds)
	c.Mapping("GetBreedsByID", c.GetBreedsByID)
	c.Mapping("GetImagesByBreed", c.GetImagesByBreed)
}

func (c *BreedController) GetBreeds() {
	apiURL := "https://api.thecatapi.com/v1/breeds"

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go FetchData(apiURL, resultChan, errChan)

	// Set a timeout for the API request
	select {
	case result := <-resultChan:
		if breeds, ok := result.([]interface{}); ok {
			filteredBreeds := []map[string]string{}
			for _, breed := range breeds {
				if breedMap, ok := breed.(map[string]interface{}); ok {
					id, idOk := breedMap["id"].(string)
					name, nameOk := breedMap["name"].(string)
					if idOk && nameOk {
						filteredBreeds = append(filteredBreeds, map[string]string{
							"id":   id,
							"name": name,
						})
					}
				}
			}
			c.Data["json"] = filteredBreeds
		} else {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Data["json"] = map[string]string{"error": "Unexpected data structure"}
		}
		c.ServeJSON()

	case <-errChan:
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

func (c *BreedController) GetBreedsByID() {
	breedID := c.Ctx.Input.Param(":breed_id")
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/breeds/%s", breedID)

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go FetchData(apiURL, resultChan, errChan)

	select {
	case result := <-resultChan:
		if breedMap, ok := result.(map[string]interface{}); ok {
			filteredBreed := map[string]interface{}{
				"id":            breedMap["id"],
				"name":          breedMap["name"],
				"origin":        breedMap["origin"],
				"description":   breedMap["description"],
				"wikipedia_url": breedMap["wikipedia_url"],
			}

			c.Data["json"] = filteredBreed
		} else {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Data["json"] = map[string]string{"error": "Unexpected data structure"}
		}
		c.ServeJSON()

	case <-errChan:
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

func (c *BreedController) GetImagesByBreed() {
	breedID := c.Ctx.Input.Param(":breed_id")
	apiKey := beego.AppConfig.String("CatAPIKey")

	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/images/search?limit=5&breed_ids=%s&api_key=%s", breedID, apiKey)

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go FetchData(apiURL, resultChan, errChan)

	// Set a timeout for the API request
	select {
	case result := <-resultChan:
		if images, ok := result.([]interface{}); ok {
			imageURLs := []string{}
			for _, image := range images {
				if imageMap, ok := image.(map[string]interface{}); ok {
					if url, ok := imageMap["url"].(string); ok {
						imageURLs = append(imageURLs, url)
					}
				}
			}

			c.Data["json"] = imageURLs
		} else {
			c.Ctx.Output.SetStatus(http.StatusInternalServerError)
			c.Data["json"] = map[string]string{"error": "Unexpected data structure"}
		}
		c.ServeJSON()

	case <-errChan:
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
