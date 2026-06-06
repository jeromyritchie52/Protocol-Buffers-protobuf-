package main

import (
	"fmt"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/reflect/protoreflect"
)

// MaxRecursionDepth defines the safety limit for nested message parsing.
const MaxRecursionDepth = 100

// SafeUnmarshal deserializes a protobuf message with recursion depth protection.
func SafeUnmarshal(data []byte, m proto.Message) error {
	options := proto.UnmarshalOptions{
		RecursionLimit: MaxRecursionDepth,
		DiscardUnknown: true,
	}
	return options.Unmarshal(data, m)
}

func main() {
	fmt.Println("Protocol Buffers optimization layer initialized.")
}