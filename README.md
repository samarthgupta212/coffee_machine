Coffee Machine
Code takes input.json as input file and parses it and serves the beverages in parallel corresponding to outlets defined in file.
I have also created input1.json file to test the refilling functionality. It also serves beverages after refilling is done in parallel.

Local Setup
Prerequisites 
Go 1.13 or above is installed

1. Clone the repository
2. Run go build command. One file with name coffee_machine.exe in case of windows.
   coffee_machine in case of mac/linux machine should be created.
3. Then run ./coffee_machine or coffee_machine.exe to make the code run.
4. To run test cases run go test -v ./...
