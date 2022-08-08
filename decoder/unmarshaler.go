package decoder

import (
	"log"
	"reflect"
	"strconv"
	"time"
)

// Unmarshal the data in the `Record` into the `target` struct; the field names in the record must
// match the target struct fields exactly ( case sensitive) and the values must be of the correct
// type.
//
// Use the `datefmt` annotation to specify the date format for date fields.
func Unmarshal(record Record, target interface{}) {
	for f, v := range record {
		fv := reflect.ValueOf(target).Elem().FieldByName(f)
		t, found := reflect.TypeOf(target).Elem().FieldByName(f)
		var datefmt string
		if found {
			datefmt = t.Tag.Get("datefmt")
		}
		if fv.IsValid() {
			switch fv.Kind() {
			case reflect.String:
				fv.Set(reflect.ValueOf(v))
			case reflect.Uint:
				if s, err := strconv.ParseUint(v, 10, 64); err == nil {
					fv.Set(reflect.ValueOf(uint(s)))
				}
			case reflect.Uint32:
				if s, err := strconv.ParseUint(v, 10, 32); err == nil {
					fv.Set(reflect.ValueOf(uint32(s)))
				}
			case reflect.Uint64:
				if s, err := strconv.ParseUint(v, 10, 64); err == nil {
					fv.Set(reflect.ValueOf(s))
				}
			case reflect.Float32:
				if s, err := strconv.ParseFloat(v, 32); err == nil {
					fv.Set(reflect.ValueOf(float32(s)))
				}
			case reflect.Float64:
				if s, err := strconv.ParseFloat(v, 64); err == nil {
					fv.Set(reflect.ValueOf(s))
				}
			case reflect.Struct:
				// We currently only support time.Time beyond the standard types
				if fv.Type().Name() == "Time" {
					if datefmt != "" {
						if t, err := time.Parse(datefmt, v); err == nil {
							fv.Set(reflect.ValueOf(t))
						}
					}
				} else {
					log.Printf("Unsupported struct type %s", fv.Type().Name())
				}
			default:
				log.Printf("Unsupported field type (%v): %s\n", fv.Kind(), f)
			}

		}
	}
}
