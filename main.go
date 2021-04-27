package main

import "flag"

func main() {
	csvFilename := flag.String("csv", "problems.csv", "csv file in format of questions.answer")
	flag.Parse()
	_ = csvFilename
}
