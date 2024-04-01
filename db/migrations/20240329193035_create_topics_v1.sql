-- +goose Up
-- +goose StatementBegin
create table if not exists topics(
    id bigserial primary key,
    userid integer not null,
    name text not null,
    subid integer,
	constraint fk_topics foreign key (subid) references topics (id)
);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin

drop table topics;
-- +goose StatementEnd
