// watches for a file to appear

package main

import (
	"encoding/csv"
	"fmt"
	"io/ioutil"
	"os"
	"runtime"
	"strconv"
	"strings"
	"time"
)

const watchedPath = "./source"

func main() {
	runtime.GOMAXPROCS(4)
	for {
		d, _ := os.Open(watchedPath)
		files, _ := d.Readdir(-1) // returns a slice of file info objects
		for _, fi := range files {
			filePath := watchedPath + "/" + fi.Name()
			f, _ := os.Open(filePath)
			data, _ := ioutil.ReadAll(f)
			f.Close()
			os.Remove(filePath) // for demo only in production you would want to confirm the file was processed

			go func(data string) {
				reader := csv.NewReader(strings.NewReader(data))
				records, _ := reader.ReadAll()
				for _, r := range records {
					invoice := new(Invoice)
					invoice.Number = r[0]
					invoice.Amount, _ = strconv.ParseFloat(r[1], 64)
					invoice.PurchaseOrderNumber, _ =
						strconv.Atoi(r[2])
					unixTime, _ := strconv.ParseInt(r[3], 10, 64)
					invoice.InvoiceDate = time.Unix(unixTime, 0)

					fmt.Printf("Received invoice '%v for $%.2f and submitted for processing\n",
						invoice.Number, invoice.Amount)
				}

			}(string(data)) // pass data from for loop in to ensure we have the correct data object
		}
	}
}

type Invoice struct {
	Number              string
	Amount              float64
	PurchaseOrderNumber int
	InvoiceDate         time.Time
}
