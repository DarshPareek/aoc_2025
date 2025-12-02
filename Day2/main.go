package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func readInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)
	if err != nil {
		fmt.Printf("error reading the file %v\n", err)
		return nil, err
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		return nil, err
	}
	return lines, nil
}

func main() {
	inp := "2558912-2663749,1-19,72-85,82984-100358,86-113,193276-237687,51-69,779543-880789,13004-15184,2768-3285,4002-4783,7702278-7841488,7025-8936,5858546565-5858614010,5117615-5149981,4919-5802,411-466,126397-148071,726807-764287,7454079517-7454227234,48548-61680,67606500-67729214,9096-10574,9999972289-10000034826,431250-455032,907442-983179,528410-680303,99990245-100008960,266408-302255,146086945-146212652,9231222-9271517,32295166-32343823,32138-36484,4747426142-4747537765,525-652,333117-414840,13413537-13521859,1626-1972,49829276-50002273,69302-80371,8764571787-8764598967,5552410836-5552545325,660-782,859-1056"
	//inp := "11-22,95-115,998-1012,1188511880-1188511890,222220-222224,1698522-1698528,446443-446449,38593856-38593862,565653-565659,824824821-824824827,2121212118-2121212124"
	lines := strings.Split(inp, ",")
	ans := 0
	for _, value := range lines {
		values := strings.Split(value, "-")
		start, err := strconv.Atoi(values[0])
		end, err := strconv.Atoi(values[1])
		if err != nil {
			return
		}
		for i := start; i <= end; i++ {
			d := strconv.Itoa(i)
			n := len(d) / 2
			for j := 1; j <= n; j++ {
				silly := "" + d[:j]
				is_silly := 1
				var parts []string
				parts = append(parts, silly)
				k := 0
				for k = j; k < len(d)-j; k += j {
					parts = append(parts, d[k:k+j])
				}
				parts = append(parts, d[k:])
				for q := 0; q < len(parts)-1; q++ {
					if parts[q] != parts[q+1] {
						is_silly = 0
						break
					}
				}
				if is_silly == 1 {
					fmt.Println(parts)
					ans += i
					break
				}
			}
		}
	}
	fmt.Println(ans)
}
