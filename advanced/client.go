package advanced

import (
	"fmt"
	"io"
	"net/http"
)

func HTTPClientFunc() {
	client := &http.Client{}

	resp, err := client.Get("https://swapi.dev/api/people/1")
	if err != nil {
		fmt.Println("Error during GET request", err)
		return
	}

	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println("Error during reading response", err)
		return
	}

	fmt.Println(string(body))
}
