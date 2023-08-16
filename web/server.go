package web

import (
    "net/http"
    "news_channel_ai_ml/web/handlers"
)

func StartServer() {
    http.HandleFunc("/news", handlers.GetNews)
    // Other route handlers
    http.ListenAndServe(":8080", nil)
}
