-- original table name in sqlite schema was sqlite_sequence, rest of def is identical
create table repro(id, name, seq);

-- name: GetRepro :one
select * from repro where id = ? limit 1;