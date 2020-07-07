# Objetivo

Desenvolva um serviço de manipulação de dados e persistência em base de dados relacional.

Requisitos:

- Criar um serviço em GO que receba um arquivo csv/txt de entrada (Arquivo Anexo)
- Este serviço deve persistir no banco de dados relacional (postgresql) todos os dados contidos no arquivo
  Obs: O arquivo não possui um separador muito convencional
 
- Deve-se fazer o split dos dados em colunas no banco de dados
 Obs: pode ser feito diretamente no serviço em GO ou em sql
 
- Realizar higienização dos dados após persistência (sem acento, maiúsculo, etc)
- Validar os CPFs/CNPJs contidos (válidos e não válidos numericamente)
- Todo o código deve estar disponível em repositório público do GIT
 
Desejável:

- Utilização das linguagen GOLANG para o desenvolvimento do serviço
- Utilização do DB Postgres
- Docker Compose , com orientações para executar (arquivo readme) 

Você será avaliado por:

- Utilização de melhores práticas de desenvolvimento (nomenclatura, funções, classes, etc);
- Utilização dos recursos mais recentes das linguagens;
- Boa organização lógica e documental (readme, comentários, etc);
- Cobertura de todos os requisitos obrigatórios.


Nota:

 - Todo a estrutura relacional deve estar documentada (criação das tabelas, etc)
 - Criação de um arquivo README com as instruções de instalação         juntamente com as etapas necessárias para configuração.
 - Você pode escolher sua abordagem de arquitetura e solução técnica.


# Tabelas

O arquivo de inicialização das tabelas encontre-se [aqui](./resource/ddl/init.sql). Ao subir os containers ele é executado automaticamente.

> **As tabelas são criadas automaticamente**: Não precisa rodar os scripts!

### Shopping

```sql
CREATE TABLE shopping (
	id serial NOT NULL,
	customer_id varchar(15) NULL,
	private int4 NULL,
	incomplete int4 NULL,
	last_shop date NULL,
	avg_ticket float8 NULL,
	last_ticket_shop float8 NULL,
	most_frequented_store varchar(15) NULL,
	last_store varchar(15) NULL
);
```

 | id  | customer_id | private | incomplete | last_shop  | avg_ticket | last_ticket_shop | most_frequented_store | last_store     |
 | --- | ----------- | ------- | ---------- | ---------- | ---------- | ---------------- | --------------------- | -------------- |
 | 1   | 08903682955 | 0       | 0          | 2013-06-12 | 53.25      | 53.25            | 79379491000850        | 79379491000850 |
 | 2   | 34524472860 | 0       | 0          | 2013-06-12 | 91.28      | 91.28            | 79379491000850        | 79379491000850 |
 | 3   | 34850830030 | 0       | 0          | 2013-06-12 | 72.05      | 72.05            |


### Inconsistency

```sql
CREATE TABLE inconsistency (
	id serial NOT NULL,
	filename text NULL,
	error_message text NULL
);
```

 | id  | filename             | error_message                              |
 | --- | -------------------- | ------------------------------------------ |
 | 1   | ./resource/teste.txt | CNPJ is invalid! [last_store:04209828840]. |


# Metodologia

- Como não era requisito manter a ordem de inserção e dado que o **custo** de _uma inserção_ no **PostGres** é quase o mesmo que a inserção de _multiplas linhas_ resolvi paralelizar a inserção dos dados. O PostGres tem um limite de campos a serem preenchidos em inserções de multiplas linhas, este limite é de: 
  
  > 65535 - Postgress limit parameters

  Para evitar o erro: *Postgress limit parameters*, busquei um denominador que executasse um número maior de inserções garantido uma performance desejável e sem que o erro ocorresse. Basicamente dividindo o número de parameters pelo número de atributos do modelo. Para facilitar a compreensão, arredondei o limite para 65000. Veja:

  > 65000 / 8 = 8125 -> Denominador == 8125
    
		Dado que há 49999 linhas
		E um modelo de 8 atributos
		Quando 49999 é dividido por 8125 
		Então o resultado é 6.

		Portanto, neste cenário posso paralelizar a inserção em 6 chamadas de 8125 linhas e executar mais um lanço somente com os itens restantes. 

- Caso alguma linha contenha dados inválidos os registro serão salvos na tabela de inconsistências.
- Antes da importação um Truncate é executado para limpar a tabela de stage.

- Temos um modulo server para consulta dos dados importados 
- Temos um modulo cmd para importar os dados.
- Ambos os modulos estão conteinerizados, porém o cmd o ciclo de vida é por execução.

# Enviroments

```env
  	DB_DRIVER=postgres
	DB_HOST=postgres
	DB_USER=postgres 
	DB_NAME=import-data
	DB_PORT=5432
	DB_PASSWORD=db@123A
	SERVER_PORT=8085
```

# Repository

 - **SaveMany:** _Grava uma coleção de dados._

 - **GetAll:** _Retorna uma coleção de dados._

 - **Truncate:** _Limpa a tabela e reinicia a pk._


# Endpoints  
  

#### Consultar itens importados

```bash
curl --location --request GET 'http://localhost:8085/shoppings' 
```

#### Consultar inconsistências

```bash
curl --location --request GET 'http://localhost:8085/inconsistencies' 
```

# Passos para testar

Dependências:
 - Docker 
 - Docker Compose
 - Make

Caso queira visualizar os *comandos* disponíveis com mais detalhe, acesse [aqui](./makefile).

#### Baixar Repositório
```bash
git clone https://github.com/desafios-job/import-data.git && cd import-data
```

#### Build
```bash
make docker-compose
```

#### Executar o EXTRACT

##### Argumentos

 - **VOLUME** -> indica onde será referenciado o volume para acesso do container ao filesystem.

	Ex: */Users/samuelsantos/git/import-data/resource*

 - **FILENAME** -> Nome do arquivo para importação. 
	
	Ex: *base_test.txt*. 

> Será necessário avaliar se o docker tem permissão de acesso ao volume informado no argumento *VOLUME*.
```
make docker-run VOLUME=$(VOLUME) FILENAME=$(FILENAME)
```

