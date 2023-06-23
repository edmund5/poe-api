package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Query struct {
	Content string `json:"content"`
}

type Message struct {
	Query []Query `json:"query"`
}

type Reply struct {
	Text string `json:"text"`
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == http.MethodPost {
			body, _ := ioutil.ReadAll(r.Body)
			var message Message
			json.Unmarshal(body, &message)
			lastContent := message.Query[len(message.Query)-1].Content

			w.WriteHeader(http.StatusOK)
			w.Header().Set("Content-Type", "text/event-stream")

			fmt.Fprint(w, formatEvent("meta", `{"content_type": "text/markdown", "linkify": false, "suggested_replies": true}`))
			fmt.Fprint(w, formatEvent("text", fmt.Sprintf(`{"text": "%s"}`, lastContent)))

			replies := []string{"Hi", "Hello", "Hey"}
			for _, reply := range replies {
				replyData, _ := json.Marshal(Reply{Text: reply})
				fmt.Fprint(w, formatEvent("suggested_reply", string(replyData)))
			}

			fmt.Fprint(w, formatEvent("done", "{}"))
		}
	})

	http.ListenAndServe(":8080", nil)
}

func formatEvent(eventType string, data string) string {
	return fmt.Sprintf("event: %s\ndata: %s\n\n", eventType, data)
}