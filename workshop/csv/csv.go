package csv

import (
	"bee-playground/utils"
	"bufio"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

func Workshop() {
	file := "file.csv"
	createCSV(file)
	readCSV(file)
}

func readCSV(file string) {
	f, err := os.Open(file)
	if err != nil {
		log.Println(err)
		return
	}

	r := csv.NewReader(f)
	for {
		line, err := r.Read()
		if err != nil {
			log.Println(err)
			break
		}

		utils.Dump(line)
	}
}

func createCSV(file string) {
	f, err := os.Create(file)
	if err != nil {
		log.Println(err)
		return
	}
	defer f.Close()

	w := bufio.NewWriter(f)

	content := `year,industry_code_ANZSIC,industry_name_ANZSIC,rme_size_grp,variable,value,unit
2011,A,"Agriculture, Forestry and Fishing",a_0,Activity unit,46134,COUNT
2011,A,"Agriculture, Forestry and Fishing",a_0,Rolling mean employees,0,COUNT
2011,A,"Agriculture, Forestry and Fishing",a_0,Salaries and wages paid,279,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",a_0,"Sales, government funding, grants and subsidies",8187,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",a_0,Total income,8866,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",a_0,Total expenditure,7618,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",a_0,Operating profit before tax,770,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",a_0,Total assets,55700,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",a_0,Fixed tangible assets,32155,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",b_1-5,Activity unit,21777,COUNT
2011,A,"Agriculture, Forestry and Fishing",b_1-5,Rolling mean employees,38136,COUNT
2011,A,"Agriculture, Forestry and Fishing",b_1-5,Salaries and wages paid,1435,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",b_1-5,"Sales, government funding, grants and subsidies",13359,DOLLARS(millions)
2011,A,"Agriculture, Forestry and Fishing",b_1-5,Total income,13771,DOLLARS(millions)
`
	count, err := w.WriteString(content)
	if err != nil {
		log.Println(err)
		return
	}
	err = w.Flush()
	if err != nil {
		return
	}

	fmt.Printf("characters write %d\n", count)
}
