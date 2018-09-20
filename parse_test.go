package parsecommand

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
)

func TestParse(t *testing.T) {
	var tmpLine string
	var args []string
	var err error
	Convey("Testing the Parse func", t, func() {
		Convey("Basic with double-quotes", func() {
			tmpLine = `psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir"`

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 8)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `-p`)
			So(args[5], ShouldEqual, `mypassword`)
			So(args[6], ShouldEqual, `copy`)
			So(args[7], ShouldEqual, `c:\path to my dir`)
		})

		Convey("Another basic with double-quotes", func() {
			tmpLine = `psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir" hallo`

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 9)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `-p`)
			So(args[5], ShouldEqual, `mypassword`)
			So(args[6], ShouldEqual, `copy`)
			So(args[7], ShouldEqual, `c:\path to my dir`)
			So(args[8], ShouldEqual, `hallo`)
		})

		Convey("Basic with single-quotes", func() {
			tmpLine = `psexec \\machine -u MYDOMAIN\myuser -p mypassword copy 'c:\path to my dir'`

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 8)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `-p`)
			So(args[5], ShouldEqual, `mypassword`)
			So(args[6], ShouldEqual, `copy`)
			So(args[7], ShouldEqual, `c:\path to my dir`)
		})

		Convey("Another basic with single-quotes", func() {
			tmpLine = `psexec \\machine -u MYDOMAIN\myuser -p mypassword copy 'c:\path to my dir' hallo`

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 9)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `-p`)
			So(args[5], ShouldEqual, `mypassword`)
			So(args[6], ShouldEqual, `copy`)
			So(args[7], ShouldEqual, `c:\path to my dir`)
			So(args[8], ShouldEqual, `hallo`)
		})

		Convey("With surrounding spaces", func() {
			tmpLine = `  psexec \\machine -u MYDOMAIN\myuser -p mypassword copy "c:\path to my dir"  `

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 8)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `-p`)
			So(args[5], ShouldEqual, `mypassword`)
			So(args[6], ShouldEqual, `copy`)
			So(args[7], ShouldEqual, `c:\path to my dir`)
		})

		Convey("With multiple spaces between args", func() {
			tmpLine = `psexec   \\machine   -u MYDOMAIN\myuser    -p mypassword   copy    "c:\path to my dir"`

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 8)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `-p`)
			So(args[5], ShouldEqual, `mypassword`)
			So(args[6], ShouldEqual, `copy`)
			So(args[7], ShouldEqual, `c:\path to my dir`)
		})

		Convey("With non-ACSII ending args", func() {
			tmpLine = `psexec   \\machine   -u MYDOMAIN\myuser    测试字符串`

			args, err = Parse(tmpLine)
			So(err, ShouldBeNil)
			So(len(args), ShouldEqual, 5)

			So(args[0], ShouldEqual, `psexec`)
			So(args[1], ShouldEqual, `\\machine`)
			So(args[2], ShouldEqual, `-u`)
			So(args[3], ShouldEqual, `MYDOMAIN\myuser`)
			So(args[4], ShouldEqual, `测试字符串`)
		})
	})
}
