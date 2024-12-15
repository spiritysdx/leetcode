func encryptionCalculate(dataA int, dataB int) int {
    if dataB == 0 {
		return dataA
	}
	return encryptionCalculate(dataA^dataB, (dataA&dataB)<<1)
}
