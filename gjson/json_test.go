package gjson

import (
	"encoding/json"
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

func TestUnmarshal(t *testing.T) {
	PatchConvey("TestUnmarshal", t, func() {
		PatchConvey("use number", func() {
			data := `{"number":1234}`
			PatchConvey("with options", func() {
				res, err := Unmarshal[map[string]interface{}](data, WithUseNumber())
				So(err, ShouldBeNil)
				So(res, ShouldResemble, map[string]interface{}{
					"number": json.Number("1234"),
				})
			})
			PatchConvey("without options", func() {
				res, err := Unmarshal[map[string]interface{}](data)
				So(err, ShouldBeNil)
				So(res, ShouldResemble, map[string]interface{}{
					"number": float64(1234),
				})
			})
		})

		PatchConvey("disable unknown fields", func() {
			data := `{"number":1234,"unknown":1234}`
			type TestStruct struct {
				Number int64 `json:"number"`
			}
			PatchConvey("with options", func() {
				res, err := Unmarshal[TestStruct](data, WithDisableUnknownFields())
				So(err, ShouldNotBeNil)
				So(err.Error(), ShouldContainSubstring, "json: unknown field")
				So(res, ShouldResemble, TestStruct{Number: 1234})
			})
			PatchConvey("without options", func() {
				res, err := Unmarshal[TestStruct](data)
				So(err, ShouldBeNil)
				So(res, ShouldResemble, TestStruct{Number: 1234})
			})
		})
	})
}

func TestMarshal(t *testing.T) {
	PatchConvey("TestMarshal", t, func() {
		v := map[string]interface{}{
			"html":  "&",
			"field": "test",
		}
		PatchConvey("with options", func() {
			data, err := Marshal[string](v, WithEscapeHtml(false), WithIndent("-", "  "))
			So(err, ShouldBeNil)
			So(data, ShouldEqual, `{
-  "field": "test",
-  "html": "&"
-}
`)
		})
		PatchConvey("without options", func() {
			data, err := Marshal[string](v)
			So(err, ShouldBeNil)
			So(data, ShouldEqual, `{"field":"test","html":"\u0026"}`)
		})
	})
}
