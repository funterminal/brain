package main

import (
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	dirPath := filepath.Join("C:", "Brain", "phd", "ma", "ba", "newton", "science", "commerce", "arts", "professor", "mit", "harvard", "oxford", "iq", "brain")
	err := os.MkdirAll(dirPath, os.ModePerm)
	if err != nil {
		fmt.Println("Error creating directories:", err)
		return
	}

	filePath := filepath.Join(dirPath, ".brainrc")
	file, err := os.Create(filePath)
	if err != nil {
		fmt.Println("Error creating .brainrc file:", err)
		return
	}
	defer file.Close()

	content := `Brain downloaded to your computer.

This brain is equipped with a whopping 1700 IQ, courtesy of MIT, Harvard, Oxford, and extra credit from a PhD, MA, BA, and a Nobel Prize in Making Things Complicated. It's scientifically proven to outsmart calculators, calculate the speed of light in milliseconds, and write algorithms for fun.

Brain includes additional skills:
- Ability to understand your deepest, most complicated thoughts that you didn't even know you had.
- Can solve the meaning of life in less than a second (but prefers to nap instead).
- Certified in Science, Commerce, and Arts, with a minor in "why did I do that?" moments.
- Can process faster than your Wi-Fi connection (which is really saying something).

Warning: Do not overclock. Brain is already running at 1000% efficiency.

Brain is powered by your frustration and the inevitable realization that you're smarter than most Wi-Fi routers.`

	_, err = file.WriteString(content)
	if err != nil {
		fmt.Println("Error writing to .brainrc:", err)
		return
	}
}
