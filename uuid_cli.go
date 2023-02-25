//
//  UUID Command Line Tool
//
//  go run uuid_cli.go 
//  OgTe8PZjQdSYTfFEjT8Gdg==
//
//  go run uuid_cli.go OgTe8PZjQdSYTfFEjT8Gdg==
//  3a04def0-f663-41d4-984d-f1448d3f0676
//
//  go run uuid_cli.go 3a04def0-f663-41d4-984d-f1448d3f0676
//  OgTe8PZjQdSYTfFEjT8Gdg==
//

package main

import (
	"encoding/base64"
	"fmt"
	"os"

	"github.com/google/uuid"
)

func convertUuidToUrlEncoding(input uuid.UUID) string {
	return base64.URLEncoding.EncodeToString(input[:])
}

func convertUrlEncodingToUuid(input string) (uuid.UUID, error) {
	decodedBytes := make([]byte, base64.URLEncoding.DecodedLen(len(input)))
	byteCount, err := base64.URLEncoding.Decode(decodedBytes, []byte(input))
	if err != nil {
		return uuid.Nil, err
	}
	decodedBytes = decodedBytes[:byteCount]
	return uuid.FromBytes(decodedBytes)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println(convertUuidToUrlEncoding(uuid.New()))
		return
	}

	input := os.Args[1]
	if len(input) < 24 {
		fmt.Println("Error: expected input to be a uuid string or base64.URLEncoding encoded 16 byte string")
	} else if len(input) == 24 {
		output, err := convertUrlEncodingToUuid(input)
		if err != nil {
			fmt.Println("Error creating uuid from base64 string:", err)
			return
		}
		fmt.Println(output.String())
	} else {
		parsedUuid, err := uuid.Parse(input)
		if err != nil {
			fmt.Println("Error parsing uuid from string:", err)
			return
		}
		fmt.Println(convertUuidToUrlEncoding(parsedUuid))
	}
}
