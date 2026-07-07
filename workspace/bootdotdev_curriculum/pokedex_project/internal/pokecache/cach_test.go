package pokecache
import(
	"testing"
	"fmt"
	"time"
)
func TestAddGet(t *testing.T) {
	const interval = 5 * time.Second
	cases := []struct{
		key string
		val []byte
	} {
		{
			key:"https://example.com",
			val: []byte("testdata"),
		}, 
		{
			key:"https://example2.com",
			val: []byte("new testdata"),
		},
	}

	for idx, entery := range cases {
		t.Run(fmt.Sprintf("test %d for %s", idx + 1, entery.key), func(t *testing.T) {
			cache := NewCahse(interval)
			cache.Add(entery.key, entery.val)
			val, ok := cache.Get(entery.key)
			if !ok {
				t.Errorf("data didn't get in the cache")
				return
			}
			if string(val) != string(entery.val) {
				t.Errorf("the returned value is not the same as expected %s %s", string(val), string(entery.val))
				return
			}

		})
	}
}


func TestReapLoop(t *testing.T) {
	const baseTime = 5 * time.Millisecond
	const waitTime = baseTime + 5 * time.Millisecond
	cache := NewCahse(baseTime)
	cache.Add("https://example.com", []byte("data 1"))

	_, ok := cache.Get("https://example.com")
	if !ok {
		t.Error("expected to find key")
		return
	}
	time.Sleep(waitTime)
	_, ok = cache.Get("https://example.com")
	if ok {
		t.Error("expected to not find key")
		return
	}
}