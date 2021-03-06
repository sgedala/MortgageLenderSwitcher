package main

import (
	"encoding/json"
	"errors"
	"fmt"
	

	"github.com/hyperledger/fabric/core/chaincode/shim"
	
)



// MORTGAGE is a high level smart contract that MORTGAGEs together business artifact based smart contracts
type MORTGAGE struct {

}

// UserDetails is for storing user credentials

type UserDetails struct{	
	UserID string `json:"uid"`
	Password string `json:"password"`
	UserType string   `json:"userType"`
	
	
}

// BorrowerDetails is for storing User Details

type BorrowerDetails struct{	
	UserID string `json:"uid"`	
	Gender string `json:"gender"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Dob string `json:"dob"`
	Email string `json:"email"`
	Phone string `json:"phone"`	
	Address string `json:"address"`
	City string `json:"city"`
	Zip string `json:"zip"`
	PropertyAddress string `json:"propertyAddress"`
	PropertyCity string `json:"propertyCity"`
	PropertyZip string `json:"propertyZip"`
	
	LenderId string `json:"lenderId"`
	LenderName string `json:"lenderName"`
	ProductName string `json:"productName"`
	LoanAmount string `json:"loanAmount"`
	InterestRate string `json:"interestRate"`	
	DocumentsSubmitted string `json:"documentsSubmitted"`
	SwitchUserFlag string `json:"switchUserFlag"`
	SwitchLenderId string `json:"switchLenderId"`
	SwitchLenderName string `json:"switchLenderName"`
	Status string `json:"status"`
}

// LenderDetails is for storing Lender Details

type LenderDetails struct{	
	LenderId string `json:"lenderId"`
	LenderName string `json:"lenderName"`
	ProductId string   `json:"productId"`
	ProductName string `json:"productName"`
	ProductType string `json:"productType"`
	Rate string `json:"rate"`
	
	
}
// TiTleInfo is to store title information of a particular property

type TitleInfo struct{	
	
	Name string `json:"name"`
	Address string `json:"address"`
	City string `json:"city"`
	State string   `json:"state"`
	Zip string `json:"zip"`
	Info string `json:"info"`
	
	
}
// Transaction is for storing transaction Details


// to return the verify result
type VerifyU struct{	
	Result string `json:"result"`
}
	
func createUserDetails(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "uid", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "password", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "userType", Type: shim.ColumnDefinition_STRING, Key: false}
	
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("UserDetails", columnDefsTableOne)
}

func createTitleInfo(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "name", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableOneDef := shim.ColumnDefinition{Name: "state", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixTableOneDef := shim.ColumnDefinition{Name: "info", Type: shim.ColumnDefinition_STRING, Key: false}
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("TitleInfo", columnDefsTableOne)
}

func createLenderDetails(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "productId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableOneDef := shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "productType", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixTableOneDef := shim.ColumnDefinition{Name: "rate", Type: shim.ColumnDefinition_STRING, Key: false}
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("LenderDetails", columnDefsTableOne)
}

func createBorrowerDetails(stub shim.ChaincodeStubInterface) error {
	// Create table one
	var columnDefsTableOne []*shim.ColumnDefinition
	columnOneTableOneDef := shim.ColumnDefinition{Name: "uid", Type: shim.ColumnDefinition_STRING, Key: true}
	columnTwoTableOneDef := shim.ColumnDefinition{Name: "gender", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThreeTableOneDef := shim.ColumnDefinition{Name: "firstName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFourTableOneDef := shim.ColumnDefinition{Name: "lastName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFiveTableOneDef := shim.ColumnDefinition{Name: "dob", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixTableOneDef := shim.ColumnDefinition{Name: "email", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSevenTableOneDef := shim.ColumnDefinition{Name: "phone", Type: shim.ColumnDefinition_STRING, Key: false}
	columnEightTableOneDef := shim.ColumnDefinition{Name: "address", Type: shim.ColumnDefinition_STRING, Key: false}
	columnNineTableOneDef := shim.ColumnDefinition{Name: "city", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTenTableOneDef := shim.ColumnDefinition{Name: "zip", Type: shim.ColumnDefinition_STRING, Key: false}
															
	columnElevenTableOneDef := shim.ColumnDefinition{Name: "propertyAddress", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTwelveTableOneDef := shim.ColumnDefinition{Name: "propertyCity", Type: shim.ColumnDefinition_STRING, Key: false}
	columnThirteenTableOneDef := shim.ColumnDefinition{Name: "propertyZip", Type: shim.ColumnDefinition_STRING, Key: false}
	
	columnFourteenTableOneDef := shim.ColumnDefinition{Name: "lenderId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnFifteenTableOneDef := shim.ColumnDefinition{Name: "lenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSixteenTableOneDef := shim.ColumnDefinition{Name: "productName", Type: shim.ColumnDefinition_STRING, Key: false}
	columnSeventeenTableOneDef := shim.ColumnDefinition{Name: "loanAmount", Type: shim.ColumnDefinition_STRING, Key: false}
	columnEighteenTableOneDef := shim.ColumnDefinition{Name: "interestRate", Type: shim.ColumnDefinition_STRING, Key: false}
	columnNineteenTableOneDef := shim.ColumnDefinition{Name: "documentsSubmitted", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTwentyTableOneDef := shim.ColumnDefinition{Name: "switchUserFlag", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTwentyOneTableOneDef := shim.ColumnDefinition{Name: "switchLenderId", Type: shim.ColumnDefinition_STRING, Key: false}
	columnTwentyTwoTableOneDef := shim.ColumnDefinition{Name: "switchLenderName", Type: shim.ColumnDefinition_STRING, Key: false}
	

	
		
	columnDefsTableOne = append(columnDefsTableOne, &columnOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwoTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThreeTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFiveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSevenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnEightTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnNineTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnElevenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwelveTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnThirteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFourteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnFifteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSixteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnSeventeenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnEighteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnNineteenTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwentyTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwentyOneTableOneDef)
	columnDefsTableOne = append(columnDefsTableOne, &columnTwentyTwoTableOneDef)
	fmt.Println("creating table !!!!! %s",columnDefsTableOne)
	return stub.CreateTable("BorrowerDetails", columnDefsTableOne)
}
// Init initializes the smart contracts
func (t *MORTGAGE) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

fmt.Println("creating table !!!!! ")
	// Check if table already exists
	_, err := stub.GetTable("LenderDetails")
	fmt.Println("err  table !!!!! &b",err)
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createLenderDetails(stub)
	
	
	
	
	fmt.Println("creating table created!!!!! %b",err)
	if err != nil {
		return nil, errors.New("Failed creating UserDetails.")
	}
		
	
	
	
	
	
	// Check if table already exists
	_, err = stub.GetTable("BorrowerDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createBorrowerDetails(stub)
	
	if err != nil {
		return nil, errors.New("Failed creating BorrowerDetails.")
	}
	
	// Check if table already exists
	_, err = stub.GetTable("TitleInfo")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createTitleInfo(stub)
	
	if err != nil {
		return nil, errors.New("Failed creating TitleInfo.")
	}
	
	// Check if table already exists
	_, err = stub.GetTable("UserDetails")
	if err == nil {
		// Table already exists; do not recreate
		return nil, nil
	}

	// Create application Table
	err = createUserDetails(stub)
	
	if err != nil {
		return nil, errors.New("Failed creating TitleInfo.")
	}
	
	return nil, nil
}
	
//registerUser to register a user
func (t *MORTGAGE) registerUserDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

		fmt.Println("registerUserDetails" )

		
	
		uid:=args[0]
		password:=args[1]
		userType:=args[2]
		
		

		
		
		


		// Insert a row
		ok, err := stub.InsertRow("UserDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: uid}},
				&shim.Column{Value: &shim.Column_String_{String_: password}},
				&shim.Column{Value: &shim.Column_String_{String_: userType}},
				
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}
	
//registerUser to register a user
func (t *MORTGAGE) registerLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

fmt.Println("registerLender" )

		
	
		lenderId:=args[0]
		lenderName:=args[1]
		productId:=args[2]
		productName:=args[3]
		productType:=args[4]
		rate:=args[5]
		

		
		
		


		// Insert a row
		ok, err := stub.InsertRow("LenderDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: lenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderName}},
				&shim.Column{Value: &shim.Column_String_{String_: productId}},
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: productType}},
				&shim.Column{Value: &shim.Column_String_{String_: rate}},
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}
//registerUser to register a user
func (t *MORTGAGE) registerBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("registerBorrower" )


	uid:=args[0]
	gender:=args[1]
	firstName:=args[2]
	lastName:=args[3]
	dob:=args[4]
	email:=args[5]
	phone:=args[6]
	address:=args[7]
	city:=args[8]
	zip:=args[9]
	propertyAddress:=args[10]
	propertyCity:=args[11]
	propertyZip:=args[12]
	lenderId:=args[13]
	lenderName:=args[14]
	productName:=args[15]
	loanAmount:=args[16]
	interestRate:=args[17]
	documentsSubmitted:=args[18]
	switchUserFlag:=args[19]
	switchLenderId:=args[20]
	switchLenderName:=args[21]
	
		
		


		// Insert a row
		ok, err := stub.InsertRow("BorrowerDetails", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: uid}},
				&shim.Column{Value: &shim.Column_String_{String_: gender}},
				&shim.Column{Value: &shim.Column_String_{String_: firstName}},
				&shim.Column{Value: &shim.Column_String_{String_: lastName}},
				&shim.Column{Value: &shim.Column_String_{String_: dob}},
				&shim.Column{Value: &shim.Column_String_{String_: email}},
				&shim.Column{Value: &shim.Column_String_{String_: phone}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: propertyAddress}},
				&shim.Column{Value: &shim.Column_String_{String_: propertyCity}},
				&shim.Column{Value: &shim.Column_String_{String_: propertyZip}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: lenderName}},
				&shim.Column{Value: &shim.Column_String_{String_: productName}},
				&shim.Column{Value: &shim.Column_String_{String_: loanAmount}},
				&shim.Column{Value: &shim.Column_String_{String_: interestRate}},
				&shim.Column{Value: &shim.Column_String_{String_: documentsSubmitted}},
				&shim.Column{Value: &shim.Column_String_{String_: switchUserFlag}},
				&shim.Column{Value: &shim.Column_String_{String_: switchLenderId}},
				&shim.Column{Value: &shim.Column_String_{String_: switchLenderName}},
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

//Registering property and their title info
func (t *MORTGAGE) registerTitleInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("registerTitleInfo")


	name:=args[0]
	address:=args[1]
	city:=args[2]
	state:=args[3]
	zip:=args[4]
	info:=args[5]

	// Insert a row
		ok, err := stub.InsertRow("TitleInfo", shim.Row{
			Columns: []*shim.Column{
				&shim.Column{Value: &shim.Column_String_{String_: name}},
				&shim.Column{Value: &shim.Column_String_{String_: address}},
				&shim.Column{Value: &shim.Column_String_{String_: city}},
				&shim.Column{Value: &shim.Column_String_{String_: state}},
				&shim.Column{Value: &shim.Column_String_{String_: zip}},
				&shim.Column{Value: &shim.Column_String_{String_: info}},
				
				
			}})

			fmt.Println("Iserted rows !!!!! %b",ok)
		if err != nil {
			return nil, err 
		}
		if !ok && err == nil {
			return nil, errors.New("Row already exists.")
		}
			
		return nil, nil

}

// to get the deatils of a borrowers against lenderid 
func (t *MORTGAGE) getBorowersWithLenderId(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("getBorowersWithLenderId")
	

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting uid to query")
	}

	lenderId := args[0]
	

	// Get the row pertaining to this Email
	var columns []shim.Column	
	rows, err := stub.GetRows("BorrowerDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	res2E := []*BorrowerDetails{}
	for row := range rows {	
		newApp:= new(BorrowerDetails)
		newApp.UserID = row.Columns[0].GetString_()
		newApp.Gender = row.Columns[1].GetString_()
		newApp.FirstName = row.Columns[2].GetString_()
		newApp.LastName = row.Columns[3].GetString_()
		newApp.Dob = row.Columns[4].GetString_()
		newApp.Email = row.Columns[5].GetString_()
		newApp.Phone = row.Columns[6].GetString_()
		newApp.Address = row.Columns[7].GetString_()
		newApp.City = row.Columns[8].GetString_()
		newApp.Zip = row.Columns[9].GetString_()
		newApp.PropertyAddress = row.Columns[10].GetString_()
		newApp.PropertyCity = row.Columns[11].GetString_()
		newApp.PropertyZip = row.Columns[12].GetString_()
		newApp.LenderId = row.Columns[13].GetString_()
		newApp.LenderName = row.Columns[14].GetString_()
		newApp.ProductName = row.Columns[15].GetString_()
		newApp.LoanAmount = row.Columns[16].GetString_()
		newApp.InterestRate = row.Columns[17].GetString_()
		newApp.DocumentsSubmitted = row.Columns[18].GetString_()
		newApp.SwitchUserFlag = row.Columns[19].GetString_()
		newApp.SwitchLenderId = row.Columns[20].GetString_()
		newApp.SwitchLenderName = row.Columns[21].GetString_()
		
		fmt.Println("Hello, World!%s",newApp.LenderId)
		fmt.Println("Hello, World!%s",lenderId)
		fmt.Println("Hello, World!%s",newApp.SwitchUserFlag)
		fmt.Println("Hello, World!%s",newApp.SwitchLenderId)
		if (newApp.LenderId == lenderId || (newApp.SwitchUserFlag == "Yes" && newApp.SwitchLenderId ==lenderId)){
		fmt.Println("--------------------------------------------------------------------------------")
		res2E=append(res2E,newApp)		
		}	
	}
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
// to get the deatils of a product against productId 
func (t *MORTGAGE) getProductRatesFromLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("getProductRatesFromLender")
	

	productId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this productId
	var columns []shim.Column

	rows, err := stub.GetRows("LenderDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the product with id  " + productId + "\"}"
		return nil, errors.New(jsonResp)
	}
	
	res2E :=  []*LenderDetails{}
	for row := range rows { 
	newApp:= new(LenderDetails)
	newApp.LenderId = row.Columns[0].GetString_()
	newApp.LenderName = row.Columns[1].GetString_()
	newApp.ProductId = row.Columns[2].GetString_()
	newApp.ProductName = row.Columns[3].GetString_()
	newApp.ProductType = row.Columns[4].GetString_()
	newApp.Rate = row.Columns[5].GetString_()
	if newApp.ProductId == productId{
		res2E=append(res2E,newApp)		
		}	
		
		}
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// to get the deatils of a product against lenderId 
func (t *MORTGAGE) getLenderDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("getLenderDetails")
	

	lenderId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this lenderId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: lenderId}}
	
	columns = append(columns, col1)

	row, err := stub.GetRow("LenderDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the lender with id  " + lenderId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	

	
	res2E :=  LenderDetails{}
	
	res2E.LenderId = row.Columns[0].GetString_()
	res2E.LenderName = row.Columns[1].GetString_()
	res2E.ProductId = row.Columns[2].GetString_()
	res2E.ProductName = row.Columns[3].GetString_()
	res2E.ProductType = row.Columns[4].GetString_()
	res2E.Rate = row.Columns[5].GetString_()
	
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
// to get the deatils of a borrower against email (for internal testing, irrespective of org)
func (t *MORTGAGE) getBorrower(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting uid to query")
	}

	Email := args[0]
	

	// Get the row pertaining to this Email
	var columns []shim.Column	
	rows, err := stub.GetRows("BorrowerDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	res2E := []*BorrowerDetails{}
	for row := range rows {	
		newApp:= new(BorrowerDetails)
		newApp.UserID = row.Columns[0].GetString_()
		newApp.Gender = row.Columns[1].GetString_()
		newApp.FirstName = row.Columns[2].GetString_()
		newApp.LastName = row.Columns[3].GetString_()
		newApp.Dob = row.Columns[4].GetString_()
		newApp.Email = row.Columns[5].GetString_()
		newApp.Phone = row.Columns[6].GetString_()
		newApp.Address = row.Columns[7].GetString_()
		newApp.City = row.Columns[8].GetString_()
		newApp.Zip = row.Columns[9].GetString_()
		newApp.PropertyAddress = row.Columns[10].GetString_()
		newApp.PropertyCity = row.Columns[11].GetString_()
		newApp.PropertyZip = row.Columns[12].GetString_()
		newApp.LenderId = row.Columns[13].GetString_()
		newApp.LenderName = row.Columns[14].GetString_()
		newApp.ProductName = row.Columns[15].GetString_()
		newApp.LoanAmount = row.Columns[16].GetString_()
		newApp.InterestRate = row.Columns[17].GetString_()
		newApp.DocumentsSubmitted = row.Columns[18].GetString_()
		newApp.SwitchUserFlag = row.Columns[19].GetString_()
		newApp.SwitchLenderId = row.Columns[20].GetString_()
		newApp.SwitchLenderName = row.Columns[21].GetString_()
		
		if newApp.Email == Email{
		res2E=append(res2E,newApp)		
		}	
	}
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
// to get the deatils of a product against borrowerId 
func (t *MORTGAGE) fetchBorrowerDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("fetchBorrowerDetails")
	

	borrowerId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this borrowerId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: borrowerId}}
	
	columns = append(columns, col1)
	fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("BorrowerDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the borrower with id  " + borrowerId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	fmt.Println("number of columsn :!!!!!!!!!!!!", len(row.Columns))

	
	res2E :=  BorrowerDetails{}
	fmt.Println("user id  :!!!!!!!!!!!!%s", row.Columns[0].GetString_())
	
	res2E.UserID = row.Columns[0].GetString_()
	res2E.Gender = row.Columns[1].GetString_()
	res2E.FirstName = row.Columns[2].GetString_()
	res2E.LastName = row.Columns[3].GetString_()
	res2E.Dob = row.Columns[4].GetString_()
	res2E.Email = row.Columns[5].GetString_()
	res2E.Phone = row.Columns[6].GetString_()
	res2E.Address = row.Columns[7].GetString_()
	res2E.City = row.Columns[8].GetString_()
	res2E.Zip = row.Columns[9].GetString_()
	res2E.PropertyAddress = row.Columns[10].GetString_()
	res2E.PropertyCity = row.Columns[11].GetString_()
	res2E.PropertyZip = row.Columns[12].GetString_()
	res2E.LenderId = row.Columns[13].GetString_()
	res2E.LenderName = row.Columns[14].GetString_()
	res2E.ProductName = row.Columns[15].GetString_()
	res2E.LoanAmount = row.Columns[16].GetString_()
	res2E.InterestRate = row.Columns[17].GetString_()
	res2E.DocumentsSubmitted = row.Columns[18].GetString_()
	res2E.SwitchUserFlag = row.Columns[19].GetString_()
	res2E.SwitchLenderId = row.Columns[20].GetString_()
	res2E.SwitchLenderName = row.Columns[21].GetString_()
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}

// to get the deatils of a user
func (t *MORTGAGE) fetchUserDetails(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("fetchUserDetails")
	

	userId := args[0]
	fmt.Println("Hello, World!!")

	// Get the row pertaining to this userId
	var columns []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: userId}}
	
	columns = append(columns, col1)
	fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("UserDetails", columns)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the user with id  " + userId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	fmt.Println("number of columsn :!!!!!!!!!!!!", len(row.Columns))

	
	res2E := UserDetails{}
	fmt.Println("user id  :!!!!!!!!!!!!%s", row.Columns[0].GetString_())
	
	res2E.UserID = row.Columns[0].GetString_()
	res2E.Password = row.Columns[1].GetString_()
	res2E.UserType = row.Columns[2].GetString_()
	
	
	
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil

}
func (t *MORTGAGE) switchLenders(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("switchLenders")
	

	
	switchLenderId := args[0]
	borrowerId := args[1]
	switchLenderName := args[2]
	fmt.Println("Hello, World!!")

	
	
		
	var colum []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: borrowerId}}
	
	colum = append(colum, col1)
	fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("BorrowerDetails", colum)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the borrower with id  " + borrowerId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	fmt.Println("number of columsn :!!!!!!!!!!!!", len(row.Columns))
	
	
	
		
	//row, err := stub.GetRow("BorrowerDetails", columns)
	
	var columns []*shim.Column
		col1 = shim.Column{Value: &shim.Column_String_{String_: row.Columns[0].GetString_()}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[1].GetString_()}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[2].GetString_()}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[3].GetString_()}}
		col5 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[4].GetString_()}}
		col6 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[5].GetString_()}}
		col7 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[6].GetString_()}}
		col8 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[7].GetString_()}}
		col9 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[8].GetString_()}}
		col10 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[9].GetString_()}}
		col11 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[10].GetString_()}}
		col12 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[11].GetString_()}}
		col13 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[12].GetString_()}}
		col14 := shim.Column{Value: &shim.Column_String_{String_: switchLenderId}}
		col15 := shim.Column{Value: &shim.Column_String_{String_: switchLenderName}}
		col16 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[15].GetString_()}}
		col17 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[16].GetString_()}}
		col18 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[17].GetString_()}}
		col19 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[18].GetString_()}}
		col20 := shim.Column{Value: &shim.Column_String_{String_: "No"}}
		col21 := shim.Column{Value: &shim.Column_String_{String_: ""}}
		col22 := shim.Column{Value: &shim.Column_String_{String_: ""}}
		
		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)
		columns = append(columns, &col7)
		columns = append(columns, &col8)
		columns = append(columns, &col9)
		columns = append(columns, &col10)
		columns = append(columns, &col11)
		columns = append(columns, &col12)
		columns = append(columns, &col13)
		columns = append(columns, &col14)
		columns = append(columns, &col15)
		columns = append(columns, &col16)
		columns = append(columns, &col17)
		columns = append(columns, &col18)
		columns = append(columns, &col19)
		columns = append(columns, &col20)
		columns = append(columns, &col21)
		columns = append(columns, &col22)
		
		row = shim.Row{Columns: columns}
		
	ok, err := stub.ReplaceRow("BorrowerDetails", row)
		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}
		
		
		return nil,nil



}


func (t *MORTGAGE) updateSwitchLender(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("updateSwitchLender")
	

	
	switchLenderId := args[0]
	borrowerId := args[1]
	switchLenderName := args[2]
	fmt.Println("Hello, World!!")

	
	
		
	var colum []shim.Column
	col1 := shim.Column{Value: &shim.Column_String_{String_: borrowerId}}
	
	colum = append(colum, col1)
	fmt.Println("Hello, World!!!")
	row, err := stub.GetRow("BorrowerDetails", colum)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the borrower with id  " + borrowerId + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist
	
	fmt.Println("number of columsn :!!!!!!!!!!!!", len(row.Columns))
	
	
	
		
	//row, err := stub.GetRow("BorrowerDetails", columns)
	
	var columns []*shim.Column
		col1 = shim.Column{Value: &shim.Column_String_{String_: row.Columns[0].GetString_()}}
		col2 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[1].GetString_()}}
		col3 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[2].GetString_()}}
		col4 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[3].GetString_()}}
		col5 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[4].GetString_()}}
		col6 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[5].GetString_()}}
		col7 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[6].GetString_()}}
		col8 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[7].GetString_()}}
		col9 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[8].GetString_()}}
		col10 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[9].GetString_()}}
		col11 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[10].GetString_()}}
		col12 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[11].GetString_()}}
		col13 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[12].GetString_()}}
		col14 := shim.Column{Value: &shim.Column_String_{String_: switchLenderId}}
		col15 := shim.Column{Value: &shim.Column_String_{String_: switchLenderName}}
		col16 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[15].GetString_()}}
		col17 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[16].GetString_()}}
		col18 := shim.Column{Value: &shim.Column_String_{String_: row.Columns[17].GetString_()}}
		col19:= shim.Column{Value: &shim.Column_String_{String_: row.Columns[18].GetString_()}}
		col20 := shim.Column{Value: &shim.Column_String_{String_: "No"}}
		col21 := shim.Column{Value: &shim.Column_String_{String_: ""}}
		col22 := shim.Column{Value: &shim.Column_String_{String_: ""}}
		
		columns = append(columns, &col1)
		columns = append(columns, &col2)
		columns = append(columns, &col3)
		columns = append(columns, &col4)
		columns = append(columns, &col5)
		columns = append(columns, &col6)
		columns = append(columns, &col7)
		columns = append(columns, &col8)
		columns = append(columns, &col9)
		columns = append(columns, &col10)
		columns = append(columns, &col11)
		columns = append(columns, &col12)
		columns = append(columns, &col13)
		columns = append(columns, &col14)
		columns = append(columns, &col15)
		columns = append(columns, &col16)
		columns = append(columns, &col17)
		columns = append(columns, &col18)
		columns = append(columns, &col19)
		columns = append(columns, &col20)
		columns = append(columns, &col21)
		columns = append(columns, &col22)
		
		row = shim.Row{Columns: columns}
		
	ok, err := stub.ReplaceRow("BorrowerDetails", row)
		if err != nil {
			return nil, fmt.Errorf("replaceRowTableOne operation failed. %s", err)
		}
		if !ok {
			return nil, errors.New("replaceRowTableOne operation failed. Row with given key does not exist")
		}

		
	 return nil,nil


}

// Checks the government data to verify whether the title information for this particular address is clean
func (t *MORTGAGE) verifyTitleInfo(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("verifyTitleInfo")
	


	
	address := args[0]
	
		
	var colum []shim.Column
	
	fmt.Println("Hello, World!!!")
	rows, err := stub.GetRows("TitleInfo", colum)
	fmt.Println("Hello, World!")
	fmt.Println(err)
	if err != nil {
		jsonResp := "{\"Error\":\"Failed getting the details of the address  " + address + "\"}"
		return nil, errors.New(jsonResp)
	}

	// GetRows returns empty message if key does not exist

	res2E := TitleInfo{}
	for row := range rows { 
	newApp:= new(TitleInfo)
	newApp.Name = row.Columns[0].GetString_()
	newApp.Address = row.Columns[1].GetString_()
	newApp.City = row.Columns[2].GetString_()
	newApp.State = row.Columns[3].GetString_()
	newApp.Zip = row.Columns[4].GetString_()
	newApp.Info = row.Columns[5].GetString_()
	if newApp.Address == address{
	fmt.Println("inside cond")
		 mapB, _ := json.Marshal(newApp)
		 return mapB, nil
		
		}	
		
		}
	
	
	
   mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))

	return mapB, nil
}

// Checks the government data to verify whether the title information for this particular address is clean
func (t *MORTGAGE) getInterestedUsers(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {

	fmt.Println("getInterestedUsers")
	


	
	
	
		
	// Get the row pertaining to this Email
	var columns []shim.Column	
	rows, err := stub.GetRows("BorrowerDetails", columns)
	if err != nil {
		return nil, fmt.Errorf("Failed to retrieve row")
	}
	res2E := []*BorrowerDetails{}
	for row := range rows {	
		newApp:= new(BorrowerDetails)
		newApp.UserID = row.Columns[0].GetString_()
		newApp.Gender = row.Columns[1].GetString_()
		newApp.FirstName = row.Columns[2].GetString_()
		newApp.LastName = row.Columns[3].GetString_()
		newApp.Dob = row.Columns[4].GetString_()
		newApp.Email = row.Columns[5].GetString_()
		newApp.Phone = row.Columns[6].GetString_()
		newApp.Address = row.Columns[7].GetString_()
		newApp.City = row.Columns[8].GetString_()
		newApp.Zip = row.Columns[9].GetString_()
		newApp.PropertyAddress = row.Columns[10].GetString_()
		newApp.PropertyCity = row.Columns[11].GetString_()
		newApp.PropertyZip = row.Columns[12].GetString_()
		newApp.LenderId = row.Columns[13].GetString_()
		newApp.LenderName = row.Columns[14].GetString_()
		newApp.ProductName = row.Columns[15].GetString_()
		newApp.LoanAmount = row.Columns[16].GetString_()
		newApp.InterestRate = row.Columns[17].GetString_()
		newApp.DocumentsSubmitted = row.Columns[18].GetString_()
		newApp.SwitchUserFlag = row.Columns[19].GetString_()
		newApp.SwitchLenderId = row.Columns[20].GetString_()
		newApp.SwitchLenderName = row.Columns[21].GetString_()
		
		if newApp.LenderId == "" {
		res2E=append(res2E,newApp)		
		}	
	}
    mapB, _ := json.Marshal(res2E)
    fmt.Println(string(mapB))
	
	return mapB, nil
}

// Invoke invokes the chaincode
func (t *MORTGAGE) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "registerLender" {
		t := MORTGAGE{}
		return t.registerLender(stub, args)	
	}else if function == "registerBorrower" {
		t := MORTGAGE{}
		return t.registerBorrower(stub, args)	
	}else if function == "switchLenders" {
		t := MORTGAGE{}
		return t.switchLenders(stub, args)	
	}else if function == "registerTitleInfo" {
		t := MORTGAGE{}
		return t.registerTitleInfo(stub, args)	
	}else if function == "registerUserDetails" {
		t := MORTGAGE{}
		return t.registerUserDetails(stub, args)	
	}else if function == "updateSwitchLender" {
		t := MORTGAGE{}
		return t.updateSwitchLender(stub, args)	
	}	

	return nil, errors.New("Invalid invoke function name.")

}

// query queries the chaincode
func (t *MORTGAGE) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	if function == "getProductRatesFromLender" { 
		t := MORTGAGE{}
		return t.getProductRatesFromLender(stub, args)
	}else if function == "getLenderDetails" {
		t := MORTGAGE{}
		return t.getLenderDetails(stub, args)	
	}else if function == "fetchBorrowerDetails" {
		t := MORTGAGE{}
		return t.fetchBorrowerDetails(stub, args)	
	}else if function == "getBorrower" {
		t := MORTGAGE{}
		return t.getBorrower(stub, args)	
	}else if function == "verifyTitleInfo" {
		t := MORTGAGE{}
		return t.verifyTitleInfo(stub, args)	
	}else if function == "getBorowersWithLenderId" {
		t := MORTGAGE{}
		return t.getBorowersWithLenderId(stub, args)	
	}else if function == "fetchUserDetails" {
		t := MORTGAGE{}
		return t.fetchUserDetails(stub, args)	
	}else if function == "getInterestedUsers" {
		t := MORTGAGE{}
		return t.getInterestedUsers(stub, args)	
	}	
	

	
	return nil, nil
}

func main() {
	
	err := shim.Start(new(MORTGAGE))
	if err != nil {
		fmt.Printf("Error starting MORTGAGE: %s", err)
	}
} 