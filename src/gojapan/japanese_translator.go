package gojapan

import (
	"net/http"
	"fmt"
	"io/ioutil"
	"encoding/json"
)

const translateUrl = "https://translate.google.co.uk/translate_a/single?client=t&sl=auto&tl=ja&hl=en&dt=at&client=tw-ob&q="

type translator interface {
	translate(u string) string
}

type translation struct {
	english, japanese string
}

// Translate an English string into Japanese
func translate(s string) (translation) {
	httpGetResponse := safeHttpGet(translateUrl + s)
	body, err := ioutil.ReadAll(httpGetResponse.Body)
	defer httpGetResponse.Body.Close()
	if err != nil {
		printAndPanic(err, "error translating: " + s)
	}
	fmt.Println(string(body[:]))
	result := translation{english:s, japanese:extractResultFromJson(string(body[:]), s)}
	return result
}

// Safely invoke an HTTP GET call on the supplied URL
func safeHttpGet(u string) (resp *http.Response) {
	resp, err := http.Get(u);
	if err != nil {
		printAndPanic(err, "error getting: " + u)
	}
	return resp
}

// Delve into Google's translation response and fetch the result
func extractResultFromJson(input string, search string) (string) {
	// The json returned by google is 'almost' valid. Just a little tweak needed
	byt := []byte(`{"` + search + `" :` + input + "}")
	var dat map[string]interface{}

	if err := json.Unmarshal(byt, &dat); err != nil {
		printAndPanic(err, "error unmarshalling")
	}
	return dat[search].([]interface{})[5].([]interface{})[0].([]interface{})[2].([]interface{})[0].([]interface{})[0].(string)
}

// Utility method to print an error's details and
func printAndPanic(err error, details string) {
	// Panic will print to stderr but it might be useful to have the error in the configured log, should we ever add one
	fmt.Printf("There's been a terrible error! %s %s", details, err)
	panic(err)
}