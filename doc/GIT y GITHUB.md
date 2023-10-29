# Manejo de Git y GitHub

En este proyecto, se implementan varias reglas y flujos de trabajo para mantener un orden y asegurar una colaboración efectiva. A continuación, se detallan las pautas y configuraciones utilizadas en Git y GitHub.

## Pautas de Commits

- Los commits deben ser redactados en tercera persona, ser descriptivos del cambio realizado y mantenerse concisos.
  - Ejemplo: "Ajusta comentarios en README.md"

## Protección de la Rama Main

Una vez que el repositorio se haga público, la rama `main` será protegida con las siguientes configuraciones:

- **Archivo CODEOWNERS:** Algunos archivos solo pueden ser modificados con la aprobación de los CODEOWNERS.
- **Pull Requests:** Se requiere crear un pull request para poder mergear a `main`.
- **Status Checks:** Todos los status checks deben pasar las pruebas antes de poder mergear a `main`.
- **Discard Approvals:** Los approvals serán descartados y considerados obsoletos cuando se envíen nuevos commits.

## Configuración del Repositorio con settings.yml

El archivo `settings.yml` en la raíz del repositorio se utiliza para definir configuraciones específicas de GitHub, como la protección de ramas, etiquetas, colaboradores y más. Este archivo ayuda a mantener un estado deseado en la configuración del repositorio, y puede ser versionado y revisado como cualquier otro archivo en el repositorio.

- **Etiquetas (Labels):** Define etiquetas personalizadas para categorizar issues y pull requests.
- **Protección de Ramas:** Configura las reglas de protección para las ramas importantes.
- **Colaboradores y Permisos:** Define colaboradores y sus niveles de permisos en el repositorio.

Se puede encontrar más detalles sobre las configuraciones disponibles y cómo se aplican en el documento [Configuración de settings.yml](./path-to-your/settings-yml-documentation).

# GitHub Actions

Se definen dos workflows en GitHub Actions, los cuales serán implementados en fases posteriores del proyecto:

## Validate Workflow

- **Test de Integración:** Verifica que el código integrado funcione correctamente.
- **Estilo de Código:** Comprueba que el estilo del código sea el adecuado.

## Deploy to Testing Workflow

- **Serverless Framework:** Ejecuta el serverless framework que despliega tanto la infraestructura como el código en la cuenta de AWS.
