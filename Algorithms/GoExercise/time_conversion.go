package main

import (
  "fmt"
  "time"
)

/**
 * Created by Dynastymasra
 * Name     : Dimas Ragil T
 * Email    : dynastymasra@gmail.com
 * LinkedIn : http://www.linkedin.com/in/dynastymasra
 * Github   : https://github.com/dynastymasra
 * Mobile and Backend Developer
 */
const (
  ANSIC       = "Mon Jan _2 15:04:05 2006"
  UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
  RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
  RFC822      = "02 Jan 06 15:04 MST"
  RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
  RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
  RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
  RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
  RFC3339     = "2006-01-02T15:04:05Z07:00"
  RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
  Kitchen     = "3:04PM"
  // Handy time stamps.
  Stamp      = "Jan _2 15:04:05"
  StampMilli = "Jan _2 15:04:05.000"
  StampMicro = "Jan _2 15:04:05.000000"
  StampNano  = "Jan _2 15:04:05.000000000"
)

func staticDate(date string, format string) string {
  t, err := time.Parse(time.UnixDate, date)
	if err != nil {
		panic(err)
	}

  return t.Format(format)
}

func dynamicTime(date string, format string) string {
  t, err := time.Parse("3:04:05PM", date)
  if err != nil {
    panic(err)
  }

  return t.Format(format)
}

func staticTimeNow(date time.Time, format string) string {
  return date.Format(format)
}

func main() {
  fmt.Println("Static Input Date")
  fmt.Println("Date Static Input \t\t:  Sat Mar 7 18:06:39 PST 2015")
  fmt.Println("Format ANSIC \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", ANSIC))
  fmt.Println("Format UnixDate \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", UnixDate))
  fmt.Println("Format RubyDate \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RubyDate))
  fmt.Println("Format RFC822 \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC822))
  fmt.Println("Format RFC822Z \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC822Z))
  fmt.Println("Format RFC850 \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC850))
  fmt.Println("Format RFC1123 \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC1123))
  fmt.Println("Format RFC1123Z \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC1123Z))
  fmt.Println("Format RFC3339 \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC3339))
  fmt.Println("Format RFC3339Nano \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", RFC3339Nano))
  fmt.Println("Format Kitchen \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", Kitchen))
  fmt.Println("Format Stamp \t\t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", Stamp))
  fmt.Println("Format StampMilli \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", StampMilli))
  fmt.Println("Format StampMicro \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", StampMicro))
  fmt.Println("Format StampNano \t\t: ", staticDate("Sat Mar 7 18:06:39 PST 2015", StampNano))

  fmt.Println("Input Static Date Now")
  now := time.Now()
  fmt.Println("Date Staic Time Now \t\t: ", now)
  fmt.Println("Format ANSIC \t\t\t: ", staticTimeNow(now, ANSIC))
  fmt.Println("Format UnixDate \t\t: ", staticTimeNow(now, UnixDate))
  fmt.Println("Format RubyDate \t\t: ", staticTimeNow(now, RubyDate))
  fmt.Println("Format RFC822 \t\t\t: ", staticTimeNow(now, RFC822))
  fmt.Println("Format RFC822Z \t\t\t: ", staticTimeNow(now, RFC822Z))
  fmt.Println("Format RFC850 \t\t\t: ", staticTimeNow(now, RFC850))
  fmt.Println("Format RFC1123 \t\t\t: ", staticTimeNow(now, RFC1123))
  fmt.Println("Format RFC1123Z \t\t: ", staticTimeNow(now, RFC1123Z))
  fmt.Println("Format RFC3339 \t\t\t: ", staticTimeNow(now, RFC3339))
  fmt.Println("Format RFC3339Nano \t\t: ", staticTimeNow(now, RFC3339Nano))
  fmt.Println("Format Kitchen \t\t\t: ", staticTimeNow(now, Kitchen))
  fmt.Println("Format Stamp \t\t\t: ", staticTimeNow(now, Stamp))
  fmt.Println("Format StampMilli \t\t: ", staticTimeNow(now, StampMilli))
  fmt.Println("Format StampMicro \t\t: ", staticTimeNow(now, StampMicro))
  fmt.Println("Format StampNano \t\t: ", staticTimeNow(now, StampNano))

  var time string
  fmt.Println("Input Date Dynamic")
  fmt.Print("Input Time Date [HH:mm:ss(PM/AM)]: ")
  fmt.Scan(&time)
  fmt.Println("Date Input Dynamic : ", time)
  fmt.Println("Format [HH:mm:ss] \t\t\t: ", dynamicTime(time, "15:04:05"))
}
