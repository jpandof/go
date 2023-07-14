package main

import (
	"fmt"
)

// El patrón CQRS (Command Query Responsibility Segregation) es un patrón de diseño de software que separa las operaciones de lectura (Query)
// y escritura (Command) en un sistema. En lugar de tener un modelo único que maneje tanto las operaciones de lectura como las de escritura,
// CQRS propone tener dos modelos separados, cada uno especializado en una tarea específica.
//
// Las operaciones de escritura, o comandos, modifican el estado de un sistema pero no devuelven ningún dato.
// Las operaciones de lectura, o consultas, devuelven datos pero no cambian el estado del sistema.
//
// Este enfoque puede traer varios beneficios:
//
// Desacoplamiento: las operaciones de lectura y escritura suelen tener requisitos diferentes,
// por lo que puede ser beneficioso separarlas para simplificar su implementación y mantenimiento.
//
// Escalabilidad: los sistemas a menudo tienen muchos más lecturas que escrituras,
// por lo que separar estos dos tipos de operaciones permite escalarlos independientemente.
//
// Rendimiento y consistencia: al separar las operaciones de lectura y escritura,
// puede optar por diferentes enfoques para el manejo de datos, lo que puede mejorar el rendimiento y proporcionar una vista consistente de los datos.

type BlogPost struct {
	ID    string
	Title string
	Body  string
}

func main() {
	db := NewPostDatabase()

	// Create a new post
	{
		cmd := &CreatePostCommand{
			Post: &BlogPost{
				ID:    "1",
				Title: "Hello, world!",
				Body:  "This is my first post.",
			},
			Store: db,
		}
		cmd.Execute()
	}

	{
		// Get the post
		read := &GetPostQuery{
			ID:    "1",
			Store: db,
		}
		post := read.Execute().(*BlogPost)
		fmt.Printf("Post ID: %s\n", post.ID)
		fmt.Printf("Post Title: %s\n", post.Title)
		fmt.Printf("Post Body: %s\n", post.Body)
	}

}
