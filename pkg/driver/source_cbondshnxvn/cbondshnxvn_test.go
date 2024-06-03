package source_cbondshnxvn

import (
	"testing"
)

func _TestDataSourceXsktcomvn_DownloadHTML(t *testing.T) {
	var s DataSourceCbonds
	data, err := s.DownloadHTML()
	if err != nil {
		t.Fatalf("error DataSourceCbonds.DownloadHTML: %v", err)
	}
	_ = data
	println(string(data))
}

func TestDataSourceCbonds2_DownloadHTML(t *testing.T) {
	var s DataSourceCbonds2
	data, err := s.DownloadHTML()
	if err != nil {
		t.Fatalf("error DataSourceCbonds.DownloadHTML: %v", err)
	}
	_ = data
	println(string(data))
}
