
CREATE TABLE segments(
    segment_name varchar(256) primary key
);

CREATE TABLE users(
    user_id integer,
    segment_name varchar(256) references segments (segment_name) on delete cascade,
    added date,
    expiry date default date('01-01-3023'),
    CONSTRAINT pk_users primary key (user_id, segment_name, added)
);

CREATE TABLE operations(
    user_id integer,
    segment_name varchar(256),
    operation_type varchar(256),
    operation_date timestamp
);


CREATE OR REPLACE FUNCTION process_users_audit() RETURNS TRIGGER AS $emp_audit$
    BEGIN
    --
    -- Добавление строки в emp_audit, которая отражает операцию, выполняемую в emp;
    -- для определения типа операции применяется специальная переменная TG_OP.
    --
    IF (TG_OP = 'DELETE') THEN
        INSERT INTO operations SELECT OLD.user_id, OLD.segment_name, 'DELETE', NOW();
        RETURN OLD;
    ELSIF (TG_OP = 'INSERT') THEN
        INSERT INTO operations SELECT NEW.user_id, NEW.segment_name, 'ADD', NOW();
        RETURN NEW;
    END IF;
    RETURN NULL; -- возвращаемое значение для триггера AFTER игнорируется
    END;
$emp_audit$ LANGUAGE plpgsql;

CREATE TRIGGER users_audit
    AFTER INSERT OR UPDATE OR DELETE ON users
    FOR EACH ROW EXECUTE PROCEDURE process_users_audit();

--При истечении срока действия
CREATE FUNCTION check_expired_users() RETURNS void AS $$
BEGIN
    INSERT INTO operations (user_id, segment_name, operation_type, operation_date)
    SELECT user_id, segment_name , 'удаление', NOW() --localtime?
    FROM users
    WHERE users.expiry <= CURRENT_DATE;

    DELETE FROM users
    WHERE users.expiry <= CURRENT_DATE;
END;
$$ LANGUAGE plpgsql;
