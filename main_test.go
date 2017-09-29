package pubsub

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestDefaultHub(t *testing.T) {

	Convey("Test default Hub", t, func() {
		topics := []string{"top", "ics"}
		So(Subscribe(topics...).Topics(), ShouldResemble, topics)
		So(Subscribe(topics...).Unsubscribe("top").Topics(), ShouldResemble, []string{"ics"})
		So(Subscribe(topics...).Subscribe(topics...).Topics(), ShouldResemble, topics)
	})

}
