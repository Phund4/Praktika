package internal

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
)

func getVacancies(ctx context.Context) (*vacancies, error) {
	url := "https://api.hh.ru/vacancies"
	method := "GET"

	client := &http.Client{}
	req, err := http.NewRequestWithContext(ctx, method, url, nil)
	if err != nil {
		return nil, fmt.Errorf("error in creating request to hh api: %v", err)
	}

	req.Header.Add("User-Agent", "api-test-agent")
	req.Header.Add("Content-Type", "application/json")

	res, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error in doing request to hh api: %v", err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		return nil, fmt.Errorf("error in reading from response body: %v", err)
	}

	var vacanciesData vacancies
	err = json.Unmarshal(body, &vacanciesData)
	if err != nil {
		return nil, fmt.Errorf("error in unmarshalling json data: %v", err)
	}

	return &vacanciesData, nil
}