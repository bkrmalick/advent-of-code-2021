package main

type Fish struct {
	timer int
}

// passes one day and returns a reference to a new fish
// if this fish spawned a new one
func (f *Fish) passDay() *Fish{
	if f.timer == 0{
		f.timer = 6
		return &Fish{8}

	} else {
		f.timer--
		return nil
	}
}
