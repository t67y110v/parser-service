package parser

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/t67y110v/parser-service/internal/client"
	"github.com/t67y110v/parser-service/internal/models"

	"net/http"
	"net/url"

	"strings"

	"github.com/sirupsen/logrus"
)

const basePathCyberleninka = "https://cyberleninka.ru"

type CyberLeninka struct {
	client client.Client
	logger *logrus.Logger
}

func NewCyberleninkaParser(l *logrus.Logger) *CyberLeninka {
	c := client.NewHTTPClient(basePathCyberleninka)
	return &CyberLeninka{
		logger: l,
		client: c,
	}
}

func (c CyberLeninka) parseJSON(jsonData []byte) (models.Response, error) {
	var response models.Response
	err := json.Unmarshal(jsonData, &response)
	if err != nil {
		return models.Response{}, err
	}

	for i := range response.Articles {
		response.Articles[i].Link = c.client.BasePath + response.Articles[i].Link
	}
	return response, nil
}

func (c CyberLeninka) Page(query string, page int) (*models.Response, error) {

	decodedQuerry, _ := url.QueryUnescape(query)
	var data = strings.NewReader(fmt.Sprintf(`{"year_to":2023,"year_from":2023,"mode":"articles","q":"%s","size":10,"from":0}`, decodedQuerry))
	req, err := http.NewRequest("POST", c.client.BasePath+"/api/search", data)
	if err != nil {
		c.logger.Error("dial creating request ", err)
		return nil, err
	}
	c.setHeaders(req)
	start := time.Now()

	bodyText, err := c.client.DoRequest(req)
	if err != nil {
		c.logger.Error("dial doing request", err)
	}

	response, err := c.parseJSON(bodyText)
	if err != nil {
		c.logger.Error("Error parsing JSON:", err)
		return nil, err
	}
	c.logger.Info("parsing done for ", query, " by ", time.Until(start).String())

	return &response, nil

}

func (c *CyberLeninka) All(query string) (*models.Response, error) {

	decodedQuerry, _ := url.QueryUnescape(query)
	var data = strings.NewReader(fmt.Sprintf(`{"year_to":2023,"year_from":2023,"mode":"articles","q":"%s","size":1000,"from":0}`, decodedQuerry))
	req, err := http.NewRequest("POST", c.client.BasePath+"/api/search", data)
	if err != nil {
		c.logger.Error("dial creating request ", err)
		return nil, err
	}
	c.setHeaders(req)
	start := time.Now()

	bodyText, err := c.client.DoRequest(req)
	if err != nil {
		c.logger.Error("dial doing request", err)
	}

	response, err := c.parseJSON(bodyText)
	if err != nil {
		c.logger.Error("Error parsing JSON:", err)
		return nil, err
	}
	c.logger.Info("parsing done for ", query, " collect ", len(response.Articles), " articles by ", time.Until(start).String())

	return &response, nil

}

func (c *CyberLeninka) setHeaders(req *http.Request) {
	req.Header.Set("Accept", "*/*")
	req.Header.Set("Accept-Language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("Connection", "keep-alive")
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Origin", "https://cyberleninka.ru")
	req.Header.Set("Referer", `https://cyberleninka.ru/search?q=%D0%90%D0%B1%D0%B0%D0%BA%D0%B0%D0%B2%D0%B8%D1%80&page=1`)
	req.Header.Set("Sec-Fetch-Dest", "empty")
	req.Header.Set("Sec-Fetch-Mode", "cors")
	req.Header.Set("Sec-Fetch-Site", "same-origin")
	req.Header.Set("User-Agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)

}
