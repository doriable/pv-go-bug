package main

import (
	"fmt"

	pb "pv/preview/gen"

	"github.com/bufbuild/protovalidate-go"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
	"google.golang.org/protobuf/reflect/protoregistry"
	"google.golang.org/protobuf/types/dynamicpb"
	// "google.golang.org/protobuf/types/known/wrapperspb"
)

func main() {
	protoregistry.GlobalFiles = new(protoregistry.Files)
	protoregistry.GlobalTypes = new(protoregistry.Types)
	msg := &pb.TestInt32Rules{
		Test: proto.Int32(5),
	}
	desc := msg.ProtoReflect().Descriptor()
	dynamicMsg := dynamicpb.NewMessage(desc)
	msg.ProtoReflect().Range(func(fd protoreflect.FieldDescriptor, v protoreflect.Value) bool {
		dynamicMsg.Set(fd, v)
		return true
	})

	v, err := protovalidate.New()
	if err != nil {
		fmt.Println("failed to initialize validator:", err)
	}

	if err = v.Validate(dynamicMsg); err != nil {
		fmt.Println("validation failed:", err)
	} else {
		fmt.Println("validation succeeded")
	}
}
