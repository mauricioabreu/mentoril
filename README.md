# Mentoril

Essa aplicação é um estudo de caso do Mentoril que acontece no Ubatuba's server.

A ideia é estudar go e construir um sistema um pouco diferente do comum, saindo um pouco de APIs de CRUD com poucas regras de negócio.

## O que é?

A aplicação tem como objetivo gerenciar configurações de um webserver existente que roda em servidores em um cluster. Cada instância roda uma cópia dessa aplicação.

Para configurar esse webserver, precisamos coletar informações de uma API externa ao webserver.

O webserver conta com uma API usada para instalar, atualizar, listar e deletar essas configurações.

## Sobre o webserver

Vamos criar um webserver básico apenas para fim de demonstrar como aplicar as configurações.

Esse webserver tem duas principais features:
* Controlar o tempo de cache dos objetos
* Controlar redirects dinâmicos

## Requisitos

* O webserver deve estar sempre com a configuração mais atualizada
* O webserver precisa subir com as configurações instaladas e atualizadas
    * Caso uma instância caia, ela precisa voltar com as configurações corretas
    * Caso o número de instâncias do webserver escale (de 5 instâncias pra 6, por exemplo), a instância nova precisa subir com as configurações corretas
* As aplicações precisam ter métricas:
    * Latência
    * Disponibilidade
