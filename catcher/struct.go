package catcher

type Sample struct {
	name string
	age  int
}

func (sample *Sample) GetName() string {
	return sample.name
}

func (sample *Sample) GetAge() int {
	return sample.age
}

func (sample *Sample) SetName(name string) *Sample {
	sample.name = name
	return sample
}

func (sample *Sample) SetAge(age int) *Sample {
	sample.age = age
	return sample
}
