package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

/*
Write a program which reads a set of lines which shows all your email messages, with date as first column, sender as second column & subject as final column.
Find what’s the average emails per day you receive by sender.
The total number of days is based on the range of days which you have given. For example, if you’re given emails from 2021-01-01 to 2021-01-03 only, the total number of days is three.
Finally, print the list of senders which are likely to be spammers - those which send you more than 3 emails per day on average.
Print the results sorted based on the average emails per day in descending order.
*/

func main() {
	getNum()
}
func getNum() {
	r := bufio.NewReader(os.Stdin)
	line, err := r.ReadString('\n')
	if err != nil {
		panic(err)
	}
	num, err := strconv.Atoi(strings.TrimSpace(line))
	if err != nil {
		panic(err)
	}

	minTime := time.Unix(-2208988800, 0)
	maxTime := minTime.Add(1<<63 - 1)
	minDate, maxDate := maxTime, minTime
	avgEmailFreq := map[string]float64{}
	for i := 0; i < num; i++ {
		line, err := r.ReadString('\n')
		if err != nil {
			panic(err)
		}
		lineSli := strings.Split(strings.TrimSpace(line), " ")
		email := lineSli[1]
		date, err := time.Parse("2006-05-02", lineSli[0])
		if err != nil {
			panic(err)
		}
		avgEmailFreq[email]++
		if date.Before(minDate) {
			minDate = date
		}
		if date.After(maxDate) {
			maxDate = date
		}

	}
	totalDays := maxDate.Add(24*time.Hour).Sub(minDate).Hours() / 24
	var emails []string
	for email := range avgEmailFreq {
		avgEmailFreq[email] /= totalDays
		emails = append(emails, email)
	}
	sort.Slice(emails, func(i, j int) bool {
		return avgEmailFreq[emails[i]] > avgEmailFreq[emails[j]]
	})
	var spammers []string
	for _, email := range emails {
		fmt.Printf("%s -> %.2f\n", email, avgEmailFreq[email])
		if avgEmailFreq[email] >= 3 {
			spammers = append(spammers, email)
		}
	}
	fmt.Println("\nThe spammers are: ")
	for _, spEmail := range spammers {
		fmt.Println(spEmail)
	}
}
