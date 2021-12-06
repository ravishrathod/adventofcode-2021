package main

type LanternFish struct {
	counter int
}

func CreateFish(initCounter int) *LanternFish {
	return &LanternFish{
		counter: initCounter,
	}
}

func (this *LanternFish) DayPassed() *LanternFish {
	if this.counter == 0 {
		this.counter = 6
		return &LanternFish{
			counter: 8,
		}
	} else {
		this.counter--
		return nil
	}
}

func (this *LanternFish) GetCounter() int {
	return this.counter
}
