package faas

import (
	"context"
	"fmt"

	"github.com/aceld/kis-flow/kis"
)

// type FaaS func(context.Context, Flow) error

func JumpFuncHandler(ctx context.Context, flow kis.Flow) error {
	fmt.Println("---> Call JumpFuncHandler ----")

	for _, row := range flow.Input() {
		str := fmt.Sprintf("In FuncName = %s, FuncId = %s, row = %s", flow.GetThisFuncConf().FName, flow.GetThisFunction().GetID(), row)
		fmt.Println(str)
	}

	return flow.Next(kis.ActionJumpFunc("funcName1"))
}
