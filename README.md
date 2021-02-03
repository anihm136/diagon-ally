# diagon-ally
Yes I know it is a terrible pun at best

## Build instructions
Get `stuffbin` -
```sh
go get -u github.com/knadh/stuffbin/...
```
Then build the binary with
```sh
go build . -o temp
```
Finally, stuff the binary with the templates -
```sh
stuffbin -a stuff -in temp -out diagon default_insert.txt default_template.svg
```
