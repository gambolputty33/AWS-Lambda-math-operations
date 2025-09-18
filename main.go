/*

sample JSON input data

{"number1": 5.234123, "number2": 7.89723489}

*/

package main

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/aws/aws-lambda-go/lambda"
)

// Input struct
type Input_struct struct {
	Number1 float64 `json:"number1"`
	Number2 float64 `json:"number2"`
}

// Output struct
type Output_struct struct {
	Sum        float64 `json:"sum"`
	Difference float64 `json:"difference"`
	Product    float64 `json:"product"`
	Quotient   float64 `json:"quotient"`
}

// Handler function for AWS Lambda
func handler(ctx context.Context, raw json.RawMessage) (json.RawMessage, error) {
	var input Input_struct

	// Decode the raw JSON into our struct
	if err := json.Unmarshal(raw, &input); err != nil {
		return nil, fmt.Errorf("failed to parse input: %w", err)
	}

	// Perform math operations on input data
	sum := input.Number1 + input.Number2
	difference := input.Number1 - input.Number2
	product := input.Number1 * input.Number2
	quotient := input.Number1 / input.Number2

	output := Output_struct{Sum: sum, Difference: difference, Product: product, Quotient: quotient}

	// Encode the result back into JSON
	resultBytes, err := json.Marshal(output)
	if err != nil {
		return nil, fmt.Errorf("failed to marshal result: %w", err)
	}

	return resultBytes, nil

}

func main() {
	// Start the Lambda handler
	lambda.Start(handler)
}
