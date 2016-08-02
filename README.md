# runss

Run shell script command using Amazon Simple Systems Management Service.

[![Build Status](https://travis-ci.org/winebarrel/runss.svg?branch=master)](https://travis-ci.org/winebarrel/runss)

## Usage

```
Usage of runss:
  -command string
      shell script command
  -instance-ids string
      comma separated instance ids
  -prompt
      show prompt
```

```
$ runss -command hostname -instance-ids i-xxx...,i-yyy...
- InstanceId: i-xxx...
  Status: Success
  Output: |
    ip-10-10-10-10
- InstanceId: i-yyy...
  Status: Success
  Output: |
    ip-10-10-10-11
```

### Show prompt

```
$ runss -instance-ids i-xxx...  -prompt
> hostname
- InstanceId: i-xxx...
  Status: Success
  Output: |
    ip-10-10-10-10

> whoami
- InstanceId: i-xxx...
  Status: Success
  Output: |
    root
```

## Installation

```
brew install https://raw.githubusercontent.com/winebarrel/runss/master/homebrew/runss.rb
```
