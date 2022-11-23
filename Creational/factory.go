package Creational

import (
	"errors"
	"fmt"
	"log"
)

type RobotType string

const (
	teacherType RobotType = "teacherRobot"
	fighterType RobotType = "fighterRobot"
)

type imRobot interface {
	whoAmi() string
}

type robot struct {
	name     string
	rType    RobotType
	autonomy float32
}

func (r robot) whoAmi() string {
	return fmt.Sprintf("Name: %s\n"+
		"Kind: %s\n"+
		"Autonomy: %2.2f",
		r.name,
		r.rType,
		r.autonomy)
}

type teacherRobot struct {
	robot
}

func newTeacherRobot(name string, autonomy float32) teacherRobot {
	return teacherRobot{
		robot: robot{
			name:     name,
			rType:    teacherType,
			autonomy: autonomy,
		},
	}
}

type fighterRobot struct {
	robot
}

func newfFighterRobot(name string, autonomy float32) fighterRobot {
	return fighterRobot{
		robot: robot{
			name:     name,
			rType:    fighterType,
			autonomy: autonomy,
		},
	}
}

func createRobot(rType RobotType) (imRobot, error) {
	if rType == teacherType {
		return newTeacherRobot("Teacher", 90), nil
	}

	if rType == fighterType {
		return newfFighterRobot("Fighter", 70), nil
	}

	return nil, errors.New(fmt.Sprintf("could not create a robot with: %s", rType))
}

func start() {
	robotA, err := createRobot(teacherType)

	if err != nil {
		log.Fatalln(err)
	}

	robotB, err := createRobot(fighterType)

	if err != nil {
		log.Fatalln(err)
	}

	wrongRobot, err := createRobot("invalid")

	fmt.Println(robotA.whoAmi())
	fmt.Println("-------------------")
	fmt.Println(robotB.whoAmi())
	fmt.Println("------------------")

	fmt.Println("wrongRobot: ", wrongRobot)
	fmt.Println("wrongRobot err: :", err.Error())
}
