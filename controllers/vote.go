package controllers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"

	"github.com/astaxie/beego"
)

// VoteController operations for Vote
type VoteController struct {
	beego.Controller
}

// URLMapping ...
func (c *VoteController) URLMapping() {
	c.Mapping("GetVotes", c.GetVotes)
	c.Mapping("Vote", c.Vote)
}

func (c *VoteController) GetVotes() {
	// Read the API key from the Beego configuration
	apiKey := beego.AppConfig.String("CatAPIKey")
	subID := beego.AppConfig.String("SubID")

	// Construct the API URL
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/votes?sub_id=%s&order=DESC&api_key=%s", subID, apiKey)

	// Create channels for results and errors
	resultChan := make(chan interface{})
	errChan := make(chan error)

	// Start the API call in a separate goroutine
	go FetchData(apiURL, resultChan, errChan)

	// Set a timeout for the API request
	select {
	case result := <-resultChan:
		// Handle successful response
		c.Data["json"] = result
		c.ServeJSON()

	case err := <-errChan:
		// Handle error
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

func (c *VoteController) Vote() {
	body, err := io.ReadAll(c.Ctx.Request.Body)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Failed to read request body"}
		c.ServeJSON()
		return
	}

	var requestBody map[string]interface{}
	if err := json.Unmarshal(body, &requestBody); err != nil {
		c.Ctx.Output.SetStatus(http.StatusBadRequest)
		c.Data["json"] = map[string]string{"error": "Invalid request body"}
		c.ServeJSON()
		return
	}

	apiKey := beego.AppConfig.String("CatAPIKey")
	subID := beego.AppConfig.String("SubID")
	apiURL := fmt.Sprintf("https://api.thecatapi.com/v1/votes?sub_id=%s&api_key=%s", subID, apiKey)

	// Add sub_id to requestBody
	requestBody["sub_id"] = subID

	body, err = json.Marshal(requestBody)
	if err != nil {
		c.Ctx.Output.SetStatus(http.StatusInternalServerError)
		c.Data["json"] = map[string]string{"error": "Error marshaling request body"}
		c.ServeJSON()
		return
	}

	resultChan := make(chan interface{})
	errChan := make(chan error)

	go PostData(apiURL, body, resultChan, errChan)

	select {
	case result := <-resultChan:
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
