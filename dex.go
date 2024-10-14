package bot

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/url"
)

type TokenProfile struct {
	URL          string `json:"url"`
	ChainID      string `json:"chainID"`
	TokenAddress string `json:"tokenAddress"`
	Icon         string `json:"icon"`
	Header       string `json:"header"`
	Description  string `json:"description"`
}

func sendMessage(text string) {
	params := url.Values{}
	params.Add("chat_id", "7158688111")
	params.Add("text", text)
	url := "https://api.telegram.org/bot7577355680:AAGnOdCr0LFCyqdY8VWJrRT1yif8BLfM0Lw/sendMessage?" + params.Encode()
	_, err := http.Get(url)
	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}
}

func Run() {
	resp, err := http.Get("https://api.dexscreener.com/token-profiles/latest/v1")
	if err != nil {
		fmt.Println("Error fetching data:", err)
		return
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading response body:", err)
		return
	}

	var tokendProfiles []TokenProfile
	if err := json.Unmarshal(body, &tokendProfiles); err != nil {
		fmt.Println("Error parsing JSON:", err)
		return
	}

	for _, token := range tokendProfiles {
		if token.ChainID == "solana" {
			fmt.Println(token.TokenAddress)
		}
	}
	
	sendMessage("asdfgasdf")
}