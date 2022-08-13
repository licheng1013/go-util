/*
 *                                  Apache License
 *                            Version 2.0, January 2004
 *                         http://www.apache.org/licenses/
 */

package test

import (
	"gitee.com/licheng1013/go-util/util"
	"testing"
)

func TestRandom(t *testing.T) {
	t.Log(util.RandomNumber(10))
	t.Log(util.RandomString(26))
	t.Log(util.RandomRangeNumPlus(10, 30))
}
