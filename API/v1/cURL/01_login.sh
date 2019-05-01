curl -k -X POST   https://localhost:8085/login/teste01   -H 'Content-Type: application/json'   -H 'Postman-Token: c99336ef-a3e9-4fbb-b88e-34c625edd3ca'   -H 'cache-control: no-cache'   -d '{"usuario":"teste01",  "senha":"123456"}' > token.json
# -k ignora tls bad certificate
