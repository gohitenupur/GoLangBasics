package main

import "fmt"

const LoginToken string="thisistoken" // is indicates the LoginToken is Public its starts with Capital 
func main() {
	var username string ="nupur"
	fmt.Println("username", username)

	fmt.Printf("Variable is  type : %T \n", username)

	var isLoggedIn bool =false
	fmt.Println("isLoggedIn", isLoggedIn)
	fmt.Printf("Variable is of type : %T \n", isLoggedIn);

	var smallVal int =256
	fmt.Println("smallVal", smallVal)
    fmt.Printf("Variable is of type : %T \n", smallVal);
	
    var bigVal uint16 =256			
	fmt.Println("bigVal", bigVal)
	fmt.Printf("Variable is of type : %T \n", bigVal);

	// uint8 =0 to 255
	//uint16 =0 to 65535
	//uint32 =0 to 4294967295
    //uint64 =0 to 18446744073709551615
    //int8 =-128 to 127
    //int16 =-32768 to 32767
    //int32 =-2147483648 to 2147483647
    //int64 =-9223372036854775808 to 9223372036854775807
    //float32 =-3.40282347e+38 to 3.40282347e+38
    //float64 =-1.7976931348623157e+308 to 1.7976931348623157e+308
    //uint8 =0 to 255
    //uint16 =0 to 65535
    //uint32 =0 to 4294967295
    //uint64 =0 to 18446744073709551615
	// byte =alias for uint8

	var floatVal float32 =255.543257382
	fmt.Println("floatVal", floatVal)
	fmt.Printf("Variable is of type : %T \n", floatVal);
	
    var doubleVal float64 =255.543257382			
	fmt.Println("doubleVal", doubleVal)
	fmt.Printf("Variable is of type : %T \n", doubleVal);


	//default values and some aliasas
	var anatherVariable int 
	fmt.Println("anatherVariable", anatherVariable)
	fmt.Printf("Variable is of type : %T \n", anatherVariable);

	// implicit type 
	var website ="www.google.com"
	fmt.Println("website",website)
	fmt.Printf("Variable is of type : %T \n", website);

	// no var style 
	noOfUsers:=30000.00
	fmt.Println("noOfUsers",noOfUsers)
	fmt.Printf("Variable is of type : %T \n", noOfUsers);

	fmt.Println(LoginToken)
	fmt.Printf("Variable is of type : %T \n", LoginToken);



}