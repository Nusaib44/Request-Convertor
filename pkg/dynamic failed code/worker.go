package failed

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"sync"
// 	"task/pkg/common"

// 	"github.com/gin-gonic/gin"
// )

// var mu sync.Mutex

// func worker(g *gin.Context) {
// 	var convertedReq ConvertedRequest

// 	func() {
// 		convertedReq = convertRequest()
// 	}()

// 	sendRequest(g, convertedReq, "https://webhook.site/0c8958b4-c5a4-4639-97d0-9c0dfc87fa5d")
// 	fmt.Println("the end")

// }

// func sendRequest(g *gin.Context, req ConvertedRequest, url string) {

// 	jStr, err := json.Marshal(req)
// 	if err != nil {
// 		common.ErrorMessage(g, "Error marshaling converted request:", err, 500)
// 		return
// 	}

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jStr))
// 	if err != nil {
// 		common.ErrorMessage(g, "Error sending request:", err, 500)
// 		return
// 	}

// 	defer resp.Body.Close()

// 	common.SurcessMessage(g, "Request sent successfully to the desired webhook URL", resp.Body)

// }
