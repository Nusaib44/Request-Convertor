package failed

// import (
// 	"net/http"
// 	"task/pkg/common"

// 	"github.com/gin-gonic/gin"
// )

// var reqMap map[string]string

// func Incoming(g *gin.Context) {

// 	go worker(g)

// 	if err := g.Bind(&reqMap); err != nil {
// 		common.ErrorMessage(g, "failed to bind", err.Error(), http.StatusBadRequest)
// 	}
// }

// // func sendRequest1(g *gin.Context, req map[string]interface{}, url string) {

// // 	jStr, err := json.Marshal(req)
// // 	if err != nil {
// // 		common.ErrorMessage(g, "Error marshaling converted request:", err, 500)
// // 		return
// // 	}

// // 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jStr))
// // 	if err != nil {
// // 		common.ErrorMessage(g, "Error sending request:", err, 500)
// // 		return
// // 	}

// // 	defer resp.Body.Close()

// // 	common.SurcessMessage(g, "Request sent successfully to the desired webhook URL", req)

// // }
