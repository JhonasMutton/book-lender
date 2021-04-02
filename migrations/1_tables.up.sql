CREATE TABLE IF NOT EXISTS USER
(
    USER_ID    INT NOT NULL AUTO_INCREMENT,
    NAME       VARCHAR(100),
    EMAIL      VARCHAR(300) UNIQUE,
    CREATED_AT DATE,
    PRIMARY KEY (USER_ID)
);

CREATE TABLE IF NOT EXISTS BOOK
(
    BOOK_ID    INT NOT NULL AUTO_INCREMENT,
    FK_USER_ID INT NOT NULL,
    TITLE      VARCHAR(100),
    PAGES      INT,
    CREATED_AT DATE,
    PRIMARY KEY (BOOK_ID)
    FOREIGN KEY (FK_USER_ID) REFERENCES USER (USER_ID),
);

CREATE TABLE IF NOT EXISTS LOAN_BOOKS
(
    LOAN_BOOKS_ID   INT NOT NULL AUTO_INCREMENT,
    FK_BOOK_ID      INT NOT NULL,
    FK_FROM_USER_ID INT NOT NULL,
    FK_TO_USER_ID   INT NOT NULL,
    LENT_AT         DATE,
    RETURNED_AT     DATE,
    IS_ACTIVE       BOOLEAN,
    PRIMARY KEY (LOAN_BOOKS_ID),
    FOREIGN KEY (FK_BOOK_ID) REFERENCES BOOK (BOOK_ID),
    FOREIGN KEY (FK_FROM_USER_ID) REFERENCES USER (USER_ID),
    FOREIGN KEY (FK_TO_USER_ID) REFERENCES USER (USER_ID)
);