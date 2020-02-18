drop table fornecedor_materia_prima;
drop table produto_materia_prima;
drop table formulacao_produto;
drop table produto;
drop table formulacao;
drop table fornecedor;
drop table materia_prima;
drop table usuario;

CREATE TABLE produto (
  id SERIAL,
  nome VARCHAR(40) NULL,
  codigo VARCHAR(10) NULL,
  PRIMARY KEY(id)
);

CREATE TABLE formulacao (
  id SERIAL,
  qtd_estimada_produzida FLOAT NULL,
  icm FLOAT NULL,
  pis FLOAT NULL,
  cofins FLOAT NULL,
  comissao FLOAT NULL,
  margem_lucro FLOAT NULL,
  despesa_fixa FLOAT NULL,
  outros_descontos FLOAT NULL,
  custo_estimado FLOAT NULL,
  obs VARCHAR(100) NULL,
  valor_venda FLOAT NULL,
  codigo VARCHAR(10) NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE formulacao_produto (
  id SERIAL,
  formulacao_id INTEGER NOT NULL,
  produto_id INTEGER NOT NULL,
  quantidade FLOAT NULL,
  valor FLOAT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY (formulacao_id) REFERENCES formulacao(id),
  FOREIGN KEY (produto_id) REFERENCES produto(id)
);

CREATE TABLE fornecedor (
  id SERIAL,
  cnpj VARCHAR(14) NULL,
  nome VARCHAR(50) NOT NULL,
  empresa VARCHAR(50) NOT NULL,
  telefone1 VARCHAR(16) NULL,
  telefone2 VARCHAR(16) NULL,
  email VARCHAR(40) NULL,
  PRIMARY KEY(id)
);

CREATE TABLE materia_prima (
  id SERIAL,
  codigo VARCHAR(10) NOT NULL,
  nome VARCHAR(20) NOT NULL,
  unidade_compra VARCHAR(1) NOT NULL,
  valor FLOAT NOT NULL,
  PRIMARY KEY(id)
);

CREATE TABLE fornecedor_materia_prima (
  id SERIAL,
  fornecedor_id INTEGER  NOT NULL,
  materia_prima_id INTEGER  NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(fornecedor_id) REFERENCES fornecedor(id),
  FOREIGN KEY(materia_prima_id) REFERENCES materia_prima(id)
);


CREATE TABLE produto_materia_prima (
  id SERIAL,
  materia_prima_id INTEGER NOT NULL,
  produto_id INTEGER NOT NULL,
  PRIMARY KEY(id),
  FOREIGN KEY(materia_prima_id) REFERENCES materia_prima(id),
  FOREIGN KEY(produto_id) REFERENCES produto(id)
);


CREATE TABLE usuario (
  id SERIAL,
  nome VARCHAR(50) NOT NULL,
  cpf VARCHAR(16) NOT NULL,
  email VARCHAR(30) NOT NULL,
  senha VARCHAR(30) NOT NULL,
  PRIMARY KEY(id)
);