package api

type Datatier interface {
	Name() string
	Alias() string
}

type DatatierAlias struct {
	Branch   string
	Datarier string
}

var DatatierMatrix = make([]DatatierAlias, 0)

func DatatierAliasFromFlag(flag string) (DatatierAlias, error) {
	alias := DatatierAlias{}

	return alias, nil
}

func DatatierMatrixFromFlag(flag string) ([]DatatierAlias, error) {
	matrix := make([]DatatierAlias, 0)

	DatatierMatrix = matrix

	return matrix, nil
}
