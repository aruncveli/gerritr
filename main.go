package main

import (
	"gerritr/cmd"
	"gerritr/pkg/review"
)

func main() {
	review.InitConfig()
	cmd.Execute()
}
