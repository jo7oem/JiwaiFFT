package jiwaifft

import (
	"encoding/csv"
	"io"
	"log"
	"os"
	"path/filepath"
	"strconv"
)

// ReadCSVData1d is input imageJ like csv data
func ReadCSVData1d(fileName string, skipLine int) ([][]float64, error) {
	absPath, err := filepath.Abs(fileName)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	f, err := os.Open(absPath)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer f.Close()
	reader := csv.NewReader(f)
	res := make([][]float64, 1, 1)
	res[0] = make([]float64, 0, 2048)
	for i := 0; ; i++ {
		if i < skipLine {
			_, err := reader.Read()
			if err == io.EOF {
				return res, nil
			}
			if err == csv.ErrFieldCount {
				log.Println(err)
			}
			if err != nil {
				log.Fatalln(err)
				return res, err
			}
			continue
		}
		record, err := reader.Read()
		if err == io.EOF || len(record) < 2 {
			return res, nil
		}
		if err != nil {
			log.Fatalln(err)
			return res, err
		}
		val, err := strconv.ParseFloat(record[1], 64)
		if err != nil {
			log.Println(err, record[1])
			continue
		}
		res[0] = append(res[0], val)
	}
	return res, nil
}
