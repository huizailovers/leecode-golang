package day09

func quickSort(arr []int ,start ,end int){
	par := pratition(arr, start ,end)
	quickSort(arr,start,par-1)
	quickSort(arr,par+1,end)
}
func pratition(arr []int ,start ,end int ) int{
	value := arr[end]
	i := start
	for j :=i; j<= end; j++ {
		if arr[j] < value{
			if i != j {
				arr[i],arr[j] = arr[j] ,arr [i]

			}
			i++
		}

	}
	return i

}
