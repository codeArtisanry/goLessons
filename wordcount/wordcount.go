package wordcount

import (
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
	"sort"
	"strings"
	"sync"
	"unicode"
	"unicode/utf8"
)

type Pair struct {
	Word  string
	Count int
}

// PariList实现了sort接口，可以使用sort.Sort对其排序

type Pairs []Pair

func (p Pairs) Swap(i, j int)      { p[i], p[j] = p[j], p[i] }
func (p Pairs) Len() int           { return len(p) }
func (p Pairs) Less(i, j int) bool { return p[j].Count < p[i].Count } // 逆序

// SplitOnNonLetters 提取单词
func SplitOnNonLetters(s string) []string {
	notALetter := func(char rune) bool {
		return !unicode.IsLetter(char)
	}
	return strings.FieldsFunc(s, notALetter)
}

/*
   基于map实现了类型WordCount, 并对期实现了Merge(), Report(), SortReport(), UpdateFreq(),
   WordFreqCounter() 方法
*/

type WordCounts map[string]int

func updateWordCount(results map[string]int, counts WordCounts) {

	for word, v := range counts {
		_, ok := results[word]
		if ok {
			results[word] += v
		} else {
			results[word] = v
		}
		fmt.Println(word, " ", results[word])
	}
}

func Update() (map[string]int, chan<- WordCounts) {
	counts := make(chan WordCounts)
	wordCount := make(map[string]int)
	go func() {
		for {
			select {
			case s := <-counts:
				updateWordCount(wordCount, s)
			}
		}
	}()

	return wordCount, counts
}

// WordCount returns a map of the counts of each “word” in the string s.
// 固定并发
func Map(src <-chan *string, counts chan<- WordCounts) {
	for s := range src {
		// 1. 分词
		words := strings.Fields(s.str)
		// 2. 当前分词结果
		countMap := make(map[string]int)
		for _, word := range words {
			_, ok := countMap[word]
			if ok {
				countMap[word]++
			} else {
				countMap[word] = 1
			}
		}
		counts <- countMap
	}
}

func Reduced(wg *sync.WaitGroup, counts chan<- WordCounts) {
	// 1. 分词
	words := strings.Fields(s.str)
	// 2. 当前分词结果
	countMap := make(map[string]int)
	for _, word := range words {
		_, ok := countMap[word]
		if ok {
			countMap[word]++
		} else {
			countMap[word] = 1
		}
	}
	counts <- countMap
	wg.Done()
}

func Do() {

	input := make(chan *string, numWordCounters)
	complate := make(chan *string, numWordCounters)
	// 生成输入
	go func() {
		for _, str := range textStrings {
			input <- ""
		}
	}()

	var wg = new(sync.WaitGroup)
	// 单线程更新防止并发问题
	results, counts := Update()
		// 并发执行
	for i := 0; i < numWordCounters; i++ {
		wg.Add(1)
		go Reduced(wg,counts)
		// go Map(input, counts)
	}
	wg.Wait()
	fmt.Println(results)
	// 等待MapReduce结束 或者实时显示结果

}


func Freq(){
	// 1. 转换
	// 2.
}
// 用于合并两个WordCount
func (source WordCount) Merge(wordcount WordCount) WordCount {
	for k, v := range wordcount {
		source[k] += v
	}

	return source
}

// 打印词频统计情况
func (wordcount WordCount) Report() {
	words := make([]string, 0, len(wordcount))
	wordWidth, frequencyWidth := 0, 0
	for word, frequency := range wordcount {
		words = append(words, word)
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	sort.Strings(words)
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")
	for _, word := range words {
		fmt.Printf("%-*s %*d\n", wordWidth, word, frequencyWidth,
			wordcount[word])
	}
}

// 从多到少打印词频
func (wordcount WordCount) SortReport() {
	p := make(Pairs, len(wordcount))
	i := 0
	for k, v := range wordcount { // 将wordcount map转换成Pairs
		p[i] = Pair{k, v}
		i++
	}

	sort.Sort(p) // 因为 Pairs 实现了排序接口，所以可以使用 sort.Sort() 对其排序

	wordWidth, frequencyWidth := 0, 0
	for _, pair := range p {
		word, frequency := pair.Key, pair.Value
		if width := utf8.RuneCountInString(word); width > wordWidth {
			wordWidth = width
		}
		if width := len(fmt.Sprint(frequency)); width > frequencyWidth {
			frequencyWidth = width
		}
	}
	gap := wordWidth + frequencyWidth - len("Word") - len("Frequency")
	fmt.Printf("Word %*s%s\n", gap, " ", "Frequency")

	for _, pair := range p {
		fmt.Printf("%-*s %*d\n", wordWidth, pair.Key, frequencyWidth,
			pair.Value)
	}

}

// 从文件中读取单词，并更新其出现的次数
func (wordcount WordCount) UpdateFreq(filename string) {
	var file *os.File
	var err error

	if file, err = os.Open(filename); err != nil {
		log.Println("failed to open the file: ", err)
		return
	}
	defer file.Close() // 本函数退出之前时，关闭文件

	reader := bufio.NewReader(file)
	for {
		line, err := reader.ReadString('\n')
		for _, word := range SplitOnNonLetters(strings.TrimSpace(line)) {
			if len(word) > utf8.UTFMax ||
				utf8.RuneCountInString(word) > 1 {
				wordcount[strings.ToLower(word)] += 1
			}
		}
		if err != nil {
			if err != io.EOF {
				log.Println("failed to finish reading the file: ", err)
			}
			break
		}
	}
}

// 并发统计单词频次
func (wordcount WordCount) WordFreqCounter(files []string) {

	results := make(chan Pair, len(files))  // goroutine 将结果发送到该channel
	done := make(chan struct{}, len(files)) // 每个goroutine工作完成后，发送一个空结构体到该channel，表示工作完成

	for i := 0; i < len(files); { // 有多少个文件就开启多少个goroutine, 使用匿名函数的方式
		go func(done chan<- struct{}, results chan<- Pair, filename string) {
			wordcount := make(WordCount)
			wordcount.UpdateFreq(filename)
			for k, v := range wordcount {
				pair := Pair{k, v}
				results <- pair
			}
			done <- struct{}{}
		}(done, results, files[i])

		i++
	}

	for working := len(files); working > 0; { // 监听通道，直到所有的工作goroutine完成任务时才退出
		select {
		case pair := <-results: // 接收发送到通道中的统计结果
			wordcount[pair.Key] += pair.Value

		case <-done: // 判断工作goroutine是否全部完成
			working--

		}
	}

DONE: // 再次启动for循环处理通道中还未处理完的值
	for {
		select {
		case pair := <-results:
			wordcount[pair.Key] += pair.Value
		default:
			break DONE
		}
	}

	close(results)
	close(done)

}
