package unitest

import (
	"fmt"
	"testing"
)

//单元测试
/*
 命令：
 go test  在写了_test.go的名录下执行，自动调用Test开头的方法
 go test -cover  打印出单元测试覆盖率
 go test -cover -coverprofile -res.out  // go提供-coverfrofile参数把结果输出到文件
 go tool cover -html=res.out   可以打开输出文件
 */
func TestSum(t *testing.T){
	got := snm(10)
	fmt.Println(got)
	want := 55
	if want != got{
		t.Fail()
	}
}

/*
基准测试 测试执行b.n（一千万次）的性能，算出每次执行的平均时间
命令：
go test -bench=Sum
go test -bench=Sum -bench=mem 可以查看每次操作的内存分配情况
 */

func BenchmarkSum(b *testing.B){
	for i:=0; i<b.N; i++{
		snm(100)
	}
}