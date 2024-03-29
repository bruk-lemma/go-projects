// package main

// import (
// 	"fmt"
// 	"net/http"

// 	"github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
// 	ReadBufferSize:  1024,
// 	WriteBufferSize: 1024,
// }

// func main() {
// 	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
// 		conn, _ := upgrader.Upgrade(w, r, nil)

// 		for {
// 			msgType, msg, err := conn.ReadMessage()
// 			if err != nil {
// 				return
// 			}

// 			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

//				if err = conn.WriteMessage(msgType, msg); err != nil {
//					return
//				}
//			}
//		})
//		http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
//			http.ServeFile(w, r, "websockets.html")
//		})
//		http.ListenAndServe(":8080", nil)
//	}
package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/websocket"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return true // Allow connections from any origin
	},
}

func main() {
	http.HandleFunc("/echo", func(w http.ResponseWriter, r *http.Request) {
		conn, err := upgrader.Upgrade(w, r, nil)
		if err != nil {
			log.Printf("Error upgrading to WebSocket: %v", err)
			return
		}
		defer conn.Close()

		for {
			msgType, msg, err := conn.ReadMessage()
			if err != nil {
				if websocket.IsUnexpectedCloseError(err, websocket.CloseGoingAway, websocket.CloseAbnormalClosure) {
					log.Printf("Unexpected WebSocket closure: %v", err)
				}
				break
			}

			fmt.Printf("%s sent: %s\n", conn.RemoteAddr(), string(msg))

			if err = conn.WriteMessage(msgType, msg); err != nil {
				log.Printf("Error writing message: %v", err)
				break
			}
		}
	})

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.ServeFile(w, r, "websockets.html")
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
