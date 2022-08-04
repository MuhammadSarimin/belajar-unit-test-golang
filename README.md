# belajar-unit-test-golang

### Menjalan kan semua testing.T dari folder root
```
go test -v ./...
```

### Masuk ke folder untuk menjalankan test ke spesifik folder
```
cd helper
go test -v
```
### Menjalan kan spesifik func Test
```
go test -v -run=NamaFunctionTest
```
### menjalan kan testing.B untuk benchmark function
```
go test -run=Bench -bench .
```