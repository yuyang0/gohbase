package hrpc

import (
	"context"

	"github.com/golang/protobuf/proto"
	"github.com/yuyang0/gohbase/pb"
)

// ListNamespace represents a ListNamespace HBase call
type ListNamespace struct {
	base
}

// Name returns the name of this RPC call.
func (cn *ListNamespace) Name() string {
	return "ListNamespaceDescriptors"
}

// ToProto converts the RPC into a protobuf message
func (cn *ListNamespace) ToProto() proto.Message {
	return &pb.ListNamespaceDescriptorsRequest{}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (cn *ListNamespace) NewResponse() proto.Message {
	return &pb.ListNamespaceDescriptorsResponse{}
}

// NewListNamespace creates a new ListNamespace request that will create the given
// namespace in HBase.
// For use by the admin client.
func NewListNamespace(ctx context.Context,
	options ...func(*ListNamespace)) *ListNamespace {
	cn := &ListNamespace{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
	}
	for _, option := range options {
		option(cn)
	}
	return cn
}


// ListTableName represents a ListTableName HBase call
type ListTableName struct {
	base
	namespace string
}

// Name returns the name of this RPC call.
func (cn *ListTableName) Name() string {
	return "ListTableNamesByNamespace"
}

// ToProto converts the RPC into a protobuf message
func (ln *ListTableName) ToProto() proto.Message {
	return &pb.ListTableNamesByNamespaceRequest{
		NamespaceName: &ln.namespace,
	}
}

// NewResponse creates an empty protobuf message to read the response of this
// RPC.
func (lt *ListTableName) NewResponse() proto.Message {
	return &pb.ListTableNamesByNamespaceResponse{}
}

// NewListNamespace creates a new ListNamespace request that will create the given
// namespace in HBase.
// For use by the admin client.
func NewListTableName(ctx context.Context, namespace string,
	options ...func(*ListTableName)) *ListTableName {
	lt := &ListTableName{
		base: base{
			ctx:      ctx,
			resultch: make(chan RPCResult, 1),
		},
		namespace:namespace,
	}
	for _, option := range options {
		option(lt)
	}
	return lt
}
