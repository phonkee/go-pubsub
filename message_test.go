package pubsub

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestMessage(t *testing.T) {

	Convey("Test New Message", t, func() {
		topic := "some:topic"
		body := []byte("this is body of message")
		message := NewMessage(topic, body)

		So(message.Body(), ShouldResemble, body)
		So(message.Topic(), ShouldEqual, topic)
	})

	Convey("Test New Message JSON Marshal", t, func() {
		data := []struct {
			data     interface{}
			expected []byte
		}{
			{
				struct{ Test string }{"value"},
				[]byte("{\"Test\":\"value\"}"),
			},
		}

		for _, item := range data {
			So(NewMessage("", item.data).Body(), ShouldResemble, item.expected)
		}
	})

	Convey("Test New String Message", t, func() {
		data := []string{
			"this is example message",
			"other example message",
		}

		for _, item := range data {
			So(NewMessage("", item).Body(), ShouldResemble, []byte(item))
		}
	})
}
