package algorithm

func insertDiseaseToDatabase(filename, diseasename string) bool {
	data, success := ReadFile(filename)
	if success == true {
		if IsSanitizedRegex(data) {
			//TODO insert into mysql
		}
	}
	return false
}
