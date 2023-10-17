<a name="readme-top"></a>

<!-- PROJECT LOGO -->
<br />
<div align="center">
  <h3 align="center">Excemplate</h3>
</div>

<!-- TABLE OF CONTENTS -->
<details>
  <summary>Table of Contents</summary>
  <ol>
    <li>
      <a href="#about-the-project">About The Project</a>
      <ul>
        <li><a href="#built-with">Built With</a></li>
      </ul>
    </li>
    <li>
      <a href="#getting-started">Getting Started</a>
      <ul>
        <li><a href="#prerequisites">Prerequisites</a></li>
        <li><a href="#installation">Installation</a></li>
      </ul>
    </li>
    <li><a href="#usage">Usage</a></li>
    <li><a href="#maintainers">Maintainers</a></li>
  </ol>
</details>

<!-- ABOUT THE PROJECT -->
## About The Project

A simple and performant templating engine for working with Excel files.

* Define multiple templates in a convenient manner
* Define items to generator
* Render
 
This service can either be used as a standalone program or as a library for integrating in to other GoLang applications

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Built With

* [Go](https://go.dev/)
* [Excelize](about:blank)
* [Cobra](https://cobra.dev/)

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- GETTING STARTED -->
## Getting Started

If you are not compiling the application for yourself, head to the releases page and download a the terminal client.

If you are going to be making modifications to the application or using it as a library, read on.

### Usage

```shell
# Will create an output.xlsx file in the current working directory
excemplate input.xlsx

# Custom output file
excemplate input.xlsx goforit.xlsx

# See help for more configuration options
excemplate -h
```


### Installation

The regular Go compile can be used if you are only intending to run the application on your architecture

```shell
go build -o excemplate .
```

#### Testing

The library code is located in the `lib` path; the rest of the code is Cobra boilerplate

```shell
go test ./lib
```

#### Compiling for targeted architecture

```shell
# Mac
GOOS=darwin GOARCH=amd64 go build -o ./excemplate .

# Windows
GOOS=windows GOARCH=amd64 go build -o ./excemplate.exe .

# Linux
GOOS=linux GOARCH=386 go build -o ./excemplate .
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- USAGE EXAMPLES -->
## Usage

The source workbook should have the following sheets (can be configured through Viper) present, see the included example for a working example using default settings

### Sheet Control (optional)

Determines the initial row state to start writing to. Will resort to first blank row otherwise

* `Sheet Name` - Name of the sheet to setup control for
* `Starting Row` - Row to start writing data from

### Template

Define some reusable templates. Each template should start with an identifier name and finish with `END` using the first column

* `Template` - Identifier for the template which will be used in the generator section
* `Content` - Columns and rows to be rendered (including the template tags) when generated
  * Dynamic variables should be defined using the following syntax `{{ .VARIABLE_NAME }}`. More info on variable name in the generators section

### Things To Generate

Items that are to be generated.

* `Template` - Template to use
* `Output Sheet` - Either use an existing sheet and start writing from the first empty row or, create a new sheet within your workbook

The program will automatically treat any further columns as template variables (identified by the value given in the header row).

* Empty values are allowed and will simply be left empty at runtime
* Variable names should be single words

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## 
<!-- MAINTAINERS -->

## Maintainers

An up-to-date list of people involved with development / support of the project

Simon Ball - [contact@simonball.me](mailto:contact@simonball.me)

<p align="right">(<a href="#readme-top">back to top</a>)</p>