package source_cbondshnxvn

import (
	"encoding/csv"
	"fmt"
	"os"

	"github.com/daominah/crawler_xsmb/pkg/core"
	"github.com/mywrap/textproc"
	"github.com/tealeg/xlsx/v3"
	"golang.org/x/net/html"
)

func (s DataSourceCbonds) ParseToResultStockNews(rawHTML string) (rows []core.ResultStockNews, err error) {
	rootNode := textproc.HTMLParseToNode(rawHTML)
	resultBoxs, err := textproc.HTMLXPath(rootNode, `//*[@id="tbExtraordinaryResult"]`)
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

func getHTMLAttr(node *html.Node, attrKey string) string {
	for _, attr := range node.Attr {
		if attr.Key != attrKey {
			continue
		}
		return attr.Val
	}
	return ""
}

func WriteToCsv(rows []core.ResultStockNews, outputPath string) error {
	outputFile, err := os.OpenFile(outputPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return fmt.Errorf("OpenFile %v", err)
	}
	w := csv.NewWriter(outputFile)

	firstLine := []string{"Ngay dang", "Ten cong ty", "Ma trai phieu", "Tieu de tin"}
	w.Write(firstLine)
	for _, r := range rows {
		w.Write([]string{r.DateISO, r.CompanyName, r.StockSymbol, r.NewsTitle})
	}
	w.Flush()
	if err := w.Error(); err != nil {
		return fmt.Errorf("csv.Writer.Error %v", err)
	}
	outputFile.Close()
	return nil
}

func WriteToExcel(rows []core.ResultStockNews, outputPath string) error {
	f := xlsx.NewFile()
	sheet, _ := f.AddSheet("Cbonds")

	row := sheet.AddRow()
	firstLine := []string{"Nguon", "Ngay dang", "Ten cong ty", "Ma trai phieu", "Tieu de tin"}
	for _, v := range firstLine {
		cell := row.AddCell()
		cell.Value = v
	}

	for _, r := range rows {
		row := sheet.AddRow()
		for _, v := range []string{"Cbonds", r.DateISO, r.CompanyName, r.StockSymbol, r.NewsTitle} {
			cell := row.AddCell()
			cell.Value = v
		}
	}
	return f.Save(outputPath)
}
