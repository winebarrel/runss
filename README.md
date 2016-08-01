# rumcmd

Run shell script command using Amazon Simple Systems Management Service.

# Usage

```
Usage of runss:
  -command string
      shell script command
  -instance-ids string
      comma separated instance ids
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
