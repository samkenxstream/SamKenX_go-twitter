package twitter

import (
	"fmt"
	"net/http"
	"reflect"
	"testing"
)

func TestStatusService_Show(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"id": "589488862814076930", "include_entities": "false"}, r)
		fmt.Fprintf(w, `{"user": {"screen_name": "dghubble"}, "text": ".@audreyr use a DONTREADME file if you really want people to read it :P"}`)
	})

	client := NewClient(httpClient)
	params := &StatusShowParams{Id: 5441, IncludeEntities: Bool(false)}
	tweets, _, err := client.Statuses.Show(589488862814076930, params)
	if err != nil {
		t.Errorf("Statuses.Show error %v", err)
	}
	expected := &Tweet{User: &User{ScreenName: "dghubble"}, Text: ".@audreyr use a DONTREADME file if you really want people to read it :P"}
	if !reflect.DeepEqual(expected, tweets) {
		t.Errorf("Statuses.Show expected:\n%+v, got:\n %+v", expected, tweets)
	}
}

func TestStatusService_ShowHandlesNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/show.json", func(w http.ResponseWriter, r *http.Request) {
		assertQuery(t, map[string]string{"id": "589488862814076930"}, r)
	})
	client := NewClient(httpClient)
	client.Statuses.Show(589488862814076930, nil)
}

func TestStatusService_Lookup(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/lookup.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "GET", r)
		assertQuery(t, map[string]string{"id": "20,573893817000140800", "trim_user": "true"}, r)
		fmt.Fprintf(w, `[{"id": 20, "text": "just setting up my twttr"}, {"id": 573893817000140800, "text": "Don't get lost #PaxEast2015"}]`)
	})

	client := NewClient(httpClient)
	params := &StatusLookupParams{Id: []int64{20}, TrimUser: Bool(true)}
	tweets, _, err := client.Statuses.Lookup([]int64{573893817000140800}, params)
	if err != nil {
		t.Errorf("Statuses.Lookup error %v", err)
	}
	expected := []Tweet{Tweet{Id: 20, Text: "just setting up my twttr"}, Tweet{Id: 573893817000140800, Text: "Don't get lost #PaxEast2015"}}
	if !reflect.DeepEqual(expected, tweets) {
		t.Errorf("Statuses.Lookup expected:\n%+v, got:\n %+v", expected, tweets)
	}
}

func TestStatusService_LookupHandlesNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/1.1/statuses/lookup.json", func(w http.ResponseWriter, r *http.Request) {
		assertQuery(t, map[string]string{"id": "20,573893817000140800"}, r)
	})
	client := NewClient(httpClient)
	client.Statuses.Lookup([]int64{20, 573893817000140800}, nil)
}

func TestStatusService_Update(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()

	mux.HandleFunc("/1.1/statuses/update.json", func(w http.ResponseWriter, r *http.Request) {
		assertMethod(t, "POST", r)
		assertQuery(t, map[string]string{}, r)
		assertPostForm(t, map[string]string{"status": "very informative tweet", "media_ids": "123456789,987654321"}, r)
		fmt.Fprintf(w, `{"id": 581980947630845953, "text": "very informative tweet"}`)
	})

	client := NewClient(httpClient)
	params := &StatusUpdateParams{MediaIds: []int64{123456789, 987654321}}
	tweet, _, err := client.Statuses.Update("very informative tweet", params)
	if err != nil {
		t.Errorf("Statuses.Update error %v", err)
	}
	expected := &Tweet{Id: 581980947630845953, Text: "very informative tweet"}
	if !reflect.DeepEqual(expected, tweet) {
		t.Errorf("Statuses.Update expected:\n%+v, got:\n %+v", expected, tweet)
	}
}

func TestStatusService_UpdateHandlesNilParams(t *testing.T) {
	httpClient, mux, server := testServer()
	defer server.Close()
	mux.HandleFunc("/1.1/statuses/update.json", func(w http.ResponseWriter, r *http.Request) {
		assertPostForm(t, map[string]string{"status": "very informative tweet"}, r)
	})
	client := NewClient(httpClient)
	client.Statuses.Update("very informative tweet", nil)
}