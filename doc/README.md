# Manejo de Git y Github

#### Para este proyecto implemento varias reglas de git y github para ordenar el trabajo:

- Los commits deben ser en tercera persona, descriptivos del cambio y concisos.
    - Por ejemplo: "Ajusta comentarios en README.md"

- La rama main sera protegida una vez que se haga publico el repositorio y controlará:
    - Algunos archivos solo pueden ser modificados con aprobacion de CODEOWNERS.
    - Se requiere pull request para mergear a main.
    - Se requiere que todos los status checks pasen las pruebas antes de mergear a main.
    - Descartar approvals seran obsoletas cuando se envían nuevas commits.

# Github Actions

#### Se definenn 2 workflows que seran implementados mas adelante:
- Validate: Test de integracion de codigo y estilo de codigo.
- Deploy to testing: Ejecuta el serverless framework que deploya en la cuenta de AWS la infraestructura y el codigo.