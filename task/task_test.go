package task

import (
	"fmt"
	"testing"
)

func TestNew(t *testing.T) {
	task := New()
	task.Do()
	fmt.Println(task.GetResult())
}
