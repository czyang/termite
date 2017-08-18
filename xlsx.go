package main

import (
	"github.com/tealeg/xlsx"
	"log"
)

func createXLSX(items []MovieItem, userId string) {
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("看过的电影")
	if err != nil {
		log.Fatal(err)
	}
	row := sheet.AddRow()
	cell := row.AddCell()
	cell.Value = "电影片名"
	cell = row.AddCell()
	cell.Value = "我的评分"
	cell = row.AddCell()
	cell.Value = "我的评语"
	cell = row.AddCell()
	cell.Value = "评分日期"
	cell = row.AddCell()
	cell.Value = "影片描述"
	cell = row.AddCell()
	cell.Value = "豆瓣链接"

	sheet.SetColWidth(0, 0, 40)
	sheet.SetColWidth(2, 2, 40)
	sheet.SetColWidth(3, 3, 10)
	sheet.SetColWidth(4, 4, 20)

	for _, v := range items {
		row = sheet.AddRow()
		cell = row.AddCell()
		cell.Value = v.Title
		cell = row.AddCell()
		cell.Value = v.MyRate
		cell = row.AddCell()
		cell.Value = v.MyReview
		cell = row.AddCell()
		cell.Value = v.RateTime
		cell = row.AddCell()
		cell.Value = v.MovieDesc
		cell = row.AddCell()
		cell.Value = v.DoubanLink
	}

	err = file.Save("./" + userId + "_movies.xlsx")
	if err != nil {
		log.Fatal(err)
	}
}