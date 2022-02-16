package main

/*
	Requirements:
	- establish connection
	- subscribe to published feeds
	- handle messages
	- handle interupptions

*/

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"net/url"
	"os"
	"os/signal"
	"reflect"
	"strings"
	"time"

	"github.com/gorilla/websocket"
)

var sig os.Signal

// this the function that translates websocket response from a byte response into a JSON struct
func getResponseFromWs(conn *websocket.Conn) (Response, error) {
	var res Response

	err := conn.ReadJSON(&res)

	if err != nil {
		return Response{}, err
	}

	return res, nil
}

// this function gets the current supply of coin deonms in lastest websocket response
func getSupplyPricesFromWs(res Response) []string {
	// supplyJson gets the current amount in the total supply and maps to a JSON string to pretty print
	supplyJSON := make([]string, 0)
	for _, element := range res.BlockData.Supply {
		supply, err := json.MarshalIndent(element, "", "  ")
		if err != nil {
			log.Fatalf(err.Error())
		}
		supplyJSON = append(supplyJSON, string(supply))
	}
	return supplyJSON
}

// this function gets all the actions that have occurred in lastest websocket response
func getPastActionsFromWs(res Response) []string {
	mapPastActions := make([]string, 1)
	counter := 0
	// this pulls all the actionable messages from the
	for i := 0; i < len(res.BlockData.Transactions); i++ {
		pastActions := res.BlockData.Transactions[i].Body.Messages[0].ExecuteMsg
		if pastActions != nil {
			b := reflect.ValueOf(pastActions)

			for _, element := range b.MapKeys() {
				mapPastActions = append(mapPastActions, element.String())
				counter++
			}

		}
	}
	return mapPastActions

}

// this is an helper function for future webpage
func loadPage(title string) *Page {
	filename := title + ".txt"
	body, _ := os.ReadFile(filename)
	return &Page{Title: title, Body: body}
}

// this is an helper function for future webpage
func handler(w http.ResponseWriter, r *http.Request) {
	p := loadPage("Home")
	t, _ := template.ParseFiles("pages/mainPage.html")

	t.Execute(w, p)
}

func main() {

	// when it comes to using a goroutine,
	// i thought about adding a 'context'
	// variable. however, for the purpose of this
	// project, i will use the time.ticker functionality
	// primarily because I want the websocket call to be agnostic
	// and for ease of integration.

	// vars to hold responses
	// keeping os.Signal for terminal hardstops for now
	interrupt := make(chan os.Signal, 1)
	signal.Notify(interrupt, os.Interrupt)

	// connection to url
	u := url.URL{Scheme: "wss", Host: "observer.terra.dev", Path: "/"}
	log.Printf("connecting to %s", u.String())

	// establish connection (note: defer close is within goroutines)
	c, resp, err := websocket.DefaultDialer.Dial(u.String(), nil)
	if err != nil {
		log.Printf("handshake failed with status %d", resp.StatusCode)
		log.Fatal("dial: ", err)
	}

	defer resp.Body.Close()

	jsonString := &Subscribe{
		Subscribe: "new_block",
		ChainID:   "columbus-5",
	}

	err = c.WriteJSON(jsonString)

	if err != nil {
		log.Println(err)
		return
	}

	// Define how often we want to refresh, and create the channels we want to operate on
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()
	done := make(chan bool)
	interruption := make(chan bool)

	// function to enable CLI stop
	// Use ^C. It will cause an interruption in the process and stop the program.
	go func() {
		sig = nil
		sig = <-interrupt
		if sig != nil {
			interruption <- true
			defer c.Close()
		}
	}()

	// function to stop after 10 minutes
	go func() {
		time.Sleep(10 * time.Minute)
		done <- true
		defer c.Close()
	}()

	// start the webpage concurrently
	go func() {
		http.HandleFunc("/", handler)
		if err := http.ListenAndServe(":8080", nil); err != nil {
			log.Fatal("ListenAndServe:", err)
		}
	}()

	// handle channel responses and updates
	for {
		select {
		case <-interruption:
			fmt.Println("\nInterruption occurred. Connection has been closed.")
			return
		case <-done:
			fmt.Println("\nConnection has been open for 10 minutes. Refresh to reconnect.")
			return
		case t := <-ticker.C:
			res, err := getResponseFromWs(c)
			if err != nil {
				break
			}
			fmt.Println("Past Actions: ", strings.Join(getPastActionsFromWs(res), "\n"))
			supplyJSONPrint := getSupplyPricesFromWs(res)
			fmt.Println("\nCurrent Supply: ", supplyJSONPrint)
			fmt.Println("Response refreshed at ", t)
		}
	}

}
