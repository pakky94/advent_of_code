package day06

func Run() {
	testRaces := []Race{
		{
			time:     7,
			distance: 9,
		},
		{
			time:     15,
			distance: 40,
		},
		{
			time:     30,
			distance: 200,
		},
	}
	_ = testRaces

	races := []Race{
		{
			time:     42,
			distance: 308,
		},
		{
			time:     89,
			distance: 1170,
		},
		{
			time:     91,
			distance: 1291,
		},
		{
			time:     89,
			distance: 1467,
		},
	}
	_ = races

	p1(races)
	p2(Race{
		time:     42899189,
		distance: 308117012911467,
		//time:     71530,
		//distance: 940200,
	})
}

func p2(race Race) {
	println("p2", winChances(race))
}

func p1(races []Race) {
	res := 1
	for _, race := range races {
		res *= winChances(race)
	}
	println("p1", res)
}

func winChances(race Race) int {
	count := 0
	for i := 1; i <= race.time; i++ {
		if i*(race.time-i) > race.distance {
			count += 1
		}
	}
	return count
}

type Race struct {
	time     int
	distance int
}

var testContent = `Time:      7  15   30
Distance:  9  40  200`
