package main

import . "github.com/kavehmz/fitgo"

func main() {
	Lines("https://www.cloudflare.com/ips-v4").Grep(`131\.`).Echo().Count()
	Lines("/tmp/my_text").Grep(`Word`).Count().Grep(`2`).Echo()
}
