// A demo to make a screenshot for readme file.
package main

import (
	"errors"

	"github.com/mrumyantsev/logx/log"
)

func main() {
	log.Info("service started")
	log.Debug("using stage database")
	log.Info("134.242.44.77\tPOST /api/user/36452\tuser login")
	log.Info("167.32.121.2\tPOST /api/product/119879283\tadd product to cart")
	log.Warn("167.32.121.2\tGET /api/product/119879283\tslow request")
	log.Info("201.87.189.21\tGET /api/product/119879283\tget product info")
	log.Error("201.87.189.21\tGET /api/product/119879283\tdatabase connection lost", errors.New("connection reset by peer"))
}
