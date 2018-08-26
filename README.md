Web Assembly

```
GOARCH=wasm GOOS=js go build -o test.wasm test.go
cp $(go env GOROOT)/misc/wasm/wasm_exec.{html,js} .
```

Serve the generated files
