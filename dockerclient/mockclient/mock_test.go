package mockclient

import (
	"reflect"
	"testing"
	"github.com/gaia-adm/mr-burns/dockerclient"
)

func TestMockInterface(t *testing.T) {
	iface := reflect.TypeOf((*dockerclient.Client)(nil)).Elem()
	mock := NewMockClient()

	if !reflect.TypeOf(mock).Implements(iface) {
		t.Fatalf("Mock does not implement the Client interface")
	}
}