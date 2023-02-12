// Persistence and CLI

// print : $ go run *.go printchain
// add : $ go run *.go addblock -data "Send 1 BTC to Joy"
package main

func main() {
	bc := NewBlockchain()
	defer bc.db.Close()

	cli := CLI{bc}
	cli.Run()
}
