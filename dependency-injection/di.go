package main

import "fmt"

// for multiple rocks with difeernet safety, we migh store kind field
// and then switch over it in placeSafety
// 2. create a new struct and factor out the switch
type RockClimber struct {
	// kind int
	rocksClimbed int
	sp Safety
}

// 2. 
type Safety struct {
	kind	int
}

func (sp *Safety) placeSafety() {
	switch sp.kind {
		// ...
	}
}

func (rc *RockClimber) climbRocks() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		// rc.placeSafety()
		rc.sp.placeSafety()
	}
}

func (rc *RockClimber) placeSafety() {
	fmt.Println("placing safety on rock climber")
}

func di() {
	rc := &RockClimber{}
	for i :=0; i < 11 ; i++ {
		rc.climbRocks()
	}
}

// 3. Use interface
type SafetyPlacer interface {
	placeSafety()
}

type NOPSafetyPlacer struct {}

func (sp NOPSafetyPlacer) placeSafety() {
	fmt.Println("Nop safety placer....")
}

type ICESafetyPlacer struct {}

func (sp ICESafetyPlacer) placeSafety() {
	fmt.Println("Ice safety placer....")
}

type RockClimber2 struct {
	rocksClimbed int
	sp SafetyPlacer
}

func (rc *RockClimber2) climbRocks() {
	rc.rocksClimbed++
	if rc.rocksClimbed == 10 {
		// rc.placeSafety()
		rc.sp.placeSafety()
	}
}

func newRockClimber(sp SafetyPlacer) *RockClimber2 {
	return &RockClimber2{
		sp: sp,
	}
}

func di2() {
	sp := NOPSafetyPlacer{}
	rc := newRockClimber(sp)
	for i :=0; i < 11 ; i++ {
		rc.climbRocks()
	}
}

func di3() {
	sp := ICESafetyPlacer{}
	rc := newRockClimber(sp)
	for i :=0; i < 11 ; i++ {
		rc.climbRocks()
	}
}