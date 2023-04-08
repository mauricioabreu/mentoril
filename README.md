# Mentoril

Essa aplicação é um estudo de caso do Mentoril que acontece no Ubatuba's server.

A ideia é estudar go e construir um sistema um pouco diferente do comum, saindo um pouco de APIs de CRUD com poucas regras de negócio.

## O que é?

A aplicação tem como objetivo gerenciar configurações de um software existente que roda em servidores em um cluster. Cada instância roda uma cópia dessa aplicação.

Para configurar esse software, precisamos coletar informações de uma API externa ao software.

O software conta com uma API usada para instalar, atualizar, listar e deletar as configurações.

## Requisitos

* O software deve estar sempre com a configuração mais atualizada
* O software precisa subir com as configurações instaladas e atualizadas
    * Caso uma instância caia, ela precisa voltar com as configurações corretas
    * Caso o número de instâncias do software escale (de 5 instância pra 6, por exemplo), a instância nova precisa subir com as configurações corretas
* As aplicações precisam ter métricas:
    * Latência
    * Disponibilidade
