# golang-account
A Go app to process transactions which load money to prepaid credit cards

## Run Instructions
Clone repo into Go workspace `src/` folder set by `GOPATH`.  I.e. `$GOPATH/src/golang-account`

Install with command:
```
go install golang-account
```

Run with command:
```
$GOPATH/bin/golang-account < data/input.txt [> data/results.txt]
```

Run unit tests (from project root folder):
```
go test
```

Check test coverage:
```
go test -coverprofile cp.out
```

The program is developed and tested with Go 1.13.  It makes use of the standard library only; no 3rd party package installs are required.

## Solution
This section details notable points about the solution.

### Overall Design
The solution is broken down into several source files.  Each represents a struct and methods for operating on it.  Complex logic is broken down into separate functions for ease of comprehension and testability.  Each source file has an associated test file.

The main module is kept as minimal as possible.  It is only responsible for IO and calling the backend APIs.

An attempt to follow Go language conventions are followed.  A best effort was made by a non-Go developer but module design may show evidence of implementation with an OO mindset.  Nonetheless, use of pointers, Godoc function comments, and error handling are implemented in true Go fashion.

### Testing
Unit tests are widely used, with a Test-Driven Development approach often taken.  Test coverage is at 80% which is healthy.  In addition, testing is fairly robust with many different input configurations tested in functions containing complex logic - something not captured in the test coverage metric.  This ensures the code base is well guarded against regressions.

### Solution Correctness
The solution is validated to work with the provided `input.txt` and `output.txt` files.  However, due to the relatively complex nature of the program and the side effects associated with some of the function designs it is possible bugs are present that are only shown when operating at greater scale.

It is believed the provided `output.txt` file contains an error.  It has one less line of output than `input.txt`.  In addition, the app processes and outputs this additional line without error, matching the 1000 input lines.

## Future Considerations
Given more time here are areas to improve upon:

### Go Packages
The flat structure of the project allows for simple compilation and sharing between package and testing source files.  However a cleaner, more scalable approach would utilize Go packages.

### Threaded solution
This program is a single threaded app.  Given the computing-intensive nature of the problem, a faster & more scalable solution would utilize Go threads.  One potential approach is with seprate threads for the following tasks:
* Reading from stdin & deserializing to the Deposit structs
* Processing the transactions by executing core business logic, updating the account and deposit data structures, and returning the output object
* Serializing the output object and printing to stdout
The threads would be implemented as Goroutines with data passed in Channels.

### Funcitonal / System Testing
A simple script (written in say Bash or Python) could be written to run the app, gather output and perform validation by comparing against expected output.  Functional or system level testing is a critical tool which complements unit testing and in this case is very simple to implement.

### Enhance Unit Test Code
Unit test coverage is reasonable but much of the code is duplicated, e.g. in [account_limits_test.go](account_limits_test.go#L25).  A better approach is to follow the design in [account_test.go](account_test.go#L7) which defines a test struct containing test data, and looping over it executing the same test code each time.  It is more readable and much easier to modify.  _The easier it is for developers to add tests the more tests they'll write!_

# Original Problem
In finance, it's common for accounts to have so-called "velocity limits". In this task, you'll write a program that accepts or declines attempts to load funds into customers' accounts in real-time.

Each attempt to load funds will come as a single-line JSON payload, structured as follows:

```json
{
  "id": "1234",
  "customer_id": "1234",
  "load_amount": "$123.45",
  "time": "2018-01-01T00:00:00Z"
}
```

Each customer is subject to three limits:

- A maximum of $5,000 can be loaded per day
- A maximum of $20,000 can be loaded per week
- A maximum of 3 loads can be performed per day, regardless of amount

As such, a user attempting to load $3,000 twice in one day would be declined on the second attempt, as would a user attempting to load $400 four times in a day.

For each load attempt, you should return a JSON response indicating whether the fund load was accepted based on the user's activity, with the structure:

```json
{ "id": "1234", "customer_id": "1234", "accepted": true }
```

You can assume that the input arrives in ascending chronological order and that if a load ID is observed more than once for a particular user, all but the first instance can be ignored. Each day is considered to end at midnight UTC, and weeks start on Monday (i.e. one second after 23:59:59 on Sunday).

Your program should process lines from `input.txt` and return output in the format specified above, either to standard output or a file. Expected output given our input data can be found in `output.txt`.

You're welcome to write your program in a general-purpose language of your choosing, but as we use Go on the back-end and TypeScript on the front-end, we do have a preference towards solutions written in Go (back-end) and TypeScript (front-end).

We value well-structured, self-documenting code with sensible test coverage. Descriptive function and variable names are appreciated, as is isolating your business logic from the rest of your code.
