package ai

import "battleships-board/types"

func calculateMove(b types.Board) types.Coord {
	return types.Coord{}
}

func remainingShips() types.Ships {
	return types.Ships{}
}

func remainingMovies() []types.Coord {
	return []types.Coord{}
}

/*
// if no hits
	// determine coord with largest surrounding area
// if hits
	// loop over hits determine probability
		// determine possible hits in surrounding coords
			// depends on remaining ships
			// ships could be touching
			// bounds of map prevent
*/
