package failed

// import (
// 	"bytes"
// 	"encoding/json"
// 	"fmt"
// 	"net/http"
// 	"sync"

// 	"github.com/gin-gonic/gin"
// )

// type Request struct {
// 	Ev        string      `json:"ev"`
// 	Et        string      `json:"et"`
// 	ID        string      `json:"id"`
// 	UID       string      `json:"uid"`
// 	MID       string      `json:"mid"`
// 	T         string      `json:"t"`
// 	P         string      `json:"p"`
// 	L         string      `json:"l"`
// 	SC        string      `json:"sc"`
// 	Attrs     []Attribute `json:"-"`
// 	UserAttrs []Trait     `json:"-"`
// }

// type ConvertedRequest struct {
// 	Event           string               `json:"event"`
// 	EventType       string               `json:"event_type"`
// 	AppID           string               `json:"app_id"`
// 	UserID          string               `json:"user_id"`
// 	MessageID       string               `json:"message_id"`
// 	PageTitle       string               `json:"page_title"`
// 	PageURL         string               `json:"page_url"`
// 	BrowserLanguage string               `json:"browser_language"`
// 	ScreenSize      string               `json:"screen_size"`
// 	Attributes      map[string]Attribute `json:"attributes"`
// 	Traits          map[string]Trait     `json:"traits"`
// }

// type Attribute struct {
// 	Value string `json:"value"`
// 	Type  string `json:"type"`
// }

// type Trait struct {
// 	Value string `json:"value"`
// 	Type  string `json:"type"`
// }

// var wg sync.WaitGroup
// var reqChan chan Request

// func Incoming(c *gin.Context) {
// 	var req Request

// 	if err := c.BindJSON(&req); err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}

// 	wg.Add(1)
// 	go func() {
// 		processRequests()
// 		defer wg.Done()
// 		reqChan <- req
// 	}()

// 	c.JSON(http.StatusOK, gin.H{"message": "Request received"})
// }

// func processRequests() {
// 	for req := range reqChan {
// 		wg.Add(1)
// 		go func(req Request) {
// 			defer wg.Done()
// 			convertedReq := ConvertRequest(req)
// 			SendRequest(convertedReq, "https://webhook.site/0c8958b4-c5a4-4639-97d0-9c0dfc87fa5d")
// 		}(req)
// 	}
// 	wg.Wait()
// 	close(reqChan)
// }

// func ConvertRequest(req Request) ConvertedRequest {
// 	convertedReq := ConvertedRequest{
// 		Event:           req.Ev,
// 		EventType:       req.Et,
// 		AppID:           req.ID,
// 		UserID:          req.UID,
// 		MessageID:       req.MID,
// 		PageTitle:       req.T,
// 		PageURL:         req.P,
// 		BrowserLanguage: req.L,
// 		ScreenSize:      req.SC,
// 		Attributes:      make(map[string]Attribute),
// 		Traits:          make(map[string]Trait),
// 	}

// 	for i, attr := range req.Attrs {
// 		// key := fmt.Sprintf("atrk%d", i+1)
// 		convertedReq.Attributes[attr.Value] = Attribute{
// 			Value: attr.Value,
// 			Type:  req.UserAttrs[i].Value,
// 		}
// 	}

// 	for i, attr := range req.UserAttrs {
// 		// key := fmt.Sprintf("uatrk%d", i+1)
// 		convertedReq.Traits[attr.Value] = Trait{
// 			Value: attr.Value,
// 			Type:  req.UserAttrs[i].Value,
// 		}
// 	}

// 	return convertedReq
// }

// func SendRequest(req ConvertedRequest, url string) {
// 	jsonData, err := json.Marshal(req)
// 	if err != nil {
// 		fmt.Println("Error marshaling converted request:", err)
// 		return
// 	}

// 	resp, err := http.Post(url, "application/json", bytes.NewBuffer(jsonData))
// 	if err != nil {
// 		fmt.Println("Error sending request:", err)
// 		return
// 	}
// 	defer resp.Body.Close()

// 	fmt.Println("Request sent successfully to the desired webhook URL")
// }
