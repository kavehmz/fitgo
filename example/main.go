package main

import . "github.com/kavehmz/fitgo"

func main() {
	Lines("https://www.cloudflare.com/ips-v4").Grep(`131\.`).Echo().Count()
	Lines("/tmp/multi_ret").Grep(`2`).Count().Echo()
}
