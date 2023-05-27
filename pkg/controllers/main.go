package controllers

import (
	"net/http"
	"task/pkg/common"

	"github.com/gin-gonic/gin"
)

type Body struct {
	Ev    string `json:"ev"`
	Et    string `json:"et"`
	ID    string `json:"id"`
	UID   string `json:"uid"`
	MID   string `json:"mid"`
	T     string `json:"t"`
	P     string `json:"p"`
	L     string `json:"l"`
	SC    string `json:"sc"`
	AtRk1 string `json:"atrk1"`
	AtRv1 string `json:"atrv1"`
	AtRt1 string `json:"atrt1"`
	AtRk2 string `json:"atrk2"`
	AtRv2 string `json:"atrv2"`
	AtRt2 string `json:"atrt2"`
	AtRk3 string `json:"atrk3"`
	AtRv3 string `json:"atrv3"`
	AtRt3 string `json:"atrt3"`
	AtRk4 string `json:"atrk4"`
	AtRv4 string `json:"atrv4"`
	AtRt4 string `json:"atrt4"`
	UaRk1 string `json:"uatrk1"`
	UaRv1 string `json:"uatrv1"`
	UaRt1 string `json:"uatrt1"`
	UaRk2 string `json:"uatrk2"`
	UaRv2 string `json:"uatrv2"`
	UaRt2 string `json:"uatrt2"`
	UaRk3 string `json:"uatrk3"`
	UaRv3 string `json:"uatrv3"`
	UaRt3 string `json:"uatrt3"`
	UaRk4 string `json:"uatrk4"`
	UaRv4 string `json:"uatrv4"`
	UaRt4 string `json:"uatrt4"`
	UaRk5 string `json:"uatrk5"`
	UaRv5 string `json:"uatrv5"`
	UaRt5 string `json:"uatrt5"`
	UaRk6 string `json:"uatrk6"`
	UaRv6 string `json:"uatrv6"`
	UaRt6 string `json:"uatrt6"`
}

var Req Body

func Incoming(g *gin.Context) {

	requests := make(chan Body)

	go worker(g, requests)

	if err := g.Bind(&Req); err != nil {
		common.ErrorMessage(g, "failed to bind", err.Error(), http.StatusBadRequest)
	}
	requests <- Req
}
