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

SELECT DISTINCT * FROM boards B LEFT OUTER JOIN board_members M ON B.boardID = M.boardID WHERE B.adminID = 1 OR  M.userID = 1;


SELECT DISTINCT * FROM boards B LEFT OUTER JOIN board_members M ON 
B.boardID = M.boardID WHERE (B.adminID = 2 OR  M.userID = 2) AND boardID = 1;


SELECT * FROM boards B 
JOIN cards C on C.boardID = B.boardID
LEFT OUTER JOIN board_members M ON B.boardID = M.boardID 
WHERE (B.adminID = 2 OR  M.userID = 2) AND cardID = 1;