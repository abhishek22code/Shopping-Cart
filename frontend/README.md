# Shopping Cart – Fullstack Assignment

This project is a fullstack “Shopping Cart” application built using Go (Gin, GORM) for the backend and React for the frontend.  
The project follows the requirements given in the assignment PDF.



## Features Implemented

User login using token-based authentication  
View list of store items  
Add items to cart  
View current cart  
Checkout (cart → order)  
View order history  
Logout  
Toast notifications for feedback messages  
SQLite database (no external DB needed)  
Code hosted on GitHub  
Postman Collection included  
README with setup & API usage instructions included ✔ (this file)



## 1. Project Structure

```text
shopping-cart/
  backend/
    main.go
    models.go
    user_handlers.go
    item_handlers.go
    cart_handlers.go
    order_handlers.go
    auth_helpers.go
    go.mod
    go.sum

  frontend/
    package.json
    public/
      index.html
      favicon.ico (or custom logo)
    src/
      App.js
      App.css
      api.js
      index.js
      pages/
        LoginPage.jsx
        ItemsPage.jsx

## How to Run Backend (Go API)

Open terminal  
Move into backend folder:
cd backend


Install dependencies:
go mod tidy


Start server:
go run .


Backend runs at:
http://localhost:8080


Test API in browser/Postman:
http://localhost:8080/ping

Response:
{"message":"pong"}

## 💻 How to Run Frontend (React App)

Open another terminal  
Move into frontend folder:
cd frontend

Install required node modules:
npm install

Start the React development server:
npm start

Frontend runs at:
http://localhost:3000

---

## Authentication Flow

- After successful login, backend returns a token
- Frontend stores it in `localStorage`
- Token must be attached to all authorized requests:

X-Auth-Token: <token_here>

Only one active token is maintained per user.

---

## Postman Collection

File included:
shopping-cart.postman_collection.json


## How to Interact with the API

There are two ways:

Using Postman → for developer testing  
Using React UI → for normal user usage

---

### Using Postman

Import this file:
shopping-cart.postman_collection.json

Then follow:

#### Create user → /users
json
{
  "username": "abhishek",
  "password": "1234"
}

#### Login → /users/login  
Copy the returned "token"
Example response:
json
{
  "token": "random-generated-token"
}


#### Add token to requests
Header:
X-Auth-Token: <token_here>


#### Add items → /items  
Example:
json
{
  "name": "Laptop",
  "status": "available"
}


Add more items for testing.

#### Add to cart → /carts
json
{
  "item_id": 1
}


#### View cart → /carts  
Token required 

####  Checkout → /orders
json
{
  "cart_id": 1
}


#### View order history → /orders  
Token required 

---

### Using the React UI

Enter username + password (must be created via Postman first)  
Add items  
Checkout  
See success messages via toast  
Check order history  
Logout anytime  

---

## Database

SQLite file automatically created:
backend/shopping_cart.db

No database setup needed
Auto migration done via GORM

---

##  Author

Abhishek Kumar Sahani 
Student — GCET CSE IVth Year