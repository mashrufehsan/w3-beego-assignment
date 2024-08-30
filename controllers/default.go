package controllers

import (
	"bytes"
	"encoding/json"
	"io"
	"net/http"

	"github.com/astaxie/beego"
)

type MainController struct {
	beego.Controller
}

func (c *MainController) Get() {
	c.TplName = "index.tpl"
}

type ShowBreedsController struct {
	beego.Controller
}

func (c *ShowBreedsController) Get() {
	c.TplName = "breeds.tpl"
}

type ShowFavsController struct {
	beego.Controller
}

func (c *ShowFavsController) Get() {
	c.TplName = "favs.tpl"
}

type ShowMyVotesController struct {
	beego.Controller
}

func (c *ShowMyVotesController) Get() {
	c.TplName = "myVotes.tpl"
}

// FetchData fetches data from the given URL asynchronously using channels.
func FetchData(apiURL string, resultChan chan<- interface{}, errChan chan<- error) {
	resp, err := http.Get(apiURL)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}

	var result interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		errChan <- err
		return
	}

	resultChan <- result
}

func PostData(apiURL string, requestBody []byte, resultChan chan<- interface{}, errChan chan<- error) {
	req, err := http.NewRequest("POST", apiURL, bytes.NewReader(requestBody))
	if err != nil {
		errChan <- err
		return
	}
	req.Header.Set("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}

	var jsonResponse map[string]interface{}
	if json.Unmarshal(body, &jsonResponse) == nil {
		resultChan <- jsonResponse
	} else {
		resultChan <- string(body)
	}
}

func DeleteData(apiURL string, resultChan chan<- interface{}, errChan chan<- error) {
	req, err := http.NewRequest("DELETE", apiURL, nil)
	if err != nil {
		errChan <- err
		return
	}

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		errChan <- err
		return
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		errChan <- err
		return
	}

	var result interface{}
	// Try to unmarshal the response into a general interface{}
	err = json.Unmarshal(body, &result)
	if err != nil {
		// If unmarshalling fails, treat the body as a string
		resultChan <- string(body)
		return
	}

	resultChan <- result
}
