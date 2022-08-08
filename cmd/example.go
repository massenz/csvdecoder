package main

import (
	"fmt"
	"github.com/massenz/csvdecoder/decoder"
	"time"
)

type Account struct {
	AccountID         string
	AccountNumber     string
	PlanID            uint32
	TranID            uint64
	OutstdPrincipal   float32
	PlanSegCreateDate time.Time `datefmt:"2006-01-02 15:04:05"`
}

func (a Account) String() string {
	return fmt.Sprintf("AccountID [%s] AccountNumber [%s] PlanID [%d] "+
		"OutstdPrincipal [%f] on [%s]",
		a.AccountID, a.AccountNumber, a.PlanID, a.OutstdPrincipal, a.PlanSegCreateDate)
}

func main() {
	fname := "testdata/account_data.psv"
	records, err := decoder.ReadRecordsFromFile(fname)

	if err != nil {
		panic(err)
	}

	for _, record := range records {
		var a Account
		decoder.Unmarshal(record, &a)
		fmt.Println(a)
	}
}
