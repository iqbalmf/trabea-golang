create
database trabeadb;
use trabeadb;
create table sample
(
    id   varchar(100) not null,
    name varchar(100) not null,
    primary key (id)
) engine = InnoDB;
select  * from sample;
desc sample;

