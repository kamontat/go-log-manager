package om_test

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestSettingOutput(t *testing.T) {
	Convey("Given output type", t, func() {

		Convey("When check 2 possible output type", func() {

			Convey("Then shouldn't be equal", nil)

		})

	})

	Convey("Given output of log and print", t, func() {

		Convey("When create new default output", func() {

			Convey("Then Stdout will be true", nil)

			Convey("Then File will be true", nil)

			Convey("Then FileName should contain /tmp folder", nil)

		})

		Convey("When create new output with file disable", func() {

			Convey("And FileName not exist", func() {

				Convey("Then get file will be nil and cause error", nil)

			})

			Convey("And FileName still exist", func() {

				Convey("Then the file result should contain error anyway", nil)

			})

		})

	})

}
