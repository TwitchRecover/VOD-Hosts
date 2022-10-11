package main

import (
	"fmt"
	"github.com/TwitchRecover/TwitchVHR"
)

var (
	options = twitchvhr.Options{
		Timeout:          10,
		AllowSource:      true,
		AllowSpectre:     true,
		AllowAudioOnly:   true,
		IncludeFramerate: true,
	}
)

func main() {
	vodIDs := []int{1606624726}
	hosts := twitchvhr.Retrieve(vodIDs, options)
	fmt.Println(hosts)
}
