package stranglerfig

// El patrón de la Higuera Estranguladora (Strangler Fig Pattern) se utiliza para migrar gradualmente
// un sistema monolítico a una arquitectura de microservicios.
// Este patrón consiste en reemplazar gradualmente partes del monolito con nuevos microservicios.
// A medida que se desarrolla nueva funcionalidad o se actualiza la existente,
// se construye como un microservicio independiente. El tráfico se redirige desde el monolito hacia
// los nuevos microservicios hasta que finalmente el monolito es "estrangulado" y reemplazado por completo.
//
// El proceso de implementación del patrón de la Higuera Estranguladora generalmente involucra los siguientes pasos:
//
// Identificar los límites: Analizar el sistema monolítico e identificar componentes o funcionalidades distintas
// que puedan ser extraídas y convertidas en microservicios.
//
// Diseñar nuevos microservicios: Crear nuevos microservicios que repliquen la funcionalidad de los componentes o funcionalidades identificadas.
// Cada microservicio debe tener límites bien definidos y un propósito claro.
//
// Reemplazo gradual: Enrutar las solicitudes entrantes o el tráfico del sistema monolítico
// hacia los microservicios correspondientes para las funcionalidades identificadas.
// Esto se puede hacer utilizando técnicas como pasarelas de API o reglas de enrutamiento.
// Comenzar con una pequeña porción del tráfico y aumentarlo gradualmente a medida que se gana confianza en los microservicios.
//
// A medida que se realiza el reemplazo gradual, el monolito se va "estrangulando" lentamente,
// ya que más y más funcionalidades son manejadas por los nuevos microservicios.
// Este enfoque permite una migración controlada y reduce los riesgos asociados con una reescritura completa del sistema monolítico.
