-- Vai criar o banco devbook se ele não existir
CREATE DATABASE IF NOT EXISTS devbook;
-- Para garantir que todos os scripts sejam executados dentro do banco devbook
USE devbook;
-- Se tiver uma tabela usuarios, eu vou excluir ela
DROP TABLE IF EXISTS usuarios;


-- unique => Não pode ter dois usuarios ou mais com o mesmo nick, se eu tentar inserir dois usuarios com o mesmo nick ele vai travar
-- default current_timestamp() => O default será a data que criarmos o usuarios no banco e o proprio banco colocará essa data para nós
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(20) not null unique,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;