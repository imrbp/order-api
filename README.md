# ASSIGNMENT 2

## Detail Database

# ADJUSTMENT

Ini branch Adjusment untuk menyesuaikan request dan response yang ada di classroom. 

Nama Database : orders_by  

Daftar Table:

- items
  - id -> untuk update tidak digunakan. maka items akan menambah. dan tidak ditampilkan di response
  - item_code 
  - description
  - quantity
  - order_id
- orders
  - id
  - order_at
  - customer_name

# Test

Testing yang sesuai dengan classroom ada di folder test/adjustment.http


# Package yang digunakan
- go-gin
- validator
- pgx
- viper
- gorm
- testify

