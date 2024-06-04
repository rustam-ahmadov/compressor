package main

import "compressor/cmd"

func main() {
	// compressor pack -vlc <path to file> -out /path/to/packed-file
	cmd.Execute()
	//text := vlc.Decode("20 30 3C 18 77 4A E4 4D 28")
	//fmt.Println(text)
}
