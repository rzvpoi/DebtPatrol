package main

import (
	"encoding/csv"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"os"
	"strconv"
	"time"
)

type ContactDetails struct {
	Uid1_data string
	Uid2_data string
	Uid       string
	Name      string
	Title     string
	Sum       string
	Time      string
	Logs      []ContactDetails
}

func main() {
	tmpl := template.Must(template.ParseFiles("patrolu.html"))

	fs := http.FileServer(http.Dir(".\assets\\"))

	http.Handle("/assets/", http.StripPrefix("/static", fs))

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		logs, uid1_data, uid2_data := readData()
		data := ContactDetails{
			Uid1_data: uid1_data,
			Uid2_data: uid2_data,
			Logs:      logs,
		}
		if r.Method != http.MethodPost {
			tmpl.Execute(w, data)
			return
		}

		currentTime := time.Now()
		currentDate := currentTime.Format("02/01/2006")

		details := ContactDetails{
			Uid:   r.FormValue("uid"),
			Title: r.FormValue("title"),
			Sum:   r.FormValue("sum"),
			Time:  currentDate,
		}

		// do something with details
		writeData(details)

		logs, uid1_data, uid2_data = readData()
		data = ContactDetails{
			Uid1_data: uid1_data,
			Uid2_data: uid2_data,
			Logs:      logs,
		}

		tmpl.Execute(w, data)
	})

	http.Handle("/assets/images/", http.StripPrefix("/assets/images/", http.FileServer(http.Dir("assets"))))

	log.Print("Listening on :8082...")
	http.ListenAndServe(":8082", nil)
}

func writeData(detail ContactDetails) {

	file, err := os.OpenFile("data.csv", os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV writer
	writer := csv.NewWriter(file)

	uid, _ := strconv.Atoi(detail.Uid)
	if uid == 1 {
		detail.Name = "razvan"
	} else {
		detail.Name = "vlad"
	}

	// Write data rows
	row := []string{detail.Uid, detail.Name, string(detail.Title), string(detail.Sum), string(detail.Time)}
	err = writer.Write(row)
	if err != nil {
		panic(err)
	}

	// Flush any buffered data to the underlying writer (the file).
	writer.Flush()

	// Check for any errors during the flush.
	if err = writer.Error(); err != nil {
		panic(err)
	}

	log.Print("New data entry added! Values: {" + detail.Name + ", " + detail.Sum + " RON }")
}

func readData() ([]ContactDetails, string, string) {
	file, err := os.Open("data.csv")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	// Create a CSV reader
	reader := csv.NewReader(file)

	// Read all the records from the CSV file
	records, err := reader.ReadAll()
	if err != nil {
		panic(err)
	}

	// Process the records
	var detail []ContactDetails
	for idx, record := range records {
		if len(record) != 5 {
			fmt.Println("Malformed record:", record)
			continue
		}
		if idx == 0 {
			continue
		}

		uid := record[0]
		name := record[1]
		title := record[2]
		sum := record[3]
		time := record[4]

		data := ContactDetails{Uid: uid, Name: name, Title: title, Sum: sum, Time: time}
		detail = append(detail, data)
	}

	// Print the data
	sum_id1 := 0
	sum_id2 := 0
	for _, data := range detail {
		uid, _ := strconv.Atoi(data.Uid)
		sum, _ := strconv.Atoi(data.Sum)
		if uid == 1 {
			sum_id1 += sum
		} else {
			sum_id2 += sum
		}

	}

	return detail, strconv.Itoa(sum_id2 / 2), strconv.Itoa(sum_id1 / 2)
}
