package source_cbondshnxvn

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"time"
)

type DataSourceDangkykinhdoanhgovvn struct{}

func (s DataSourceDangkykinhdoanhgovvn) DownloadMasothue(companyName string) ([]byte, error) {
	httpClient := &http.Client{Timeout: 64 * time.Second}
	u := "https://masocongty.vn/search?by=all&pro=all&name=" + url.QueryEscape(companyName)
	r, err := http.NewRequest("POST", u, nil)
	log.Printf("r: %v", r.URL)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	w, err := httpClient.Do(r)
	if err != nil {
		return nil, fmt.Errorf("httpClient.Do: %w", err)
	}
	defer w.Body.Close()
	b, err := io.ReadAll(w.Body)
	if err != nil {
		return nil, fmt.Errorf("io.ReadAll: %w", err)
	}
	return b, err
}

func (s DataSourceDangkykinhdoanhgovvn) SearchMasothue(companyName string) (string, error) {
	body, err := s.DownloadMasothue(companyName)
	if err != nil {
		return "", err
	}
	_ = body
	return "hohohaha", nil
}
