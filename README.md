# ASSIGNMENT 2

## Detail Database

Nama Database : orders_by  

Daftar Table:

- items
  - item_id
  - item_code
  - description
  - quantity
  - order_id
- orders
    - order_id
    - order_at
    - customer_name

## a. Create Order
Path: /orders 

Method: Post 

Request Body :

## b. Get Orders
Path: /orders 

Method: GET

## c. Update Order
Path: /orders/:orderId 

Method:  PUT

Request Body:

## d. Delete Order 
Path: /orders/:orderId
 
Method: DELETE


# Package yang digunakan
- go-gin
- validator
- pgx
- viper
- gorm
- testify


# TODO
- Error Handling
- Unit Testing
- When Updating Order, If user want to add new Items. its possible?