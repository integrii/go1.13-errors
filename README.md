# go1.13-errors
Testing the new features in go 1.13 for wrapping/unwrapping errors.


Output:

```
Plan println:
a bad codepath was called: This is some error that happened down the stack.  A specific one at that.
Error is of type SomeError
Unwrapped:
This is some error that happened down the stack.  A specific one at that.
```
