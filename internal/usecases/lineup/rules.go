package lineup

import (
	"fmt"
)

func ValidateLineup(lineup Lineup, players []Player) error {
	playerMap := map[int]Player{}
	for _, p := range players {
		playerMap[p.Id] = p
	}

	playerPositionCount := map[int]map[Position]int{}
	lastPositions := map[int]Position{}
	consecOutOrBench := map[int]int{}
	infieldPlayed := map[int]bool{}

	for inning := 1; inning <= 6; inning++ {
		infielders := 0
		outfielders := 0

		assignments := lineup.Defense[inning]
		for _, a := range assignments {
			playerId := a.PlayerId
			pos := a.Position

			if _, ok := playerPositionCount[playerId]; !ok {
				playerPositionCount[playerId] = map[Position]int{}
			}
			playerPositionCount[playerId][pos]++

			if playerPositionCount[playerId][pos] > 2 {
				return fmt.Errorf("player %s played %s more than 2 times", playerMap[playerId].Name, pos)
			}

			if isInfield(pos){
				infielders++
				infieldPlayed[playerId] = true
			} else if isOutfield(pos) {
				outfielders++
			}

			last := lastPositions[playerId]
			if(isOutfield(pos) || isBench(pos)) && (isOutfield(last) || isBench(last)){
				consecOutOrBench[playerId]++
				if(consecOutOrBench[playerId] >= 2) {
					return fmt.Errorf("player %s has played Outfield or Bench last inning", playerMap[playerId].Name)
				}
			} else {
				consecOutOrBench[playerId] = 0
			}
			lastPositions[playerId] = last
		}

		if infielders > 7 {
			return fmt.Errorf("too many infielders in inning %d: %d", inning, infielders)
		}

		if outfielders > 5 {
			return fmt.Errorf("too many outfielders in inning %d: %d", inning, outfielders)
		}
	}

	for _, p := range players {
		if !infieldPlayed[p.Id] {
			return fmt.Errorf("%s did not play the infield at least once",p.Name)
		}
	}
	return nil
}

func isInfield(pos Position) bool {
	infields := map[Position]bool{"P": true, "C": true, "1B": true, "2B": true, "SS": true, "3B": true}
	return infields[pos]
}

func isOutfield(pos Position) bool {
	outfields := map[Position]bool{"LF": true, "RF": true, "LCF": true, "RCF": true}
	return outfields[pos]
}

func isBench(pos Position) bool {
	bench := map[Position]bool{"Bench": true}
	return bench[pos]
}
