package main

import (
	"bufio"
	"context"
	"errors"
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"sync"

	"golang.org/x/sync/semaphore"
)

func main() {
	flag.BoolVar(&hideError, "hideerror", false, "hide query ERROR")
	flag.BoolVar(&hideNoEntry, "hidenoentry", false, "hide NOENTRY")
	flag.Int64Var(&concurrency, "concurrency", 500, "concurrency of DNS query")
	flag.Parse()
	if concurrency < 0 {
		log.Fatalf("concurrency should be positive: %d", concurrency)
	}

	err := masslookup(context.Background(), os.Stdout, os.Stdin)
	if err != nil {
		log.Fatal(err)
	}
}

var (
	hideError   bool
	hideNoEntry bool
	concurrency int64
)

func masslookup(ctx context.Context, w io.Writer, r io.Reader) error {
	br := bufio.NewReader(r)
	var wg sync.WaitGroup
	sem := semaphore.NewWeighted(concurrency)
	for {
		line, err := br.ReadString('\n')
		if err != nil {
			if errors.Is(err, io.EOF) {
				break
			}
			return err
		}
		err = sem.Acquire(ctx, 1)
		if err != nil {
			return err
		}
		line = strings.TrimSpace(line)
		wg.Add(1)
		go func(q string) {
			defer wg.Done()
			defer sem.Release(1)
			names, err := net.LookupAddr(q)
			if err != nil {
				if !hideError {
					fmt.Fprintf(w, "%s\tERROR\t%s\n", q, err)
				}
				return
			}
			if len(names) == 0 {
				if !hideNoEntry {
					fmt.Fprintf(w, "%s\tNOENTRY\n", q)
				}
				return
			}
			fmt.Fprintf(w, "%s\tFOUND\t%s\n", q, strings.Join(names, " "))
		}(line)
	}
	wg.Wait()
	return nil
}
