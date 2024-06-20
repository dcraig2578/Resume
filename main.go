package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	var totCa float64
	var totJo float64
	run := 0
	greeting()
	for i := 0; ; i++ {
		getMeasurements(&totJo, &totCa, run)
		run++
	}
}
func greeting() {
	var Reset = "\033[0m"
	var Green = "\033[32m"
	fmt.Print(Green)
	fmt.Println("Hello and welcome to your personal virtual calories coach.")
	time.Sleep(1 * time.Second)
	fmt.Println("I'm here to help you figure out roughly how many calories you burn per exercise selected.")
	fmt.Println("Calories burned will be dependent on your anthropometrics, sets and reps performed.")
	time.Sleep(1 * time.Second)
	fmt.Println("Each exercise selected will use your given data to calculate joules and calories burned.")
	fmt.Print(Reset)
	fmt.Println()
}

func checkNotebookForNotes() bool {
	fi, err := os.Stat("file.txt")
	if err != nil {
		return false
	}
	size := fi.Size()
	if size <= 0 {
		return false
	} else {
		return true
	}

}

func getMeasurements(totJo *float64, totCa *float64, run int) []float64 {
	var Reset = "\033[0m"
	var Green = "\033[32m"
	if _, err := os.Stat("file.txt"); err == nil {
		var hasNotes bool = checkNotebookForNotes()
		if hasNotes == true {
			if run == 0 {
				fmt.Print(Green)
				fmt.Println("Welcome back!")
				fmt.Print(Reset)
			}
			zip := readMeasurementFile()
			sel, kg, reps := movements()
			equation(totJo, totCa, sel, kg, reps)
			return zip
		}
	}
	getMeasure()
	sel, kg, reps := movements()
	equation(totJo, totCa, sel, kg, reps)
	return nil
}

func readMeasurementFile() []float64 {
	var som []float64
	file, err := os.Open("file.txt")
	if err != nil {
		return som
	}
	defer file.Close()
	var ftFloat []float64
	var files []string
	reader := bufio.NewReader(file)

	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				if len(line) > 0 {
					// fmt.Println("fff 1545")

					files = append(files, line)
				}
				break
			}
		}
		files = append(files, line)
	}

	for _, line := range files {
		value, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Println("Error converting line to float64:", err)
			continue
		}
		ftFloat = append(ftFloat, value)
	}
	return ftFloat
}
func getName() []string {
	names, err := os.OpenFile("names.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	}
	defer names.Close()
	name := ""
	writer2 := bufio.NewWriter(names)
	fmt.Println("Let's get your name!")
	fmt.Scanf("%s\n", &name)
	_, err = writer2.WriteString(name + "\n")
	if err != nil {
	}
	err = writer2.Flush()
	file, err := os.Open("names.txt")
	if err != nil {
		fmt.Errorf("failed to open file: %v", err)

	}
	defer file.Close()
	reader := bufio.NewReader(file)
	var name1 []string
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				if len(line) > 0 {
					name1 = append(name1, line)
				}
				break
			}
			fmt.Println("Error reading from buffer:", err)

		}
		name1 = append(name1, line)

	}
	return name1
}
func getMeasure() {
	var Reset = "\033[0m"
	var Green = "\033[32m"
	os.Create("file.txt")
	file, err := os.OpenFile("file.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	fmt.Print(Green)
	time.Sleep(1 * time.Second)
	fmt.Println("Pick one side of the body to measure and keep to that side for consistency.")
	fmt.Println("Accuracy is import please measure to at least the hundreths place, .00.")
	time.Sleep(1 * time.Second)
	fmt.Println("Enter the distance in meters of your foot to the center of your knee.")
	var footKnee float64
	fmt.Scanf("%f\n", &footKnee)
	ftStri := strconv.FormatFloat(footKnee, 'f', -2, 64)
	_, err = writer.WriteString(ftStri + "\n")
	time.Sleep(1 * time.Second)
	fmt.Println("Enter the distance in meters from center of your knee to your hip.")
	var kneeHip float64
	fmt.Scanf("%f\n", &kneeHip)
	khStri := strconv.FormatFloat(kneeHip, 'f', -2, 64)
	_, err = writer.WriteString(khStri + "\n")
	var pelShoulder float64
	time.Sleep(1 * time.Second)
	fmt.Println("Enter the distance in meters from center of your pelvis to the center of your shoulder.")
	fmt.Scanf("%f\n", &pelShoulder)
	psStri := strconv.FormatFloat(pelShoulder, 'f', -2, 64)
	_, err = writer.WriteString(psStri + "\n")
	var shoulElb float64
	time.Sleep(1 * time.Second)
	fmt.Println("Enter the distance in meters from center of your shoulder to the center of your elbow.")
	fmt.Scanf("%f\n", &shoulElb)
	seStri := strconv.FormatFloat(shoulElb, 'f', -2, 64)
	_, err = writer.WriteString(seStri + "\n")
	var elbHand float64
	time.Sleep(1 * time.Second)
	fmt.Println("Enter the distance in meters from center of your elbow to the center of your hand (mid palm).")
	fmt.Scanf("%f\n", &elbHand)
	fmt.Print(Reset)
	ehStri := strconv.FormatFloat(elbHand, 'f', -2, 64)
	_, err = writer.WriteString(ehStri + "\n")
	err = writer.Flush()
}
func movements() (int, float64, int) {
	var Reset = "\033[0m"
	var Green = "\033[32m"
	sel := 0
	reps := 0
	var mass int
	fmt.Println()
	fmt.Print(Green)
	time.Sleep(1 * time.Second)
	lifts := []string{
		"Front/Back/Goblet Squats",
		"Deadlift",
		"Bench Press",
		"Pull Up",
		"Chin Up",
		"Push Up",
		"Clean",
		"Clean & Jerk",
		"Snatch",
		"Barbell Rows",
		"DB Rows",
		"Shoulder Press",
		"Lunges",
		"Bulgarian Split Squat",
	}
	for {
		for i := 0; i < len(lifts); i++ {
			fmt.Println(i+1, ":", lifts[i])
		}
		fmt.Print("00 : I'm done\n")

		fmt.Scanf("%d\n", &sel)
		if sel == 1 {
			break
		} else if sel == 2 {
			break
		} else if sel == 3 {
			break
		} else if sel == 4 {
			break
		} else if sel == 5 {
			break
		} else if sel == 6 {
			break
		} else if sel == 7 {
			break
		} else if sel == 8 {
			break
		} else if sel == 9 {
			break
		} else if sel == 10 {
			break
		} else if sel == 11 {
			break
		} else if sel == 12 {
			break
		} else if sel == 13 {
			break
		} else if sel == 14 {
			break
		} else if sel == 00 {
			fmt.Print(Green)
			fmt.Println("Awesome sounds good! Now get some recovery in my friend!")
			fmt.Print(Reset)
			os.Exit(0)
		} else {
			time.Sleep(1 * time.Second)
			fmt.Print(Green)
			fmt.Println("That wasn't a valid selection, please pick again.")
			// fmt.Print(Reset)
			time.Sleep(1 * time.Second)
			continue
		}
	}
	fmt.Print(Green)
	fmt.Println("Now enter in weight(lbs) being used. If it's not a weight loaded movement please enter in your bodyweight(lbs).")
	fmt.Scanf("%d\n", &mass)
	kg := (float64(mass) * float64(0.45359237))
	fmt.Println("How many overall reps (sets*reps) did you perform?")
	fmt.Scanf("%d\n", &reps)
	fmt.Print(Reset)
	fmt.Println()
	return sel, kg, reps
}

func equation(totJo *float64, totCa *float64, sel int, kg float64, reps int) {
	var Reset = "\033[0m"
	var Magenta = "\033[35m"
	file, err := os.Open("file.txt")
	if err != nil {
		fmt.Errorf("failed to open file: %v", err)

	}
	defer file.Close()
	var ftFloat []float64
	var files []string
	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		if err != nil {
			if err.Error() == "EOF" {
				if len(line) > 0 {
					files = append(files, line)
				}
				break
			}
			fmt.Println("Error reading from buffer:", err)
			return
		}
		files = append(files, line)
	}
	for _, line := range files {
		value, err := strconv.ParseFloat(strings.TrimSpace(line), 64)
		if err != nil {
			fmt.Println("Error converting line to float64:", err)
			continue
		}
		ftFloat = append(ftFloat, value)
	}
	if sel == 1 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		squat := (float64(kg) * float64(9.8)) * ftFloat[1] * float64(reps)
		fmt.Print(squat, "\n")
		*totJo += squat
		time.Sleep(1 * time.Second)
		cal := squat * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()

		fmt.Println("Total Joules: ", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 2 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		dead := (float64(kg) * float64(9.8)) * (ftFloat[0] + ftFloat[1]) * float64(reps)
		*totJo += dead
		fmt.Print(dead, "\n")
		time.Sleep(1 * time.Second)
		cal := dead * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()

		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 3 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		bench := (float64(kg) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += bench
		fmt.Print(bench, "\n")
		time.Sleep(1 * time.Second)
		cal := bench * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 4 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		pullUp := (float64(kg) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += pullUp
		fmt.Print(pullUp, "\n")
		time.Sleep(1 * time.Second)
		cal := pullUp * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 5 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		chinUp := (float64(kg) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += chinUp
		fmt.Print(chinUp, "\n")
		time.Sleep(1 * time.Second)
		cal := chinUp * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 6 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		pushUp := (float64(kg) * float64(.6) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += pushUp
		fmt.Print(pushUp, "\n")
		time.Sleep(1 * time.Second)
		cal := pushUp * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 7 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		clean := (float64(kg) * float64(9.8)) * (ftFloat[0] + ftFloat[1] + ftFloat[2]) * float64(reps)
		*totJo += clean
		fmt.Print(clean, "\n")
		time.Sleep(1 * time.Second)
		cal := clean * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 8 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		cleanJerk := (float64(kg) * float64(9.8)) * (ftFloat[0] + ftFloat[1] + ftFloat[2] + ftFloat[3] + ftFloat[4]) * float64(reps)
		*totJo += cleanJerk
		fmt.Print(cleanJerk, "\n")
		time.Sleep(1 * time.Second)
		cal := cleanJerk * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 9 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		snatch := (float64(kg) * float64(9.8)) * (ftFloat[0] + ftFloat[1] + ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += snatch
		fmt.Print(snatch, "\n")
		time.Sleep(1 * time.Second)
		cal := snatch * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 10 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		bbRows := (float64(kg) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += bbRows
		fmt.Print(bbRows, "\n")
		time.Sleep(1 * time.Second)
		cal := bbRows * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 11 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		dbRows := (float64(kg) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += dbRows
		fmt.Print(dbRows, "\n")
		time.Sleep(1 * time.Second)
		cal := dbRows * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 12 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		shPress := (float64(kg) * float64(9.8)) * (ftFloat[4] + ftFloat[3]) * float64(reps)
		*totJo += shPress
		fmt.Print(shPress, "\n")
		time.Sleep(1 * time.Second)
		cal := shPress * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 13 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		lunge := (float64(kg) * float64(9.8)) * (ftFloat[0]) * float64(reps)
		*totJo += lunge
		fmt.Print(lunge, "\n")
		time.Sleep(1 * time.Second)
		cal := lunge * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	} else if sel == 14 {
		time.Sleep(1 * time.Second)
		fmt.Print(Magenta)
		fmt.Print("Joules of energy used: ")
		bulgLunge := (float64(kg) * float64(9.8)) * (ftFloat[0]) * float64(reps)
		*totJo += bulgLunge
		fmt.Print(bulgLunge, "\n")
		time.Sleep(1 * time.Second)
		cal := bulgLunge * float64(0.239006)
		fmt.Printf("Calories expended: %.02f\n", cal)
		*totCa += cal
		time.Sleep(3 * time.Second)
		fmt.Println()
		fmt.Println("Total Joules:", *totJo)
		fmt.Println("Total Calories:", *totCa)
		fmt.Print(Reset)
	}
}
