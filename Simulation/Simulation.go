package Simulation

import (
	"fmt"
	"github.com/go-gl/mathgl/mgl32"
	"math/rand"
	"os"
	"strconv"
)

const (
	particleCount       = 4
	postionBounds       = 10
	startVelocityBounds = 0

	g                              = 1
	collisionDistance              = 1
	doCollision                    = false
	collisionElasticEnergy float32 = 0.5
)

var outFilePath string
var file os.File

var particles []particle
var frameCount int

func SetUpSimulation(_frameCount int, absPath string) {

	frameCount = _frameCount
	outFilePath = absPath + "/simulationData.txt"

	newfile, err := os.Create(outFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	file = *newfile
	file.WriteString("info " + strconv.Itoa(particleCount) + " " + strconv.Itoa(frameCount) + "\n")

	particles = make([]particle, particleCount)

	for i := 0; i < particleCount; i++ {
		particle := particle{
			position: mgl32.Vec3{
				(rand.Float32()*2 - 1) * postionBounds,
				(rand.Float32()*2 - 1) * postionBounds,
				(rand.Float32()*2 - 1) * postionBounds,
			},
			velocity: mgl32.Vec3{
				(rand.Float32()*2 - 1) * startVelocityBounds,
				(rand.Float32()*2 - 1) * startVelocityBounds,
				(rand.Float32()*2 - 1) * startVelocityBounds,
			},
			mass: 1,
		}

		particles[i] = particle
	}
}

func UpdateSimulation(frame int) {
	file.WriteString("f " + strconv.Itoa(frame) + "\n")

	for i, currentParticle := range particles {
		// Console Print
		fmt.Printf("Calculating Particle %d of %d in Frame %d of %d \r", i, len(particles), frame, frameCount)

		currentParticle.setUpForNewFrame()

		currentParticle.applyGravityForce()

		currentParticle.applyForcesToVelocity()

		particles[i] = currentParticle
	}

	for i, currentParticle := range particles {

		currentParticle.applyVelocityToPosition()

		// Writing pos to file
		file.WriteString("p " + strconv.FormatInt(int64(i), 10) + " " +
			strconv.FormatFloat(float64(currentParticle.position[0]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(currentParticle.position[1]), 'f', -1, 64) + " " +
			strconv.FormatFloat(float64(currentParticle.position[2]), 'f', -1, 64) + "\n")

		particles[i] = currentParticle
	}
}

func EndSimulation() {
	file.Close()
}
