package api

import (
	"encoding/json"
	"net/http"

	"skipr/internal/usecases/lineup"
	"slices"
)

var players []lineup.Player
var playerId int = 1

func PlayerHandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		json.NewEncoder(w).Encode(players)
	case http.MethodPost:
		var p lineup.Player
		if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}
        p.Id = playerId
        playerId++
		players = append(players, p)
		w.WriteHeader(http.StatusCreated)
    case http.MethodDelete:
        var p lineup.Player
        if err := json.NewDecoder(r.Body).Decode(&p); err != nil {
            http.Error(w, err.Error(), http.StatusBadRequest)
            return
        }

        for i, player := range players {
            if player.Id == p.Id {
                players = slices.Delete(players, i, i+1)
                w.WriteHeader(http.StatusNoContent)
                return
            }
        }

        http.Error(w, "Player not found", http.StatusNotFound)

    default:
        http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func PlayersHandler(w http.ResponseWriter, r *http.Request) {
    var p lineup.Player
    p.Id = playerId
    p.Name = "Donny"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Landon"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Noah"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Caleb"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Nolan"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Michael"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Bennet"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Nate"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Carter"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Tommy"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Daniel"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Rocky"
    playerId++
    players = append(players, p)

    p.Id = playerId
    p.Name = "Braydon"
    playerId++
    players = append(players, p)
    w.WriteHeader(http.StatusCreated)
}

func LineupHandler(w http.ResponseWriter, r *http.Request){
	if r.Method != http.MethodPost {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	plan, err := lineup.GenerateLineup(players, 6)	
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	
	json.NewEncoder(w).Encode(plan)
}
