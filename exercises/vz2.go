package main

import(
	"fmt"
	"os"
	"bufio"
	"strconv"
	"strings"
	"math/rand"
	"sync"
)


type Club struct{
	name string
	nickname string
	year int
}

type Match struct{
	id int
	home string
	guest string
	result string
}

type Clocky struct{
	globalMax int
	localMax int
	localMin int
	Mutex sync.Mutex
}

func writeClub(filename string, c Club){
	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, os.ModeAppend)
	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla")
	}
	defer file.Close()

	line := fmt.Sprintf("%s,%s,%d\n", c.name, c.nickname, c.year)
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(line)
	if err!=nil{
		fmt.Println("Neuspesan upis kluba u fajl!")
	}else{
		fmt.Println("Klub uspesno unet u fajl!")
	}
	writer.Flush()
}

func writeMatch(filename string, m Match){

	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_APPEND, os.ModeAppend)
	
	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla")
	}
	defer file.Close()

	line := fmt.Sprintf("%d,%s,%s,%s\n", m.id, m.home, m.guest, m.result)
	writer := bufio.NewWriter(file)
	_, err = writer.WriteString(line)
	if err!=nil{
		fmt.Println("Neuspesan upis meca u fajl!")
	}else{
		fmt.Println("Mec uspesno unet u fajl!")
	}
	writer.Flush()
}

func writeMatches(filename string, arr []Match){

	file, err := os.OpenFile(filename, os.O_WRONLY | os.O_TRUNC, os.ModeAppend)

	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla")
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	for _, m := range arr{
		line := fmt.Sprintf("%d,%s,%s,%s\n", m.id, m.home, m.guest, m.result)		
		_, err = writer.WriteString(line)
		if err!=nil{
			fmt.Println("Neuspesan upis meca u fajl!")
		}else{
			fmt.Println("Mec uspesno unet u fajl!")
		}
	}
	writer.Flush()
}


func readClub(filename string) []Club{
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeExclusive)
	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla!")
	}
	defer file.Close()

	clubs := []Club{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		parts := strings.Split(line, ",")
		yearInt, _ := strconv.Atoi(parts[2])
		club := Club{
			name: parts[0],
			nickname: parts[1],
			year: yearInt,
		}
		clubs = append(clubs, club)
	}
	return clubs
}

func printClubs(arr []Club){
	for _, c := range arr{
		fmt.Printf("%s %s %d", c.name, c.nickname, c.year)
		fmt.Println()
	}
}

func readMatch(filename string) []Match{
	file, err := os.OpenFile(filename, os.O_RDONLY, os.ModeExclusive)
	if err!=nil{
		fmt.Println("Greska pri otvaranju fajla!")
	}
	defer file.Close()

	matches := []Match{}
	scanner := bufio.NewScanner(file)
	for scanner.Scan(){
		line := scanner.Text()
		parts := strings.Split(line, ",")
		idInt, _ := strconv.Atoi(parts[0])
		match := Match{
			id: idInt,
			home: parts[1],
			guest: parts[2],
			result: parts[3],
		}
		matches = append(matches, match)
	}
	return matches
}

func printMatches(arr []Match){
	for _, m := range arr{
		fmt.Printf("%d %s %s %s", m.id, m.home, m.guest, m.result)
		fmt.Println()
	}
}

func generateMatches(clubs []Club, globalMax int, localMax int, localMin int) bool{

	clocky := Clocky{
		globalMax: globalMax,
		localMax: localMax,
		localMin: localMin,
	}
	c := make(chan Match, len(clubs) * 2)
	id := 0
	for _, club := range clubs{
		for _, otherClub := range clubs{
			if club.name != otherClub.name{
				go generateMatch(id, club.name, otherClub.name, &clocky, c)
				id++
			}
		}
	}

	for i:= 0; i < len(clubs) * 2; i++{
		match := <-c
		writeMatch("mecevi.txt", match)
	}
	return true
}

func findMatchById(arr []Match, id int) (*Match, bool) {
    for i := range arr {
        if arr[i].id == id {
            return &arr[i], true 
        }
    }
    return nil, false
}

func generateMatch(id int, home string, guest string, clocky *Clocky, c chan Match){
	goals:= rand.Intn(clocky.localMax-clocky.localMin+1)+clocky.localMin

	clocky.Mutex.Lock()

	if clocky.globalMax <= goals{
		goals = clocky.globalMax
		clocky.globalMax = 0
	}else{
		clocky.globalMax -= goals
	}

	defer clocky.Mutex.Unlock()

	delimiter := rand.Intn(goals+1)
	hg:= goals - delimiter
	gg:= delimiter

	result:=fmt.Sprintf("%d:%d", hg, gg)
	match:= Match{
		id: id,
		home: home,
		guest: guest,
		result: result,
	}
	c <- match
}

func main(){

	scanner := bufio.NewScanner(os.Stdin)
	running := true
	generated := false

	var maxGoalsGlobal int
	var maxGoalsLocal int
	var minGoalsLocal int
 
	for running {
		fmt.Println("1.Unos kluba\n2.Unos parametara\n")
		scanner.Scan()
		text := scanner.Text()
		if text == "1"{
			fmt.Println("Naziv:")
			scanner.Scan()
			name := scanner.Text()
			fmt.Println("Skraceni naziv:")
			scanner.Scan()
			nickname := scanner.Text()
			fmt.Println("Godina:")
			scanner.Scan()
			yearString := scanner.Text()
			year, err := strconv.Atoi(yearString)
			if err!=nil{
				fmt.Println("Godina mora biti broj!")
				continue
			}
			club := Club{
				name: name,
				nickname: nickname,
				year: year,
			}
			writeClub("klubovi.txt", club)

		}else if text == "2"{
			fmt.Println("Idemo dalje...")

			fmt.Println("MAX golova globalno:")
			scanner.Scan()
			maxGoalsResult, err:= strconv.Atoi(scanner.Text())
			if err!=nil{
				fmt.Println("Morate uneti broj!")
				continue
			}
			maxGoalsGlobal = maxGoalsResult
			fmt.Println("MAX golova po utakmici:")
			scanner.Scan()
			maxGoalsLocal, err= strconv.Atoi(scanner.Text())
			if err!=nil{
				fmt.Println("Morate uneti broj!")
				continue
			}
			fmt.Println("MIN golova po utakmici:")
			scanner.Scan()
			minGoalsLocal, err= strconv.Atoi(scanner.Text())
			if err!=nil{
				fmt.Println("Morate uneti broj!")
				continue
			}
			fmt.Printf("MAX GLOBAL:%d MAX LOCAL:%d MIN LOCAL:%d", maxGoalsGlobal, maxGoalsLocal, minGoalsLocal)
			fmt.Println()
			running = false
		}else{
			fmt.Println("Ponovite unos!")
		}
	}	
	running = true
	for running {
		fmt.Println("1.Ispis klubova\n2.Generisi meceve\n3.Ispisi meceve\n4.Izmeni mec\n5.Izlaz\n")
		scanner.Scan()
		text := scanner.Text()
		if text == "1"{
			clubs := readClub("klubovi.txt")
			printClubs(clubs)
		}else if text == "2"{
			if generated{
				fmt.Println("Mecevi su vec generisani!")
				continue
			}
			clubs := readClub("klubovi.txt")
			generated = generateMatches(clubs, maxGoalsGlobal, maxGoalsLocal, minGoalsLocal)
		}else if text == "3"{
			matches := readMatch("mecevi.txt")
			printMatches(matches)
		}else if text == "4"{
			matches := readMatch("mecevi.txt")
			fmt.Println("Unesite id meca:")
			scanner.Scan()
			idString := scanner.Text()
			id, err := strconv.Atoi(idString)
			if err!=nil{
				fmt.Println("Morate uneti broj!")
				continue
			}

			mec, isFound:=findMatchById(matches, id)
			if isFound==false{
				fmt.Println("Mec nije pronadjen!")
				continue
			}		
			fmt.Println("Unesite broj golova domacina:")
			scanner.Scan()
			domacin := scanner.Text()
			fmt.Println("Unesite broj golova gosta:")
			scanner.Scan()
			gost := scanner.Text()
			
			line := fmt.Sprintf("%s:%s", domacin, gost)
			mec.result = line
			writeMatches("mecevi.txt", matches)

		}else if text == "5"{
			running = false
		}else{
			fmt.Println("Ponovite unos!")
		}
	}

}