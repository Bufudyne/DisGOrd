package main

import (
	"fmt"
	"github.com/bwmarrin/discordgo"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

func (app *application) VirtualTerminal(w http.ResponseWriter, r *http.Request) {
	if err := app.renderTemplate(w, r, "terminal", nil); err != nil {
		app.errorLog.Println(err)
	}
}

var upgradeConnection = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin:     func(r *http.Request) bool { return true },
}

var wsChan = make(chan WsPayload)
var clients = make(map[WebSocketConnection]string)

type WebSocketConnection struct {
	*websocket.Conn
}

func (app *application) ListenForWS(conn *WebSocketConnection) {
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error", fmt.Sprintf("%v", r))
		}
	}()

	var payload WsPayload

	for {
		err := conn.ReadJSON(&payload)
		if err != nil {
			// do nothing
		} else {
			payload.Conn = *conn
			wsChan <- payload
		}
	}
}

func (app *application) ListenToWsChannel() {
	//var response WsJsonResponse
	for {
		e := <-wsChan
		switch e.Action {
		case "send_to_channel":
			app.errorLog.Println(e.Message)
			app.sendMessage(e.ChannelMessage)
			break
		}
	}
}

// broadcastToAll sends ws response to all connected clients
func (app *application) broadcastToAll(response WsJsonResponse) {
	for client := range clients {
		err := client.WriteJSON(response)
		if err != nil {
			// the user probably left the page, or their connection dropped
			log.Println("websocket err")
			_ = client.Close()
			delete(clients, client)
		}
	}
}

//WsJsonResponse define the response sent back from websocket
type WsJsonResponse struct {
	Action      string                  `json:"action"`
	Message     string                  `json:"message"`
	ChatMessage discordgo.MessageCreate `json:"chat_message"`
	GuildList   []ChannelData           `json:"guild_list"`
	//aGuildList  []*discordgo.Guild      `json:"guild_list"`
	MessageType string `json:"message_type"`
}

type ChannelData struct {
	GuildID     string               `json:"guild_id"`
	ChannelList []*discordgo.Channel `json:"channel_list"`
}

type WsPayload struct {
	Action         string              `json:"action"`
	Message        string              `json:"message"`
	ChannelMessage ChannelMessage      `json:"channel_message"`
	Conn           WebSocketConnection `json:"-"`
}
type ChannelMessage struct {
	Guild   string `json:"guild"`
	Channel string `json:"channel"`
	Message string `json:"message"`
}

//WsEndpoint upgrades connection to websocket
func (app *application) WsEndpoint(w http.ResponseWriter, r *http.Request) {
	ws, err := upgradeConnection.Upgrade(w, r, nil)
	if err != nil {
		app.errorLog.Println(err)
	}
	log.Println("Client connected to endpoint")

	var response WsJsonResponse
	response.Action = "debug"
	response.Message = "Connected to WS"
	response.MessageType = "log_message"
	conn := WebSocketConnection{Conn: ws}
	clients[conn] = ""
	err = ws.WriteJSON(response)
	if err != nil {
		app.errorLog.Println(err)
	}

	response.Action = "guild_list"
	response.GuildList = app.getServerList()
	err = ws.WriteJSON(response)
	if err != nil {
		app.errorLog.Println(err)
	}

	go app.ListenForWS(&conn)
}
