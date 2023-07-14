package registrydiscovery

import (
	"context"
	"fmt"
	"github.com/micro/go-micro/v2"
	"github.com/micro/go-micro/v2/client"
)

// En una arquitectura de microservicios, tienes varios servicios individuales ejecutándose, cada uno de los cuales realiza una tarea o conjunto de tareas específicas.
// Estos servicios necesitan comunicarse entre sí para lograr los objetivos del sistema en su conjunto.
// Para hacer esto, un servicio necesita saber dónde se encuentra otro servicio. Aquí es donde entra en juego el patrón de registro y descubrimiento de servicios.
//
// El patrón de registro y descubrimiento de servicios se compone de dos partes principales:
//
// Registro de servicios: Cuando un servicio se inicia, se "registra" en una especie de directorio o "registro de servicios".
// Esto generalmente implica decirle al registro cuál es su nombre y cómo se puede acceder a él (por ejemplo, su dirección IP y puerto).
//
// Descubrimiento de servicios: Cuando un servicio necesita comunicarse con otro, puede buscar el nombre del otro servicio en el registro de servicios.
// El registro luego proporciona la información necesaria para que el primer servicio se conecte al segundo.

type ExampleService struct{}

func (e *ExampleService) Method(ctx context.Context, req *client.Request, rsp *client.Response) error {

	return nil
}

func newService() {
	// Crea un servicio
	service := micro.NewService(
		micro.Name("example.service"),
	)

	// Inicializa el servicio
	service.Init()

	// Registra el manejador
	micro.RegisterHandler(service.Server(), new(ExampleService))

	// Corre el servicio
	if err := service.Run(); err != nil {
		fmt.Println(err)
	}
}

func myService() {
	// Crea un servicio
	//service := micro.NewService(micro.Name("example.caller"))

	// Crea un cliente del servicio
	//exampleService := NewExampleService("example.service", service.Client())

	// Llama al método del servicio
	//rsp, err := exampleService.Method(context.TODO(), &example.Request{
	//	Name: "John",
	//})

	//if err != nil {
	//	fmt.Println(err)
	//	return
	//}
	//
	//fmt.Println(rsp.Msg)
}
