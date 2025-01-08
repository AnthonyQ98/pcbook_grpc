package sample

import (
	"pcbook/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

// NewKeyboard returns a new sample Keyboard
func NewKeyboard() *pb.Keyboard {
	keyboard := &pb.Keyboard{
		Layout:  randomKeyboardLayout(),
		Backlit: randomBool(),
	}
	return keyboard
}

// NewCPU returns a new sample CPU
func NewCPU() *pb.CPU {
	brand := randomCPUBrand()
	name := randomCPUName(brand)
	numberCores := randomInt(2, 8)
	numberThreads := randomInt(numberCores, 12)
	minGhz := randomFloat32(2.0, 3.5)
	maxGhz := randomFloat32(minGhz, 5.0)
	cpu := &pb.CPU{
		Brand:         brand,
		Name:          name,
		NumberCores:   uint32(numberCores),
		NumberThreads: uint32(numberThreads),
		MinGhz:        minGhz,
		MaxGhz:        maxGhz,
	}
	return cpu
}

// NewGPU returns a new sample GPU
func NewGPU() *pb.GPU {
	brand := randomGPUBrand()
	name := randomGPUName(brand)
	minGhz := randomFloat32(1.0, 1.5)
	maxGhz := randomFloat32(minGhz, 2.0)
	memory := &pb.Memory{Value: uint64(randomInt(2, 6)), Unit: pb.Memory_GIGABYTE}
	gpu := &pb.GPU{
		Brand:  brand,
		Name:   name,
		MinGhz: minGhz,
		MaxGhz: maxGhz,
		Memory: memory,
	}
	return gpu
}

// NewRAM returns a new sample RAM
func NewRAM() *pb.Memory {
	mem := &pb.Memory{Value: uint64(randomInt(4, 64)), Unit: pb.Memory_GIGABYTE}
	return mem
}

// NewSSD returns a new sample SSD
func NewSSD() *pb.Storage {
	mem := &pb.Memory{Value: uint64(randomInt(128, 1024)), Unit: pb.Memory_GIGABYTE}
	return &pb.Storage{Driver: pb.Storage_SSD, Memory: mem}
}

// NewHDD returns a new sample HDD
func NewHDD() *pb.Storage {
	mem := &pb.Memory{Value: uint64(randomInt(1, 6)), Unit: pb.Memory_TERABYTE}
	return &pb.Storage{Driver: pb.Storage_HDD, Memory: mem}
}

// NewScreen returns a new sample Screen
func NewScreen() *pb.Screen {
	return &pb.Screen{
		SizeInch: randomFloat32(13, 17),
		Resolution: &pb.Screen_Resolution{
			Width:  uint32(randomInt(1280, 1920)),
			Height: uint32(randomInt(720, 1080)),
		},
		Panel:      randomScreenPanel(),
		Multitouch: randomBool(),
	}
}

// NewLaptop returns a new sample Laptop
func NewLaptop() *pb.Laptop {
	brand := randomLaptopBrand()
	name := randomLaptopName(brand)
	laptop := &pb.Laptop{
		Id:          randomID(),
		Brand:       brand,
		Name:        name,
		Cpu:         NewCPU(),
		Ram:         NewRAM(),
		Gpus:        []*pb.GPU{NewGPU()},
		Storages:    []*pb.Storage{NewSSD(), NewHDD()},
		Screen:      NewScreen(),
		Keyboard:    NewKeyboard(),
		PriceUsd:    float64(randomFloat32(1500, 3000)),
		ReleaseYear: uint32(randomInt(2015, 2019)),
		UpdatedAt:   timestamppb.Now(),
	}
	return laptop
}
