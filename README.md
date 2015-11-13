# Logger

A [Go](http://golang.org/) logging library able to log in MongoDB, standard outputs and files

[![Build Status](https://travis-ci.org/VALERE91/logger.svg?branch=master)](https://travis-ci.org/VALERE91/logger)
[![Coverage Status](https://coveralls.io/repos/VALERE91/logger/badge.svg?branch=master&service=github)](https://coveralls.io/github/VALERE91/logger?branch=master)
[![License MIT](https://img.shields.io/npm/l/express.svg)](http://opensource.org/licenses/MIT)

## Installation

```bash
go get github.com/valere91/logger
```

The library has one dependency : mgo.v2

```bash
go get gopkg.in/mgo.v2
```

## Basic Usage

### To standard outputs

```go
StdoutLogger.Log("Something to log")
StderrLogger.Log("Something to log")
```

### To file

```go
//First initialise the log file 
ConfigureFileLogger("./output.log")

//You can use it after
FileLogger.Log("Something to log")
```

### To MongoDB

```go
//First initialise the MongoDB logger 
ConfigureMongoLogger("localhost","database","log_collection")

//You can use it after
MongoDBLogger.Log("Something to log")
```

## Documentation

See docs on [Wiki](https://github.com/VALERE91/logger/wiki)

## License

MIT License

Copyright (c) 2015 Valère Plantevin


Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:


The above copyright notice and this permission notice shall be included in
all copies or substantial portions of the Software.


THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.  IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
THE SOFTWARE.
