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
Created new file: 20230330.txt
```

3. Start adding phrases

```bashscript
$ ./add 'I want to remember this line'
Added phrase to file: 20230330.txt
```

4. Combine files into a single file (e.g. `combined.txt`)

```bashscript
$ cat *.txt > combined.txt
```

5. Navigate through the file with fzf, select a line, and copy to clipboard

```bashscript
$ fzf < combined.txt | tee >(pbcopy)
```

<img width="880" alt="CleanShot 2023-07-24 at 16 35 17@2x" src="https://github.com/nancyel/add/assets/76415082/f6a90ad8-3371-4048-a024-a72ba6897faf">


6. Or simply, preview files from your terminal (save to `zprofile` for quick access)

```bashscript
# $HOME/zprofile

# Preview with fzf
preview() {
  fzf --preview 'cat {}'
}
```


### Contributing
If you'd like to contribute to Add, please follow these guidelines:

- Fork the repository and create a new branch for your changes.
- Make your changes and write tests to ensure they work as expected.
- Run go fmt and go vet to ensure your code follows Go conventions.
- Submit a pull request with your changes. Be sure to provide a clear and concise description of your changes and why they are valuable.
