# API Documentation

A simple Go application that provides an API for managing a basket of items.

## Installation

1. Clone the repository.
2. Install the required dependencies.
3. Build the application.

## Usage

1. Run the application.
2. Access the API endpoints to manage the basket items.

## Note

This is a simplified example for learning purposes.

## Usage

### Adding Items to the Basket

To add items to the basket, make a `POST` request to `/api/basket` endpoint with a JSON payload containing the items. The payload should have the following structure:

```json
{
  "items": [
    {
      "id": 1,
      "name": "Item 1",
      "price": 10,
      "quantity": 2
    },
    {
      "id": 2,
      "name": "Item 2",
      "price": 15,
      "quantity": 1
    }
  ]
}

`## Usage

### Adding Items to the Basket

To add items to the basket, make a `POST` request to `/api/basket` endpoint with a JSON payload containing the items. The payload should have the following structure:

```json
{
  "items": [
    {
      "id": 1,
      "name": "Item 1",
      "price": 10,
      "quantity": 2
    },
    {
      "id": 2,
      "name": "Item 2",
      "price": 15,
      "quantity": 1
    }
  ]
} `

For example, using cURL:

`curl -X POST -H "Content-Type: application/json" -d '{"items":[{"id":1,"name":"Item 1","price":10,"quantity":2},{"id":2,"name":"Item 2","price":15,"quantity":1}]}' http://localhost:8000/api/basket`

### Retrieving Items from the Basket

To retrieve all items from the basket, make a `GET` request to `/api/basket/all` endpoint. You can also apply filters by adding query parameters:

-   `lower`: Filter items with a price greater than this value.
-   `upper`: Filter items with a price lower than this value.
-   `name`: Filter items with a name containing the specified value.

For example, using cURL:


`curl -X GET "http://localhost:8000/api/basket/all?lower=0&upper=20&name=item"`

Please note that this is a simplified example, and in a production environment, you might need to handle authentication, error handling, and other security measures.