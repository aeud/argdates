package argdates

import (
	"fmt"
	"log"
	"testing"
)

func TestDays(t *testing.T) {
	start := ""
	end := ""
	duration := uint(0)
	offset := uint(1)
	days, err := Days(&start, &end, &duration, &offset)
	if err != nil {
		t.Error(err)
	}
	log.Println(len(days))
	for _, d := range days {
		fmt.Println(d)
	}

}
