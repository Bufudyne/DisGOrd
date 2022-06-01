package main

import (
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

//WsJsonResponse define the response sent back from websocket
type WsJsonResponse struct {
	Action      string `json:"action"`
	Message     string `json:"message"`
	MessageType string `json:"message_type"`
}

//WsEndpoint upgrades connection to websocket
func (app *application) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		app.errorLog.Println(err)
	}
	log.Println("Client connected to endpoint")

	var response WsJsonResponse
	response.Message = `<em><mall>Connected to server</small></em>`

	err = ws.WriteJSON(response)
	if err != nil {
		app.errorLog.Println(err)
	}
}
