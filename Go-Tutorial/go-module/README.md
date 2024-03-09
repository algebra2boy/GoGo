## Important
In hello directory, if we want to reference the code from other module, we need to get the relative path based on the current folder. For example,
```bash
go mod edit -replace example.com/greetings=../greetings
```

update/synchronize a module's dependencies
```bash
go mod tidy
```