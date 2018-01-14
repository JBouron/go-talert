package talert

import "net/http"
import "net/url"

/* Build time variables. */
const API_KEY = "<Your Telegram API KEY>"
const ME = "<Your Telegram user ID>"

func makeUrl(t string) string {
	if API_KEY == "<Your Telegram API KEY>" {
		panic("API key not defined, define it during build.")
	} else {
		return "https://api.telegram.org/bot" + API_KEY + "/" + t
	}
}

/**
 * Send message.
 * @param dest : Destination chat/user ID.
 * @param message : The text message.
 */
func SendMessageTo(dest string, message string) {
	u := makeUrl("sendMessage")
	values := url.Values{}
	values.Set("chat_id", dest)
	values.Set("text", message)
	http.PostForm(u, values)
}

/**
 * Send message to yourself.
 * @param message : The text message.
 */
func SendMessageSelf(message string) {
	if ME == "<Your Telegram user ID>" {
		panic("ME id not defined, define it during build.")
	} else {
		SendMessageTo(ME, message)
	}
}
