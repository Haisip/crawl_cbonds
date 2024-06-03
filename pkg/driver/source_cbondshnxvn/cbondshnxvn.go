package source_cbondshnxvn

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
	"strings"
	"time"
)

// DataSourceCbonds
type DataSourceCbonds struct{}

func (s DataSourceCbonds) DownloadHTML() ([]byte, error) {
	r, err := http.NewRequest("GET", "https://cbonds.hnx.vn/", nil)
	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	httpClient := &http.Client{Timeout: 64 * time.Second}
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

// DataSourceCbonds
type DataSourceCbonds2 struct{}

func (s DataSourceCbonds2) DownloadHTML() ([]byte, error) {
	page := strconv.Itoa(1)
	pageSize := strconv.Itoa(100)
	body := strings.NewReader(`keysSearch%5B%5D=&keysSearch%5B%5D=&keysSearch%5B%5D=&keysSearch%5B%5D=&keysSearch%5B%5D=&keysSearch%5B%5D=&keysSearch%5B%5D=&currentPages%5B%5D=1&currentPages%5B%5D=1&currentPages%5B%5D=` + page + `&currentPages%5B%5D=1&numberRecord%5B%5D=10&numberRecord%5B%5D=10&numberRecord%5B%5D=` + pageSize + `&numberRecord%5B%5D=10`)
	r, err := http.NewRequest("POST", "https://cbonds.hnx.vn/to-chuc-phat-hanh/tin-cong-bo-x", body)
	r.Header.Set("CP-TOKEN", `CfDJ8PdG7etAbaxFiQmMzHBRJ_2FL4jRiYDX7DsmJXxowB9CbsbPkYbgxllYY32PNxHCYU3bwBt6rh8odSEfgHzqzj838_7TyTD3W7tTgQttraCAO1TWOjBiVfV5fbots0sV_oAn0lnfe1Zu93pzXYtQMTM`)
	r.Header.Set("Cookie", `CP-TOKEN-COOKIE=CfDJ8PdG7etAbaxFiQmMzHBRJ_2SVNIR7NXM0z3TFSpXwPPTX8mExIgzjIRkJBFAtFXNbPaYbc-VZkPfwuSiTZBn3eajqJ-aYqZv22NPaXaMz3FlNCP8DLU78zypJH3jbjIOUj20QaWnM4DDryaHH1ekq0Q`)
	r.Header.Set("Origin", "https://cbonds.hnx.vn")
	r.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	if err != nil {
		return nil, fmt.Errorf("http.NewRequest: %w", err)
	}
	httpClient := &http.Client{Timeout: 64 * time.Second}
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
