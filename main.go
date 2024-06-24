package main

import "compressor/cmd"

func main() {
	cmd.Execute()

	//m := map[rune]string{
	//	'l': "0",
	//	'e': "10",
	//	'h': "110",
	//	'o': "111",
	//}
	//
	//fmt.Println(m)
	//
	//var tableBuf bytes.Buffer
	//if err := gob.NewEncoder(&tableBuf).Encode(m); err != nil {
	//	log.Fatal("can't serialize the table")
	//}
	//fmt.Println(tableBuf)
	//
	//r := bytes.NewReader(tableBuf.Bytes())
	//var mp map[rune]string
	//gob.NewDecoder(r).Decode(&mp)
	//fmt.Println(mp)
}
