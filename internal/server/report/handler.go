package report

import (
	"fmt"
	excel "github.com/360EntSecGroup-Skylar/excelize"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func Download(ctx *gin.Context) {
	f := excel.NewFile()

	// 创建一个样式，并设置背景色为黄色
	styleID, err := f.NewStyle(`{"fill":{"type":"pattern","color":["#CCCCCC"],"pattern":1}}`)
	if err != nil {
		fmt.Println("无法创建样式:", err)
		return
	}

	// 在默认的Sheet中设置单元格的值
	f.SetCellValue("Sheet1", "A1", "Hello")
	f.SetCellValue("Sheet1", "B1", "World!")
	alphaString := excel.ToAlphaString(1223532453)

	fmt.Println(alphaString)
	// 在 A1 单元格应用样式
	f.SetCellStyle("Sheet1", "A1", "A1", styleID)

	buf, err := f.WriteToBuffer()

	name := "报表管理"
	fileName := fmt.Sprintf("%s%d.xlsx", name, time.Now().UnixMilli())
	headers := map[string]string{
		"Content-Type":              "application/octet-stream",
		"Content-Disposition":       "attachment; filename=" + fileName,
		"Content-Transfer-Encoding": "binary",
	}
	ctx.DataFromReader(http.StatusOK, int64(buf.Len()), "application/octet-stream", buf, headers)
}
