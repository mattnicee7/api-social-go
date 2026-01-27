INSERT INTO usuarios (nome, nick, email, senha)
VALUES
("Usuario 1", "usuario_1", "usuario1@gmail.com", "$2a$10$L2d6eyIFE8c9GeK3/5kFt.Ydans77okYP9swQnPvizX0C8TO3Zv2i"),
("Usuario 2", "usuario_2", "usuario2@gmail.com", "$2a$10$L2d6eyIFE8c9GeK3/5kFt.Ydans77okYP9swQnPvizX0C8TO3Zv2i"),
("Usuario 3", "usuario_3", "usuario3@gmail.com", "$2a$10$L2d6eyIFE8c9GeK3/5kFt.Ydans77okYP9swQnPvizX0C8TO3Zv2i");

INSERT INTO seguidores (usuario_id, seguidor_id)
VALUES
(1, 2),
(3, 1),
(1, 3);