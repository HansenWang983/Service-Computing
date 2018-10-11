package main

import (
    "fmt"
    "io"
    "os"
	"os/exec"
	"bufio"
	flag "github.com/spf13/pflag"
)
/*================================= types =========================*/
type selpg_args struct {
    start_page  int //起始页码
    end_page    int //终止页码
    in_filename  string //输入文件名
    page_len    int //每页行数
    form_deli   bool //是否按分页符分页
    print_dest string //打印机目的地
}

/*================================= globals =======================*/
var progname string

/*================================= prototypes ====================*/
// func Usage();
// func Init(args *selpg_args);
// func process_args(args *selpg_args);
// func process_input(args *selpg_args);
// func print_write(args *selpg_args, line string, stdin io.WriteCloser);

/*================================= main()=========================*/
func main() {
	progname = os.Args[0]
	var args selpg_args
	Init(&args)
	process_args(&args)
	process_input(&args)
}

/*================================= Init(args *selpg_args)=========================*/
//Initial flags
func Init(args *selpg_args) {
	flag.Usage = Usage
	flag.IntVar(&args.start_page, "s", -1, "Start page number.")
	flag.IntVar(&args.end_page, "e", -1, "End page number.")
	flag.IntVar(&args.page_len, "l", 72, "Line number per page.")
	flag.BoolVar(&args.form_deli, "f", false, "Determine form-feed-delimited")
	flag.StringVar(&args.print_dest, "d", "", "specify the printer")
	flag.Parse()
	// fmt.Printf("args=%s, num=%d\n", flag.Args(), flag.NArg())
}

/*================================= process_args(args *selpg_args)=========================*/

func process_args(args *selpg_args) {
	// print(args.start_page)
	// print(args.end_page)

	if args.start_page == -1 || args.end_page == -1 {
		fmt.Fprintf(os.Stderr, "Error! %s: Not enough arguments\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if os.Args[1][0] != '-' || os.Args[1][1] != 's' {
		fmt.Fprintf(os.Stderr, "Error! %s: 1st arg should be -s=start_page_number\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	end_index := 2
	if len(os.Args[1]) == 2 {
		end_index = 3
	}

	if os.Args[end_index][0] != '-' || os.Args[end_index][1] != 'e' {
		fmt.Fprintf(os.Stderr, "Error! %s: 2st arg should be -e=end_page_number\n\n", progname)
		flag.Usage()
		os.Exit(1)
	}

	if args.start_page > args.end_page || args.start_page < 1 || args.end_page < 1 {
		fmt.Fprintln(os.Stderr, "Error! Page number invalid\n\n")
		flag.Usage()
		os.Exit(1)
	}
}

/*================================= process_input(args *selpg_args)=========================*/

func process_input(args *selpg_args) {
	var stdin io.WriteCloser
	var err error
	var cmd *exec.Cmd

	result := ""
	//记录总行数
	line_count := 0
	//记录页码
	page_count := 1

	if args.print_dest != "" {
		cmd = exec.Command("lp", "-d"+args.print_dest)
		stdin, err = cmd.StdinPipe()
		if err != nil {
			fmt.Println(err)
		}
	} else {
		stdin = nil
	}
	
	// fmt.Printf("args=%s, num=%d\n", pflag.Args(), pflag.NArg())
	
	if flag.NArg() > 0 {
		//读输入文件
		args.in_filename = flag.Arg(0)
		output, err := os.Open(args.in_filename)
		if err != nil {
			fmt.Println(err)
			// fmt.Fprintf(os.Stderr, "Error! No such no such file or directory\n\n")
			os.Exit(1)
		}
		// 最小的缓存容器是 16 字节
		reader := bufio.NewReader(output)
		// 按分页符’\f‘分页
		if args.form_deli {
			for pageNum := 0; pageNum <= args.end_page; pageNum++ {
				line, err := reader.ReadString('\f')
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				page_count++
				result += line
			}
		}else { //按行数分页
			for {
				// line, _, err := reader.ReadLine()
				line, err := reader.ReadString('\n')
				if err != io.EOF && err != nil {
					fmt.Println(err)
					os.Exit(1)
				}
				if err == io.EOF {
					break
				}
				line_count++
				if(line_count > args.page_len){
					page_count++
					line_count = 1
				}
				if(page_count >= args.start_page && page_count <= args.end_page){
					result += line
				}
			}
		}
	} else {
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			line := scanner.Text()
			line += "\n"
			line_count++
			if(line_count > args.page_len){
				page_count++
				line_count = 1
			}
			if page_count >= args.start_page && page_count <= args.end_page {
				result += line
			}
		}
	}
	if(page_count < args.start_page || page_count < args.end_page){
		fmt.Fprintln(os.Stderr, "\nError! Page number exceed the total number of the pages\n\n")
		flag.Usage()
		os.Exit(1)
	}
	print_write(args, string(result), stdin)

	if args.print_dest != "" {
		stdin.Close()
		cmd.Stdout = os.Stdout
		cmd.Run()
	}
}


func print_write(args *selpg_args, line string, stdin io.WriteCloser) {
	if args.print_dest != "" {
		stdin.Write([]byte(line + "\n"))
	} else {
		fmt.Println(line)
	}
}

/*================================= Usage() =======================*/

//show the usage of the command selpg
func Usage() {
	fmt.Printf("Usage:\n\n")
	fmt.Printf("\tselpg -s=Number -e=Number [optional_opts] [other_args|filename]\n\n")
	fmt.Printf("The arguments are:\n\n")
	fmt.Printf("\t-s=Number\tStart from Page <Number>.\n")
	fmt.Printf("\t-e=Number\tEnd to Page <Number>.\n")
	fmt.Printf("\t-l=Number\t[optional_opts]Specify the number of line per page.Default is 72.\n")
	fmt.Printf("\t-f\t\t[optional_opts]Specify that the pages are sperated by \\f.\n")
	fmt.Printf("\t-d=Destination\t\t[optional_opts]Specify the printer.\n")
	fmt.Printf("\t[other_args|filename]\t[other_args|filename]Read input from the file or some non-optional_opts.\n\n")
}
