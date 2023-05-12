package main

import (
    "fmt"
    "encoding/json"
    "io/ioutil"
    "math/rand"
    "log"
    "github.com/gofiber/fiber/v2"
    "github.com/gofiber/template/html"
    "strconv"
    "strings"
)

func main() {

    // Initialize standard Go html template engine
    engine := html.New("./Views", ".html")

    app := fiber.New(fiber.Config{
        Views: engine,
    })
    
    app.Get("/", func(c *fiber.Ctx) error {
        
        jsonFile := "./public/db/svenska-ord.json" 
	    fileInput, err := ioutil.ReadFile(jsonFile)
        var strSlice []string
        if err != nil {
		fmt.Println("Error when opening file: ", err)
	    }
        numberOfWords := 3
        bindChar := "_-+."
        bindChar = string(bindChar[rand.Intn(len(bindChar))])
        number := strconv.Itoa(rand.Intn(10000))
        numberPos := rand.Intn(2)
        json.Unmarshal([]byte(fileInput), &strSlice)

        //Add Words
        var word string
        for i:=0; i<numberOfWords; i++{
            var temp string
            temp = strings.Title(strSlice[rand.Intn(len(strSlice))])
            word += temp
            
            if i<numberOfWords-1 {
                word += bindChar
            }
        }

        //Add number
        if numberPos == 0 {
            word += bindChar + number
         } else if numberPos == 1 {
            word = number + bindChar + word
         }

        return c.Render("index", fiber.Map{
            "Name": word,
            "NumberOfWords": numberOfWords,
            "BindChar": bindChar,
            "Number": number,
            "NumberPos": numberPos,
        })
    })

     app.Post("/", func(c *fiber.Ctx) error {
        
         jsonFile := "./public/db/svenska-ord.json" 
	     fileInput, err := ioutil.ReadFile(jsonFile)
         var strSlice []string
         numberOfWords, _ := strconv.Atoi(c.FormValue("numberOfWords")) 
         bindChar := "_-+."
         bindChar = string(bindChar[rand.Intn(len(bindChar))])
         number := strconv.Itoa(rand.Intn(10000))
         numberPos := rand.Intn(2)
         if err != nil {
		 fmt.Println("Error when opening file: ", err)
	     }

         json.Unmarshal([]byte(fileInput), &strSlice)

         // Add Words
         var word string
        for i:=0; i<numberOfWords; i++{
            var temp string
            temp = strings.Title(strSlice[rand.Intn(len(strSlice))])
            word += temp
            
            if i<numberOfWords-1 {
                word += bindChar
            }
        }

        // Add number 
        if numberPos == 0 {
            word += bindChar + number
         } else if numberPos == 1 {
            word = number + bindChar + word
         }

        

         return c.Render("index", fiber.Map{
             "Name": word,
             "NumberOfWords": numberOfWords,
             "BindChar": bindChar,
             "Number": number,
             "NumberPos": numberPos,
         })
     })

    app.Static("/", "./public")

    log.Fatal(app.Listen(":3000"))
}
