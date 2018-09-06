package main

import (
	"medichain/deploy"
)

func main() {
	//deploy.DeployEasyCns(false)
	deploy.DeployAllByDefaultEasyCnsAddress(true)
}