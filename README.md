# Feira do rolo APP

## Intuíto e motivação

Aplicativo feito para finalização da disciplina de frontend: angular da 
pós do uniesp. O intuíto do aplicativo é simular uma "feira do rolo" 
como popularmente dito, feira onde as pessoas vão para trocar, comprar
ou vender coisas. Aqui, foram feitos cruds e visualizações apenas 
com o intuíto de aplicarmos a parte de rotas, exibição e fetch em API'S dentro
do angular. 

## Tecnologias utilizadas 

### NPM. Versão: 8.11.0
### GOLang versão 1.17.6
### Docker 20.10.11-build
## Editores utilizados 
### Visual Studio Code 1.17.2
### Goland 2022.2.3 (Opcional, é possível ver o código em go através
### do visual studio code)

## Setup e execução 
Abaixo estão os passos para a correta execução do projeto em si. Lembrando 
que serão necessárias as instalações das tecnologias acima utilizadas 
para a correta instalação e os editores para melhor visualização do 
código produzido. 

## Client
1) para o cliente: Navegue até a raiz do projeto e utilze o comando npm install para 
instalação de todas as dependências presentes no projeto. Não foram utilizadas bibliotecas
como bootstrap ou similares para facilitação de execução do código. 
2) Após isso, inicie o projeto através do comando npm start, ainda na raiz do mesmo.
## Servidor

1) Para o servidor, navegue até a pasta chamada server na raiz do projeto.
2) dentro da pasta, use o comando docker-compose up -d para inicializar
o container do postgres em sua máquina. 
3) Após isso, utilize o comando: go build ./cmd/feira-do-rolo/main ainda
na raiz do projeto. Ele irá buildar e criar o executável chamado main. 
4) basta o main.exe construído no passo anterior. 

## Melhorias 
Existem uma série de melhorias ainda a serem feitas ao projeto, que seguem: 
1) melhor tratamento de erros na parte do client. 
2) Aumento da cobertura de testes unitários no servidor e client. 
3) Evolução do projeto em si. 

Isso não foi feito, pois o escopo da disciplina focou apenas no básico de
ensino do angular e, com o que foi feito até o momento, a entrega está 
satisfatória, cobrindo basicament tudo que foi feito na disciplina em si. 





