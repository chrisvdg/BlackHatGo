# Credential harvester

The goal of this exercise is to harvest the login attemps from a roundcube instance

The Go server mymics the login screen of a roundcube server and logs a login attempt to a file.

To run the roundcube server to get the source files from execute the following script:
```sh
./run_roundcube
```

Run the harvester server
```sh
go run main.go
```
