package serializer

import (
	"fmt"
	"io/ioutil"

	"google.golang.org/protobuf/proto"
)

func WriteProtobufToJSONFile(message proto.Message, filename string) error {
	data, err := ProtobufToJson(message)
	if err != nil {
		return fmt.Errorf("error marshaling to JSON: %v", err)
	}
	err = ioutil.WriteFile(filename, []byte(data), 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func WriteProtobufToBinaryFile(message proto.Message, filename string) error {
	data, err := proto.Marshal(message)
	if err != nil {
		return fmt.Errorf("error marshaling to binary: %v", err)
	}
	err = ioutil.WriteFile(filename, data, 0644)
	if err != nil {
		return fmt.Errorf("error writing to file: %v", err)
	}
	return nil
}

func ReadProtobufFromBinaryFile(filename string, message proto.Message) error {
	data, err := ioutil.ReadFile(filename)
	if err != nil {
		return fmt.Errorf("error reading from file: %v", err)
	}
	err = proto.Unmarshal(data, message)
	if err != nil {
		return fmt.Errorf("error unmarshaling from binary: %v", err)
	}
	return nil
}
