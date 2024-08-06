-- +goose Up
INSERT INTO counter (id, count) VALUES (1, 0);

-- +goose Down
DELETE FROM counter WHERE id = 1;
