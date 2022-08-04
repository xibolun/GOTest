package basic

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSort(t *testing.T) {

	ast := assert.New(t)

	arr := []string{"/dev/sda1", "/dev/sda10", "/dev/nvme01", "/dev/sda2", "/dev/sda", "/dev/sdb", "/dev/sdc", "/dev/sdd"}
	sort.Slice(arr, func(i, j int) bool {
		itemI := len(arr[i])
		itemJ := len(arr[j])

		if itemI == itemJ {
			return arr[i] < arr[j]
		}
		return itemI < itemJ
	})

	ast.Equal(arr[0], "/dev/sda")
	ast.Equal(arr[1], "/dev/sda1")
	ast.Equal(arr[2], "/dev/sda2")
	ast.Equal(arr[3], "/dev/sda10")
	ast.Equal(arr[len(arr)-1], "/dev/nvme01")
}
