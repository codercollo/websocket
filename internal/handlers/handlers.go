package handlers

import (
	"log"
	"net/http"

	"github.com/CloudyKit/jet/v6"
	"github.com/gorilla/websocket"
)

// Initialize Jet template set
var views = jet.NewSet(
	jet.NewOSFileSystemLoader("./html"),
	jet.InDevelopmentMode(),
)

// upgradeConnection configures WebSocket upgrade settings
var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,                                       //size of read buffer in bytes
	WriteBufferSize: 1024,                                       //size of write buffer in bytes
	CheckOrigin:     func(r *http.Request) bool { return true }, //allow all origins
}

// Home handler renders the home page
func Home(w http.ResponseWriter, r *http.Request) {
	//Dender home.jet template
	err := renderPage(w, "home.jet", nil)
	if err != nil {
		log.Println(err)

	}

}

// WsJsonResponse defines the structure of JSON messages sent over WebSocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

// WsEndPoint handles WebSocket connections
func WsEndPoint(w http.ResponseWriter, r *http.Request) {
	//Upgrade the HTTP connection to a WebSocket
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
	}

	//log successful connection
	log.Println("Client connected to endpoint")

	//Prepare initial JSON response
	var response WsJsonResponse
	response.Message = `<em><small>Connected to server </small><em>`

	//Send initial message to client
	err = ws.WriteJSON(response)
	if err != nil {
		log.Println(err)
	}

}

// renderPage loads and executes a Jet template
func renderPage(w http.ResponseWriter, tmpl string, data jet.VarMap) error {
	//Get the template by name
	view, err := views.GetTemplate(tmpl)
	if err != nil {
		log.Println(err)
		return err
	}

	//Execute template and write to response
	err = view.Execute(w, data, nil)
	if err != nil {
		log.Println(err)
		return err
	}

	return nil
}
