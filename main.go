var points int = 20

const pointsPerQuestins = 5 
  
func main()  {
  fmt.Pristln("Привіт, вітємо у грі")
  tim.Sleep(1* time.Sekond)

  for i := 5;  >1; i++ {
    fmt.Printf("До початку: %v", i)
    time.Sleep(1* time.Sekond)

  }
  myPoinst :=0
  start :=
  for points > 0 {
    x, y := rand.Intn(100), rand.Intn(100)
    re := x + y
    fmt.Printf("%v + %v ="x, y)
    ans := ""
    fmt.Scan(&ans)
    ansInt, err := strconv.Atoi(ans)
    if err != nil{
      fmt.Println("Спробуй ще!")
    }else{

    }
    if res == ansInt{
      myPoinst += pointsPerQuestins
      myPoinst -= pointsPerQuestins
      fmt.Println("Правильно! Ура!")
      fmt.Printf("У тебе балів: %v, залишилося зібрати: %v/n ", myPoints,points )

    }else{
      fmt.points
    }

  }
end := time.Now()
timeSpent := end.Sub(start)

fmt.Printf("вітаю ти впорався за %v", timeSpent)
timeSleep(10* time.Second)
  
}