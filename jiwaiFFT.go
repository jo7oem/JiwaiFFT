package jiwaifft

import (
	"flag"
	"fmt"
	"path"
	"strings"
)

var (
	isPngMode     bool
	isCsvMode     bool
	directoryScan bool
	csvSkipLine   uint
	useColor      int
)

// Gray が標準で使用する色のモード
const (
	Gray = 0
	R
	G
	B
)

// Run entry point
func Run() error {
	flag.BoolVar(&isPngMode, "png", true, "Read input data as PNG or JPEG.")
	flag.BoolVar(&isCsvMode, "csv", true, "Read input data as CSV.")
	if isCsvMode == true {
		isPngMode = false
	}
	flag.BoolVar(&directoryScan, "d", false, "Scan directory")
	flag.UintVar(&csvSkipLine, "skip", 1, "Skip line in CSV")
	color := strings.ToLower(*flag.String("color", "Gray", "Specify the color to use.(R,G,B or Gray)"))
	switch color {
	case "r":
		useColor = R
	case "g":
		useColor = G
	case "b":
		useColor = B
	default:
		useColor = Gray
	}
	_ = useColor

	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		flag.PrintDefaults()
		return nil
	}
	var fileList []string
	if directoryScan {
		fileList = append(fileList, path.Dir(path.Clean(args[0])))
	}
	fmt.Println(len(args))
	for _, i := range fileList {
		fmt.Println(i)
	}
	return nil
}
