func diStringMathch(s string) []int {
	n := len(s)
	var output,output2 []int
	for i:=0;i<n+1 ;i++{
		output = append(output, i)
	}
	fmt.Println(n,output)
	str := strings.Split(s,"")
	fmt.Println(str)
	for _,j := range str{
		fmt.Println(j)
		if j == "I" {
		output2	= append(output2,output[0])
			output= output[1:]
		} else {
			fmt.Println("output[n] = ",output[n])
		   output2 = append(output2,output[n])
		   output = output[:n]
		   fmt.Println(output)
		}
		n--
	}
	output2 = append(output2, output[0])
	return output2
}
