package cloudcraft

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
	"net/url"
)

//RequestResponse Recieve Request and Respond with Response Body.
//The response body will contain valid error or valid response.
func (c *Client) RequestResponse(method, path string, reqbody, out interface{}) error {

	//create a request
	req, err := c.createRequest(method, path, reqbody)
	if err != nil {
		log.Printf("[ERROR] Create Request Failed %s", err)
		return err
	}
	//send the request
	log.Printf("[DEBUG] Request Path %s", path)
	log.Printf("[DEBUG] Request Body %s", reqbody)

	resp, err := c.sendRequest(req)

	if err != nil {
		log.Printf("[ERROR] Request Failed %s", err)
		return err
	}
	//close the response body
	defer resp.Body.Close()

	//check if the status code is outside a succcessful response and if it is populate body with custom error msg
	if resp.StatusCode < 200 || resp.StatusCode > 299 {
		body, err := ioutil.ReadAll(resp.Body)
		if err != nil {
			log.Printf("[ERROR] Error Reading error resp body")
			return err
		}
		return fmt.Errorf("[ERROR] API error %s: %s", resp.Status, body)
	}

	respbody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	if len(respbody) == 0 {
		respbody = []byte{'{', '}'}
	}

	// If they don't care about the body, then we don't care to give them one,
	// so bail out because we're done.
	if out == nil {
		return nil
	}

	return json.Unmarshal(respbody, &out)
}

func (c *Client) createRequest(method, path string, body interface{}) (*http.Request, error) {
	rel := &url.URL{Path: path}
	u := c.baseURL.ResolveReference(rel)

	//Check the request body is not nill and the method is not a GET
	var bodyReader io.Reader
	if method != "GET" && body != nil {
		bjson, err := json.Marshal(body)
		if err != nil {
			return nil, err
		}
		bodyReader = bytes.NewReader(bjson)
	}

	log.Printf("[DEBUG] URL Endpoint " + u.String())
	req, err := http.NewRequest(method, u.String(), bodyReader)
	if err != nil {
		return nil, err
	}

	req.Header.Set("Authorization", "Bearer "+c.apiKey)
	req.Header.Add("Content-Type", "application/json")
	return req, nil
}

func (c *Client) sendRequest(req *http.Request) (*http.Response, error) {
	resp, err := c.HTTPClient.Do(req)
	if err != nil {
		log.Printf("[ERROR] Request Failed. Check Credentials %s", err)
		return nil, err
	}
	log.Printf("[DEBUG] Finished Send Request")

	return resp, err
}
