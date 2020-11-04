/*SELECT DISTINCT B.boardID, B.boardName, B.theme, B.star FROM boards B JOIN board_members M
ON B.boardID = M.boardID WHERE B.adminID = 1 OR  M.userID = 1;


 boardid | boardname | theme | star 
---------+-----------+-------+------
       1 | board_1   | dark  | f
       2 | board_2   | dark  | f
(3 строки)


SELECT adminID, boardName, theme, star, M.userID FROM boards JOIN board_members M on M.boardID = boardID WHERE boardID = 1;*/

/*DROP TABLE IF EXISTS cards;
CREATE TABLE cards (
                       cardID    SMALLSERIAL PRIMARY KEY,
                       boardID     BIGSERIAL REFERENCES boards(boardID) ON DELETE CASCADE,
                       cardName    TEXT,
                       cardOrder SMALLSERIAL
);
*/

