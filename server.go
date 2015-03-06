package main
import (
	"github.com/zenazn/goji"
	"github.com/shiwork/ppserver/server"
)


func main() {
	goji.Get("/apps/ppserver/demo", server.SampleApi)
	goji.Serve()
}
