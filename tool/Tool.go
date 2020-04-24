package tool

import (
	caller "adjust/caller"
	converter "adjust/converter"
	"strings"
	"sync"
)

var (
	mu                  = &sync.Mutex{}
	MAX_THREADS_ALLOWED = 50
)

type AdjustTool struct {
}

func (a *AdjustTool) Run(urlArray []string, isParallel bool, parallelCount uint) []ResultDTO {
	conFac := new(converter.ConverterFactory)
	cnv := conFac.GetConverter("md5")
	hc := new(caller.HttpCaller)
	resultMap := make([]ResultDTO, 0)
	if isParallel {
		a.RunWithParallel(urlArray, parallelCount, cnv, hc, &resultMap)
	} else {
		a.printHashes(urlArray, cnv, hc, &resultMap, false)
	}
	for _, result := range resultMap {
		result.Render()
	}
	return resultMap
}

func (a *AdjustTool) RunWithParallel(urlArray []string, parallelCount uint, cnv converter.Converter, hc *caller.HttpCaller, resultMap *[]ResultDTO) {
	var stride int
	if int(parallelCount) > MAX_THREADS_ALLOWED {
		parallelCount = uint(MAX_THREADS_ALLOWED)
	}
	if int(parallelCount) > len(urlArray) {
		stride = 1
		parallelCount = uint(len(urlArray))
	} else {
		stride = len(urlArray) / int(parallelCount)
	}
	lastGoRoutine := parallelCount - 1
	var wg sync.WaitGroup
	wg.Add(int(parallelCount))
	for g := 0; g < int(parallelCount); g++ {
		go func(g int) {
			start := g * stride
			end := start + stride
			if g == int(lastGoRoutine) {
				end = len(urlArray)
			}
			a.printHashes(urlArray[start:end], cnv, hc, resultMap, true)
			wg.Done()
		}(g)
	}
	wg.Wait()
}

func (a *AdjustTool) printHashes(urlArray []string, cnv converter.Converter, hc *caller.HttpCaller, resultMap *[]ResultDTO, isParallel bool) {
	for _, url := range urlArray {
		hasHttpPrefix := strings.HasPrefix(url, "http://")
		if !hasHttpPrefix {
			url = "http://" + url
		}
		respHash := cnv.GetHash(hc.Call(url))
		if isParallel {
			mu.Lock()
			result := ResultDTO{url: url, hash: respHash}
			*resultMap = append(*resultMap, result)
			mu.Unlock()
		} else {
			result := ResultDTO{url: url, hash: respHash}
			*resultMap = append(*resultMap, result)
		}
	}
}
