package main

import (
	"fmt"

	"github.com/dannov91/multitype/common"
	"github.com/dannov91/multitype/models"
)

func main() {
	//----- with RegularUser model type
	rUser := models.RegularUser{Name: "Ralph"}

	name, err := common.UserName(rUser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name ==> %v\n", name)

	// SHOULD ALWAYS BE POINTER
	err = common.SetUserName(&rUser, "Tom")
	if err != nil {
		panic(err)
	}

	fmt.Printf("rUser.Name ==> %v\n", rUser.Name)

	// ---- with PrivilegedUser model type
	pUser := models.RegularUser{Name: "Goth"}

	name, err = common.UserName(pUser)
	if err != nil {
		panic(err)
	}

	fmt.Printf("name ==> %v\n", name)

	// SHOULD ALWAYS BE POINTER
	err = common.SetUserName(&pUser, "Seth")
	if err != nil {
		panic(err)
	}

	fmt.Printf("pUser.Name ==> %v\n", pUser.Name)
}
