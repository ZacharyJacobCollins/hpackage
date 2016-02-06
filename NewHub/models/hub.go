package models

import(
	"net/http"
	"text/template"
)

//Uses templates to insert unique websocket on page
var homeTempl = template.Must(template.ParseFiles("./html/home.html"))

//hub is similar to a "channel" in slack.
type hub struct {
	// Registered broadcasts.
	connections map[*connection]bool

	// Inbound messages from the connections.
	broadcast chan []byte

	// Register requests from the connections.
	register chan *connection

	// Unregister requests from connections.
	unregister chan *connection

	//message channel for persistance of switching between channels
	messages []byte
}

func (h *hub)serveHub(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", 405)
		return
	}
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	homeTempl.Execute(w, r.Host)
}

func (h *hub) run() {
	for {
		select {
		case c := <-h.register:
			h.connections[c] = true
		case c := <-h.unregister:
			if _, ok := h.connections[c]; ok {
				delete(h.connections, c)
				close(c.send)
			}
		//if there is message in channel broadcast, send to c of all h.connections
		case m := <-h.broadcast:
			//DATABASE ADD HERE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE TODO DATABASE 
			m = append(m, []byte("\n")...)
			h.messages = append(h.messages, m...)
			for c := range h.connections {
				select {
				case c.send <- m:
				default:
					close(c.send)
					delete(h.connections, c)
				}
			}
		}
	}
}

func newHub() hub{
	h := hub{
  	broadcast:   make(chan []byte),
  	register:    make(chan *connection),
  	unregister:  make(chan *connection),
  	connections: make(map[*connection]bool),
		//slice of messages for testing
		messages:    make([]byte, 0),
  }
	return h
}
