package main

import "sync"

func main() {
	mutex()
}

// sync.Mutex (Mutex simple): Es una primitiva de exclusión mutua que se utiliza para proteger el acceso a los datos en un entorno multihilo.
// Cuando un goroutine adquiere un Mutex, todos los demás goroutines deben esperar hasta que el Mutex sea liberado.
// Por lo tanto, sólo una goroutine puede tener acceso a los datos protegidos por el Mutex a la vez.
//
// sync.RWMutex (Mutex de lectura/escritura): Este es un tipo de Mutex más especializado que permite la lectura concurrente de datos.
// Esto significa que varias goroutines pueden leer los datos al mismo tiempo (siempre que no haya ninguna escritura en curso),
// pero sólo una goroutine puede escribir en los datos a la vez. Además, cuando una goroutine está escribiendo,
// ninguna otra puede leer o escribir hasta que la escritura se complete. Esta es una ventaja cuando los datos se leen mucho más a menudo de lo que se escriben.
//
//En resumen, deberías usar sync.Mutex cuando los datos se leen y se escriben a aproximadamente la misma frecuencia,
// y sync.RWMutex cuando los datos se leen mucho más a menudo de lo que se escriben,
// pero cada caso debe ser analizado cuidadosamente para decidir cuál de las dos primitivas de sincronización es la más adecuada.

func mutex() {
	var counter int
	var mutex sync.Mutex

	const goRoutines = 2
	var wg sync.WaitGroup
	wg.Add(goRoutines)
	for i := 0; i < goRoutines; i++ {
		go func() {
			for count := 0; count < 2; count++ {
				mutex.Lock()
				//defer mutex.Unlock()
				{
					value := counter
					value++
					counter = value
				}
				mutex.Unlock()
			}
			wg.Done()
		}()
	}
	wg.Wait()

}
