package main

import (
    "fmt"
    "strings"
    "sort"
    "io/ioutil"
    "os"
    "runtime"
)

func IndexSep(s string, char string, offset int) int {
    suffix := s[offset:]
    suffixIndex := strings.Index(suffix, char)

    if (suffixIndex == -1) {
        return -1
    }

    return offset + suffixIndex
}

func FindEndOfMatch(s string, chars []string, first_index int) int {
    last_index := first_index

    for i := 0; i < len(chars); i++ {
        c := chars[i]
        index := IndexSep(s, c, last_index + 1)

        if index == -1 {
            return -1
        }

        last_index = index
    }

    return last_index
}

func FindCharInString(s string, char string) []int {
    index := 0
    indexes := make([]int, 0)

    for index > -1 {
        index = IndexSep(s, char, index)
        
        if index > -1 {
            indexes = append(indexes, index)
            index = index + 1
        }
    }

    return indexes
}

func ComputeMatchLength(s string, chars []string) float64 {
    first_char := chars[0]
    rest := chars[1:]
    first_indexes := FindCharInString(s, first_char)
    result := make([]float64, 0)

    for i := 0; i < len(first_indexes); i++ {
        first_index := first_indexes[i]
        last_index := FindEndOfMatch(s, rest, first_index)
        if last_index > -1 {
            result = append(result, float64(last_index - first_index + 1))
        }
    }

    if len(result) == 0 {
        return 0.0
    }

    sort.Float64s(result)

    return result[0]
}

func Score(choice string, query string) float64 {
    if len(query) == 0 {
        return 1.0
    }

    if len(choice) == 0 {
        return 0.0
    }

    choice = strings.ToLower(choice)
    query = strings.ToLower(query)

    match_length := ComputeMatchLength(choice, strings.Split(query, ""))

    if match_length <= 0 {
        return 0.0
    }

    score := float64(len(query)) / match_length
    return score / float64(len(choice))
}

type Fscore struct {
    score float64
    choice string
}

type FscoreArr []Fscore

func (a FscoreArr) Len() int {
    return len(a)
}

func (a FscoreArr) Swap(i, j int) {
    a[i], a[j] = a[j], a[i]
}

func (a FscoreArr) Less(i, j int) bool {
    return a[i].score < a[j].score
}

func Match(files []string, query string) []string {
    scores := make(FscoreArr, 0)

    for _, f := range files {
        scores = append(scores, Fscore{Score(f, query), f})
    }

    sort.Sort(scores)

    result := make([]string, 0)

    for _, f := range scores {
        if f.score > 0 {
            result = append(result, f.choice)
        }
    }

    return result
}

func main() {

    runtime.GOMAXPROCS(runtime.NumCPU())

    bytes, _ := ioutil.ReadAll(os.Stdin)
    files := strings.Split(string(bytes), "\n")

    if len(os.Args) == 1 {
        for _, f := range files {
            fmt.Println(f)
        }

        return
    }

    query := os.Args[1]
    files = Match(files, query)

    for _, f := range files {
        fmt.Println(f)
    }

}
