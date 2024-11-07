package v2

func partition(arr []int, low int, high int)([]int, int){
    pivot := arr[high]
    i := low
    for j:= low; j<high; j++{
        if arr[j] < pivot{
            arr[i], arr[j] = arr[j], arr[i]
            i++
        }
    }
    arr[i], arr[high] = arr[high], arr[i]
    return arr, i

}

func quickSort(arr []int, low int, high int){
    if low < high{
        var p int
        arr, p = partition(arr, low, high)
        quickSort(arr, low, p-1)
        quickSort(arr, p+1, high)
    }
}


func StartQuickSort(arr []int){
    quickSort(arr, 0, len(arr)-1)
}