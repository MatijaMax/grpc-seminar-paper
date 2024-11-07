package v2

func Sort(array []int){
	for i:= range(len(array)){
		for j:= range(len(array)){
			if(array[i] > array[j]){
				temp:= array[j]
				array[j] = array[i]
				array[i] = temp
			}
		}
	}	
}