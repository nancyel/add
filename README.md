# Add

### Description
Add is a simple program written in Go that allows you to capture memorable lines in an instant. Use it to quickly jot down your moments of inspiration without the hassle of opening a text editor.

### Getting started

To get started with `add`, follow these simple steps:


1. Install

```bashscript
$ git clone https://github.com/nancyel/add.git
$ cd add
$ go install .
$ go build -o add
```

2. Initialize a new file

```bashscript
$ ./add init
Created new file: 20230330.md
```

3. Start adding phrases

```bashscript
$ ./add 'I want to remember this line'
Added phrase to file: 20230330.md
```


### Contributing
If you'd like to contribute to Add, please follow these guidelines:

- Fork the repository and create a new branch for your changes.
- Make your changes and write tests to ensure they work as expected.
- Run go fmt and go vet to ensure your code follows Go conventions.
- Submit a pull request with your changes. Be sure to provide a clear and concise description of your changes and why they are valuable.
