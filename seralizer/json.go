package serializer

import (
	"fmt"

	"google.golang.org/protobuf/encoding/protojson"
	proto "google.golang.org/protobuf/proto"
)

func ProtobufToJson(message proto.Message) (string, error) {
	jsonpbMarshaler := protojson.MarshalOptions{
		EmitUnpopulated: true,
		UseProtoNames:   true,
	}
	data, err := jsonpbMarshaler.Marshal(message)
	if err != nil {
		return "", fmt.Errorf("error marshaling to JSON: %v", err)
	}
	return string(data), nil
}
