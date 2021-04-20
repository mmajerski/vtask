### vtask - Simple Go CLI program

The program is using [Cobra](https://github.com/spf13/cobra) library to help execute and test commands.

##

**About**

```
vtask is a simple CLI.
  Usage:
  vtask [command]

  Available Commands:
    version             Prints current version of the CLI
    help                Views more information about this program
    run [-f | --file]   Starts the HTTP web server and serves HTML file provided as an argument
```

**How to install and run**

Run every command from root directory.

```
$ git clone https://github.com/userq11/vtask.git
$ cd vtask
$ go run main.go run --file sampledata/index.html
Server with file has started on http://localhost:9090 (^C to exit)
```

You can just run straightforward

```
$ go run main.go run --file sampledata/index.html
```

Or you can build and then run from binary file

```
$ go build .
$ ./vtask run --file sampledata/index.html
Server with file has started on http://localhost:9090 (^C to exit)
```

Or you can first install and then run commands

```
$ go install .
$ vtask run --file sampledata/index.html
```

To run test:

```
$ go test ./...
?       github.com/userq11/vtask        [no test files]
ok      github.com/userq11/vtask/cmd    0.002s
```

**Example**

```
$ vtask version
Current version: v0.1

$ vtask help
Usage:
vtask [command]

Available Commands:
        version                   Prints current version of the CLI
        help                      Views more information about this program
        run [-f | --file]         Starts the HTTP web server and serves HTML file provided as an argument

$ vtask run
Please check help and provide a HTML file as an argument with --file flag

$ vtask run --file notexist.html
Cannot read provided file(file might not exist)

$ vtask run --file sampledata/index.html
Server with file has started on http://localhost:9090 (^C to exit)
```
