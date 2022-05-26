-- CREATE TABLE clients(
-- 	id SERIAL PRIMARY KEY NOT NULL,
--     nome VARCHAR(100) NOT NULL,
--     documento VARCHAR(14) NOT NULL,
--     tipo VARCHAR(15) NOT NULL,
--     email VARCHAR(100) NOT NULL,
--     telefone VARCHAR(14)    
-- );

-- CREATE TABLE accounts(
-- 	id SERIAL PRIMARY KEY NOT NULL,
--     numero VARCHAR(5) NOT NULL,
--     agencia VARCHAR(4) NOT NULL,
--     saldo INT NOT NULL,
--     cliente_id INT  NOT NULL,
--     FOREIGN KEY(cliente_id) REFERENCES clients(id)
-- );

-- CREATE TABLE operations(
-- 	id SERIAL NOT NULL PRIMARY KEY,
--     conta_origem INT NOT NULL,
--     conta_destino INT,
--     valor INT NOT NULL,
--     taxa INT NOT NULL,
--     tipo VARCHAR(15) NOT NULL,
--     FOREIGN KEY (conta_origem) REFERENCES accounts(id),
--     FOREIGN KEY (conta_destino) REFERENCES accounts(id)
-- );

create table Clients(
    Id serial,
    Name varchar,
    Document varchar,
    Email varchar,
    PRIMARY KEY (Id)
);

create table Accounts(
    Id serial,
    Client_Id integer,
    Branch_Number varchar,
    Account_Number varchar,
    Balance bigint,
    PRIMARY KEY (Id),
    FOREIGN KEY (Client_Id) REFERENCES Client (Id)
);

create table Operations(
    Id serial,
    Type_Operation varchar,
    Origin_Account_Id integer,
    Destination_Account_Id integer,
    Amount bigint,
    Service_Charge integer,
    Date_Operation date,
    PRIMARY KEY (Id),
    FOREIGN KEY (Origin_Account_Id) REFERENCES Account (Id),
    FOREIGN KEY (Destination_Account_Id) REFERENCES Account (Id)
);
