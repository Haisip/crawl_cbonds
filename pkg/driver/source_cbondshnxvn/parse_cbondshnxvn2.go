package source_cbondshnxvn

import (
	"fmt"
	"github.com/daominah/crawler_xsmb/pkg/core"
	"github.com/mywrap/textproc"
)

func (s DataSourceCbonds2) ParseToResultStockNews(rawHTML string) (rows []core.ResultStockNews, err error) {
	rootNode := textproc.HTMLParseToNode(rawHTML)
	resultBoxs, err := textproc.HTMLXPath(rootNode, `//*[@id="tbInconstant"]`)
	if err != nil || len(resultBoxs) == 0 {
		return nil, fmt.Errorf("HTMLXPath tbExtraordinaryResult, err %v", err)
	}
	resultBox := resultBoxs[0]
	trs, err := textproc.HTMLXPath(resultBox, `//tr`)
	if err != nil || len(trs) == 0 {
		return rows, fmt.Errorf("HTMLXPath tr, err %v", err)
	}
	for _, tr := range trs {
		tds, _ := textproc.HTMLXPath(tr, `//td`)
		if len(tds) < 5 {
			continue
		}
		row := core.ResultStockNews{}
		row.DateISO = textproc.HTMLGetText(tds[1])
		row.CompanyName = textproc.HTMLGetText(tds[2])
		row.StockSymbol = textproc.HTMLGetText(tds[3])
		row.NewsTitle = textproc.HTMLGetText(tds[4])
		rows = append(rows, row)
	}

	return rows, nil
}
