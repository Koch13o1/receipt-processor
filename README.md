# Receipt Processor

A simple Go-based web service for processing receipts and calculating points based on a set of rules.

## Overview

This application accepts receipt data via a POST request, validates and converts the input, computes points based on several rules, and stores the receipt with precomputed points. A GET endpoint retrieves the points for a saved receipt using its unique ID.

## Technologies Used

This application used:
1. Go: Primary language of backend development
2. Gin: Used as our HTTP web framework to handle routing


## Getting Started

### Prerequisites

- [Go 1.23+](https://golang.org/dl/)


### Installation

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/Koch13o1/receipt-processor.git
   cd receipt-processor
   ```

2. **Installing Dependencies:**
     ```sh
     go mod tidy
      ```


### API Endpoints
- POST /receipts/process
- GET /receipts/{id}/points
(Available at http://localhost:8080)


### Testing
Unit tests are provided in the tests/ directory. Run the tests using:
```sh
go test -v ./...
```


### Architecture & Design Choices
- Points calculation is done at the time of receipt storage instead of retrieval to improve performance.
- The application is designed to handle receipts efficiently using in-memory storage for quick access.
- It has the controllers, services, storage, and tests separate, for scalability and reusability reasons.


### Challenges Faced
- Ensuring accurate and efficient receipt parsing while handling potential formatting inconsistencies.


### Future Scope
- Persist receipts in a database instead of in-memory storage to allow scalability.
- Add authentication and user-based receipt management.
- Also, in the receipt.go, I could add a validation logic to check if sum of all the prices of items equals the total as we are given both the values.


###


