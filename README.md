# golang-bb

![Peek 2021-02-21 23-46](https://user-images.githubusercontent.com/35817813/108651463-8e8b1100-74a0-11eb-88dc-95bc2beae092.gif)


 - Configure o postgress, crie um banco com o nome de go_test_gabriel_bubble
 
   - na pasta api/src/infra/migrations existem scripts SQL que irão auxiliar na criação do database.
  
   -  Há uma configuração docker do banco postgres para facilitar o uso do banco.
 
- Para start a api bastar dar um go run src/main/main.go

- Para startar o frontend 
   - Instale todos os pacotes do npm (npm install)
   - Execute o comando npm run serve
  
- Considerações
  - Projeto iniciado com layers de responsabilidades bem separados.
  
  - Inversão de dependências em alguns pontos chaves da aplicação
  
  - Facilidade de alterar regras de negócio

- Pontos para melhorias futuras
  - Melhor reaproveitamento de código para as responses https.
  
  - Validações de regras (ex; não criar dois usernames iguais).

  - Melhoria/padronização nas respostas https e respostas do infra layer.
 
  - Tratar erros adequadamente 
  
  - Confiuração adequada de seeders/migrations
