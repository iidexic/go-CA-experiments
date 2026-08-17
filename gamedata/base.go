package gamedata

import "os"

var (
	windowName  = "CellularAutomaton"
	windowSize  = [2]int{1280, 720}
	windowScale = 2
)

func WindowName() string {
	return windowName
}
func WindowSize() [2]int {
	return windowSize
}
func WindowScale() int {
	return windowScale
}
func GameSize() [2]int {
	return [2]int{windowSize[0] / windowScale, windowSize[1] / windowScale}
}

func CWD() string {
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return cwd
}

func GetResource(name, category string) string { return "" }
