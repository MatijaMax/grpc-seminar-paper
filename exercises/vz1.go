package main

import(
	"fmt"
	"os"
	"bufio"
	"strings"
	"strconv"
)

func exCheck(arr []Student, indeks string) bool{
	for _, s := range arr{
		if s.indeks == indeks{
			return true
		}
	}
	return false
}

func write(filepath string, s Student){
	file, err := os.OpenFile(filepath, os.O_APPEND | os.O_WRONLY, os.ModeAppend)
	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla!")
	}
	defer file.Close()

	writer := bufio.NewWriter(file)
	line := fmt.Sprintf("%s,%s,%s,%.2f\n", s.indeks, s.ime, s.prezime, s.prosek)
	_, err = writer.WriteString(line)
	if err != nil {
		fmt.Println("Greska pri upisu u fajl!")
		return
	}
    writer.Flush()
}

func read(filepath string) []Student{
	arr := []Student{}
	file, err := os.OpenFile(filepath, os.O_RDONLY, os.ModeExclusive)
	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla!")
	}
	defer file.Close()

	fileInfo, err := file.Stat()
    if err != nil {
        fmt.Println("Error reading file info!")
        return arr
    }
    if fileInfo.Size() == 0 {
        fmt.Println("The file is empty.")
        return arr
    }

	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		text := scanner.Text()
		parts := strings.Split(text, ",")
		prosek, _ := strconv.ParseFloat(parts[3], 64)
		arr = append(arr, Student{indeks: parts[0], ime: parts[1], prezime: parts[2], prosek: prosek})
	}
	
	return arr

}

func simpleSort(arr []Student, up bool){
	for i:=0; i<len(arr); i++{
		for j:=i+1; j<len(arr); j++{
			if up{
				if arr[j].prosek < arr[i].prosek {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}
			}else{
				if arr[j].prosek > arr[i].prosek {
					temp := arr[i]
					arr[i] = arr[j]
					arr[j] = temp
				}				
			}
		}
	}
}

func globalAvg(arr []Student){
	var sum float64
	sum = 0
	if len(arr) > 0{
		for _, s := range arr{
			sum += s.prosek
		}
		avg := sum/float64(len(arr))
		fmt.Printf("%.2f", avg)

	}else{
		fmt.Println("Nema upisanih studenata!")
	}
}

func printAll(arr []Student){
	for _, s := range arr{
		fmt.Printf("%s %s %s %.2f \n", s.indeks, s.ime, s.prezime, s.prosek)
	}

}

type Student struct{
	indeks string
	ime string
	prezime string
	prosek float64
}


func main(){
	fmt.Println("APLIKACIJA STUDENTI")
	studenti := read("studenti.txt")
	scanner := bufio.NewScanner(os.Stdin)
	running := true
	for running {
		fmt.Println()
		fmt.Println("1.Unos \n2.Brisanje(Trolling ya) \n3.Rastuci ispis \n4.Opadajuci ispis \n5.Ukupni prosek \n6.Izlaz \n")
		scanner.Scan()
		text := scanner.Text()
		if text == "1"{
			fmt.Println("Indeks: ")
			scanner.Scan()
			indeks := scanner.Text()
			if exCheck(studenti, indeks){
				fmt.Println("Indeks vec postoji!")
				continue
			}
			fmt.Println("Ime: ")
			scanner.Scan()
			ime := scanner.Text()
			fmt.Println("Prezime: ")
			scanner.Scan()
			prezime := scanner.Text()
			fmt.Println("Prosek: ")
			scanner.Scan()
			prosek := scanner.Text()
			prosekC, err := strconv.ParseFloat(prosek, 64)
			if err!=nil{
				fmt.Println("Niste uneli broj kako treba covece!!!ALO!!!")
				continue
			}
			student := Student{
				indeks: indeks,
				ime: ime,
				prezime: prezime,
				prosek: prosekC,
			}
			studenti = append(studenti, student)
			write("studenti.txt", student)

		}else if text == "2"{
			fmt.Println("Nema brisanja studenta :)")
		}else if text == "3"{
			fmt.Println("Sortirani studenti rastuce")
			simpleSort(studenti, true)
			printAll(studenti)
		}else if text == "4"{
			fmt.Println("Sortirani studenti opadajuce")
			simpleSort(studenti, false)
			printAll(studenti)
		}else if text == "5"{
			fmt.Println("Ukupni prosek")
			globalAvg(studenti)
		}else if text == "6"{
			running = false
			fmt.Println("Izlaz...")
		}else {
			fmt.Println("Pogresan unos, pokusajte ponovo!")			
		}
	}

}