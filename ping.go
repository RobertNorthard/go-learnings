package main 

import (
  "flag"
  "fmt"
  "github.com/sparrc/go-ping"
)

var hostname string

func init() {
	flag.StringVar(&hostname, "hostname", "google.com", "Hostname to ping")
}

func usage() {
  flag.PrintDefaults()
}

func main() {
   flag.Parse()

   args := flag.Args()

   // checks non flagged arguments
   if len(args) != 0 {
     usage()
   }
   
    pinger, err := ping.NewPinger(hostname)

     if err != nil {
       panic(err)
     }

     pinger.Count = 3
     pinger.Run()

     fmt.Printf("PING %s (%s):\n", pinger.Addr(), pinger.IPAddr())
     fmt.Println(pinger.Statistics())
}