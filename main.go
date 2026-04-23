package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

func handler(w http.ResponseWriter, r *http.Request) {
	hostname, _ := os.Hostname()
	html := fmt.Sprintf(`
<!DOCTYPE html>
<html lang="es">
<head>
<meta charset="UTF-8">
<title>Mi Web UOC</title>
<style>
body { font-family: Arial, sans-serif; text-align: center; margin-top: 80px; }
h1 { color: #003366; }
img { width: 200px; margin-top: 20px; }
.pod { margin-top: 16px; font-weight: bold; color: #1f2937; }
</style>
</head>
<body>
<h1>Soy alumno de la UOC</h1>
<p>Bienvenido a mi primera aplicación web con Go y Docker.</p>
<img src="/logo.png" alt="Logo UOC">
<p class="pod">Hola! Atendido por el pod: %s</p>
</body>
</html>`, hostname)
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