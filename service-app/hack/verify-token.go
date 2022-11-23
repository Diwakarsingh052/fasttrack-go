package main

import (
	"fmt"
	"github.com/golang-jwt/jwt/v4"
	"log"
	"os"
)

var tokenStr = `eyJhbGciOiJSUzI1NiIsInR5cCI6IkpXVCJ9.eyJpc3MiOiJhcGkgcHJvamVjdCIsInN1YiI6IjEwMSIsImV4cCI6MTY2OTE5ODY4NywiaWF0IjoxNjY5MTk1Njg3LCJyb2xlcyI6WyJVU0VSIl19.UJPqDGVROIvNabbpPohkchzMPOJrHabVqst7hdt7lXgl1zMNrKAoLjdmAfLzpdfhkeMTjGnpee84Qd10lOjJn5KZTebaH5gEavQXTAa1R-6Z82ntncE1Z9cfMxFsMd7XXb_gWZv_XMB-iQu9N_NB-9g-uY625MipQ3b0Bxzf1A_IulRIIxWaAl6pfbIjVK0Q9M9rVuRAXKWL7iFE72JDRzQaJL3bKyPAXGsYW9jOTI10AHkR_ituJLmJzHDxDxa3sjFxz5wMLCft0z9TIGvssu3nBpuIR49895qNUZredYzDLlcWLv4pswF2o6DejOyCyLJL3VgKbDefQRjUwpsvkg`

func main() {

	type claims struct {
		jwt.RegisteredClaims
		Roles []string `json:"roles"`
	}
	PrivatePEM, err := os.ReadFile("private.pem")
	if err != nil {
		log.Fatalln("not able to read pem file")
	}

	privateKey, err := jwt.ParseRSAPrivateKeyFromPEM(PrivatePEM)

	if err != nil {
		log.Fatalln(err)
	}

	var c claims

	token, err := jwt.ParseWithClaims(tokenStr, &c, func(token *jwt.Token) (interface{}, error) {
		return privateKey.Public(), nil
	})

	if err != nil {
		fmt.Println("parsing token", err)
		return
	}

	if !token.Valid {
		fmt.Println("invalid token")
		return
	}
	fmt.Println(c)

}
