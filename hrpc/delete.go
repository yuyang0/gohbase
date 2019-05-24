// Copyright (C) 2015  The GoHBase Authors.  All rights reserved.
// This file is part of GoHBase.
// Use of this source code is governed by the Apache License 2.0
// that can be found in the COPYING file.

package hrpc

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/yuyang0/gohbase/pb"
)

// DeleteNamespace represents a DeleteTable HBase call
type DeleteNamespace struct {
	base
	namespace []byte
}

// NewDeleteNamespace creates a new DeleteNamespace request that will delete the
// given namespace in HBase. For use by the admin client.
func NewDeleteNamespace(ctx context.Context, namespace []byte) *DeleteNamespace {
	return &DeleteNamespace{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		namespace: namespace,
	}
}

// Name returns the name of this RPC call.
func (dn *DeleteNamespace) Name() string {
	return "DeleteNamespace"
}

// ToProto converts the RPC into a protobuf message
func (dn *DeleteNamespace) ToProto() proto.Message {
	ns := string(dn.namespace)
	return &pb.DeleteNamespaceRequest{
		NamespaceName: &ns,
	}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (dn *DeleteNamespace) NewResponse() proto.Message {
	return &pb.DeleteNamespaceResponse{}
}

// DeleteTable represents a DeleteTable HBase call
type DeleteTable struct {
	base
	namespace []byte
}

// NewDeleteTable creates a new DeleteTable request that will delete the
// given table in HBase. For use by the admin client.
func NewDeleteTable(ctx context.Context, table []byte) *DeleteTable {
	namespace, table := splitTableName(table)
	return &DeleteTable{
		base: base{
			table:    table,
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		namespace:namespace,
	}
}

// Name returns the name of this RPC call.
func (dt *DeleteTable) Name() string {
	return "DeleteTable"
}

// ToProto converts the RPC into a protobuf message
func (dt *DeleteTable) ToProto() proto.Message {
	return &pb.DeleteTableRequest{
		TableName: &pb.TableName{
			Namespace: dt.namespace,
			Qualifier: dt.table,
		},
	}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (dt *DeleteTable) NewResponse() proto.Message {
	return &pb.DeleteTableResponse{}
}
