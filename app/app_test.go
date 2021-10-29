package app

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"os"
	"testing"

	"goweb/web22-1/model"

	"github.com/stretchr/testify/assert"
)

func TestMembers(t *testing.T) { //code Refactoring을 위한 테스팅 페이지 코딩 시작
	os.Remove("./test.db")
	assert := assert.New(t)
	ah := MakeHandler("./test.db") //테스팅 서버 구축을 위한 MakeHandler()
	defer ah.Close()

	ts := httptest.NewServer(ah)
	defer ts.Close()
	resp, err := http.PostForm(ts.URL+"/members", url.Values{"id": {"testmember"}, "name": {"test member"}}) // add는 Post전송인데, addMemberHandler로 데이터전송시 r.FormValue로 받기 때문에 POST가 아니라 PostForm메소드
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	var member model.Member                          //아래 JSON포맷 값을 읽어오는데 Decode할 member 개체가 필요하기 떄문에 값을 채워준다
	err = json.NewDecoder(resp.Body).Decode(&member) //addMemberHandler에서 JSON포맷으로 보내주는 member를 읽어온다.
	assert.NoError(err)
	assert.Equal(member.Name, "test member")
	id1 := member.ID                                                                                          //서버가 추가한 ID
	resp, err = http.PostForm(ts.URL+"/members", url.Values{"id": {"testmember2"}, "name": {"test member2"}}) // addMemberHandler로 데이터전송시 r.FormValue로 받기 때문에 POST가 아니라 PostForm메소드
	assert.NoError(err)
	assert.Equal(http.StatusCreated, resp.StatusCode)
	//var member Member									//아래 JSON포맷 값을 읽어오는데 Decode할 member 개체가 필요하기 떄문에 값을 채워준다
	err = json.NewDecoder(resp.Body).Decode(&member)
	assert.NoError(err)
	assert.Equal(member.Name, "test member2")
	id2 := member.ID
	//Get전송 테스트시작부분 / complete-member test
	resp, err = http.Get(ts.URL + "/members")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	members := []*model.Member{} //선언대입문으로 *Member{}(인스턴스)를 members에 넣어준다
	err = json.NewDecoder(resp.Body).Decode(&members)
	assert.NoError(err)
	assert.Equal(len(members), 2)
	for _, t := range members { // 첫 번째 인자가 index 인데 무시 _, 두 번째 인자가 값 t.ID가 나온다
		if t.ID == id1 {
			assert.Equal("test member", t.Name)
		} else if t.ID == id2 {
			assert.Equal("test member2", t.Name)
		} else {
			assert.Error(fmt.Errorf("testID should be id1 or id2"))
		}
	}

	//DELETE 메소드 만들어서 테스트
	rep, _ := http.NewRequest("DELETE", ts.URL+"/members/"+id1, nil)
	resp, err = http.DefaultClient.Do(rep)
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	//Get 전송 테스트
	resp, err = http.Get(ts.URL + "/members")
	assert.NoError(err)
	assert.Equal(http.StatusOK, resp.StatusCode)
	members = []*model.Member{} // 대입문으로 바꿔준다
	err = json.NewDecoder(resp.Body).Decode(&members)
	assert.NoError(err)
	assert.Equal(len(members), 1) //사이즈가 1로 줄어야 한다.
	for _, t := range members {   // member list 읽은 것이
		assert.Equal(t.ID, id2) // t.ID는 id2와 같아야 한다.
	}
}
