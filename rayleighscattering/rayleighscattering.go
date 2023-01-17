package main

import (
	"fmt"
	"math"
)

func rayleighScattering(wavelength, particleDiameter, refractiveIndex, numberOfParticles float64) (float64, float64) {
	// Calculate k using the approximation for small particles
	k_small := (3 * math.Pi * refractiveIndex * math.Pow(particleDiameter, 2)) / (2 * numberOfParticles)

	// Alternatively, calculate k using the approximation for large particles
	k_large := (3 * math.Pi * math.Pow(refractiveIndex, 2) * math.Pow(particleDiameter, 6)) / (2 * numberOfParticles)

	// calculate the intensity of the scattered light
	intensity_small := k_small * math.Pow(wavelength, -4)
	intensity_large := k_large * math.Pow(wavelength, -4)

	return intensity_small, intensity_large
}

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Rayleigh Scattering")

	/*
		The Rayleigh scattering law is a formula used to calculate the scattering of light by particles in a medium. It states that the scattering intensity is proportional to the fourth power of the wavelength of the light. In other words, shorter wavelengths (such as blue or violet) are scattered more intensely than longer wavelengths (such as red).

		The formula for the Rayleigh scattering law is:

		I = k * λ^-4

			where:

			I .. is the intensity of the scattered light
			k .. is a constant that depends on the size and composition of the particles
			λ .. is the wavelength of the light

		This is code snippet that calculates the scattering for a given wavelength.
	*/

	// wavelength of the light, in nanometers
	wavelength := 500.0

	/*
		The constant k depends on the size and composition of the particles that are scattering the light. It is generally difficult to calculate k accurately, as it requires knowledge of the exact size and composition of the particles. However, the following approximations that can be made:

		(1) If the particles are much smaller than the wavelength of the light, then k can be approximated as: k = (3 * π * n * d^2) / (2 * N)
		(2) If the particles are much larger than the wavelength of the light, then k can be approximated as: k = (3 * π * n^2 * d^6) / (2 * N)
			n .. is the refractive index of the particles
			d .. is the diameter of the particles
			N .. is the number of particles per unit volume
	*/

	// diameter of the particles, in nanometers
	particleDiameter := 100.0

	// refractive index of the particles
	refractiveIndex := 1.5

	// number of particles per unit volume
	numberOfParticles := 1e9

	small, large := rayleighScattering(wavelength, particleDiameter, refractiveIndex, numberOfParticles)

	fmt.Println("Wavelength in nanometer:", wavelength)
	fmt.Println("Intensity of scattered light for small particles:", small)
	fmt.Println("Intensity of scattered light for large particles:", large)
}
