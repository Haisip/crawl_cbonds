package source_cbondshnxvn

import (
	"testing"
)

func TestDataSourceXsktcomvn_DownloadHTML(t *testing.T) {
	var s DataSourceDangkykinhdoanhgovvn
	data, err := s.DownloadMasothue("Nova Saigon Royal")
	if err != nil {
		t.Fatalf("error DataSourceCbonds.DownloadHTML: %v", err)
	}
	_ = data
	println(string(data))
}
