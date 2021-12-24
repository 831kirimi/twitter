package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

func getTwitterUserId(screen_name string, bearer string) uint64 {
	client := &http.Client{}
	url := "https://api.twitter.com/1.1/users/lookup.json"

	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)
	req.Header.Set("Authorization", bearer)
	params := req.URL.Query()
	params.Add("screen_name", screen_name)
	req.URL.RawQuery = params.Encode()

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var f interface{}
	json.Unmarshal(body, &f)
	m := f.([]interface{})[0].(map[string]interface{})
	idStr := m["id_str"].(string)
	id, _ := strconv.ParseUint(idStr, 10, 64)
	return id
}

func getTwitterScreenName(userId uint64, bearer string) string {
	client := &http.Client{}
	url := "https://api.twitter.com/1.1/users/lookup.json"

	req, _ := http.NewRequest(
		"GET",
		url,
		nil,
	)
	req.Header.Set("Authorization", bearer)
	params := req.URL.Query()
	userIdStr := strconv.Itoa(int(userId))
	params.Add("user_id", userIdStr)
	req.URL.RawQuery = params.Encode()

	resp, _ := client.Do(req)
	body, _ := ioutil.ReadAll(resp.Body)
	defer resp.Body.Close()

	var f interface{}
	json.Unmarshal(body, &f)
	m := f.([]interface{})[0].(map[string]interface{})
	screenName := m["screen_name"].(string)
	return screenName
}

func main() {
	screenName := "Twitter"
	userId := uint64(783214)
	bearer := "******"
	fmt.Println(getTwitterUserId(screenName, bearer))
	fmt.Println(getTwitterScreenName(userId, bearer))
}
