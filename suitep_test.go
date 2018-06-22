package app

import (
	"testing"

	. "gopkg.in/check.v1"
)

// Hook up gocheck into the "go test" runner.
func Test(t *testing.T) { TestingT(t) }

type ParallelSuitesTestSuite struct {
	tag string
}

var _ = Suite(&ParallelSuitesTestSuite{"a"})
var _ = Suite(&ParallelSuitesTestSuite{"b"})
var _ = Suite(&ParallelSuitesTestSuite{"cde"})
var _ = Suite(&ParallelSuitesTestSuite{"xyzzy"})
var _ = Suite(&ParallelSuitesTestSuite{"supercalifragilisticexpialadocious"})

func (s *ParallelSuitesTestSuite) Test(c *C) {

	for i := 0; i < 40000; i++ {
		c.Logf("%s", s.tag)

	}
}
