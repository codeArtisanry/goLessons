package goblin

import (
	"fmt"
	"testing"

	"github.com/franela/goblin"
)

func TestGoblin(t *testing.T) {
	g := goblin.Goblin(t)

	g.Describe("top", func() {
		g.Before(func() {
			fmt.Println("01. top Before 在测试开始之前,仅运行一次,做初始化工作")
		})
		g.After(func() {
			fmt.Println("1. top After  在测试结束之后,仅运行一次,做清理工作")
		})
		g.BeforeEach(func() {
			fmt.Println("00. top BeforeEach  在每一个It及子项目测试开始之前,都运行一次,做独立子项目初始化工作")
		})
		g.AfterEach(func() {
			fmt.Println("00. top AfterEach  在每一个It及子项目测试结束之后,都运行一次,做独立子项目清理工作")
		})
		g.It(" Should Post test1", func() {
			fmt.Println("03. top test1 测试项目执行")
		})
		g.It(" Should Post test It 2", func() {
			fmt.Println("1. top test It 2")
		})
		g.Describe("2. Should Post Nested-level", func() {
			g.Before(func() {
				fmt.Println("2. Nested-level Before 在测试开始之前,仅运行一次,做初始化工作")
			})
			g.After(func() {
				fmt.Println("2. Nested-level After 在测试结束之后,仅运行一次,做清理工作")
			})
			g.BeforeEach(func() {
				fmt.Println("2. Nested-level BeforeEach 在每一个It及子项目测试开始之前,都运行一次,做独立子项目初始化工作")
			})
			g.AfterEach(func() {
				fmt.Println("2. Nested-level AfterEach 在每一个It及子项目测试结束之后,都运行一次,做独立子项目清理工作")
			})
			g.It("test1", func() {
				fmt.Println("2. Nested-level test1")
			})
			g.It("test2", func() {
				fmt.Println("2. Nested-level test2")
			})
		})
		g.It("test2", func() {
			fmt.Println("1. top test2")
		})

	})
}
