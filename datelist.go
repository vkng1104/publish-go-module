package datelist

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"strconv"
)

func urlencode(data map[string]string) string {
	var buf bytes.Buffer
	for k, v := range data {
		buf.WriteString(url.PathEscape(k))
		buf.WriteByte('=')
		buf.WriteString(url.PathEscape(v))
		buf.WriteByte('&')
	}
	s := buf.String()
	return s[0 : len(s)-1]
}

type DatelistClient struct {
	api_key string
}

func (c DatelistClient) ListCalendars(filters map[string]string) ([]map[string]interface{}, error) {
	url := "https://datelist.io/api/calendars"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var val []map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &val)
	return val, err
}

func (c DatelistClient) ListProducts(filters map[string]string) ([]map[string]interface{}, error) {
	url := "https://datelist.io/api/products?"
	if filters != nil {
		url = url + urlencode(filters)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var val []map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &val)
	return val, err
}

func (c DatelistClient) ListBookedSlots(filters map[string]string) ([]map[string]interface{}, error) {
	url := "https://datelist.io/api/booked_slots?"
	if filters != nil {
		url = url + urlencode(filters)
	}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var val []map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &val)
	return val, err
}

func (c DatelistClient) UpdateBookedSlot(id float64, data map[string]string) (map[string]interface{}, error) {
	url := "https://datelist.io/api/booked_slots/" + strconv.Itoa(int(id))

	r, w := io.Pipe()
	go func() {
		json.NewEncoder(w).Encode(data)
		w.Close()
	}()

	req, err := http.NewRequest("PATCH", url, r)
	if err != nil {
		return nil, err
	}

	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/json")
	req.Header.Add("Authorization", c.api_key)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}

	var val map[string]interface{}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	err = json.Unmarshal([]byte(body), &val)
	return val, err
}

func Client(api_key string) DatelistClient {
	return DatelistClient{
		api_key: api_key,
	}
}
