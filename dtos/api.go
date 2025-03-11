package dtos

// Example 1:
// {
//  "status": "error",
//  "code": 404,
//  "message": "User not found.",
//  "errors": [
//    { "field": "id", "message": "No user exists with the provided ID." }
//  ]
// }

// Example 2:
// {
//  "status": "success",
//  "code": 200,
//  "message": "Fetched data successfully.",
//  "data": [ /* Items */ ],
//  "meta": {
//    "page": 1,
//    "total_pages": 5,
//    "total_items": 100
//  },
//  "errors": null
// }

// APIResponse represents the structure of an API response
type APIResponse struct {
	Status  string     `json:"status"`            // success or error
	Code    int        `json:"code"`              // HTTP status code
	Message *string    `json:"message,omitempty"` // Descriptive message
	Data    any        `json:"data,omitempty"`    // Response data (any type)
	Errors  []APIError `json:"errors,omitempty"`  // List of errors, optional
	Meta    any        `json:"meta,omitempty"`    // Metadata, optional
}

// APIError represents an individual error in the response
type APIError struct {
	Field   string `json:"field"`   // Field associated with the error (e.g., "id")
	Message string `json:"message"` // Error message
}
