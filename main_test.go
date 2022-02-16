package main

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/gorilla/websocket"
	"github.com/stretchr/testify/require"
)

var upgrader = websocket.Upgrader{}

func echo(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		return
	}
	defer c.Close()
	for {
		mt, message, err := c.ReadMessage()
		if err != nil {
			break
		}
		err = c.WriteMessage(mt, message)
		if err != nil {
			break
		}

		err = c.WriteJSON(message)
		if err != nil {
			break
		}

		err = c.ReadJSON(message)
		if err != nil {
			break
		}

	}
}
func TestGetResponseFromWebsocket(t *testing.T) {
	/* var testStruct struct {
		SomeTitle string `json:"some"`
	} */

	// Create test server with the echo handler.
	s := httptest.NewServer(http.HandlerFunc(echo))
	defer s.Close()

	// Convert http://127.0.0.1 to ws://127.0.0.
	u := "ws" + strings.TrimPrefix(s.URL, "http")

	// Connect to the server
	ws, _, err := websocket.DefaultDialer.Dial(u, nil)
	if err != nil {
		return
	}

	w, err := ws.NextWriter(websocket.TextMessage)

	if err != nil {
		return
	}

	var testResponse Response
	file, _ := ioutil.ReadFile("fullJson.json")
	testJson := json.Unmarshal([]byte(file), &testResponse)

	buf, err := json.Marshal(testJson)
	if err != nil {
		return
	}
	_, err = w.Write(buf)

	w.Close()

	if err != nil {
		t.Fatalf("%v", err)
	}
	defer ws.Close()

	{
		testResponse, err := getResponseFromWs(ws)
		if err != nil {
			require.ErrorIs(t, err, nil)
		}
		require.IsType(t, Response{}, testResponse)

	}
}
