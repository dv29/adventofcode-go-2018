package main

import (
	"fmt"
	"github.com/spf13/viper"
	// "io/ioutil"
	// "flag"
	"io"
	"log"
	"net/http"
	"os"
	"strconv"
)

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func getUrl(day uint) string {
	return fmt.Sprintf("https://adventofcode.com/2018/day/%d/input", day)
}

func printString(s string) {
	fmt.Printf("%s\n", s)
}

func main() {
	// build env vars
	v := viper.New()
	v.SetEnvPrefix("aoc")
	v.BindEnv("session")
	sessionKey := v.GetString("session")
	// fmt.Print("%v\n", os.Getenv("AOC_SESSION"))
	day, err := strconv.Atoi(os.Args[1])
	check(err)

	// create file to write
	file, err := os.Create(fmt.Sprintf("%02d/input", uint(day)))
	check(err)
	defer file.Close()

	// fetch input
	req, err := http.NewRequest("GET", getUrl(uint(day)), nil)
	check(err)
	req.Header.Add("cookie", fmt.Sprintf("session=%s", sessionKey))
	res, err := http.DefaultClient.Do(req)
	defer res.Body.Close()
	bytesWritten, err := io.Copy(file, res.Body)
	fmt.Printf("Bytes Written: %d\n", bytesWritten)
	check(err)
}
