package models

import (
	"context"
	"fmt"
)

type File struct {
	FilePath string `json:"filePath"`
}

type MyContext struct {
	Arg1 int
	Arg2 string
	Arg3 bool
}

func doSomething(ctx context.Context) {
	myCtx := ctx.Value("my_context").(MyContext)
	arg1 := myCtx.Arg1
	arg2 := myCtx.Arg2
	arg3 := myCtx.Arg3
	fmt.Println(arg1, arg2, arg3)
}

func main() {
	myCtx := MyContext{1, "hello", true}
	ctx := context.WithValue(context.Background(), "my_context", myCtx)
	fmt.Println(ctx)
}
