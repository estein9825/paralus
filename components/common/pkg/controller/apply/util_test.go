package apply

import (
	"testing"

	clusterv2 "github.com/RafaySystems/rcloud-base/components/common/proto/types/controller"
)

func TestGetGVK(t *testing.T) {
	ns := clusterv2.Namespace{}

	gvk, err := GetGVK(&ns)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(gvk)
}
