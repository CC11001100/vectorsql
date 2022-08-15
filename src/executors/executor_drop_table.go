// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package executors

import (
	"databases"
	"planners"
)

// 删除表执行器 
type DropTableExecutor struct {
	ctx  *ExecutorContext
	plan *planners.DropTablePlan
}

func NewDropTableExecutor(ctx *ExecutorContext, plan planners.IPlan) IExecutor {
	return &DropTableExecutor{
		ctx:  ctx,
		plan: plan.(*planners.DropTablePlan),
	}
}

func (executor *DropTableExecutor) Execute() (*Result, error) {
	ectx := executor.ctx
	ast := executor.plan.Ast

	schema := ectx.session.GetDatabase()
	if !ast.FromTables[0].Qualifier.IsEmpty() {
		schema = ast.FromTables[0].Qualifier.String()
	}
	database, err := databases.GetDatabase(schema)
	if err != nil {
		return nil, err
	}

	table := ast.FromTables[0].Name.String()
	if err := database.Executor().DropTable(table); err != nil {
		return nil, err
	}

	result := NewResult()
	return result, nil
}

func (executor *DropTableExecutor) String() string {
	return ""
}
