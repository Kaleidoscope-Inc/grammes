// Copyright (c) 2018 Northwestern Mutual.
//
// Permission is hereby granted, free of charge, to any person obtaining a copy
// of this software and associated documentation files (the "Software"), to deal
// in the Software without restriction, including without limitation the rights
// to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
// copies of the Software, and to permit persons to whom the Software is
// furnished to do so, subject to the following conditions:
//
// The above copyright notice and this permission notice shall be included in
// all copies or substantial portions of the Software.
//
// THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
// IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
// FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
// AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
// LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
// OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
// THE SOFTWARE.

package traversal

import (
	"testing"

	. "github.com/smartystreets/goconvey/convey"
)

func TestAddStep(t *testing.T) {
	Convey("Given a graph traversal", t, func() {
		g := NewTraversal()
		Convey("When AddStep is called with []byte", func() {
			b := []byte("test")
			g.AddStep("test", b)
			Convey("Then g should equal g.test(test)", func() {
				So(g.String(), ShouldEqual, "g.test(test)")
			})
		})

		Convey("When AddStep is called with byte", func() {
			b := byte('T')
			g.AddStep("test", b)
			Convey("Then g should equal g.test(T)", func() {
				So(g.String(), ShouldEqual, "g.test(T)")
			})
		})

		Convey("When AddStep is called with int32", func() {
			i := int32(1234)
			g.AddStep("test", i)
			Convey("Then g should equal g.test(1234)", func() {
				So(g.String(), ShouldEqual, "g.test(1234)")
			})
		})

		Convey("When AddStep is called with int64", func() {
			i := int64(1234)
			g.AddStep("test", i)
			Convey("Then g should equal g.test(1234L)", func() {
				So(g.String(), ShouldEqual, "g.test(1234L)")
			})
		})

		Convey("When AddStep is called with float64", func() {
			i := float64(1.234)
			g.AddStep("test", i)
			Convey("Then g should equal g.test(1.234)", func() {
				So(g.String(), ShouldEqual, "g.test(1.234000d)")
			})
		})

		Convey("When AddStep is called with bool", func() {
			b := true
			g.AddStep("test", b)
			Convey("Then g should equal g.test(true)", func() {
				So(g.String(), ShouldEqual, "g.test(true)")
			})
		})

		Convey("When AddStep is called with string", func() {
			b := `field`
			g.AddStep("test", b)
			Convey(`Then g should equal g.test("field")`, func() {
				So(g.String(), ShouldEqual, `g.test("field")`)
			})
		})

		Convey("When AddStep is called with string with quotes", func() {
			b := `"field"`
			g.AddStep("test", b)
			Convey(`Then g should equal g.test("\"field\"")`, func() {
				So(g.String(), ShouldEqual, `g.test("\"field\"")`)
			})
		})

		Convey("When AddStep is called with string with escaped quotes", func() {
			b := `\"field\"`
			g.AddStep("test", b)
			Convey(`Then g should equal g.test("\"field\"")`, func() {
				So(g.String(), ShouldEqual, `g.test("\"field\"")`)
			})
		})
	})

}

func TestString(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'String' is called", func() {
			result := g.String()
			Convey("Then result should equal 'g'", func() {
				So(result, ShouldEqual, "g")
			})
		})
	})
}

func TestRaw(t *testing.T) {
	Convey("Given a ) String { that represents the graph's traversal", t, func() {
		g := NewTraversal()
		Convey("When 'Raw' is called", func() {
			result := g.Raw()
			Convey("Then result should equal 'g'", func() {
				So(result.String(), ShouldEqual, "g")
			})
		})
	})
}

func TestGatherInts(t *testing.T) {
	Convey("When GatherInts is called with one param", t, func() {
		result := gatherInts(1)
		Convey("Then result should equal '1'", func() {
			So(result, ShouldEqual, "1")
		})
	})

	Convey("When GatherInts is called with many params", t, func() {
		result := gatherInts(1, 2, 3)
		Convey("Then result should equal empty string", func() {
			So(result, ShouldEqual, "")
		})
	})

	Convey("When GatherInts is called with no params", t, func() {
		result := gatherInts()
		Convey("Then result should equal empty string", func() {
			So(result, ShouldEqual, "")
		})
	})
}

func TestAddListProperty(t *testing.T) {
	Convey("Given a graph traversal", t, func() {
		g := NewTraversal()
		Convey("When AddListProperty is called with a key and zero properties", func() {
			g.AddListProperty("key1", []string{})
			Convey("Then g should equal g", func() {
				So(g.String(), ShouldEqual, "g")
			})
		})

		Convey("When AddListProperty is called with a key and 2 properties", func() {
			g.AddListProperty("key1", []string{"value1", "value2"})
			Convey("Then g should equal g.property(list, \"key1\", \"value1\").property(list, \"key1\", \"value2\")", func() {
				So(g.String(), ShouldEqual, "g.property(list, \"key1\", \"value1\").property(list, \"key1\", \"value2\")")
			})
		})
	})
}

func TestAddSetProperty(t *testing.T) {
	Convey("Given a graph traversal", t, func() {
		g := NewTraversal()
		Convey("When AddSetProperty is called with a key and zero properties", func() {
			var zeroProps []interface{}
			g.AddSetProperty("key1", zeroProps)
			Convey("Then g should equal g", func() {
				So(g.String(), ShouldEqual, "g")
			})
		})

		Convey("When AddSetProperty is called with a key and 6 properties of different types", func() {
			var props []interface{}
			props = append(props, "value1", 3.4e+38, 1.7e+308, 9223372036854775807, 1, true)
			g.AddSetProperty("key1", props)
			Convey("Then g should equal g.property(set, \"key1\", \"value1\").property(set, \"key1\", 339999999999999996123846586046231871488.000000d).property(set, \"key1\", 169999999999999993883079578865998174333346074304075874502773119193537729178160565864330091787584707988572262467983188919169916105593357174268369962062473635296474636515660464935663040684957844303524367815028553272712298986386310828644513212353921123253311675499856875650512437415429217994623324794855339589632.000000d).property(set, \"key1\", 9223372036854775807).property(set, \"key1\", 1).property(set, \"key1\", true)", func() {
				So(g.String(), ShouldEqual, "g.property(set, \"key1\", \"value1\").property(set, \"key1\", 339999999999999996123846586046231871488.000000d).property(set, \"key1\", 169999999999999993883079578865998174333346074304075874502773119193537729178160565864330091787584707988572262467983188919169916105593357174268369962062473635296474636515660464935663040684957844303524367815028553272712298986386310828644513212353921123253311675499856875650512437415429217994623324794855339589632.000000d).property(set, \"key1\", 9223372036854775807L).property(set, \"key1\", 1L).property(set, \"key1\", true)")
			})
		})
	})
}
