package main

import (
	"context"
	"io"
	"log"
	"net/http"

	"github.com/aliyun/fc-runtime-go-sdk/fc"
)

func main() {
	fc.StartHttp(HandleHttpRequest)
}

func HandleHttpRequest(ctx context.Context, w http.ResponseWriter, req *http.Request) error {
	w.WriteHeader(http.StatusOK)
	w.Header().Add("Content-Type", "text/plain")
	reqBody, err := io.ReadAll(req.Body)
	log.Printf("%s", reqBody)
	if err != nil {
		log.Fatal("Err: ", err)
	}
	_, _ = w.Write(reqBody)
	return nil
}
