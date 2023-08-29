package fofa

import (
	"fmt"
	"github.com/joho/godotenv"
	"log"
	"os"
	"testing"
	"zrWorker/core/slog"
)

func TestFofa(t *testing.T) {
	err := godotenv.Load()
	if err != nil {
		log.Fatal(err)
	}
	email := os.Getenv("FOFA_EMAIL")
	key := os.Getenv("FOFA_KEY")
	fmt.Printf("===============================")
	fmt.Printf(email)
	fmt.Printf(key)
	Init(email, key)
	size, r := this.Search("www.baidu.com")
	displayResponse(r)
	slog.Printf(slog.INFO, "本次搜索，返回结果总条数为：%d，此次返回条数为：%d", size, len(r))

	fmt.Println(size, r)
}
