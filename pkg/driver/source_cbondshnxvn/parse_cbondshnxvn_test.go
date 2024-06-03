package source_cbondshnxvn

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
)

//go:embed parse_cbondshnxvn_test.html
var testDataCbonds string // result of TestDataSourceXsktcomvn_DownloadHTML

func TestParseToResultStockNews(t *testing.T) {
	if len(strings.TrimSpace(testDataCbonds)) == 0 {
		t.Fatalf("empty testDataCbonds")
	}
	rows, err := (DataSourceCbonds{}).ParseToResultStockNews(testDataCbonds)
	if err != nil {
		t.Fatalf("error ParseToResultStockNews: %v", err)
	}
	for _, row := range rows {
		beauty, _ := json.MarshalIndent(row, "", "\t")
		t.Logf("%+s\n", beauty)
	}
}
