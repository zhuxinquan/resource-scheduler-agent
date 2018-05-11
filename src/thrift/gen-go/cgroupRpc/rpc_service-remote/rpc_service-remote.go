// Autogenerated by Thrift Compiler (0.10.0)
// DO NOT EDIT UNLESS YOU ARE SURE THAT YOU KNOW WHAT YOU ARE DOING

package main

import (
        "flag"
        "fmt"
        "math"
        "net"
        "net/url"
        "os"
        "strconv"
        "strings"
        "git.apache.org/thrift.git/lib/go/thrift"
        "cgroupRpc"
)


func Usage() {
  fmt.Fprintln(os.Stderr, "Usage of ", os.Args[0], " [-h host:port] [-u url] [-f[ramed]] function [arg1 [arg2...]]:")
  flag.PrintDefaults()
  fmt.Fprintln(os.Stderr, "\nFunctions:")
  fmt.Fprintln(os.Stderr, "  string ReadAllCgroupMetric(string req)")
  fmt.Fprintln(os.Stderr, "  string ReadSingleSubsytemCgroupMetric(string path, string subSystem)")
  fmt.Fprintln(os.Stderr, "  string Exec(string req)")
  fmt.Fprintln(os.Stderr, "  string SetMetric(string req)")
  fmt.Fprintln(os.Stderr, "  string GetSysInfo()")
  fmt.Fprintln(os.Stderr, "  string GetProcessInfo()")
  fmt.Fprintln(os.Stderr, "  string GetCpuAndMemStats()")
  fmt.Fprintln(os.Stderr, "  string GetGroupList()")
  fmt.Fprintln(os.Stderr, "  string GroupAdd(string path, string subSystems, string weight)")
  fmt.Fprintln(os.Stderr, "  string GroupDelete(string path)")
  fmt.Fprintln(os.Stderr)
  os.Exit(0)
}

func main() {
  flag.Usage = Usage
  var host string
  var port int
  var protocol string
  var urlString string
  var framed bool
  var useHttp bool
  var parsedUrl url.URL
  var trans thrift.TTransport
  _ = strconv.Atoi
  _ = math.Abs
  flag.Usage = Usage
  flag.StringVar(&host, "h", "localhost", "Specify host and port")
  flag.IntVar(&port, "p", 9090, "Specify port")
  flag.StringVar(&protocol, "P", "binary", "Specify the protocol (binary, compact, simplejson, json)")
  flag.StringVar(&urlString, "u", "", "Specify the url")
  flag.BoolVar(&framed, "framed", false, "Use framed transport")
  flag.BoolVar(&useHttp, "http", false, "Use http")
  flag.Parse()
  
  if len(urlString) > 0 {
    parsedUrl, err := url.Parse(urlString)
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
    host = parsedUrl.Host
    useHttp = len(parsedUrl.Scheme) <= 0 || parsedUrl.Scheme == "http"
  } else if useHttp {
    _, err := url.Parse(fmt.Sprint("http://", host, ":", port))
    if err != nil {
      fmt.Fprintln(os.Stderr, "Error parsing URL: ", err)
      flag.Usage()
    }
  }
  
  cmd := flag.Arg(0)
  var err error
  if useHttp {
    trans, err = thrift.NewTHttpClient(parsedUrl.String())
  } else {
    portStr := fmt.Sprint(port)
    if strings.Contains(host, ":") {
           host, portStr, err = net.SplitHostPort(host)
           if err != nil {
                   fmt.Fprintln(os.Stderr, "error with host:", err)
                   os.Exit(1)
           }
    }
    trans, err = thrift.NewTSocket(net.JoinHostPort(host, portStr))
    if err != nil {
      fmt.Fprintln(os.Stderr, "error resolving address:", err)
      os.Exit(1)
    }
    if framed {
      trans = thrift.NewTFramedTransport(trans)
    }
  }
  if err != nil {
    fmt.Fprintln(os.Stderr, "Error creating transport", err)
    os.Exit(1)
  }
  defer trans.Close()
  var protocolFactory thrift.TProtocolFactory
  switch protocol {
  case "compact":
    protocolFactory = thrift.NewTCompactProtocolFactory()
    break
  case "simplejson":
    protocolFactory = thrift.NewTSimpleJSONProtocolFactory()
    break
  case "json":
    protocolFactory = thrift.NewTJSONProtocolFactory()
    break
  case "binary", "":
    protocolFactory = thrift.NewTBinaryProtocolFactoryDefault()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid protocol specified: ", protocol)
    Usage()
    os.Exit(1)
  }
  client := cgroupRpc.NewRpcServiceClientFactory(trans, protocolFactory)
  if err := trans.Open(); err != nil {
    fmt.Fprintln(os.Stderr, "Error opening socket to ", host, ":", port, " ", err)
    os.Exit(1)
  }
  
  switch cmd {
  case "ReadAllCgroupMetric":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "ReadAllCgroupMetric requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.ReadAllCgroupMetric(value0))
    fmt.Print("\n")
    break
  case "ReadSingleSubsytemCgroupMetric":
    if flag.NArg() - 1 != 2 {
      fmt.Fprintln(os.Stderr, "ReadSingleSubsytemCgroupMetric requires 2 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    fmt.Print(client.ReadSingleSubsytemCgroupMetric(value0, value1))
    fmt.Print("\n")
    break
  case "Exec":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "Exec requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.Exec(value0))
    fmt.Print("\n")
    break
  case "SetMetric":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "SetMetric requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.SetMetric(value0))
    fmt.Print("\n")
    break
  case "GetSysInfo":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetSysInfo requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetSysInfo())
    fmt.Print("\n")
    break
  case "GetProcessInfo":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetProcessInfo requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetProcessInfo())
    fmt.Print("\n")
    break
  case "GetCpuAndMemStats":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetCpuAndMemStats requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetCpuAndMemStats())
    fmt.Print("\n")
    break
  case "GetGroupList":
    if flag.NArg() - 1 != 0 {
      fmt.Fprintln(os.Stderr, "GetGroupList requires 0 args")
      flag.Usage()
    }
    fmt.Print(client.GetGroupList())
    fmt.Print("\n")
    break
  case "GroupAdd":
    if flag.NArg() - 1 != 3 {
      fmt.Fprintln(os.Stderr, "GroupAdd requires 3 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    argvalue1 := flag.Arg(2)
    value1 := argvalue1
    argvalue2 := flag.Arg(3)
    value2 := argvalue2
    fmt.Print(client.GroupAdd(value0, value1, value2))
    fmt.Print("\n")
    break
  case "GroupDelete":
    if flag.NArg() - 1 != 1 {
      fmt.Fprintln(os.Stderr, "GroupDelete requires 1 args")
      flag.Usage()
    }
    argvalue0 := flag.Arg(1)
    value0 := argvalue0
    fmt.Print(client.GroupDelete(value0))
    fmt.Print("\n")
    break
  case "":
    Usage()
    break
  default:
    fmt.Fprintln(os.Stderr, "Invalid function ", cmd)
  }
}
