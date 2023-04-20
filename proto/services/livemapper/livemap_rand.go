package livemapper

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/galexrt/fivenet/query/fivenet/model"
	"github.com/galexrt/fivenet/query/fivenet/table"
	jet "github.com/go-jet/jet/v2/mysql"
	"go.uber.org/zap"
)

func (s *Server) GenerateRandomUserMarker() {
	userIdentifiers := []string{
		"char1:fcee377a1fda007a8d2cc764a0a272e04d8c5d57",
		"char1:0ff2f772f2527a0626cac48670cbc20ddbdc09fb",
		"char2:d9793ddb457316fb3951d1b1092526183270a307",
		"char2:d7abbfba01625bec803788ee42da86461c96e0bd",
		"char1:ad4fb9f44bb784dd30effcc743a9c169db4d625d",
	}

	markers := make([]*model.FivenetUserLocations, len(userIdentifiers))

	resetMarkers := func() {
		xMin := -3300
		xMax := 4300
		yMin := -3300
		yMax := 5000
		for i := 0; i < len(markers); i++ {
			x := float64(rand.Intn(xMax-xMin+1) + xMin)
			y := float64(rand.Intn(yMax-yMin+1) + yMin)

			job := "ambulance"
			hidden := false
			markers[i] = &model.FivenetUserLocations{
				Identifier: userIdentifiers[i],
				Job:        &job,
				Hidden:     &hidden,

				X: &x,
				Y: &y,
			}
		}
	}

	moveMarkers := func() {
		xMin := -100
		xMax := 100
		yMin := -100
		yMax := 100

		for i := 0; i < len(markers); i++ {
			curX := *markers[i].X
			curY := *markers[i].Y

			newX := curX + float64(rand.Intn(xMax-xMin+1)+xMin)
			newY := curY + float64(rand.Intn(yMax-yMin+1)+yMin)

			markers[i].X = &newX
			markers[i].Y = &newY
		}
	}

	resetMarkers()

	counter := 0
	for {
		if counter >= 15 {
			resetMarkers()
			counter = 0
		} else {
			moveMarkers()
		}

		stmt := locs.
			INSERT(
				locs.Identifier,
				locs.Job,
				locs.X,
				locs.Y,
				locs.Hidden,
			).
			MODELS(markers).
			ON_DUPLICATE_KEY_UPDATE(
				locs.X.SET(jet.RawFloat("VALUES(x)")),
				locs.Y.SET(jet.RawFloat("VALUES(y)")),
			)

		_, err := stmt.Exec(s.db)
		if err != nil {
			s.logger.Error("failed to insert/ update random location to locations table", zap.Error(err))
		}

		counter++
		time.Sleep(3 * time.Second)
	}
}

func (s *Server) GenerateRandomDispatchMarker() {
	userIdentifiers := []string{
		"char1:fcee377a1fda007a8d2cc764a0a272e04d8c5d57",
		"char1:0ff2f772f2527a0626cac48670cbc20ddbdc09fb",
		"char2:d9793ddb457316fb3951d1b1092526183270a307",
		"char2:d7abbfba01625bec803788ee42da86461c96e0bd",
		"char1:ad4fb9f44bb784dd30effcc743a9c169db4d625d",
	}

	markers := make([]*model.GksphoneJobMessage, len(userIdentifiers))

	jMessage := table.GksphoneJobMessage
	resetMarkers := func() {
		xMin := -3300
		xMax := 4300
		yMin := -3300
		yMax := 5000
		rSrc := rand.NewSource(time.Now().UnixNano())
		ran := rand.New(rSrc)
		for i := 0; i < len(markers); i++ {
			x := float64(rand.Intn(xMax-xMin+1) + xMin)
			y := float64(rand.Intn(yMax-yMin+1) + yMin)

			name := fmt.Sprintf("TEST %d", i)
			number := "01234567"
			message := fmt.Sprintf("Message of Dispatch %d", i)
			photo := ""
			gps := fmt.Sprintf("GPS: %f, %f", x, y)
			jobm := "[\"ambulance\"]"

			var owner int32
			if ran.Float64() > 0.5 {
				owner = 1
			} else {
				owner = 0
			}

			t := time.Now()
			anon := "0"
			markers[i] = &model.GksphoneJobMessage{
				Name:    &name,
				Number:  &number,
				Message: &message,
				Photo:   &photo,
				Gps:     &gps,
				Owner:   owner,
				Jobm:    &jobm,
				Time:    t,
				Anon:    &anon,
			}
		}

		stmt := jMessage.
			INSERT(
				jMessage.Name,
				jMessage.Number,
				jMessage.Message,
				jMessage.Photo,
				jMessage.Gps,
				jMessage.Owner,
				jMessage.Jobm,
				jMessage.Time,
				jMessage.Anon,
			).
			MODELS(markers)

		_, err := stmt.Exec(s.db)
		if err != nil {
			s.logger.Error("failed to insert random dispatch", zap.Error(err))
		}
	}

	resetMarkers()

	counter := 0
	for {
		if counter >= 20 {
			resetMarkers()
			counter = 0
		}

		counter++
		time.Sleep(3 * time.Second)
	}
}
