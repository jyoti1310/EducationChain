/*
Copyright IBM Corp 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

import (
	"errors"
	"fmt"
	"strconv"
	"encoding/json"
	
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type DegreesCompleted struct{
		GovtID string `json:"GovtID"`
        InstituteID int `json:"InstituteID"`
        DegreeID int `json:"DegreeID"`
        DegreeName string `json:"DegreeName"`
        PassingYear string `json:"PassingYear"`
		Percentage int `json:"Percentage"`
		RollNo int `json:"Rollno"`
		Grade string `json:"Grade"`
		Type string `json:"Type"`
}

type StudentInformation struct{
        GovtID string `json:"GovtID"`
        Age int `json:"Age"`
        StudentName string `json:"StudentName"`
        DOB string `json:"DOB"`
}

type AppliedDegreeRequests struct{
		GovtID string `json:"GovtID"`
        AppliedInstituteID int `json:"InstituteID"`
        DegreeID int `json:"DegreeID"`
        DegreeName string `json:"DegreeName"`
		PreRequisiteDegree string `json:"PreRequisiteDegree"`
		CompletedInstituteID int `json:"InstituteID"`
        Approved int `json:"Percentage"`
}

//var employeeLogBog map[string]SKATEmployee
type CompletedDegreesRepository struct {
	CompletedDegrees []DegreesCompleted `json:"CompletedDegrees"`
}
/*type SKATEmployeeRepository struct{
	EmployeeList []SKATEmployee `json:"employee_list"`
}*/

/*type searchedEmployees struct{
	SearchedEmployeeList []SKATEmployee `json:"searched_employee_list"`
}*/

func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init resets all the things
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting 1")
	}

	err := stub.PutState("hello_Block", []byte(args[0]))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Invoke isur entry point to invoke a chaincode function
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "init" {
		return t.Init(stub, "init", args)
	} else if function == "write" {
		return t.write(stub, args)
	} else if  function == "addToDegreesCompleted" {
		return t.addNewCompletedDegree(stub, args)
	} /*else if  function == "updateLogBog" {
		return t.updateSKATEmployee(stub, args)
	}*/
	fmt.Println("invoke did not find func: " + function)

	return nil, errors.New("Received unknown function invocation: " + function)
}

// Query is our entry point for queries
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Println("query is running " + function)

	// Handle different functions
	if function == "read" { //read a variable
		return t.read(stub, args)
	} /*else if function == "searchLogBog" {
		return t.searchSKATEmployee(stub,args)
	}/*else if function == "searchLogBog" {
		return t.searchSKATEmployee(stub,args)
	}*/
	fmt.Println("query did not find func: " + function)

	return nil, errors.New("Received unknown function query: " + function)
}

// write - invoke function to write key/value pair
func (t *SimpleChaincode) write(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, value string
	var err error
	fmt.Println("running write()")

	if len(args) != 2 {
		return nil, errors.New("Incorrect number of arguments. Expecting 2. name of the key and value to set")
	}

	key = args[0] //rename for funsies
	value = args[1]
	err = stub.PutState(key, []byte(value)) //write the variable into the chaincode state
	if err != nil {
		return nil, err
	}
	return nil, nil
}

// read - query function to read key/value pair
func (t *SimpleChaincode) read(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var key, jsonResp string
	var err error

	if len(args) != 1 {
		return nil, errors.New("Incorrect number of arguments. Expecting name of the key to query")
	}

	key = args[0]
	valAsbytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + key + "\"}"
		return nil, errors.New(jsonResp)
	}

	return valAsbytes, nil
}

// ============================================================================================================================
// Init Employee - create a new Employee, store into chaincode state
// ============================================================================================================================
func (t *SimpleChaincode) addNewCompletedDegree(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {
	var err error
	var key string
	//,jsonResp string
	
	//   0       1       2     3
	// "asdf", "blue", "35", "bob"
	if len(args) != 8 {
		return nil, errors.New("Incorrect number of arguments. Expecting 7")
	}

	//input sanitation
	fmt.Println("- start adding new Degree")
	/*if len(args[0]) <= 0 {
		return nil, errors.New("1st argument must be a non-empty string")
	}
	if len(args[1]) <= 0 {
		return nil, errors.New("2nd argument must be a non-empty string")
	}
	if len(args[2]) <= 0 {
		return nil, errors.New("3rd argument must be a non-empty string")
	}
	if len(args[3]) <= 0 {
		return nil, errors.New("4th argument must be a non-empty string")
	}*/
	NewDegreeCompleted := DegreesCompleted{}
	NewDegreeCompleted.GovtID = args[0]
	
	NewDegreeCompleted.InstituteID, err =strconv.Atoi(args[1])
	
	if err != nil {
		return nil, errors.New("InstituteID must be a numeric string")
	}
	NewDegreeCompleted.DegreeID, err = strconv.Atoi(args[2])
	
	if err != nil {
		return nil, errors.New("DegreeID must be a numeric string")
	}
	NewDegreeCompleted.DegreeName=args[3]
	NewDegreeCompleted.PassingYear=args[4]
	NewDegreeCompleted.Percentage, err=strconv.Atoi(args[5])
	
	if err != nil {
		return nil, errors.New("Percentage must be a numeric string")
	}
	NewDegreeCompleted.RollNo, err=strconv.Atoi(args[6])

	if err != nil {
		return nil, errors.New("Roll No must be a numeric string")
	}
	NewDegreeCompleted.Grade=args[7]
	NewDegreeCompleted.Type=args[8]

	/*if len(args) == 6 {
	Employee.Comment = args[5]
  	}*/
	fmt.Println("adding Degree @ " + NewDegreeCompleted.GovtID + ", " + strconv.Itoa(NewDegreeCompleted.InstituteID));
	fmt.Println("- end add Degree 1")
	jsonAsBytes, _ := json.Marshal(NewDegreeCompleted)

	if err != nil {
		return jsonAsBytes, err
	}
	
	key = NewDegreeCompleted.GovtID + "_2"
	t.appendtoCompletedDegreeRepository(stub,key,NewDegreeCompleted)
	fmt.Println("- end add Degree 2")
	return jsonAsBytes, nil
}

//==================================================================================================================================
//Append to EmployeeRepository
//===================================================================================================================================

func (t *SimpleChaincode) appendtoCompletedDegreeRepository(stub shim.ChaincodeStubInterface,  key string,newDegree DegreesCompleted) (bool, error){

var jsonResp string 
repositoryJsonAsBytes, err := stub.GetState(key)
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + "SKATEmployeeRepository" + "\"}"
		return false, errors.New(jsonResp)
	}
	var degreeRepository CompletedDegreesRepository
	json.Unmarshal(repositoryJsonAsBytes, &degreeRepository)	

	degreeRepository.CompletedDegrees = append(degreeRepository.CompletedDegrees,newDegree)
	//update Employee Repository
	updatedRepositoryJsonAsBytes, _  := json.Marshal(degreeRepository)
	err = stub.PutState(key, updatedRepositoryJsonAsBytes)	//store employee with id as key
	if err != nil {
		return false, err
	}		
	return true, nil
}


	
