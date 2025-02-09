package function

// CheckInNumberSlice 检查一个数字是否在切片中
func CheckInNumberSlice(num uint64, slice []uint64) bool {
	for _, v := range slice {
		if v == num {
			return true
		}
	}
	return false
}

// DelEleInSlice 从切��中删除一个元素
func DelEleInSlice(num uint64, slice []uint64) []uint64 {
	for i, v := range slice {
		if v == num {
			return append(slice[:i], slice[i+1:]...)
		}
	}
	return slice
}
