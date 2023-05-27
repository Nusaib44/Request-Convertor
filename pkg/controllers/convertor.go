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
		Event:           req.Ev,
		EventType:       req.Et,
		AppID:           req.ID,
		UserID:          req.UID,
		MessageID:       req.MID,
		PageTitle:       req.T,
		PageURL:         req.P,
		BrowserLanguage: req.L,
		ScreenSize:      req.SC,
		Attributes:      make(map[string]Attribute),
		Traits:          make(map[string]Trait),
	}

	if req.AtRk1 != "" {
		convertedReq.Attributes[req.AtRk1] = Attribute{
			Value: req.AtRv1,
			Type:  req.AtRt1,
		}
	}
	if req.AtRk2 != "" {
		convertedReq.Attributes[req.AtRk2] = Attribute{
			Value: req.AtRv2,
			Type:  req.AtRt2,
		}
	}
	if req.AtRk3 != "" {
		convertedReq.Attributes[req.AtRk3] = Attribute{
			Value: req.AtRv3,
			Type:  req.AtRt3,
		}
	}
	if req.AtRk4 != "" {
		convertedReq.Attributes[req.AtRk4] = Attribute{
			Value: req.AtRv4,
			Type:  req.AtRt4,
		}
	}
	if req.UaRk1 != "" {
		convertedReq.Traits[req.UaRk1] = Trait{
			Value: req.UaRv1,
			Type:  req.UaRt1,
		}
	}
	if req.UaRk2 != "" {
		convertedReq.Traits[req.UaRk2] = Trait{
			Value: req.UaRv2,
			Type:  req.UaRt2,
		}
	}
	if req.UaRk3 != "" {
		convertedReq.Traits[req.UaRk3] = Trait{
			Value: req.UaRv3,
			Type:  req.UaRt3,
		}
	}
	if req.UaRk4 != "" {
		convertedReq.Traits[req.UaRk4] = Trait{
			Value: req.UaRv4,
			Type:  req.UaRt4,
		}
	}
	if req.UaRk5 != "" {
		convertedReq.Traits[req.UaRk5] = Trait{
			Value: req.UaRv5,
			Type:  req.UaRt5,
		}
	}
	if req.UaRk6 != "" {
		convertedReq.Traits[req.UaRk6] = Trait{
			Value: req.UaRv6,
			Type:  req.UaRt6,
		}
	}

	// 	req.AtRk1: {
	// 		Value: req.AtRv1,
	// 		Type:  req.AtRt1,
	// 	},
	// 	req.AtRk2: {
	// 		Value: req.AtRv2,
	// 		Type:  req.AtRt2,
	// 	},
	// },
	// {
	// 	req.UaRk1: {
	// 		Value: req.UaRv1,
	// 		Type:  req.UaRt1,
	// 	},
	// 	req.UaRt2: {
	// 		Value: req.UaRv2,
	// 		Type:  req.UaRt2,
	// 	},
	// 	req.UaRt3: {
	// 		Value: req.UaRv3,
	// 		Type:  req.UaRt3,
	// 	},
	// },

	return convertedReq
}
