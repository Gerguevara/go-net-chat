package chat

import (
	"flag"
	"io"
	"log"
)

var (
	port = flag.Int("p", 3090, "port")
	host = flag.String("h", "localhost", "host")
)

func CopyContent(dst io.Writer, src io.Reader) {
	_, err := io.Copy(dst, src)
	if err != nil {
		log.Fatal(err)
	}
}
