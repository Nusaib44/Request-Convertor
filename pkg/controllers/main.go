package controllers

import (
	"net/http"
	"task/pkg/common"

	"github.com/gin-gonic/gin"
)

type Body struct {
	ev   string 
	et    string 
	id   string 
	uid  string 
	mid   string 
	t     string 
	p     string 
	l     string 
	sc   string 
	atrk1 string 
	atrv1 string 
	atrt1 string 
	atrk2 string 
	atrv2 string 
	atrt2 string 
	uatrk1 string 
	uatrv1 string 
	uatrt1 string 
	uatrk2 string 
	uatrv2 string 
	uatrt2 string 
	uatrk3 string 
	uatrv3 string 
	uatrt3 string 
}

func Incoming(g *gin.Context) {

	requests := make(chan Body)
	var req Body

	go worker(g,requests)

	if err := g.Bind(&req); err != nil {
		common.ErrorMessage(g, "failed to bind", err.Error(), http.StatusBadRequest)
	}
	requests <- req
}
