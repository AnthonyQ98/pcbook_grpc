package sample

import (
	"pcbook/pb"
	"time"

	"math/rand"

	"github.com/google/uuid"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

func randomKeyboardLayout() pb.Keyboard_Layout {
	switch rand.Intn(3) {
	case 1:
		return pb.Keyboard_QWERTY
	case 2:
		return pb.Keyboard_QWERTZ
	default:
		return pb.Keyboard_AZERTY
	}
}

func randomBool() bool {
	return rand.Intn(2) == 1
}

func randomCPUBrand() string {
	return randomStringFromSet("Intel", "AMD")
}

func randomCPUName(brand string) string {
	if brand == "Intel" {
		return randomStringFromSet("Xeon E-2286M", "Core i9-9980HK", "Core i7-9750H")
	}
	return randomStringFromSet("Ryzen Threadripper 1950X", "Ryzen 5 2600", "Ryzen 7 2700")
}

func randomGPUBrand() string {
	return randomStringFromSet("NVIDIA", "AMD")
}

func randomGPUName(brand string) string {
	if brand == "NVIDIA" {
		return randomStringFromSet("RTX 2060", "RTX 2070", "GTX 1080")
	}
	return randomStringFromSet("RX 5700", "RX 5600", "RX 5600XT")
}

func randomStringFromSet(a ...string) string {
	n := len(a)
	if n == 0 {
		return ""
	}
	return a[rand.Intn(n)]
}

func randomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}

func randomFloat32(min float32, max float32) float32 {
	return min + rand.Float32()*(max-min)
}

func randomScreenPanel() pb.Screen_Panel {
	if rand.Intn(2) == 1 {
		return pb.Screen_IPS
	}
	return pb.Screen_OLED
}

func randomID() string {
	return uuid.New().String()
}

func randomLaptopBrand() string {
	return randomStringFromSet("Dell", "Lenovo", "HP", "Asus", "Microsoft", "Apple")
}

func randomLaptopName(brand string) string {
	switch brand {
	case "Dell":
		return randomStringFromSet("Latitude", "Vostro", "XPS", "Alienware")
	case "Lenovo":
		return randomStringFromSet("Thinkpad", "Legion", "IdeaPad")
	case "HP":
		return randomStringFromSet("Spectre", "Elitebook", "Omen")
	case "Asus":
		return randomStringFromSet("ROG", "TUF", "VivoBook")
	case "Microsoft":
		return randomStringFromSet("Surface", "Laptop Surface")
	case "Apple":
		return randomStringFromSet("Macbook", "Macbook Air", "Macbook Pro")
	}
	return ""
}
