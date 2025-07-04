-- name: ListUsers :many
select * from users;

-- name: GetUser :one
select * from users where id = ?;

-- name: CreateUser :exec
insert into users (id,name,email,password,created_at,updated_at) values (?,?,?,?,?,?);

-- name: UpdateUser :exec
update users set name =?, email = ?, password = ?, updated_at = ? where id =?;

-- name: DeleteUser :exec
delete from users where id =?;


-- name: ListEventTypes :many
select * from event_types;

-- name: GetEventType :one
select * from event_types where id = ?;

-- name: CreateEventType :exec
insert into event_types (id,code,description,created_at,updated_at) values (?,?,?,?,?);

-- name: UpdateEventType :exec
update event_types set code =?, description =?, updated_at = ? where id =?;

-- name: DeleteEventType :exec
delete from event_types where id =?;


-- name: ListEvents :many
select * from events;

-- name: GetEvent :one
select * from events where id = ?;

-- name: CreateEvent :exec
insert into events (id,event_type_id,user_id, created_at) values (?,?,?,?);

-- name: UpdateEvent :exec
update events set event_type_id=?,user_id=? where id =?;

-- name: DeleteEvent :exec
delete from events where id =?;


