// Copyright 2016-present The Hugo Authors. All rights reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package bufferpool

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBufferPool(t *testing.T) {
	buff1 := GetBuffer()

	buff := GetBuffer()
	fmt.Println(buff.Len())
	fmt.Printf("%p\n", buff)
	fmt.Printf("%p\n", buff1)
	buff.WriteString("do be do be do")
	assert.Equal(t, "do be do be do", buff.String())
	fmt.Println(buff.Len())
	PutBuffer(buff)

	assert.Equal(t, 0, buff.Len())
}
