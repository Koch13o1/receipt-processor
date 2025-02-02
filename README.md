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

- [Go 1.19+](https://golang.org/dl/)

### Installation

1. **Clone the Repository:**
   ```sh
   git clone https://github.com/YOUR_USERNAME/receipt-processor.git
   cd receipt-processor

2. **Installing Dependencies:**
  ```sh
  go mod tidy

### Running the Service

