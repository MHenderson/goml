package main

// go get github.com/go-gota/gota/dataframe
import ("fmt"
	    "os"
        "log"
	    "github.com/go-gota/gota/dataframe")

func main() {
	weatherFile, err := os.Open("data-raw/weatherHistory.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer weatherFile.Close()

	weatherDF := dataframe.ReadCSV(weatherFile)

	weatherSummary := weatherDF.Describe()

	fmt.Println(weatherSummary)
}
