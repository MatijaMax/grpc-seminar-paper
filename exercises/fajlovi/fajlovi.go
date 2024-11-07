package fajlovi

import(
	"os"
	"fmt"
	"bufio"
)

func Create(filename string){
	os.Create(filename)
}

func Read(filepath string){

	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModeExclusive)
	if err!= nil{
		fmt.Println(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		fmt.Println(scanner.Text())
	}
}

func Write(filepath string){
	
	file, err := os.OpenFile(filepath, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err!= nil{
		fmt.Println(err)
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	for i:= 1; i < 11; i++{
		line := fmt.Sprintf("test%d\n", i)
		_, err = writer.WriteString(line)
		if err != nil{
			fmt.Println(err)
		}
	}

	if err := writer.Flush(); err != nil {
        fmt.Println("Error flushing buffer:", err)
    }
}