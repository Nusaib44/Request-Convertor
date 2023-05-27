package failed

// import "fmt"

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

// func convertRequest() ConvertedRequest {
// 	convertedReq := ConvertedRequest{
// 		Event:           reqMap["ev"],
// 		EventType:       getStringValue(reqMap, "et"),
// 		AppID:           getStringValue(reqMap, "id"),
// 		UserID:          getStringValue(reqMap, "id"),
// 		MessageID:       getStringValue(reqMap, "mid"),
// 		PageTitle:       getStringValue(reqMap, "t"),
// 		PageURL:         getStringValue(reqMap, "p"),
// 		BrowserLanguage: getStringValue(reqMap, "l"),
// 		ScreenSize:      getStringValue(reqMap, "sc"),
// 		Attributes:      make(map[string]Attribute),
// 		Traits:          make(map[string]Trait),
// 	}

// 	println("evvv..vvv.flvv", reqMap["ev"])

// 	convertAttributes(convertedReq.Attributes, reqMap, "atrk", "atrv", "atrt")
// 	convertTraits(convertedReq.Traits, reqMap, "uatrk", "uatrv", "uatrt")
// 	return convertedReq
// }

// func getStringValue(data map[string]string, key string) string {
// 	value := data[key]
// 	println(value, "strvvvv")
// 	return value
// }

// func convertAttributes(attrMap map[string]Attribute, reqMap map[string]string, keyPrefix string, valuePrefix string, typePrefix string) {
// 	for i := 1; ; i++ {
// 		attrKey := getStringValue(reqMap, fmt.Sprintf("%s%d", keyPrefix, i))
// 		attrValue := getStringValue(reqMap, fmt.Sprintf("%s%d", valuePrefix, i))
// 		attrType := getStringValue(reqMap, fmt.Sprintf("%s%d", typePrefix, i))

// 		if attrKey == "" || attrValue == "" || attrType == "" {
// 			break
// 		}

// 		attr := Attribute{
// 			Value: attrValue,
// 			Type:  attrType,
// 		}

// 		attrMap[attrKey] = attr
// 	}
// }
// func convertTraits(ttrMap map[string]Trait, reqMap map[string]string, keyPrefix string, valuePrefix string, typePrefix string) {
// 	for i := 1; ; i++ {
// 		ttrKey := getStringValue(reqMap, fmt.Sprintf("%s%d", keyPrefix, i))
// 		ttrValue := getStringValue(reqMap, fmt.Sprintf("%s%d", valuePrefix, i))
// 		ttrType := getStringValue(reqMap, fmt.Sprintf("%s%d", typePrefix, i))

// 		if ttrKey == "" || ttrValue == "" || ttrType == "" {
// 			break
// 		}

// 		ttr := Trait{
// 			Value: ttrValue,
// 			Type:  ttrType,
// 		}

// 		ttrMap[ttrKey] = ttr
// 	}
// }
