GOOS=js GOARCH=wasm go build -o docs/red.wasm
go run red.go
python3 -m http.server 8080 &
open http://localhost:8080
