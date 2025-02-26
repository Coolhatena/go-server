package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Message string `json:"message"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	switch r.Method {
	case "GET":
		response := Response{Message: "Se recibio una solicitud GET"}
		json.NewEncoder(w).Encode(response)

	case "POST":
		response := Response{Message: "Se recibio una solicitud POST"}
		json.NewEncoder(w).Encode(response)

	case "PUT":
		response := Response{Message: "Se recibio una solicitud PUT"}
		json.NewEncoder(w).Encode(response)

	case "DELETE":
		response := Response{Message: "Se recibio una solicitud DELETE"}
		json.NewEncoder(w).Encode(response)

	default:
		http.Error(w, "Método no permitido", http.StatusMethodNotAllowed)
	}

	// 1. Parsear los valores del cuerpo de la solicitud
	err := r.ParseForm()
	if err != nil {
		http.Error(w, "Error al parsear los datos", http.StatusBadRequest)
		return
	}

	// 2. Convertir los datos a un mapa
	data := make(map[string]string)
	for key, values := range r.Form {
		data[key] = values[0] // Tomamos el primer valor en caso de que haya varios
	}

	// 3. Convertir el mapa a JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(data)
}

func main() {
	http.HandleFunc("/", handler)
	fmt.Println("Servidor corriendo en http://localhost:3333")
	http.ListenAndServe(":3333", nil)
}
