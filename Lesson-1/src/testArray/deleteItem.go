package main

func deleteItem(slice []int, index int) []int {
	return append(slice[:index], slice[index+1:]...)
}
