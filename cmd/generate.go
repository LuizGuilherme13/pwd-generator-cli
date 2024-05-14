package cmd

import (
	"fmt"
	"log"
	"math/rand"
	"time"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(generateCmd)

	generateCmd.Flags().IntP("length", "l", 6, "password length")
	generateCmd.Flags().BoolP("upper", "u", false, "include uppercase chars")
	generateCmd.Flags().BoolP("numbers", "n", false, "include numbers")
	generateCmd.Flags().BoolP("specials", "s", false, "include specials chars")
}

var generateCmd = &cobra.Command{
	Use: "generate",
	Run: generate,
}

func generate(cmd *cobra.Command, args []string) {
	allChars := "abcdefghijklmnopqrstuvwxyz"

	length, err := cmd.Flags().GetInt("length")
	if err != nil {
		log.Fatalln(err)
	}

	upper, err := cmd.Flags().GetBool("upper")
	if err != nil {
		log.Fatalln(err)
	}
	if upper {
		allChars += "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	}

	includeNum, err := cmd.Flags().GetBool("numbers")
	if err != nil {
		log.Fatalln(err)
	}
	if includeNum {
		allChars += "1234567890"
	}

	includeSpecials, err := cmd.Flags().GetBool("specials")
	if err != nil {
		log.Fatalln(err)
	}
	if includeSpecials {
		allChars += "!@#$%&*()-_=+<>,.:;|?/\\'\""
	}

	r := rand.New(rand.NewSource(time.Now().Unix()))
	newPwd := ""

	for i := 0; i < length; i++ {
		newPwd += string(allChars[r.Intn(len(allChars))])
	}
	fmt.Println(newPwd)
}
