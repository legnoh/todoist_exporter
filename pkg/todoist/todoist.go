package todoist

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

func search_query_by_filter(token string) *http.Response {
	client := &http.Client{}

	req, err := http.NewRequest("GET", "https://api.todoist.com/rest/v1/tasks", nil)
	if err != nil {
		log.Print(err)
	}
	q := req.URL.Query()
	q.Add("filter", "(due before: +0hours|overdue|no due date) & !#売却リスト & !#作りたいものリスト & subtask")
	req.URL.RawQuery = q.Encode()

	req.Header.Add("Authorization", "Bearer "+token)

	parseFormErr := req.ParseForm()
	if parseFormErr != nil {
		fmt.Println(parseFormErr)
	}

	resp, err := client.Do(req)

	if err != nil {
		fmt.Println("Failure : ", err)
	}

	respBody, _ := ioutil.ReadAll(resp.Body)
	fmt.Println(respBody)
	return resp
}
