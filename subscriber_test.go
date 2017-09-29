package pubsub

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSubscriber(t *testing.T) {

	Convey("Test Subscribe/Unsubscribe", t, func() {
		hub := New()
		topics := []string{"top", "ics"}
		So(newSubscriber(hub).Subscribe(topics...).Topics(), ShouldResemble, topics)
		So(newSubscriber(hub).Subscribe(topics...).Unsubscribe("top").Topics(), ShouldResemble, []string{"ics"})
		So(newSubscriber(hub).Subscribe(topics...).Subscribe(topics...).Topics(), ShouldResemble, topics)
	})

	Convey("Test Match", t, func() {
		data := []struct {
			topics     []string
			matches    []string
			notmatches []string
		}{
			{
				[]string{"admin:user"},
				[]string{"admin:user:1", "admin:user"},
				[]string{"admin:username", "other:things", "adm:oops"},
			},
		}

		for _, item := range data {

			subscriber := Subscribe(item.topics...)

			for _, match := range item.matches {
				So(subscriber.Match(match), ShouldBeTrue)
			}

			for _, notmatch := range item.notmatches {
				So(subscriber.Match(notmatch), ShouldBeFalse)
			}
		}
	})

	Convey("Test Publish", t, func() {
		calls := 0
		subscriber := Subscribe("topic").Do(func(message Message) {
			calls += 1
		})
		defer subscriber.Close()

		message := NewMessage("topic", struct{}{})

		So(subscriber.Publish(message), ShouldEqual, 1)
		So(Publish(message), ShouldEqual, 1)

		So(calls, ShouldEqual, 2)
	})

	Convey("Test Publish invalid subscriber func", t, func() {
		subscriber := Subscribe("topic")
		defer subscriber.Close()

		So(subscriber.Publish(NewMessage("topic", "something")), ShouldEqual, 0)
	})

}
