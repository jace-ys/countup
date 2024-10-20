package idgen

import (
	"strings"

	"github.com/jackc/pgx/v5/pgtype"
	"github.com/segmentio/ksuid"
)

type ID[T Resource] struct {
	UID ksuid.KSUID
}

func New[T Resource]() ID[T] {
	return ID[T]{UID: ksuid.New()}
}

func (id ID[T]) String() string {
	var resource T
	return resource.Prefix() + "_" + id.UID.String()
}

func FromString[T Resource](id string) (ID[T], error) {
	var resource T
	prefix := resource.Prefix() + "_"

	uid, err := ksuid.Parse(strings.TrimPrefix(id, prefix))
	if err != nil {
		return ID[T]{ksuid.Nil}, err //nolint:wrapcheck
	}

	return ID[T]{uid}, nil
}

var _ pgtype.TextValuer = (*ID[resource])(nil)

func (id ID[T]) TextValue() (pgtype.Text, error) {
	return pgtype.Text{
		String: id.UID.String(),
		Valid:  true,
	}, nil
}

var _ pgtype.TextScanner = (*ID[resource])(nil)

func (id *ID[T]) ScanText(v pgtype.Text) error {
	return id.UID.Scan(v.String) //nolint:wrapcheck
}
