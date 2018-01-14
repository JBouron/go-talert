# go-talert
Send messages on Telegram from your go applications.

This is not a Telegram bot, just a simple interface that you can use to send your-self alerts from your scripts written in Go.


# Usage
Two simple functions are available :
* SendMessageTo : Sends a text message to a particular user.
* SendMessageSelf : Sends a text message to yourself.


~~~
package main

import "github.com/JBouron/talert"

func main() {
	/* Send a message to your-self. */
	talert.SendMessageSelf("Hello me")

	/* Send a message to someone else. */
	talert.SendMessageTo("<ID>", "Hello world")
}
~~~

# Installation
First insert your API key and your user ID in talert.go under `API_KEY` and `ME` respectively.
Then run `go install`.
