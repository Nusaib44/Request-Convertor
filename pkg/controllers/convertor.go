package controllers

type ConvertedRequest struct {
	Event           string               `json:"event"`
	EventType       string               `json:"event_type"`
	AppID           string               `json:"app_id"`
	UserID          string               `json:"user_id"`
	MessageID       string               `json:"message_id"`
	PageTitle       string               `json:"page_title"`
	PageURL         string               `json:"page_url"`
	BrowserLanguage string               `json:"browser_language"`
	ScreenSize      string               `json:"screen_size"`
	Attributes      map[string]Attribute `json:"attributes"`
	Traits          map[string]Trait     `json:"traits"`
}

type Attribute struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

type Trait struct {
	Value string `json:"value"`
	Type  string `json:"type"`
}

func convertRequest(req Body) ConvertedRequest {
	convertedReq := ConvertedRequest{
		Event:           req.ev,
		EventType:       req.et,
		AppID:           req.id,
		UserID:          req.uid,
		MessageID:       req.mid,
		PageTitle:       req.t,
		PageURL:         req.p,
		BrowserLanguage: req.l,
		ScreenSize:      req.sc,
		Attributes: map[string]Attribute{
			req.atrk1: {
				Value: req.atrv1,
				Type:  req.atrt1,
			},
			req.atrk2: {
				Value: req.atrv2,
				Type:  req.atrt2,
			},
		},
		Traits: map[string]Trait{
			req.uatrk1: {
				Value: req.uatrv1,
				Type:  req.uatrt1,
			},
			req.uatrk2: {
				Value: req.uatrv2,
				Type:  req.uatrt2,
			},
			req.uatrk3: {
				Value: req.uatrv3,
				Type:  req.uatrt3,
			},
		},
	}

	return convertedReq
}
