package main

import (
 "bytes"
 "compress/zlib"
 "crypto/sha1"
 "encoding/hex"
 "flag"
 "fmt"
 "io"
 "os"
 "path/filepath"
)

func main(){ flag.Parse(); args:=flag.Args(); if len(args)==0 { usage(); return }; switch args[0] { case "init": initRepo(); case "hash-object": if len(args)<2 {usage();return}; hashObject(args[1]); case "cat-file": if len(args)<2 {usage();return}; catFile(args[1]); default: usage() } }
func usage(){ fmt.Println("go-git init | hash-object <file> | cat-file <sha>") }
func initRepo(){ os.MkdirAll(".git/objects",0755); os.MkdirAll(".git/refs/heads",0755); os.WriteFile(".git/HEAD",[]byte("ref: refs/heads/main\n"),0644); fmt.Println("Initialized repository") }
func hashObject(path string){ data,err:=os.ReadFile(path); must(err); payload:=append([]byte(fmt.Sprintf("blob %d\x00",len(data))),data...); sum:=sha1.Sum(payload); sha:=hex.EncodeToString(sum[:]); dir:=filepath.Join(".git/objects",sha[:2]); os.MkdirAll(dir,0755); var b bytes.Buffer; zw:=zlib.NewWriter(&b); _,err=zw.Write(payload); must(err); must(zw.Close()); must(os.WriteFile(filepath.Join(dir,sha[2:]),b.Bytes(),0444)); fmt.Println(sha) }
func catFile(sha string){ if len(sha)!=40 { panic("full 40-character SHA required") }; f,err:=os.Open(filepath.Join(".git/objects",sha[:2],sha[2:])); must(err); defer f.Close(); zr,err:=zlib.NewReader(f); must(err); defer zr.Close(); data,err:=io.ReadAll(zr); must(err); i:=bytes.IndexByte(data,0); if i<0 {panic("invalid object")}; os.Stdout.Write(data[i+1:]) }
func must(err error){ if err!=nil { panic(err) } }
