package dailytemperatures

import (
	"fmt"
	"reflect"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	ret := dailyTemperatures([]int{13, 12, 15, 11, 9, 12, 16})

	if !reflect.DeepEqual(ret, []int{2, 1, 4, 2, 1, 1, 0}) {
		t.Error(fmt.Sprint(ret) + " != {2,1,4,2,1,1,0}")
	}
}
