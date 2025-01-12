module github.com/NachooSR/Golang-Learning

go 1.22.1

replace github.com/NachooSR/greetings => ../greetings

require github.com/NachooSR/greetings v0.0.0-00010101000000-000000000000


//go mod edit -replace (nombre modulo (Repo))=../ (nombre carpeta)
//go mod tidy agrega los paquetes faltantes al proyecto