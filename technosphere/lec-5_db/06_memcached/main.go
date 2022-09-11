package main

import (
	"fmt"
	"time"
	"github.com/bradfitz/gomemcache/memcache"
)

var memcacheClient *memcache.Client

func getRecord(mkey string) *memcache.Item {
	println("get", mkey)
	// получаем одиночную запись
	item, err := memcacheClient.Get(mkey)
	// если записи нету, то для этого есть специальная ошибка, её надо обрабатывать отдеьно, это почти штатная ситуация, а не что-то страшное
	if err == memcache.ErrCacheMiss {
		fmt.Println("Record not found in memcache")
		return nil
	} else if err != nil {
		PanicOnErr(err)
	}
	return item
}


func main() {
	MemcachedAddresses := []string{"127.0.0.1:11211"}
	memcacheClient = memcache.New(MemcachedAddresses...)

	mkey := "record_21"

	item := getRecord(mkey)
	fmt.Printf("first get %+v\n", item)

	ttl := 5

	// Set устанавливает значение, не зависимо от того, была там запись или нет
	err := memcacheClient.Set(&memcache.Item{
		Key: mkey,
		Value: []byte("1"),
		Expiration: int32(ttl),
	})
	PanicOnErr(err)

	time.Sleep(time.Microsecond)

	item = getRecord(mkey)

	// second get &{Key:record_21 Value:[49] Flags:0 Expiration:0 casid:10977}
	// Value:[49] возвращается слайс байт!
	fmt.Printf("second get  %+v\n", item)

	// Add добавляет запись если её ещё не было
	err = memcacheClient.Add(&memcache.Item{
		Key: mkey,
		Value: []byte("2"),
		Expiration: int32(ttl),
	})

	// если запись не была добавлена, то вернётся соотвтетсувющая ошибка
	if err == memcache.ErrNotStored {
		fmt.Println("Record not stored")
	} else if err != nil {
		PanicOnErr(err)
	}

	item = getRecord(mkey)
	fmt.Printf("third get %+v\n", item)
}

//PanicOnErr panics on error
func PanicOnErr(err error) {
	if err != nil {
		panic(err)
	}
}