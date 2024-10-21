```
go install 'github.com/apstndb/filelocreader/cmd/filelocreader@latest'
```

It can read file content from `filepath:line:begin-end` format. It is compatible with `gopls` command.

```
$ loc=./testdata/src.go:5:6-10                                      
$ filelocreader ${loc}                                    
Expr
$ gopls implementation ${loc} | xargs filelocreader
Literal
Ident
```