package main

import (
	//"fmt"
	//"priprema/v2"
	//"bufio"
	//"os"
	//"strconv"
)

/*
type Student struct{
	ime string
	prezime string
	godinaUpisa int
	prosek float64
}
*/

func main(){

	///// UVOD
	//fmt.Println("PRIPREMA")

	/*fmt.Println("HELLO MONEY!")
	//var args []string = os.Args
	args := os.Args
	for _, v := range args{
		fmt.Println(v)
	} 

	i := 3
	pointer := &i
	fmt.Println(pointer)
	fmt.Println(*pointer)
	*pointer = 5
	fmt.Println(i)
	defer :)
	
	fajlovi.Create("marko");
	fajlovi.Write("marko")
	fajlovi.Read("marko")
	*/

	///// V2

	/*

	//z1
	trougao := v2.Trougao{
		A: 3,
		B: 4,
		C: 5,
	}

	heron := v2.Heron(&trougao)
	fmt.Println(heron)
	
	//z2
	array := []int{1, 5, 2, 44, 31}
	v2.Sort(array)
	fmt.Printf("%v", array)
	
	//z3
	
	scanner := bufio.NewScanner(os.Stdin)
	korisnici := make(map[string]Student)
	running := true

	for running {
		fmt.Println("Opcije: \n 1.Unos \n 2.Brisanje \n 3.Ispis\n 4.Izlaz \n")
		fmt.Print("Izaberite opciju!")

		scanner.Scan()
		option := scanner.Text()

		if option == "1"{
			fmt.Print("Unesite broj indeksa: ")
			scanner.Scan()
			indeks := scanner.Text()
			fmt.Print("Unesite ime: ")
			scanner.Scan()
			ime := scanner.Text()
			fmt.Print("Unesite prezime: ")
			scanner.Scan()
			prezime := scanner.Text()
			fmt.Print("Unesite godinu upisa: ")
			scanner.Scan()
			godinaUpisa, _ := strconv.Atoi(scanner.Text())
			fmt.Print("Unesite prosek: ")
			scanner.Scan()
			prosek, _ := strconv.ParseFloat(scanner.Text(), 64)

			student := Student{
				ime: ime,
				prezime: prezime,
				godinaUpisa: godinaUpisa,
				prosek: prosek,
			}
			korisnici[indeks] = student
			fmt.Println("Unet korisnik!")

		}else if option == "2"{
			fmt.Println("Unesite broj indeksa: ")
			scanner.Scan()
			indeks := scanner.Text()
			delete(korisnici, indeks)
			fmt.Println("Korisnik izbrisan!")
		}else if option == "3"{	
			for key, value := range korisnici{
				fmt.Printf("%s %s %s %d %.2f \n", key, value.ime, value.prezime, value.godinaUpisa, value.prosek)
			}

		}else if option == "4"{
			fmt.Println("Izlaz...")
			running = false
		}else{
			fmt.Println("Nepostojeca opcija, ponovite unos!")
		}

	}
	
	

	//z4
	
	
	unsorted := []int{4, 6, 11, 3, 2, 1, 14, 325, 123}
	fmt.Printf("%v", unsorted)
	fmt.Println("\n")
	v2.StartQuickSort(unsorted)
	fmt.Printf("%v", unsorted)
	

	//z5
	
	linkedStack := v2.LinkedStack{
		Tops: nil,
	}
	linkedStack.Push(4)
	fmt.Printf("Top value after first push: %d\n", linkedStack.Tops.Value)
	
	linkedStack.Push(5)
	fmt.Printf("Top value after second push: %d\n", linkedStack.Tops.Value)
	
	linkedStack.Push(6)
	fmt.Printf("Top value after third push: %d\n", linkedStack.Tops.Value)
	linkedStack.Pop()
	fmt.Printf("Top value after pop: %d\n", linkedStack.Tops.Value)
    */
	
} 