![URL Shortener Logo](./imagen.png)

# Servicio de Acortador de URL

![Go Version](https://img.shields.io/badge/Go-1.21-blue)
![AWS DynamoDB](https://img.shields.io/badge/AWS-DynamoDB-orange)
![AWS Lambda](https://img.shields.io/badge/AWS-Lambda-green)
![AWS ElastiCache](https://img.shields.io/badge/AWS-ElastiCache-red)

## Índice

- [Arquitectura y Tecnologías](#arquitectura-y-tecnologías)
  - [Base de Datos](#base-de-datos)
  - [APIs](#apis)
  - [Observabilidad](#observabilidad)
  - [Caché](#caché)
- [Flujo Básico](#flujo-básico)
  - [Acortar URL](#acortar-url)
  - [Recuperar URL](#recuperar-url)
  - [Eliminar URL](#eliminar-url)
- [Escalabilidad y Rendimiento](#escalabilidad-y-rendimiento)
  - [DynamoDB](#dynamodb)
  - [AWS Lambda](#aws-lambda)
  - [Caché](#caché-1)
- [Observabilidad](#observabilidad-1)
  - [Monitoreo y Alertas](#monitoreo-y-alertas)
  - [Registros](#registros)
- [Puntos de Validación](#puntos-de-validación)

## Arquitectura y Tecnologías

### Base de Datos
Opté por AWS DynamoDB por su rápido y escalable acceso a datos. DynamoDB ofrece escalabilidad sin problemas, lo que lo convierte en una opción adecuada para un servicio de acortamiento de URL donde el rendimiento y la escalabilidad son factores clave.

### APIs
Para los endpoints de la API, decidí implementar funciones Lambda en AWS escritas en Go. Los beneficios de rendimiento de Go y la facilidad de implementación en Lambda lo convirtieron en una opción ideal. Las funciones Lambda son responsables de acortar y recuperar URLs.

### Observabilidad
Estoy usando AWS CloudWatch para monitorear métricas básicas y configurar alertas. Para obtener análisis de rendimiento más detallados, he integrado New Relic ya que ofrece un excelente soporte para Go.

### Caché
Para reducir la latencia, he añadido AWS ElastiCache. Almacena URLs accedidas con frecuencia, lo que acelera su recuperación.

## Flujo Básico

### Acortar URL
Cuando se recibe una solicitud para acortar una URL, una función Lambda basada en Go genera un identificador corto único, guarda la URL completa y el identificador corto en DynamoDB, y luego devuelve la URL acortada.

### Recuperar URL
Al recibir una URL acortada, el sistema primero verifica la caché. Si no se encuentra, realiza una consulta en DynamoDB y redirige a la URL larga original.

### Eliminar URL
Eliminar una URL acortada es sencillo: simplemente se elimina el registro correspondiente en DynamoDB.

## Escalabilidad y Rendimiento

### DynamoDB
He configurado particiones en DynamoDB para gestionar un gran número de solicitudes.

### AWS Lambda
La concurrencia y la asignación de memoria en las funciones Lambda basadas en Go están ajustadas para manejar eficientemente los picos de tráfico.

### Caché
Se mantiene un TTL (Tiempo de Vida) razonable en la caché para garantizar que las URLs más populares se resuelvan más rápidamente.

## Observabilidad

### Monitoreo y Alertas
He configurado CloudWatch para métricas personalizadas y alertas relacionadas con la latencia, errores y otros signos vitales.

### Registros
Se mantienen registros detallados de las funciones Lambda para una mejor trazabilidad.

## Puntos de Validación

1. **APIs de Creación y Eliminación de URLs**: El rendimiento, la latencia y el flujo de datos hacia y desde DynamoDB se monitorean de cerca.
2. **Navegación de URL Corta a URL Larga**: Se verifica la latencia de redirección.
3. **Fallos en Componentes**: He planificado y documentado cómo el sistema se recuperará de diferentes fallos en componentes como Lambda y DynamoDB.
