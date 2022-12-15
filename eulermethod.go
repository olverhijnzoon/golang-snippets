package main

import (
	"fmt"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

// First-order linear ordinary differential equation to solve: y' = x + y
func dy_linearordinary(x, y float64) float64 { return x + y }

// Harmonic oscillator differential equation to solve: y' = -k * y/m
func dy_harmonicoscillator(x, y, dy, k, m float64) float64 { return -k * y / m }

func main() {

	fmt.Println("# Golang Snippets")
	fmt.Println("## Euler Method")

	/*
		In the example, we are solving the differential equation y' = x + y with initial conditions x0 = 0 and y0 = 1. The Euler method is a simple numerical method for approximating solutions to differential equations, and it works by taking small steps (in this case, with step size h = 0.1) and using the derivative at the current point to approximate the value of the function at the next point. There is also a second example, in which we are the harmonic oscillator equation y' = -k * y/m.

		For clarification about the harmonic oscillator representations: To derive the second equation, y'' + k/m * y = 0, from the first equation, y' = -k * y/m, we can take the derivative of both sides of the first equation with respect to time. This yields the following:

		d/dt(y') = d/dt(-k * y/m)
		y'' = -k/m * y'

		Solving for y'' in the above equation yields y'' = -k/m * y'. We can then substitute this expression for y'' into the second equation, which gives us:

		y'' + k/m * y = -k/m * y' + k/m * y
		0 = -k/m * y'

		Since the right-hand side of the equation is equal to zero, the equation is satisfied for any value of y'. This means that the second equation, y'' + k/m * y = 0, can be derived from the first equation, y' = -k * y/m, by taking the derivative of both sides of the first equation with respect to time and solving for y''.

		A nice follow-up to the Euler method would be to explore other numerical methods for solving differential equations, such as the Runge-Kutta method or the Adams-Bashforth method. Furthermore, exploring the use of advanced variants of the method, such as the Runge-Kutta-Fehlberg method or the Runge-Kutta-Chebyshev method, can improve its efficiency and accuracy in solving certain types of equations.
	*/

	// Initial conditions
	x0 := 0.0
	y0 := 1.0
	h := 0.1 // step size

	// Create a new plot
	p := plot.New()
	p.Title.Text = "Solution to y' = x + y"
	p.X.Label.Text = "x"
	p.Y.Label.Text = "y"
	pts := make(plotter.XYs, 11)
	i := 0
	pts[0].X = x0
	pts[0].Y = y0

	// Euler method
	for x := x0; x < 1; x += h {
		y0 += h * dy_linearordinary(x, y0)
		fmt.Printf("x=%f, y=%f\n", x, y0)
		pts[i].X = x
		pts[i].Y = y0
		i++
	}

	// Add the points to the plot
	line, err := plotter.NewLine(pts)
	if err != nil {
		panic(err)
	}
	line.LineStyle.Width = vg.Points(1)
	p.Add(line)

	// Save the plot to a file
	if err := p.Save(4*vg.Inch, 4*vg.Inch, "euler_olde.png"); err != nil {
		panic(err)
	}

	// Initial conditions - harmonic oscillator
	x0 = 0.0
	y0 = 1.0   // initial displacement
	dy0 := 0.0 // initial velocity
	k := 1.0   // spring constant
	m := 1.0   // mass of oscillating object
	h = 0.1    // step size

	// Create a new plot
	p2 := plot.New()
	p2.Title.Text = "Solution to y' = -k * y/m"
	p2.X.Label.Text = "x"
	p2.Y.Label.Text = "y"
	pts2 := make(plotter.XYs, 101)
	i = 0
	pts2[0].X = x0
	pts2[0].Y = y0

	// Euler method
	for x := x0; x < 10; x += h {
		y0 += h * dy0
		dy0 += h * dy_harmonicoscillator(x, y0, dy0, k, m)
		fmt.Printf("x=%f, y=%f, dy=%f\n", x, y0, dy0)
		pts2[i].X = x
		pts2[i].Y = y0
		i++
	}

	// Add the points to the plot
	line2, err := plotter.NewLine(pts2)
	if err != nil {
		panic(err)
	}
	line2.LineStyle.Width = vg.Points(1)
	p2.Add(line2)

	// Save the plot to a file
	if err := p2.Save(4*vg.Inch, 4*vg.Inch, "euler_osci.png"); err != nil {
		panic(err)
	}

}
