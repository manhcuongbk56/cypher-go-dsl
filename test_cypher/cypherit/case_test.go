package cypherit

import (
	"fmt"
	"testing"
	"time"
)

func TestTimeZone(t *testing.T) {
	utc, err := time.LoadLocation("Asia/Ho_Chi_Minh")
	if err != nil {
		fmt.Println("err: ", err.Error())
	}
	fmt.Println(utc)
}
