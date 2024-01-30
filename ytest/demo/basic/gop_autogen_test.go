package main

import (
	"github.com/goplus/yap/ytest"
	"testing"
)

type case_foo struct {
	ytest.Case
}
//line ytest/demo/basic/foo_ytest.gox:1
func (this *case_foo) Main() {
//line ytest/demo/basic/foo_ytest.gox:1:1
	server := new(foo)
//line ytest/demo/basic/foo_ytest.gox:2:1
	server.Main()
//line ytest/demo/basic/foo_ytest.gox:3:1
	this.RunMock("foo.com", server)
//line ytest/demo/basic/foo_ytest.gox:5:1
	this.Run("get /p/$id", func() {
//line ytest/demo/basic/foo_ytest.gox:6:1
		id := "123"
//line ytest/demo/basic/foo_ytest.gox:7:1
		this.Get("http://foo.com/p/" + id)
//line ytest/demo/basic/foo_ytest.gox:8:1
		this.RetWith(200)
//line ytest/demo/basic/foo_ytest.gox:9:1
		this.Json(map[string]string{"id": "12"})
	})
}
func Test_foo(t *testing.T) {
	ytest.Gopt_Case_TestMain(new(case_foo), t)
}