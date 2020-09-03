/*
 * OpenAPI Petstore
 *
 * This spec is mainly for testing Petstore server and contains fake endpoints, models. Please do not use this for any other purpose. Special characters: \" \\
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package petstore
// Cat struct for Cat
type Cat struct {
	ClassName string `json:"className"`
	Color string `json:"color,omitempty"`
	Declawed bool `json:"declawed,omitempty"`
}
