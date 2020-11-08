package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"

	"github.com/shelly-tools/coiot"
)

func main() {

	ip := flag.String("ip", "192.168.178.240", "the Shellys ip address")
	path := flag.String("path", "/cit/s", "the CoIoT path - /cit/d or /cit/s")

	flag.Parse()
	req := coiot.Message{
		Type:      coiot.Confirmable,
		Code:      coiot.GET,
		MessageID: 12345,
		Payload:   []byte(""),
	}
	req.SetPathString(*path)

	c, err := coiot.Dial("udp", *ip+":5683")
	if err != nil {
		fmt.Printf("Error dialing: %v", err)
		os.Exit(2)
	}

	rv, err := c.Send(req)
	if err != nil {
		fmt.Printf("Error sending request: %v", err)
		os.Exit(2)
	}

	if rv != nil {
		if isJSON(rv.Payload) == true {
			fmt.Printf("OK: Payload is valid JSON:\n %s", rv.Payload)
			os.Exit(0)
		} else {
			fmt.Printf("CRITICAL: Payload is invalid JSON:\n %s", rv.Payload)
			os.Exit(2)
		}
	}

}

func isJSON(s []byte) bool {
	var js map[string]interface{}
	return json.Unmarshal([]byte(s), &js) == nil
}
