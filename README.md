# gooce

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

gooce is a CLI which exports Google Calendar events to text written in Go.
This is useful for pasting events to any text field.

## Usage

```
$ gooce
```

## Installation

At first, to use gooce you should create a new credential on https://console.cloud.google.com/apis/credentials and save the credential file as `~/.gooce/client_secret.json`.

Then you can install gooce by the following command:

```
$ go get github.com/kami-zh/gooce
```

## Contributing

Bug reports and pull requests are welcome on GitHub at https://github.com/kami-zh/gooce.

## License

The gem is available as open source under the terms of the [MIT License](http://opensource.org/licenses/MIT).
