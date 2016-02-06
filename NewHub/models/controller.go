package models

import (
  "net/http"
  "strconv"
  "log"
  "flag"
)

type controller struct {
  hubs []*hub
}

func (c *controller) AddHub() {
  //new hub, function in hub model
  h := newHub()
  c.hubs = append(c.hubs, &h)
}

//architecture controller -> hubs -> connections

//Runs all hubs in goroutines
//Serves content, in a goroutine in main
//Parameter is number of hubs to run
func (c *controller) Run() {
    //Embed the websocket in each hub page, and create the routing for each hub
    for i, h := range c.hubs{
      go h.run()
      num := strconv.Itoa(i)
      //serve the home
      http.HandleFunc("/"+num, h.serveHub)
      //serve the websocket on the page
      http.HandleFunc("/ws"+"/"+num, h.serveWs)
    }
    //TODO for loop here to add the number of hubs passed in as signature
    c.AddHub()
    c.AddHub()
    c.AddHub()
}

func NewController() controller{
  var contr = controller {
    hubs : make([]*hub, 0),
  }
  return contr
}
