package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
	"unicode"

	"github.com/tealeg/xlsx"
)

type Rec struct {
	Col1 int    `json:"arshinNum"`
	Col2 string `json:"verifDate"`
	Col3 string `json:"validDate"`
	Col4 string `json:"typeSi"`
	Col5 string `json:"conclusion"`
	Col6 string `json:"verifSurname"`
	Col7 string `json:"verifName"`
	Col8 string `json:"verifLastname"`
	Col9 string `json:"verifSNILS"`
}

func main() {
	layout := "02.01.2006"
	var fn string
	var load string
	if len(os.Args) > 2 {
		if len(os.Args[1]) == 2 {
			load = os.Args[1]
			fn = os.Args[2]
		} else if len(os.Args[2]) == 2 {
			fn = os.Args[2]
			load = os.Args[1]
		} else {
			fmt.Println("Введите корректные данные: \nсsv2xml -1 file.csv [Как черновики] или \nсsv2xml -2 file.csv [Как отправленные]")
			os.Exit(1)
		}

		if strings.Contains(load, "-1") {
			load = "1"
		} else if strings.Contains(load, "-2") {
			load = "2"
		} else {
			fmt.Println("Введите корректные данные: \nсsv2xml -1 export.xlsx [Как черновики] или \nсsv2xml -2 file.xlsx [Как отправленные]")
			os.Exit(1)
		}
	}
	if strings.Contains(fn, ".xls") {

		xlsxFile, err := xlsx.OpenFile(fn)
		if err != nil {
			fmt.Printf("Error opening xlsx file: %s\n", err)
			os.Exit(1)
		}
		cv, err := os.Create(strings.TrimSuffix(fn, filepath.Ext(fn)) + ".csv")
		if err != nil {
			fmt.Printf("Error creating csv file: %s\n", err)
			return
		}
		defer cv.Close()

		writer := csv.NewWriter(cv)
		writer.Comma = ';'
		for _, sheet := range xlsxFile.Sheets {
			for _, row := range sheet.Rows {
				csvRow := make([]string, len(row.Cells))
				for i, cell := range row.Cells {
					csvRow[i] = cell.String()
				}
				writer.Write(csvRow)
			}
		}
		writer.Flush()
		layout = "01-02-06"
		fn = strings.TrimSuffix(fn, filepath.Ext(fn)) + ".csv"
	}

	if !strings.Contains(fn, ".csv") {

		fmt.Println("Введите корректные данные: \nсsv2xml -1 file.csv [Как черновики] или \nсsv2xml -2 file.csv [Как отправленные]")
		os.Exit(1)
	}

	f, err := os.Open(fn)
	if err != nil {
		fmt.Printf("Error creating csv file: %s\n", err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'
	var recs []Rec
	for i := 1; ; i++ {

		row, err := r.Read()
		if err != nil {
			break
		}
		row[0] = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, row[0])

		num, _ := strconv.Atoi(row[0])
		verifDate, err := time.Parse(layout, row[1])
		if err != nil {
			layout = "01-02-06"
			verifDate, err = time.Parse(layout, row[1])
			if err != nil {
				panic(err)
			}
		}
		validDate, err := time.Parse(layout, row[2])
		if err != nil {
			layout = "01-02-06"
			validDate, err = time.Parse(layout, row[2])
			if err != nil {
				panic(err)
			}
		}
		r := Rec{
			Col1: num,
			Col2: verifDate.Format(time.DateOnly),
			Col3: validDate.Format(time.DateOnly),
			Col4: row[3],
			Col5: row[4],
			Col6: row[5],
			Col7: row[6],
			Col8: row[7],
			Col9: row[8],
		}
		recs = append(recs, r)
	}
	jf, err := os.Create("fsa_upload.xml")
	if err != nil {
		panic(err)
	}
	defer jf.Close()
	jStr := `<?xml version="1.0" encoding="utf-8"?>
	<Message xsi:noNamespaceSchemaLocation="schema.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	  <VerificationMeasuringInstrumentData>
`
	flog, err := os.OpenFile(strings.TrimSuffix(os.Args[0], filepath.Ext(os.Args[0]))+".log", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
	currentTime := time.Now()
	tstamp := currentTime.Format("2006-01-02 15:04:05")
	if err != nil {
		fmt.Printf("Error creating log file: %s\n", err)
		return
	}
	defer flog.Close()
	for i, row := range recs {
		flog.WriteString(tstamp + " Обработано записей: " + fmt.Sprint(i, " ") + fmt.Sprint(row, "\n"))
		fmt.Println("Обработано записей: ", i, row)
		if len([]rune(row.Col5)) == 8 {

			jRow := fmt.Sprintf(`	    <VerificationMeasuringInstrument>
			<NumberVerification>%d</NumberVerification>
			<DateVerification>%s</DateVerification>
			<DateEndVerification>%s</DateEndVerification>
			<TypeMeasuringInstrument>%s</TypeMeasuringInstrument>
			<ApprovedEmployee>
			  <Name>
				<Last>%s</Last>
				<First>%s</First>
				<Middle>%s</Middle>
			  </Name>
			  <SNILS>%s</SNILS>
			</ApprovedEmployee>
			<ResultVerification>1</ResultVerification>
		  </VerificationMeasuringInstrument>
		`, row.Col1, row.Col2, row.Col3, row.Col4, row.Col6, row.Col7, row.Col8, row.Col9)
			jStr += jRow
		} else {
			jRow := fmt.Sprintf(`
		  <VerificationMeasuringInstrument>
			<NumberVerification>%d</NumberVerification>
			<DateVerification>%s</DateVerification>
			<TypeMeasuringInstrument>%s</TypeMeasuringInstrument>
			<ApprovedEmployee>
			  <Name>
				<Last>%s</Last>
				<First>%s</First>
				<Middle>%s</Middle>
			  </Name>
			  <SNILS>%s</SNILS>
			</ApprovedEmployee>
			<ResultVerification>2</ResultVerification>
		  </VerificationMeasuringInstrument>
		  `, row.Col1, row.Col2, row.Col4, row.Col6, row.Col7, row.Col8, row.Col9)
			jStr += jRow
		}

	}
	jStr += fmt.Sprintf(`  </VerificationMeasuringInstrumentData>
	<SaveMethod>%s</SaveMethod>
  </Message>`, load)
	jf.Write([]byte(jStr))
}
