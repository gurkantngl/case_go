q1.go

Explanation
The program defines a slice of strings input_strings containing various strings. It then uses the sort.Slice function to sort the strings based on two criteria:

The number of occurrences of the letter "a" in each string.
The length of the strings.
The sorting is done in descending order of the number of "a" occurrences, followed by descending order of string length.

Example Output
After running the program, you should see the sorted strings printed to the console.


q2.go

Explanation
The program defines a recursive function recursive that takes an integer n as input. If n is equal to 1, the function returns without performing any further operations. Otherwise, it recursively calls itself with n/2 as the new input and then prints the original value of n.

In the main function, the recursive function is called with the input 9, which triggers the recursive calls and prints the values of 9, 4, 2, and 1 in that order.

Example Output
After running the program, you should see the following output: 2 4 9

q3.go

Explanation
The program defines a function getMostRepeatedString that takes a slice of strings inputList as input. It then iterates through the input list and maintains a count of each string's occurrences in a map counts. The function keeps track of the string with the maximum count maxCount and updates the mostRepeatedString variable accordingly. Finally, it returns the mostRepeatedString.

In the main function, the getMostRepeatedString function is called with the input ["apple", "pie", "apple", "red", "red", "red"], which returns "red" as the most repeated string, and this result is printed to the console.

Example Output
After running the program, you should see the following output: red


q4.go

The API server will start running on http://localhost:8000.

Endpoints
GET /users: Get all users.
GET /users/{id}: Get a user by ID.
POST /users: Create a new user.
PUT /users/{id}: Update a user by ID.
DELETE /users/{id}: Delete a user by ID.
Usage
You can use tools like cURL or Postman to interact with the API endpoints. Here are some examples:

Get all users:

curl http://localhost:8000/users


Get a user by ID:

curl http://localhost:8000/users/1


Create a new user:

curl -X POST -H "Content-Type: application/json" -d '{"username":"newuser","email":"newuser@example.com"}' http://localhost:8000/users

Update a user by ID:

curl -X PUT -H "Content-Type: application/json" -d '{"username":"updateduser","email":"updateduser@example.com"}' http://localhost:8000/users/1
Delete a user by ID:

curl -X DELETE http://localhost:8000/users/1


Database
The API uses an SQLite database (users.db) to store user data. The database file will be created automatically if it does not exist.

Dependencies
The API uses the following external packages:

github.com/gorilla/mux for routing HTTP requests.
github.com/mattn/go-sqlite3 for interacting with the SQLite database.


User-Management-app

Start the development server

npm run dev
