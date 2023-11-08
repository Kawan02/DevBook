-- Vai criar o banco devbook se ele não existir
CREATE DATABASE IF NOT EXISTS devbook;
-- Para garantir que todos os scripts sejam executados dentro do banco devbook
USE devbook;

-- Se tiver uma tabela publicacoes, exclui ela
DROP TABLE IF EXISTS publicacoes;

-- Se tiver uma tabela usuarios, exclui ela
DROP TABLE IF EXISTS usuarios;

-- Se tiver uma tabela seguidores, exclui ela
DROP TABLE IF EXISTS seguidores;


-- unique => Não pode ter dois usuarios ou mais com o mesmo nick, se eu tentar inserir dois usuarios com o mesmo nick ele vai travar
-- default current_timestamp() => O default será a data que criarmos o usuarios no banco e o proprio banco colocará essa data para nós
CREATE TABLE usuarios(
    id int auto_increment primary key,
    nome varchar(50) not null,
    nick varchar(50) not null unique,
    email varchar(50) not null unique,
    senha varchar(100) not null,
    criadoEm timestamp default current_timestamp()
) ENGINE=INNODB;


CREATE TABLE seguidores(
    usuario_id int not null,
    FOREIGN KEY (usuario_id) -- Chave estrangeira => basicamente referencia uma outra tabela
    REFERENCES usuarios(id)
    ON DELETE CASCADE, -- ON DELETE CASCADE => basicamente vai excluir todas as informações do usuário se ele não existir mais na tabela usuarios
    
    seguidor_id int not null,
    FOREIGN KEY (seguidor_id)
    REFERENCES usuarios(id)
    ON DELETE CASCADE,

    primary key(usuario_id, seguidor_id)  -- Chave primária composta => A chave primária vai ser uma junção das duas colunas
) ENGINE=INNODB;


CREATE TABLE publicacoes(
    id int auto_increment primary key,
    titulo varchar(50) not null,
    conteudo varchar(300) not null,
    
    autor_id int not null,
    FOREIGN KEY (autor_id) -- Chave estrangeira => basicamente referencia uma outra tabela
    REFERENCES usuarios(id)
    ON DELETE CASCADE, -- ON DELETE CASCADE => basicamente vai excluir todas as informações do usuário se ele não existir mais na tabela usuarios
    
    curtidas int default 0,
    criadaEm timestamp default current_timestamp()
) ENGINE=INNODB;
