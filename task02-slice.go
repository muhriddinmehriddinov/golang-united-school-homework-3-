package homework

func reverse(input []int64) (result []int64) {
	for i:=0;i<len(input);i++{
		slice2=input[1:len(input)+1]
		for i:=0;i<len(input);i++{
			input[i]=slice2[len(slice2)]
			len(slice2)=len(slice2)-1
		}
	}
	return input [] 
}
