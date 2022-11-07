Neste desafio você terá que usar o que aprendemos com Multithreading e APIs para buscar o resultado mais rápido entre duas APIs distintas.

As duas requisições serão feitas simultaneamente para as seguintes APIs:

https://cdn.apicep.com/file/apicep/xxxxxx-yyy.json

http://viacep.com.br/ws/xxxxxx-yyy/json/

https://api.postmon.com.br/v1/cep/xxxxxxyyy

url = "https://www.cepaberto.com/api/v3/cep?cep=01001000"
headers = {'Authorization': 'Token token=foo'}

Os requisitos para este desafio são:

-   Acatar a API que entregar a resposta mais rápida e descartar a resposta mais lenta.

-   O resultado da request deverá ser exibido no command line, bem como qual API a enviou.

-   Limitar o tempo de resposta em 1 segundo. Caso contrário, o erro de timeout deve ser exibido.
