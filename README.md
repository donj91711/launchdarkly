Welcome to the LauchDarkly Coding Test Project - completed by Don Jackson

Instructions:
1) Download the repository from https://github.com/launchdarkly/be-coding-test-Don-Jackson

2) Open a terminal and navigate to the project directory.

3) Install the Gorilla Mux router library by running:
go get github.com/gorilla/mux

4) Start the application by running:
go run main.go
After running this command, you should see the application start and download test results into memory. Trace messages will be displayed in the terminal.

5) Test the application using Postman. The application is reachable on localhost. Here are the available endpoints:
a) Retrieve a list of students:
   - URL: http://localhost/students

b) Retrieve a specific student by ID (replace {id} with an actual student ID):
   - URL: http://localhost/students/{id}

c) Retrieve a list of exams:
   - URL: http://localhost/exams

d) Retrieve specific exam details by number (replace {number} with an actual exam number):
   - URL: http://localhost/exams/{number}

6) Available tests can be completed by running:
go test ./... 


Please let me know if you have any questions or encounter any issues during setup or testing.
