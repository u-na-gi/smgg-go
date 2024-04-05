# smgg-go

`smgg-go` stands for Struct-Merging-Code-Generator. It is a tool designed to generate setters that merge structures of the same type, such as entities within your project. The generator works across all packages within a project, providing a comprehensive solution for struct merging needs.

This tool is especially useful for projects that require the merging of struct fields from various sources or for simplifying the process of combining data from similar structs.

## Warning

This is an experimental project and may perform highly risky operations, including overwriting existing files or generating unexpected files. Please use it with caution and at your own risk.


### Features

- Generates setters for merging structs of the same type.
- Supports all packages within a project.
- Includes code written for validation purposes.

### Installation

You can easily install the CLI version of `smgg-go` using the following command:

```shell
go install github.com/u-na-gi/smggcli
```

This command downloads the CLI version of `smgg-go`, enabling you to use the struct-merging-code-generator directly from your command line.

### Usage

After installation, you can invoke the generator by simply running `smggcli` followed by the necessary arguments and options. For detailed usage instructions, please refer to the `smggcli` documentation.

### Contributing

Contributions to `smgg-go` are welcome. Whether it's improving the codebase, adding new features, or fixing bugs, your help is appreciated. Please feel free to fork the repository, make your changes, and submit a pull request.

### License

`smgg-go` is open-source software licensed under the MIT license. Please see the LICENSE file in the repository for more details.