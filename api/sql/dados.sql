INSERT INTO usuarios(nome, nick, email, senha)
values
("Usuário 1", "usuario_1", "usuario1@gmail.com", "$2a$10$3OK3b5wJIvZpgP0C6jkT2e4zdoiZv7OScFoxBr/r/Fuyn3gy2kkR6"),
("Usuário 2", "usuario_2", "usuario2@gmail.com", "$2a$10$3OK3b5wJIvZpgP0C6jkT2e4zdoiZv7OScFoxBr/r/Fuyn3gy2kkR6"),
("Usuário 3", "usuario_3", "usuario3@gmail.com", "$2a$10$3OK3b5wJIvZpgP0C6jkT2e4zdoiZv7OScFoxBr/r/Fuyn3gy2kkR6");



INSERT INTO seguidores(usuario_id, seguidor_id)
values
(1, 2),
(3, 1),
(1, 3);