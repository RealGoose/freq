package main

import (
        "sync"
        "bufio"
        "net/http"
        "fmt"
        "os"
        "strings"
        "io/ioutil"
)

func main(){
        // fmt.Println("\n")
        fmt.Println("spray tool by realGoose")
        //fmt.Println("(\____/)\n(  °  ʖ  °)\n\╭☞ \╭☞")
        // fmt.Println("\n")

        colorReset := "\033[0m"
        colorRed := "\033[31m"
        //colorGreen := "\033[32m"
        client := &http.Client{}

        sc := bufio.NewScanner(os.Stdin)

        jobs := make(chan string)
        var wg sync.WaitGroup

        for i:= 0; i < 20; i++{

                wg.Add(1)
                go func(){
                        defer wg.Done()
                        for domain := range jobs {
                                
                                req, err := http.NewRequest("GET", domain+"/robots.txt", nil)
                                req.Header.Set("User-Agent", "TESTuserAgent")
                                resp, err := client.Do(req)
                                //resp, err := http.Get(domain)


                                if err != nil{
                                        continue
                                }
                                body, err := ioutil.ReadAll(resp.Body)
                                if err != nil {
                                fmt.Println(err)
                                }
                                sb := string(body)
                                if resp.StatusCode >= 200 && resp.StatusCode <= 299{
                                        // fmt.Println(check_result)
                                        fmt.Println(string(colorRed), "RESULT:", domain, string(colorReset))
                                }

                        }

                }()

        }



        for sc.Scan(){
                domain := sc.Text()
                jobs <- domain


        }
        close(jobs)
        wg.Wait()
}
