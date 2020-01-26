package main

import (
	"io/ioutil"
	"net/http"
	"net/url"
	"strings"
	"time"
)

func checkCodes(codes []string) {
	for i := 0; i < len(codes); i++ {
		time.Sleep(time.Second * 10)
		codes[i] = url.QueryEscape(codes[i])
		codes[i] = strings.Replace(codes[i], "%0D", "", -1)
		checkCode(codes[i])
	}
}

func checkCode(code string) {
	CkURL = "https://api.nitestats.com/v1/codes/checker?codes=" + string(code)
	resp, err := http.Get(CkURL)
	if err != nil {
		logger.Error("%v", err)
		return
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Error("%v", err)
		return
	}
	s := string(body)
	if s == "You are being rate limited!" {
		//Later
	} else {
		s = strings.Replace(s, "<br>", "", -1)
		logger.Info(s)
	}
}
