package main
import  (
	"fmt"
	"flag"
	"net/http"
	"sync"
	"time"
)

var wg sync.WaitGroup

func sendReq(url string) {
	//fmt.Println(url)

	resp, err  :=  http.Get(url)
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	//fmt.Println(resp.Status)
}

func main() {
	urlPtr := flag.String("url", "Url", "Url to benchmark")
	noOfRequests := flag.Int("n", 100, "No of requests")
	concurrentConn := flag.Int("c", 10, "No of concurrent")

	flag.Parse()
	reqPerC := *noOfRequests / *concurrentConn

	start := time.Now()

	for j := 0; j<*concurrentConn;j++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for i := 0; i<reqPerC; i++ {
				sendReq(*urlPtr)
			}
		}()
	}

	wg.Wait()

	finish := time.Now()

	fmt.Println("Url: ", *urlPtr);
	fmt.Println("N: ", *noOfRequests);
	fmt.Println("C: ", *concurrentConn);
	fmt.Println("Per: ", reqPerC);

	fmt.Println("Time taken ", finish.Sub(start))


}
