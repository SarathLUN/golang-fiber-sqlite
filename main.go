package main

import (
	"github.com/SarathLUN/golang-fiber-sqlite/initializers"
	"github.com/SarathLUN/golang-fiber-sqlite/routes"
)

func init() {
	// connect to database via initializes package
	initializers.ConnectDB()

}
func main() {
	// start frontend router
	routes.StartFrontEndRouter("localhost:3000")
}
