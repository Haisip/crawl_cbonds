package source_cbondshnxvn

import (
	_ "embed"
	"encoding/json"
	"strings"
	"testing"
)

//go:embed parse_cbonds2.html
var testDataCbonds2 string // result of TestDataSourceXsktcomvn_DownloadHTML

func TestParseToResultStockNews2(t *testing.T) {
	if len(strings.TrimSpace(testDataCbonds2)) == 0 {
		t.Fatalf("empty testDataCbonds")
	}
	rows, err := (DataSourceCbonds2{}).ParseToResultStockNews(testDataCbonds2)
	if err != nil {
		t.Fatalf("error ParseToResultStockNews: %v", err)
	}
	for i, row := range rows {
		beauty, _ := json.MarshalIndent(row, "", "\t")
		t.Logf("i %v: %+s\n", i, beauty)
	}
}
