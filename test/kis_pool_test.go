package test

import (
	"context"
	"github.com/aceld/kis-flow/common"
	"github.com/aceld/kis-flow/config"
	"github.com/aceld/kis-flow/flow"
	"testing"
)

func TestNewKisPool(t *testing.T) {

	ctx := context.Background()

	// 1. Create 2 KisFunction configuration instances
	source1 := config.KisSource{
		Name: "TickTokOrder",
		Must: []string{"order_id", "user_id"},
	}

	source2 := config.KisSource{
		Name: "UserOrderErrorRate",
		Must: []string{"order_id", "user_id"},
	}

	myFuncConfig1 := config.NewFuncConfig("funcName1", common.C, &source1, nil)
	if myFuncConfig1 == nil {
		panic("myFuncConfig1 is nil")
	}

	myFuncConfig2 := config.NewFuncConfig("funcName4", common.E, &source2, nil)
	if myFuncConfig2 == nil {
		panic("myFuncConfig4 is nil")
	}

	// 2. Create a KisFlow configuration instance
	myFlowConfig1 := config.NewFlowConfig("flowName1", common.FlowEnable)

	// 3. Create a KisFlow object
	flow1 := flow.NewKisFlow(myFlowConfig1)

	// 4. Link Functions to Flow
	if err := flow1.Link(myFuncConfig1, nil); err != nil {
		panic(err)
	}
	if err := flow1.Link(myFuncConfig2, nil); err != nil {
		panic(err)
	}

	// 5. Commit raw data
	_ = flow1.CommitRow("This is Data1 from Test")
	_ = flow1.CommitRow("This is Data2 from Test")
	_ = flow1.CommitRow("This is Data3 from Test")

	// 6. Execute flow1
	if err := flow1.Run(ctx); err != nil {
		panic(err)
	}
}
