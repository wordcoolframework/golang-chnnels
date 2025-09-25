package boot

import (
	"bufio"
	"fiber-gorm-channel-ecommerce/src/interface/provider"
	"fiber-gorm-channel-ecommerce/src/pkg/configCore"
	"fiber-gorm-channel-ecommerce/src/pkg/databaseCore"
	"fmt"
	"os"
	"strings"
)

func SeedData(args []string) {

	configCore.ConfigurationSet()

	databaseCore.ConnectMysqlDB()

	var seederName string
	if len(args) > 0 {
		seederName = strings.Join(args, " ")
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		var lines []string
		for scanner.Scan() {
			lines = append(lines, scanner.Text())
		}
		seederName = strings.Join(lines, "\n")
	}

	if seederName == "" {
		fmt.Println("No input text provided")
		return
	}

	seeder, ok := provider.SeederRegister[seederName]

	if !ok {
		fmt.Println("Seeder not found:", seederName)
		return
	}

	fmt.Println("Seeding:", seeder.Run())

}
