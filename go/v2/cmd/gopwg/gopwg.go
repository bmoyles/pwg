package main

import (
	"flag"
	"fmt"
	"github.com/bmoyles/gopwg"
	"log"
	"os"
	"runtime/pprof"
)

var (
	count      = flag.Int("count", 1, "Number of passwords to generate")
	dictfile   = flag.String("dictfile", "/usr/share/dict/words", "Dictionary file")
	cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
)

func main() {
	flag.Parse()
	if *cpuprofile != "" {
		f, err := os.Create(*cpuprofile)
		if err != nil {
			log.Fatal(err)
		}
		pprof.StartCPUProfile(f)
		defer pprof.StopCPUProfile()
	}
	file, err := os.Open(*dictfile)
	if err != nil {
		log.Fatalln(err)
	}
	defer file.Close()

	fileInfo, err := file.Stat()
	if err != nil {
		log.Fatalln(err)
	}

	pwg := gopwg.NewPasswordGenerator(file, fileInfo.Size())
	passwords, err := pwg.Generate(*count)
	if err != nil {
		log.Fatalln(err)
	}
	for _, password := range passwords {
		fmt.Println(password)
	}
}
