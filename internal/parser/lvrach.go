package parser

import (
	"errors"
	"fmt"
	"net/http"
	"regexp"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"

	"github.com/t67y110v/parser-service/internal/client"
	"github.com/t67y110v/parser-service/internal/models"

	"github.com/sirupsen/logrus"
)

const basePathLvrach = "https://www.lvrach.ru"

type Lvrach struct {
	logger *logrus.Logger
	client client.Client
}

func NewLvrachParser(l *logrus.Logger) *Lvrach {
	c := client.NewHTTPClient(basePathLvrach)
	return &Lvrach{
		logger: l,
		client: c,
	}
}

func (l Lvrach) Page(query string, page int) (*models.Response, error) {
	wg := sync.WaitGroup{}

	res := models.Response{}
	res.Articles = make([]models.Article, 0)

	wg.Add(1)
	start := time.Now()
	l.makeGetRequest(page, query, &res, &wg)

	wg.Wait()
	fmt.Println()
	sort.Slice(res.Articles, func(i, j int) bool {
		return res.Articles[i].ID < res.Articles[j].ID
	})

	amount, err := l.getAmountArticles(query)
	if err != nil {
		l.logger.Error("getting amount articles ", err)
		return nil, err
	}

	res.Found = amount
	l.logger.Info("parsing done for ", query, " by ", time.Until(start).String())

	for i := range res.Articles {
		res.Articles[i].ID = 0
	}
	return &res, nil
}

func (l Lvrach) All(query string) (*models.Response, error) {

	amount, err := l.getAmountArticles(query)
	if err != nil {
		l.logger.Error("getting amount articles ", err)
		return nil, err
	}

	count := (amount / 15) - 1

	l.logger.Info("pages to parse - ", count, " by query ", query)
	wg := sync.WaitGroup{}

	res := models.Response{}
	res.Articles = make([]models.Article, 0)

	wg.Add(count)
	start := time.Now()
	for i := 1; i <= count; i++ {
		go l.makeGetRequest(i, query, &res, &wg)
	}

	wg.Wait()
	fmt.Println()
	sort.Slice(res.Articles, func(i, j int) bool {
		return res.Articles[i].ID < res.Articles[j].ID
	})
	res.Found = amount
	l.logger.Info("parsing done ", count, " pages by ", time.Until(start).String())
	for i := range res.Articles {
		res.Articles[i].ID = 0
	}

	return &res, nil
}

func (l Lvrach) getAmountArticles(query string) (int, error) {
	req, err := http.NewRequest("GET", l.client.BasePath+"/search", nil)
	if err != nil {
		l.logger.Error("dial creating get request", err)
		return 1, nil
	}

	queryUrl := req.URL.Query()
	queryUrl.Add("search_text", query)
	queryUrl.Add("p", strconv.Itoa(1))
	req.URL.RawQuery = queryUrl.Encode()

	l.setHeaders(req)
	bodyText, err := l.client.DoRequest(req)
	if err != nil {
		l.logger.Error("dial doing request", err)
	}

	re := regexp.MustCompile(`Найдено документов: (\d+)`)

	count, ok := strings.CutPrefix(string(re.Find((bodyText))), "Найдено документов: ")
	if !ok {
		return 0, errors.New("can't parse amount of query")
	}

	return strconv.Atoi(count)
}

func (l *Lvrach) makeGetRequest(page int, find string, res *models.Response, wg *sync.WaitGroup) {
	defer func() {
		if r := recover(); r != nil {
			l.logger.Panic(" Panic Recovered. Error:\n", r)
			return
		}
	}()

	defer wg.Done()

	req, err := http.NewRequest("GET", l.client.BasePath+"/search", nil)
	if err != nil {
		l.logger.Error("dial creating get  request ", err)
		return
	}

	query := req.URL.Query()

	query.Add("search_text", find)
	query.Add("p", strconv.Itoa(page))
	req.URL.RawQuery = query.Encode()

	l.setHeaders(req)

	bodyText, err := l.client.DoRequest(req)
	if err != nil {
		l.logger.Error("dial doing request", err)
	}

	var ul string

	re := regexp.MustCompile(`<ul class="news_list search_results">(?s)(.*?)</ul>`)

	matches := re.FindAllStringSubmatch(string(bodyText), -1)

	if len(matches) == 0 || len(matches[0]) == 0 {
		l.logger.Debug("can't find ul tag on page")
		return
	}

	ul = matches[0][0]

	var li []string

	{
		re := regexp.MustCompile(`<li>(?s)(.*?)</li>`)

		matches := re.FindAllStringSubmatch(string(ul), -1)

		for _, m := range matches {
			li = append(li, m[0])
		}

		if len(li) < 15 {
			l.logger.Debug("len of li less than 15")

			return
		}
		fmt.Print(page, "-")
		{
			for j, list := range li {
				re := regexp.MustCompile(`<a href="(.*?)>|<dt>(?s)(.*?)</dt>|</a>(?s)(.*?)</li>`)

				matches := re.FindAllStringSubmatch(string(list), -1)

				if len(matches) == 0 || len(matches[0]) == 0 {
					l.logger.Debug("cant find values of href name and annotaion in li tag")

					return
				}
				for i := 0; i < len(matches); i += 3 {
					href := matches[i][0]
					href, _ = strings.CutPrefix(href, `<a href="`)
					href, _ = strings.CutSuffix(href, `" target="_blank">`)
					href = l.client.BasePath + href

					mu := sync.Mutex{}
					mu.Lock()

					index, _ := strconv.Atoi(fmt.Sprintf("%d090%d", page, j))

					res.Articles = append(res.Articles, models.Article{
						Name:       matches[i+1][0][4 : len(matches[i+1][0])-5],
						Annotation: strings.TrimSpace(matches[i+2][0][4 : len(matches[i+2][0])-5]),
						Link:       href,
						ID:         index,
					})

					mu.Unlock()
				}
			}

		}
	}

}

func (l Lvrach) setHeaders(req *http.Request) {
	req.Header.Set("authority", "www.lvrach.ru")
	req.Header.Set("accept", "text/html,application/xhtml+xml,application/xml;q=0.9,image/avif,image/webp,image/apng,*/*;q=0.8,application/signed-exchange;v=b3;q=0.7")
	req.Header.Set("accept-language", "ru-RU,ru;q=0.9,en-US;q=0.8,en;q=0.7")
	req.Header.Set("cache-control", "no-cache")

	req.Header.Set("cookie", "osp_user_id=95507938; lvrachru=dng811qehbkep6bliibjepskqt;")
	req.Header.Set("pragma", "no-cache")
	req.Header.Set("referer", "https://www.lvrach.ru/search?search_text=ntvgthfnehf&drop.x=0&drop.y=0")
	req.Header.Set("sec-ch-ua", `"Not.A/Brand";v="8", "Chromium";v="114", "Google Chrome";v="114"`)
	req.Header.Set("sec-ch-ua-mobile", "?0")
	req.Header.Set("sec-ch-ua-platform", `"Linux"`)
	req.Header.Set("sec-fetch-dest", "document")
	req.Header.Set("sec-fetch-mode", "navigate")
	req.Header.Set("sec-fetch-site", "same-origin")
	req.Header.Set("sec-fetch-user", "?1")
	req.Header.Set("upgrade-insecure-requests", "1")
	req.Header.Set("user-agent", "Mozilla/5.0 (X11; Linux x86_64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/114.0.0.0 Safari/537.36")

}

// 95507962
// 95507961
// 95507938
