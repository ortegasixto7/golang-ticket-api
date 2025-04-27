# golang-ticket
Ticketing Backend in Golang

This Test Project pretends to generate tickets using Golang, taking the commit to generate and validate the ticket properly.

- Create integration
- Disable integration
- Purchase plan (tickets creations)

- Generate ticket
- Validate ticket
- Expire ticket (cron job/serverless)


Un tercero puede:
- Crear app
- Disabilitar app
- Listar apps
- Comprar planes (Para poder crear tickets)
- Crear tickets
- Validar ticket
- Listar tickets
- Expirar ticket (Se realiza automatico mediante Cron Job / Serverless)

Un admin puede:
- Ver todas las apps creadas
- Crear apps
- Desabilitar apps
- Ver los planes disponibles
- Crear nuevos planes
- Desabilitar planes
- Modificar planes
- Crear ticket
- Modificar ticket
- Validar ticket
- Expirar ticket



Install goose (For migrations)
goose create sql sql