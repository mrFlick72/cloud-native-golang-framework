package date

import (
	"fmt"
	"github.com/mrflick72/cloud-native-golang-framework/utils"
	"testing"
)

func Test_DateFrom_Parsing(t *testing.T) {
	date, err := DateFrom("1985-12-14", utils.AsPointer(DEFAULT_DATE_TIME_FORMATTER))

	if err != nil {
		panic(err)
	}

	fmt.Println(date)
}

func Test_Date_Formatting(t *testing.T) {
	date, err := DateFrom("1985-12-14", utils.AsPointer(DEFAULT_DATE_TIME_FORMATTER))

	if err != nil {
		panic(err)
	}

	fmt.Println(date)
}
