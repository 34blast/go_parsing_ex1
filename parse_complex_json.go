package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"reflect"
)

const JSON_PRODUCT1 = `{"type":"product", "field2": "field2" , "field3": "field3" , "field4": "field4" }`
const JSON_PRODUCT2 = `{"type":"product", "productId":"rms_gtin1", "price":44.44, "accountNumber":123456 }`
const JSON_PRODUCT3 = `{"type":"product", "productId":"rms_gtin1", "mfgPlants":["plant1", "plant2"]}`
const JSON_PRODUCT4 = `{"type":"product", "productId":"rms_gtin1", "price": 44.44, "lotNumber": 44, "serialNumber": "serialNumber1", "mfgPlants":["plant1", "plant2"], "drugInfo":{"drugName":"aspirin", "drugDescription":"aspirin"}, "manufacturerInfo":{"mfgName":"CompanyA", "gtin":"123456789"}}`
const JSON_PRODUCT5 = `{"type":"product", "productId":"rms_gtin5", "price": 44.44, "lotNumber": 44, "serialNumber": "serialNumber1", "mfgPlants":["plant1", "plant2"], "drugInfo":{"drugName":"aspirin", "drugDescription":"aspirin"}, "manufacturerInfo":{"mfgName":"CompanyA", "gtin":"123456789"}}`

type Container map[string]json.RawMessage
type FloatStruct struct {
	Price float64
}

func main() {
	var myInterface1 interface{}
	err1 := json.Unmarshal([]byte(JSON_PRODUCT1), &myInterface1)
	if err1 != nil {
		log.Fatal(err1)
		os.Exit(-1)
	}
	fmt.Println()
	fmt.Println("JSON_PRODUCT1")
	fmt.Printf("%s\n", myInterface1)
	fmt.Println()

	var myInterface2 interface{}
	err2 := json.Unmarshal([]byte(JSON_PRODUCT2), &myInterface2)
	if err2 != nil {
		log.Fatal("could not unmarshal interface 2,  exiting, err = ", err2)
		os.Exit(-1)
	}
	fmt.Println()
	fmt.Println("JSON_PRODUCT2")
	fmt.Printf("%s\n", myInterface2)

	fmt.Println("  --------- change price, account, and add a new element addedElement --------- ")

	// Set a new price , change account number and add new element addedElement
	i2Map := myInterface2.(map[string]interface{})
	i2Map["price"] = 99.99
	i2Map["accountNumber"] = 123456789
	if i2Map["addedElement"] == nil {
		i2Map["addedElement"] = "added element"
	}
	fmt.Printf("%s\n", myInterface2)
	fmt.Println()

	var myInterface3 interface{}
	err3 := json.Unmarshal([]byte(JSON_PRODUCT3), &myInterface3)
	if err3 != nil {
		log.Fatal(err3)
	}
	fmt.Println()
	fmt.Println("JSON_PRODUCT3")
	fmt.Printf("%s\n", myInterface3)
	bytes3, err3m := json.Marshal(myInterface3)
	if err3m != nil {
		log.Fatal(err3m)
		os.Exit(0)
	}
	fmt.Println("as a String")
	fmt.Println(string(bytes3))
	fmt.Println()

	var myInterface4 interface{}
	err4 := json.Unmarshal([]byte(JSON_PRODUCT4), &myInterface4)
	if err4 != nil {
		log.Fatal(err4)
		os.Exit(0)
	}
	fmt.Println()
	fmt.Println("JSON_PRODUCT4")
	fmt.Printf("%s\n", myInterface4)
	bytes4, err4m := json.Marshal(myInterface4)
	if err4m != nil {
		log.Fatal(err4m)
		os.Exit(0)
	}
	fmt.Println("as a String")
	fmt.Println(string(bytes4))

	prodMap := reflect.ValueOf(myInterface4)
	if prodMap.Kind() != reflect.Map {
		log.Fatal("Must be Map to continue")
		os.Exit(-1)
	}
	fmt.Println("  --------- print out map key value pairs --------- ")

	for _, key := range prodMap.MapKeys() {
		strct := prodMap.MapIndex(key)
		value := strct.Elem()

		if value.Kind() == reflect.String {
			//fmt.Println("      value is of type String")
		} else if value.Kind() == reflect.Float64 {
			//fmt.Println("      value is of type Float64")
		} else if value.Kind() == reflect.Slice {
			//fmt.Println("      value is of type Slice")
		} else if value.Kind() == reflect.Map {
			//fmt.Println("      value is of type Map")
		} else {
			fmt.Println("      UNKNOWN XXXXXXXXXXXXXXX value is of type ", value.Kind())
		}
		fmt.Println(key.Interface(), value)
	}

	var container5 Container

	err5 := json.Unmarshal([]byte(JSON_PRODUCT5), &container5)
	if err5 != nil {
		log.Fatal(err5)
		os.Exit(0)
	}
	fmt.Println()
	fmt.Println("JSON_PRODUCT5")
	fmt.Printf("%s\n", container5)

	fmt.Println("  --------- change price, drugName, drugDescription --------- ")

	// we have to unmarshall it since it is RawMessage Json, change it then marshall it back up
	// then set it

	var drugInterface interface{}
	drugErr := json.Unmarshal(container5["drugInfo"], &drugInterface)
	if drugErr != nil {
		log.Fatal(drugErr)
		os.Exit(-1)
	}
	drugMap := drugInterface.(map[string]interface{})
	fmt.Println("unmarshalled drugName = ", drugMap["drugName"])
	drugMap["drugDescription"] = "Propecia"
	drugMap["drugName"] = "Propecia"
	drugBytes, drugErr := json.Marshal(drugMap)
	if drugErr != nil {
		log.Fatal(drugErr)
		os.Exit(-1)
	}
	container5["drugInfo"] = drugBytes

	var unmarshalledPrice float64
	priceErr := json.Unmarshal(container5["price"], &unmarshalledPrice)
	if priceErr != nil {
		log.Fatal(priceErr)
		os.Exit(-1)
	}
	unmarshalledPrice = 88.88
	priceBytes, bytesErr := json.Marshal(unmarshalledPrice)
	if bytesErr != nil {
		log.Fatal(bytesErr)
		os.Exit(-1)
	}
	container5["price"] = priceBytes

	fmt.Printf("%s\n", container5)
	fmt.Println()

}

func IsInstanceOf(objectPtr, typePtr interface{}) bool {
	return reflect.TypeOf(objectPtr) == reflect.TypeOf(typePtr)
}
