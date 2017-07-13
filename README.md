abcutil
=======


[![Build Status](http://img.shields.io/travis/abcsuite/abcutil.svg)](https://travis-ci.org/abcsuite/abcutil)
[![Coverage Status](http://img.shields.io/coveralls/abcsuite/abcutil.svg)](https://coveralls.io/r/abcsuite/abcutil?branch=master)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](http://copyfree.org)
[![GoDoc](http://img.shields.io/badge/godoc-reference-blue.svg)](http://godoc.org/github.com/abcsuite/abcutil)

Package abcutil provides abcd-specific convenience functions and types.
A comprehensive suite of tests is provided to ensure proper functionality.  See
`test_coverage.txt` for the gocov coverage report.  Alternatively, if you are
running a POSIX OS, you can run the `cov_report.sh` script for a real-time
report.

This package was developed for abcd, a full-node implementation of Aero which
is under active development by Company 0.  Although it was primarily written for
dcrd, this package has intentionally been designed so it can be used as a
standalone package for any projects needing the functionality provided.

## Installation and Updating

```bash
$ go get -u github.com/abcsuite/abcutil
```

## Docker

All tests and linters may be run in a docker container using the script `run_tests.sh`.  This script defaults to using the current supported version of go.  You can run it with the major version of go you would like to use as the only arguement to test a previous on a previous version of go (generally Aero supports the current version of go and the previous one).

```
./run_tests.sh 1.7
```

To run the tests locally without docker:

```
./run_tests.sh local
```

## License

Package abcutil is licensed under the [copyfree](http://copyfree.org) ISC
License.
