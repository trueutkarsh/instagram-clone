package utils

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
	"fmt"
	"time"
)

const dateFormat = "2006-01-02"

/*
Custom Date type to store the date in
Database as time.Time but for api
purposes use yyyy-mm-dd format
*/

type Date time.Time

func (d *Date) UnmarshalJSON(b []byte) error {
	var s string
	if err := json.Unmarshal(b, &s); err != nil {
		return err
	}

	t, err := time.Parse(dateFormat, s)
	if err != nil {
		return err
	}

	t = t.In(time.UTC)

	*d = Date(t)

	return nil
}

func (d Date) MarshalJSON() ([]byte, error) {
	return json.Marshal(time.Time(d).Format(dateFormat))
}

func (d Date) String() string {
	return time.Time(d).String()
}

func (date *Date) Scan(value interface{}) (err error) {
	nullTime := &sql.NullTime{}
	err = nullTime.Scan(value)
	if err != nil {
		return err
	}
	*date = Date(nullTime.Time)
	return nil
}

func (d Date) GormDataType() string {
	return "Date"
}

func (d Date) Value() (driver.Value, error) {
	return time.Time(d), nil
}

func test_date() {

	data := []byte(`{"date":"1996-06-04"}`)

	var m struct {
		DOB Date `json:"date"`
	}

	err := json.Unmarshal(data, &m)
	if err != nil {
		fmt.Println("Error unmarshalling json:", err)
	}

	fmt.Println("Parsed date: ", m.DOB)

	b, err := json.Marshal(m)
	if err != nil {
		fmt.Println("Error marshalling object:", err)
	}

	fmt.Println("Marshaled json:", string(b))

}
