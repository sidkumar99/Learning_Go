package main

import("fmt"
		"os"
		"encoding/csv"
		"io"
		"log")


type Question struct{
	question string
	answer string
}


func main(){
	bootup()
	//Initialize slice to store the questions
	questions := parse_csv()

	correct := 0
	total := 0

	for _,q := range questions{
		
		fmt.Printf("%s = ", q.question)
		
		var attempt string
		//Read users answer
		fmt.Scanf("%s", &attempt)
		
		if attempt == q.answer{
			correct++
		}

		total ++
	}
	//Print results
	results(correct, total)
}

func results(c int, t int){
	fmt.Printf("You Got %d Out of %d Questions Right! \n",c,t)
}

func bootup(){
	fmt.Println("Welcome to the Quiz!")

}

func parse_csv() []Question {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(csvFile)
	//Create questions array to hold questions
	var questions []Question

	for{
		line, err := reader.Read()
		if err == io.EOF{
			break
		} else if err != nil{
			log.Fatal(err)
		}

		questions = append(questions, Question{
			question: line[0],
			answer: line[1],
		})
	}
	
	return questions
}
