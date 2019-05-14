package main

import("fmt"
		"os"
		"encoding/csv"
		"io"
		"log"
		"time")


type Question struct{
	question string
	answer string
}


func main(){
	bootup()
	//Initialize slice to store the questions
	total_time := get_time()
	timer1 := time.NewTimer(time.Duration(total_time)*time.Second)
	
	

	questions, total := parse_csv()

	correct := 0
	
	go func(){
		<-timer1.C
		fmt.Println("Times Up!")
		results(correct, total)
		os.Exit(0)
	}()
	
	for _,q := range questions{
		
		fmt.Printf("%s = ", q.question)
		
		var attempt string
		//Read users answer
		fmt.Scanf("%s", &attempt)
		
		if attempt == q.answer{
			correct++
		}
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

func parse_csv() ([]Question, int) {
	csvFile, _ := os.Open("problems.csv")
	reader := csv.NewReader(csvFile)
	t_questions := 0
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
		t_questions ++
	}
	
	return questions,t_questions
}

func get_time() int{
	fmt.Println("How much time do you want")
	var time int
	fmt.Scanf("%d", &time)
	return time
}