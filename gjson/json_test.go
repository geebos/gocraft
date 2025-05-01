package gjson

import (
	"testing"

	. "github.com/bytedance/mockey"
	. "github.com/smartystreets/goconvey/convey"
)

func TestUnmarshalFromPath(t *testing.T) {
	PatchConvey("TestUnmarshalFromPath", t, func() {
		// test data
		data := `{
			"number":1234567891234567891,
			"double":0.1,
			"simple_string":"string",
			"json_string":"{\"test\":1}"
		}`
		number := int64(1234567891234567891)
		double := 0.1

		PatchConvey("types", func() {
			PatchConvey("string", func() {
				res, err := UnmarshalFromPath[int64](data, "number")
				So(err, ShouldBeNil)
				So(res, ShouldEqual, number)
			})
			PatchConvey("bytes", func() {
				res, err := UnmarshalFromPath[int64]([]byte(data), "number")
				So(err, ShouldBeNil)
				So(res, ShouldEqual, number)
			})
		})

		PatchConvey("parse number", func() {
			PatchConvey("wrong value when type is interface", func() {
				res, err := UnmarshalFromPath[any](data, "number")
				So(err, ShouldBeNil)
				float, ok := res.(float64)
				So(ok, ShouldBeTrue)
				So(int64(float), ShouldNotEqual, number)
			})
			PatchConvey("parse int", func() {
				res, err := UnmarshalFromPath[int64](data, "number")
				So(err, ShouldBeNil)
				So(res, ShouldEqual, number)
			})
			PatchConvey("parse float", func() {
				res, err := UnmarshalFromPath[float64](data, "double")
				So(err, ShouldBeNil)
				So(res, ShouldEqual, double)
			})
		})

		PatchConvey("parse string", func() {
			PatchConvey("simple string", func() {
				str, err := UnmarshalFromPath[string](data, "simple_string")
				So(err, ShouldBeNil)
				So(str, ShouldEqual, "string")
			})
			PatchConvey("json string", func() {
				str, err := UnmarshalFromPath[string](data, "json_string")
				So(err, ShouldBeNil)
				So(str, ShouldEqual, "{\"test\":1}")
			})
		})
	})
}

func TestUnmarshalFromPathWithDefault(t *testing.T) {
	PatchConvey("TestUnmarshalFromPathWithDefault", t, func() {
		PatchConvey("success", func() {
			res := UnmarshalFromPathWithDefault[int](`{"data":1}`, "data", 2)
			So(res, ShouldEqual, 1)
		})
		PatchConvey("default", func() {
			res := UnmarshalFromPathWithDefault[int](`{"data":"1"}`, "data", 2)
			So(res, ShouldEqual, 2)
		})
	})
}
