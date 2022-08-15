// Copyright 2020 The VectorSQL Authors.
//
// Code is licensed under Apache License, Version 2.0.

package datatypes

import (
	"fmt"
	"github.com/CC11001100/vectorsql/src/base/binary"
	"github.com/CC11001100/vectorsql/src/base/errors"
	"github.com/CC11001100/vectorsql/src/datavalues"
	"io"

	//"base/binary"
	//"base/errors"
	//"datavalues"
)

const (
	DataTypeInt64Name = "Int64"
)

type Int64DataType struct {
}

func NewInt64DataType() IDataType {
	return &Int64DataType{}
}

func (datatype *Int64DataType) Name() string {
	return DataTypeInt64Name
}

func (datatype *Int64DataType) Serialize(writer *binary.Writer, v datavalues.IDataValue) error {
	return writer.Int64(datavalues.AsInt(v))
}

func (datatype *Int64DataType) SerializeText(writer io.Writer, v datavalues.IDataValue) error {
	_, err := writer.Write([]byte(fmt.Sprintf("%v", datavalues.AsInt(v))))
	return err
}

func (datatype *Int64DataType) Deserialize(reader *binary.Reader) (datavalues.IDataValue, error) {
	if res, err := reader.Int64(); err != nil {
		return nil, errors.Wrap(err)
	} else {
		return datavalues.ToValue(res), nil
	}
}
