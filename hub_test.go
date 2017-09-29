package pubsub

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestHub(t *testing.T) {

	Convey("Test blank hub", t, func() {
		hub := New()
		So(hub, ShouldNotBeNil)

	})

}
