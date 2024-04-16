package main

import (
  "fmt"
  "github.com/nleeper/goment"
  "flag"
  "os"
  "time"
)

var (
  year *int
  month *int
  day *int
  hour *int
  minute *int
  second *int
  nanosecond *int
  location *string
  timezone *time.Location
)

func init() {
  year = flag.Int("year", 0, "Year")
  month = flag.Int("month", 0, "Month")
  day = flag.Int("day", 0, "Day")
  hour = flag.Int("hour", 0, "Hour")
  minute = flag.Int("minute", 0, "Minute")
  second = flag.Int("second", 0, "Second")
  nanosecond = flag.Int("nanosecond", 0, "Nanosecond")
  location = flag.String("timezone", "America/Toronto", "Timezone")
}

func main() {
  // SetLocale to ensure local specific names, start of week, etc
  goment.SetLocale("en")

  flag.Parse()

  timezone, err := time.LoadLocation(*location)

  t, err := goment.New(goment.DateTime{
    Year:        *year,
    Month:       *month,
    Day:         *day,
    Hour:        *hour,
    Minute:      *minute,
    Second:      *second,
    Nanosecond:  *nanosecond,
    Location:    timezone,
  })

  if err != nil {
    fmt.Println(err)
    os.Exit(1)
  }

  fmt.Println(t.FromNow())
}
