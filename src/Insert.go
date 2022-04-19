package main

func insertDiseaseToDatabase(filename , diseasename string) bool{
	data,success := ReadFile(filename);
	if success == true {
		if isSanitizedRegex(data){
			//TODO insert into mysql 
		}
	}
	return false;
}
