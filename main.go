package main

import (
  //  "bytes"
    //"log"
    "io"
    "os"
    "os/exec"
    "bytes"
  //  "context"
    "github.com/creack/pty"
    "fmt"
    shell "github.com/kballard/go-shellquote"
    //"strings"
)
// docker run -d -p 22github.com/creack/pty22:22 -e SSH_KEY="$(cat ~/.ssh/id_rsa.pub)" krlmlr/debian-ssh:wheezy
type ExpectSubprocess struct{
	Cmd          *exec.Cmd
	buf          *buffer
	outputBuffer []byte
}

type buffer struct {
	f       *os.File
	b       bytes.Buffer
	collect bool
  collection bytes.Buffer
}
	
//const URL = "StrictHosgithub.com/creack/ptytKeyChecking=no mippbe@10.114.200.151:22@10.112.16.84"
func main() {
    //ctx, cancel := cgithub.com/creack/ptyontext.WithTimeout(context.Background(), 4*time.Second)
    //var out bytes.Bugithub.com/creack/ptyffer
    var f *os.File
    splitArgs, _ := shell.Split("ssh -o StrictHostKeyChecking=no mippbe@10.114.200.151:22@10.112.16.84")
    path, _ := exec.LookPath(splitArgs[0])
    cmd := exec.Command(path, splitArgs[1:]...)
    size, _ := pty.GetsizeFull(os.Stdin)
    fmt.Println(size)
    // TODO: child.sendline(f'export COLUMNS={columns};shopt -s checkwinsize; stty rows {rows} cols {columns}')
    err := pty.Setsize(os.Stdin, size)
    if err != nil {
            fmt.Println(err)
    }
    cmdRet, _ := pty.StartWithSize(cmd, size)
    f = cmdRet
    defer cmd.Wait()
    go io.Copy(os.Stdin, f)
    go io.Copy(f, os.Stdin)
    //go io.Copy(f, os.Stdin)
    //err := cmd.Run()
    //if err != nil {
    //    fmt.Println(fmt.Sprint(err))
    //    return
    //}
}
