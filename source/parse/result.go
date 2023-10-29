package parse

import (
	"bufio"
	"lem-in/source/errors"
	graph "lem-in/source/graph"
)

func GetResult(scanner *bufio.Scanner) (*graph.Result, error) {
	terrain := graph.CreateGraph()
	var err error
	for scanner.Scan() {
		err = terrain.ReadDataFromLine(scanner.Text())
		if err != nil {
			return nil, errors.ErrInvalidDataFormat(err)
		}
	}
	err = terrain.ValidateByFieldInfo()
	if err != nil {
		return nil, errors.ErrInvalidDataFormat(err)
	}
	err = terrain.Match()
	if err != nil {
		return nil, errors.ErrPaths(err)
	}
	return terrain.Result, nil
}
