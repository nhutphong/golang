package main

import (
	"fmt"
    "sync"
)


/*
 Điều 1: Khi khai báo một struct ở đó mutex quản lý quyền truy xuất vào một hay vài trường trong struct, 
 hãy đặt mutext lên trên cùng.
*/
var sum struct {
    sync.Mutex     // <-- this mutex protects
    i int          // <-- this integer underneath
}


/*
Điều 2: Giữ mutex lock càng ngắn càng tốt, hãy mở khóa ngay khi truy xuất xong dữ liệu. Nếu khóa mutex quá lâu, 
bạn sẽ bắt các tác vụ đồng bộ thực thi tuần tự nối đuôi nhau.
*/
// Cách dở

func doSomething(){
    mu.Lock()
    item := cache["myKey"]
    http.Get() // lệnh này chả cần nằm trong lệnh khóa Mutext
    mu.Unlock()
}
// Cách đúng
func doSomethingTwo(){
    mu.Lock()
    item := cache["myKey"]
    mu.Unlock() //Mở khóa luôn khi truy xuất bộ đệm xong
    http.Get() // Cho nó ra ngoài tốt hơn
}


/*
Điều 3: Trong một hàm có nhiều điểm trả về (multiple return paths), 
hãy sử dùng defer Mutex.Unlock() để đảm bảo bạn không quên mở khóa Mutex,
mà cũng không cần viết lệnh mở khóa trước mỗi điểm return. 
Như vậy defer giúp bạn khỏi bận tâm bảo trì mã nguồn
kể cả sau đây 3 tháng khi đồng nghiệp thêm một nhánh rẽ (new case), code chạy vẫn ổn.
*/

func doSomethingThree(){
    mu.Lock()
    defer mu.Unlock()
    err := ...
    if err != nil{
        //log error
        return  // <-- mutex sẽ mở khóa ở đây
    }
    err = ...
    if err != nil{
        //log error
        return  // <-- hoặc ở đây
    }
    return  // <-- và chắc chắn ở đây cũng được mở khóa
}

/*
Nhớ rằng defer chỉ hoạt động ở phạm vi hàm chứ không ở trong mỗi khối lệnh (block scope)

*/
func doSomethingFour(){
    for {
        mu.Lock()
        defer mu.Unlock()         
        
        // <-- defer không hoạt động ở phạm vi trong vòng lặp for
     }
   // <-- defer hoạt động khi hàm thoát thôi
}
// Code trên sẽ bị dead lock khi Mutex không được unlock đầy đủ



type DataStore struct {
	sync.Mutex // ← this mutex protects the cache below
	cache map[string]string
}
func New() *DataStore{
	return &DataStore{
		cache: make(map[string]string),
	}
}

func (ds *DataStore) set(key string, value string) {
	ds.Lock()
	defer ds.Unlock()
	ds.cache[key] = value
}
func (ds *DataStore) get(key string) string {
	ds.Lock()
	defer ds.Unlock()
		if ds.count() > 0 { <-- count() also takes a lock!
		item := ds.cache[key]
		return item
		}
	return “”
}

func (ds *DataStore) count() int {
	ds.Lock()
	defer ds.Unlock()
	return len(ds.cache)
}

func main() {
/* Running this below will deadlock because the get() method will       take a lock and will call the count() method which will also take a  lock before the set() method unlocks()
*/
	store := New()
	store.set(“Go”, “Lang”)
	result := store.get(“Go”)
	fmt.Println(result)
}