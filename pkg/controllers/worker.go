package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"task/pkg/common"

	"github.com/gin-gonic/gin"
)

func worker(g *gin.Context, req <-chan Body) {

	for r := range req {
		go func(rs Body) {
			convertedReq := convertRequest(rs)
			sendRequest(g, convertedReq, "https://webhook.site/0c8958b4-c5a4-4639-97d0-9c0dfc87fa5d")
		}(r)
	}
	fmt.Println("the end")
}

func sendRequest(g *gin.Context, req ConvertedRequest, url string) {

	jStr, err := json.Marshal(req)
	if err != nil {
		common.ErrorMessage(g, "Error marshaling converted request:", err, 500)
		return
	}

	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jStr))
	if err != nil {
		common.ErrorMessage(g, "Error sending request:", err, 500)
		return
	}

	defer resp.Body.Close()

	common.SurcessMessage(g, "Request sent successfully to the desired webhook URL", resp.Body)

}
