package main

import (
	"flag"
	"log"
	"net/http"
	"fmt"
	"io/ioutil"
	"reflect"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
    ReadBufferSize:  1024,
    WriteBufferSize: 1024,
    CheckOrigin: func(r *http.Request) bool {
    	b, _ := ioutil.ReadAll(r.Body)
    	fmt.Printf(string(b))
        return true
    },
} 


func echo(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	log.Print("Connected with:", r.URL)
	defer conn.Close()
	for {
		mt, message, err := conn.ReadMessage()
		log.Printf("Message: ", reflect.TypeOf(message))
		log.Printf("Mt", reflect.TypeOf(mt))
		log.Printf("Err", reflect.TypeOf(err))
		// if err != nil {
		// 	log.Println("read:", err)
		// 	break
		// }
		// log.Printf("recv: %s", message)
		// err = c.WriteMessage(mt, message)
		// if err != nil {
		// 	log.Println("write:", err)
		// 	break
		// }
	}
}

func main() {
	flag.Parse()
	log.SetFlags(0)
	http.HandleFunc("/websocket", echo)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
