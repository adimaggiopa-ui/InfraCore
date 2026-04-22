package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Mi Web UOC</title>
<style>
body { font-family: Arial, sans-serif; text-align: center; margin-top: 80px; }
h1 { color: #003366; }
img { width: 200px; margin-top: 20px; }
</style>
</head>
<body>
<h1>Soy alumno de la UOC</h1>
<p>Bienvenido a mi primera aplicación web con Go y Docker.</p>
<img src="/logo.png" alt="Logo UOC">
</body>
</html>`
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	fmt.Fprint(w, html)
}

func imageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "logo.png")
}

func main() {
	http.HandleFunc("/", handler)
	http.HandleFunc("/logo.png", imageHandler)

	fmt.Println("Servidor arrancado en http://localhost:8081")
	log.Fatal(http.ListenAndServe(":8081", nil))
}