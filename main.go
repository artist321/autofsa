package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
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
	var fn string
	var load string
	if len(os.Args) > 2 {
		fn = os.Args[1]
		load = os.Args[2]
		if strings.Contains(load, "-1") {
			load = "1"
		} else if strings.Contains(load, "-2") {
			load = "2"
		} else {
			fmt.Println("Введите корректный Тип сохранения данных: сsv2xml export.csv -1 [Как черновики] или сsv2xml file.csv -2 [Как отправленные]")
			os.Exit(1)
		}
	} else if len(os.Args) == 1 {
		fn = os.Args[1]
		load = "1"
	} else {
		fmt.Println("Введите название файлы [export.csv]")
		os.Exit(1)
	}

	f, err := os.Open(fn)
	if err != nil {
		panic(err)
	}

	defer f.Close()
	r := csv.NewReader(f)
	r.Comma = ';'
	r.Comment = '#'
	var recs []Rec
	for i := 1; ; i++ {
		fmt.Println(i)
		row, err := r.Read()
		if err != nil {
			break
		}
		//fmt.Println(snum, len(snum))
		row[0] = strings.Map(func(r rune) rune {
			if unicode.IsPrint(r) {
				return r
			}
			return -1
		}, row[0])
		//fmt.Println("clean", snum, len(snum))

		num, _ := strconv.Atoi(row[0])
		//slins, _ := strconv.Atoi(row[8])
		r := Rec{
			Col1: num,
			Col2: row[1],
			Col3: row[2],
			Col4: row[3],
			Col5: row[4],
			Col6: row[5],
			Col7: row[6],
			Col8: row[7],
			Col9: row[8],
		}
		recs = append(recs, r)
	}
	jf, err := os.Create(f.Name() + ".xml")
	if err != nil {
		panic(err)
	}
	defer jf.Close()
	jStr := `<?xml version="1.0" encoding="utf-8"?>
	<Message xsi:noNamespaceSchemaLocation="schema.xsd" xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance">
	  <VerificationMeasuringInstrumentData>
`
	for _, row := range recs {
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
