INSERT INTO users values (1, 'epridius@yandex.ru', 'pkaterina', 'Lalala', '', ''),
                         (2, 'katts@', 'lala', 'katts', '', ''),
                         (3, 'egor@', 'lala', 'aist', '', '');

INSERT INTO boards values (1, 1, 'board_1', 'dark', false),
                          (2, 2, 'board_2', 'dark', false);

INSERT INTO board_members values (1, 2), (1, 3), (2, 1),(2, 3);

--INSERT INTO users values (2, 'epridius@yandex.ru', 'pkaterina', 'Lalala', '', 'default/default_avatar.png');