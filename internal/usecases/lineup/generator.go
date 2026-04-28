package lineup

import (
	"errors"
	"math/rand"
	"slices"
	"time"
    "sort"
)

func GenerateLineup(players []Player, innings int) (Lineup, error){
    // Not sure if we will keep this as we can't control how many kids make it
    if len(players) < 9 {
        return Lineup{}, errors.New("not enough players")
    }

    // Initialize the game plan
    lineup := Lineup{
        Innings: innings,
        Players: players,
        Defense: make(map[int][]Assignment),
        BattingOrder: generateBattingOrder(players),
    }

    playerAssignments := make(map[int][]Assignment)

    // Assign field positions for each inning
    for inning := 1; inning <= innings; inning++ {
        assignments, err := assignPositionsForInning(players, inning, playerAssignments)
        if err != nil {
            return Lineup{}, err
        }
        lineup.Defense[inning] = assignments

        for _, a := range assignments {
            playerAssignments[a.PlayerId] = append(playerAssignments[a.PlayerId], a)
        }
    }

    // Validate game plan against rules
    // if err := ValidateLineup(lineup, players); err != nil {
    //     return Lineup{}, err
    // }

    return lineup, nil
}

func generateBattingOrder(players []Player) []Player {
    r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := make([]Player, len(players))
	copy(shuffled, players)

    r.Shuffle(len(shuffled), func(i, j int) {
        shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
    })

    return shuffled
}

var infieldPositions = []Position{"P", "C", "1B", "2B", "SS", "3B"}
var outfieldPositions = []Position{"LF", "LCF", "RCF", "RF"}
var benchPosition Position = Bench

func assignPositionsForInning(players []Player, inning int, history map[int][]Assignment) ([]Assignment, error){
	assignments := []Assignment{}
	usedPositions := map[Position]bool{}
	infieldCount := 0
	outfieldCount := 0

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	shuffled := make([]Player, len(players))
	copy(shuffled, players)
	r.Shuffle(len(shuffled), func(i, j int){
		shuffled[i], shuffled[j] = shuffled[j], shuffled[i]
	})

	for _, player := range shuffled {
		// get players past assignments
		past := history[player.Id]
		positionCounts := make(map[Position]int)
		var lastPosition Position

		for _, a := range past {
			positionCounts[a.Position]++
			if a.Inning == inning-1 {
				lastPosition = a.Position
			}
		}

		var assigned Position


		if assigned == "" && infieldCount <= 7 {
            // Fill in all the Infield Positions First
			for _, pos := range infieldPositions {
				if !usedPositions[pos] && positionCounts[pos] < 2 {
					assigned = pos
					infieldCount++
					break
				}
			}
		}

        if assigned == "" && outfieldCount <= 5 {
            for _, pos := range outfieldPositions {
                // Rules for assigning outfield positions: 
                // 1. Position must not be used already
                // 2. Player must not have played the same position more than twice in the current game
                // 3. Player's last position must not be in the outfield
				if !usedPositions[pos] && positionCounts[pos] < 2 && !slices.Contains(outfieldPositions, lastPosition) {
					assigned = pos
					outfieldCount++
					break
				}
			}
		}

        // If no position was assigned and all 12 positions have been accounted for, we will assign a bench position
		if assigned == "" && lastPosition != benchPosition && len(usedPositions) > 12 { 
			assigned = benchPosition
		}

        // If no position was assigned and there are still positions available, we will assign an outfield position
		if assigned == "" {
            if outfieldCount <= 5 {
                leftOver := outfieldPositions[outfieldCount%len(outfieldPositions)]
                if !usedPositions[leftOver] && !slices.Contains(outfieldPositions, lastPosition) {
                    assigned = leftOver
                    outfieldCount++
                }
            }
		}

        if assigned == "" {
            assigned = benchPosition
        }

		usedPositions[assigned] = true

		assignments = append(assignments, Assignment{
			PlayerId: player.Id,
            PlayerName: player.Name,
			Inning: inning,
			Position: assigned,
		})
	}
    
    sort.SliceStable(assignments, func(i, j int) bool {
        return positionPriority[assignments[i].Position] < positionPriority[assignments[j].Position]
    })

	return assignments, nil
}

var positionPriority = map[Position]int{
	"P": 1,
	"C": 2,
	"1B": 3,
	"2B": 4,
	"3B": 5,
	"SS": 6,
	"LF": 7,
	"CF": 8,
	"RF": 9,
	"LFC": 10,
	"RCF": 11,
	"SF": 12,
}

// func sortDefenseAssignments(assignments []Assignment) []Assignment{
//     sort.SliceStable(assignments, func(i, j int) bool {
//         return positionPriority[assignments[i].Position] < positionPriority[assignments[j].Position]
//     })
// }

