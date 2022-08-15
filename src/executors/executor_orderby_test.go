// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package executors

import (
	"context"
	"github.com/CC11001100/vectorsql/src/columns"
	"github.com/CC11001100/vectorsql/src/datablocks"
	"github.com/CC11001100/vectorsql/src/datatypes"
	"github.com/CC11001100/vectorsql/src/mocks"
	"github.com/CC11001100/vectorsql/src/planners"
	"github.com/CC11001100/vectorsql/src/processors"
	"github.com/CC11001100/vectorsql/src/transforms"
	"testing"



	"github.com/stretchr/testify/assert"
)

func TestOrderByExecutor(t *testing.T) {
	tests := []struct {
		name   string
		plan   planners.IPlan
		source []interface{}
		expect *datablocks.DataBlock
	}{
		{
			name: "simple",
			plan: planners.NewOrderByPlan(
				planners.Order{
					Expression: planners.NewVariablePlan("name"),
					Direction:  "asc",
				},
				planners.Order{
					Expression: planners.NewVariablePlan("age"),
					Direction:  "desc",
				},
			),
			source: mocks.NewSourceFromSlice(
				mocks.NewBlockFromSlice(
					[]*columns.Column{
						{Name: "name", DataType: datatypes.NewStringDataType()},
						{Name: "age", DataType: datatypes.NewInt32DataType()},
					},
					[]interface{}{"x", 11},
					[]interface{}{"z", 13},
					[]interface{}{"y", 12},
					[]interface{}{"y", 13},
				)),
			expect: mocks.NewBlockFromSlice(
				[]*columns.Column{
					{Name: "name", DataType: datatypes.NewStringDataType()},
					{Name: "age", DataType: datatypes.NewInt32DataType()},
				},
				[]interface{}{"x", 11},
				[]interface{}{"y", 13},
				[]interface{}{"y", 12},
				[]interface{}{"z", 13},
			),
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			mock, cleanup := mocks.NewMock()
			defer cleanup()
			ctx := NewExecutorContext(mock.Ctx, mock.Log, mock.Conf, mock.Session)

			stream := mocks.NewMockBlockInputStream(test.source)

			tctx := transforms.NewTransformContext(mock.Ctx, mock.Log, mock.Conf)
			datasource := transforms.NewDataSourceTransform(tctx, stream)

			orderby := NewOrderByExecutor(ctx, test.plan.(*planners.OrderByPlan))
			result, err := orderby.Execute()
			assert.Nil(t, err)

			sink := processors.NewSink("sink")
			pipeline := processors.NewPipeline(context.Background())
			pipeline.Add(datasource)
			pipeline.Add(result.In)
			pipeline.Add(sink)
			pipeline.Run()

			err = pipeline.Wait(func(x interface{}) error {
				actual := x.(*datablocks.DataBlock)
				expect := test.expect
				assert.True(t, mocks.DataBlockEqual(actual, expect))
				return nil
			})
			assert.Nil(t, err)
		})
	}
}
