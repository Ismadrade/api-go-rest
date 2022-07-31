
<div style="display: inline_block"><br>
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/go/go-original.svg" />
  <img align="center" alt="RabbitMQ" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/docker/docker-original-wordmark.svg" />
  <img align="center" alt="Spring" height="80" width="90" src="https://cdn.jsdelivr.net/gh/devicons/devicon/icons/postgresql/postgresql-original.svg" />
</div>

# Minha Primeira API - Go Lang
Pequena API desenvolvido em Go Lang, que permite criar, consultar, editar e deletar uma personalidade.

## Requisitos

- Go Lang
- Docker

## Configurando o Ambiente

### Docker
1) Instalar o Docker (Tutorial via Windows). [Download](https://www.docker.com/products/docker-desktop/) e  [Instalação](https://sh-tsang.medium.com/tutorial-docker-installation-in-wsl-2-of-windows-f4471fc3e1d4)

2) No diretório raiz, executar o comando abaixo:
```shell
docker-compose up
```

### Go Lang

1) [Baixar](https://go.dev/doc/install) e [instalar](https://medium.com/@rafaelmoraisdev/como-instalar-go-no-windows-10-7787faac3a7f) o Go Lang.

## Executando o projeto
Dentro da pasta raiz, rodar o seguinte comando:
```shell
go run main.go  
```
O projeto estará rodando em localhost na porta 8000
```http request
http://localhost:8000/
```

## Endpoints


  <b><span style="color:green"> GET</span></b> `http://localhost:8000/api/personalidades`

  <b><span style="color:green"> GET</span></b> `http://localhost:8000/api/personalidades/1`

  <b><span style="color:yellow"> POST</span></b> `http://localhost:8000/api/personalidades`

  <b><span style="color:blue"> PUT</span></b> `http://localhost:8000/api/personalidades/1`

  <b><span style="color:red"> DELETE</span></b> `http://localhost:8000/api/personalidades/1`

### Body

```json
{
    "nome": "Meu Nome",
    "historia": "Minha historia"
}
```


## Referências
- [Alura](https://cursos.alura.com.br/course/go-desenvolvendo-api-rest)
