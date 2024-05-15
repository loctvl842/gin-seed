// package main
//
// import (
// 	"gin-seed/routers"
// )
//
// func main() {
// 	router := routers.SetupRoute()
// 	router.Run(":8081")
// }

package main

import (
	"gin-seed/cmd"
)

func main() {
	cmd.Execute()
}
