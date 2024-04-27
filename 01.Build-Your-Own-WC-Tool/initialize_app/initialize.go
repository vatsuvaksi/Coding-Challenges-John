package initializeapp

import (
	"fmt"
	"wc-tool/service"
)

func InitialzeApplication() {
	fmt.Println("╔════════════════════════════════════════════════════════════════╗")
	fmt.Println("║                  WELCOME TO WC TOOLS                           ║")
	fmt.Println("║    The Application gives you complete knowledge of files       ║")
	fmt.Println("║    Currently, the application supports the following commands: ║")
	fmt.Println("║    [-c -l -w  -m (default)]                                    ║")
	fmt.Println("║    Please enter the command in the following format:           ║")
	fmt.Println("║    vwc -[flag] txt_file_path                                   ║")
	fmt.Println("║                                                                ║")
	fmt.Println("║                            Enjoy!                              ║")
	fmt.Println("║                                                                ║")
	fmt.Println("║                                      - Vatsal       			  ║")
	fmt.Println("╚════════════════════════════════════════════════════════════════╝")
}

func Start() {
	// Parse Input
	input, isStepFinal, err := service.FetchInput()
	if err != nil {
		fmt.Print("err-> ", err)
	}

	if isStepFinal {
		if len(input) == 4 {
			input = append(input, "placeholder.value")
		}

		var tempArr = []string{"empty", input[4], input[1]}
		fmt.Println(service.CommandFactory(tempArr)(tempArr))
	} else {
		fmt.Println(service.CommandFactory(input)(input))
	}

}
