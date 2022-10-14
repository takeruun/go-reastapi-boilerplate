package app

import "time"

func main() {
	time.Local = time.FixedZone("JST", 9*60*60)
}
