package tempus

import "fmt"

type Date struct {
	Year  int
	Month int
	Day   int
}

func (d Date) MarshalText() ([]byte, error) {
	s := fmt.Sprintf("%04d-%02d-%02d", d.Year, d.Month, d.Day)
	return []byte(s), nil
}

func (d *Date) UnmarshalText(data []byte) error {
	var year, month, day int

	s := string(data)
	_, err := fmt.Sscanf(s, "%04d-%02d-%02d", &year, &month, &day)
	if err != nil {
		return fmt.Errorf("%q - bad date (%w)", s, err)
	}

	d.Year = year
	d.Month = month
	d.Day = day

	return nil
}
