package parser

import (
	"encoding/json"
	"fmt"
	"io"
	"log"

	//"log"
	"net/http"
	"net/url"

	"strings"
)

type Article struct {
	Name       string ` json:"name"`
	Annotation string `json:"annotation"`
	Link       string `json:"link"`
}

type Response struct {
	Found    int       `json:"found,omitempty"`
	Articles []Article `json:"articles"`
}

func parseJSON(jsonData []byte) (Response, error) {
	var response Response
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return Response{}, err
	}
	return response, nil
}

func Parse(querry string) []Article {

	client := &http.Client{}
	decodedQuerry, _ := url.QueryUnescape(querry)
	var data = strings.NewReader(fmt.Sprintf(`{"year_to":2023,"year_from":2023,"mode":"articles","q":"%s","size":10,"from":0}`, decodedQuerry))
	req, err := http.NewRequest("POST", "https://cyberleninka.ru/api/search", data)
	if err != nil {
		log.Fatal(err)
	}
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	//req.Header.Set("Cookie", "_ym_uid=1686512360458954049; _ym_d=1690796675; _ym_isad=1; _ga=GA1.2.517948244.1690796675; _gid=GA1.2.1900067405.1690796675; _gat=1")
	req.Header.Set("Origin", "https://cyberleninka.ru")
	req.Header.Set("Referer", "https://cyberleninka.ru/search?q=%D0%90%D0%B1%D0%B0%D0%BA%D0%B0%D0%B2%D0%B8%D1%80&page=1")
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	resp, err := client.Do(req)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()
	bodyText, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Fatal(err)
	}
	response, err := parseJSON(bodyText)
	if err != nil {
		fmt.Println("Error parsing JSON:", err)
		return nil
	}

	return response.Articles

}
