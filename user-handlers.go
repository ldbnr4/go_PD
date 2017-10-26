package main

import (
	"encoding/json"
	"io"
	//"fmt"
	"net/http"
	"os"

	"github.com/gorilla/mux"
)

// AddUserMsg ...
type AddUserMsg struct {
	Username string
	Password string
	Email    string
	Nickname string
}

//AddUserResp ...
type AddUserResp struct {
	Error CreateUserError
	ID    string
}

//DelUserMsg ...
type DelUserMsg struct {
	ID       string
	Password string
}

//LoginMsg ...
type LoginMsg struct {
	Username string
	Password string
}

//GetFriendsResponse ...
type GetFriendsResponse struct {
	FriendReqs []SimpleUser
	Friends    []SimpleUser
}

//SimpleUser ...
type SimpleUser struct {
	Name string
	ID   string
}

//UserCreate ...
func UserCreate(w http.ResponseWriter, r *http.Request) {
	r.Header.Set("Access-Control-Allow-Origin", "*")

	msg := new(AddUserMsg)
	FillStruct(r, msg)

	insertResp := InsertUser(*msg)
	userDirLocation := PrjDir + insertResp.ID
	ifErr(os.MkdirAll(userDirLocation, os.ModePerm))
	ifErr(copyFileContents(PrjDir+"profile.png", userDirLocation+"/"+insertResp.ID))

	// w.WriteHeader(http.StatusCreated)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	ifErr(json.NewEncoder(w).Encode(insertResp))
}

//UserDelete ...
func UserDelete(w http.ResponseWriter, r *http.Request) {
	msg := new(DelUserMsg)
	FillStruct(r, msg)
	RemoveUser(*msg)
	ifErr(json.NewEncoder(w).Encode("Completed"))
}

//Login ...
func Login(w http.ResponseWriter, r *http.Request) {
	msg := new(LoginMsg)
	FillStruct(r, msg)
	w.Header().Set("Access-Control-Allow-Origin", "*")
	ifErr(json.NewEncoder(w).Encode(RetrieveUser(*msg)))
}

//ProfPic ...
func ProfPic(w http.ResponseWriter, r *http.Request) {
	UID := mux.Vars(r)["UID"]

	f, err := os.Open(PrjDir + UID + "/" + UID)
	ifErr(err)
	io.Copy(w, f)
	defer f.Close()
}

//GetFriends ...
func GetFriends(w http.ResponseWriter, r *http.Request) {
	UID := mux.Vars(r)["UID"]
	ifErr(json.NewEncoder(w).Encode(GetFriendsResponse{GetFriendReqs(UID), GetFriendsMgo(UID)}))
}

// elseif(!empty($_GET['USER_SEARCH'])){
// 	$nickname = $_GET['nickname'];
// 	$regex = new MongoDB\BSON\Regex ("^$nickname");
// 	echo json_encode(array("err"=>false,"result"=>$collection->find(array("nickname"=> $regex),array("projection"=>array("nickname"=>1)))->toArray()));
// 	exit(0);
// }

// elseif(!empty($_GET['PROF_INFO'])){
// 	$nickname = $_GET['nickname'];
// 	echo json_encode(array("err"=>false,"result"=>$collection->findOne(array("nickname"=> $nickname),array("projection"=>array("_id"=>0,"joined"=>1)))));
// 	exit(0);
// }

// elseif(!empty($_POST['ACPT_REQ'])){
// 	$uid = new MongoDB\BSON\ObjectId($_POST['UID']);
// 	$friendId = new MongoDB\BSON\ObjectId($_POST['FRND_ID']);
// 	$collection->updateOne(
// 		array( "_id" => $friendId),
// 		array( '$pull'=> array("friendReqs"=>$uid))
// 	);
// 	$result = $collection->updateOne(
// 		array( "_id" => $uid),
// 		array( '$pull'=> array("friendReqs"=>$friendId))
// 	);
// 	if($result->getModifiedCount() != 1){
// 		echo json_encode(array("err"=>true, "msg" => "Failed to remove from reqs list."));
// 		exit(0);
// 	}
// 	else{
// 		$result = $collection->updateOne(
// 			array( "_id" => $friendId),
// 			array( '$addToSet'=> array("friends"=>$uid))
// 		);
// 		if($result->getModifiedCount() != 1){
// 			echo json_encode(array("err"=>true, "msg" => "Failed to add id to friends list."));
// 			exit(0);
// 		}
// 		$result = $collection->updateOne(
// 			array( "_id" => $uid),
// 			array( '$addToSet'=> array("friends"=>$friendId))
// 		);
// 		if($result->getModifiedCount() != 1){
// 			echo json_encode(array("err"=>true, "msg" => "Failed to add id to friends list."));
// 			exit(0);
// 		}
// 		echo json_encode(array("err"=>false));
// 		exit(0);
// 	}
// }

// elseif(!empty($_POST['DEC_REQ'])){
// 	$uid = new MongoDB\BSON\ObjectId($_POST['UID']);
// 	$friendId = new MongoDB\BSON\ObjectId($_POST['FRND_ID']);
// 	$result = $collection->updateOne(
// 		array( "_id" => $uid),
// 		array( '$pull'=> array("friendReqs"=>$friendId))
// 	);
// 	if($result->getModifiedCount() != 1){
// 		echo json_encode(array("err"=>true, "msg" => "Failed to delete album."));
// 		exit(0);
// 	}
// 	else{
// 		echo json_encode(array("err"=>false));
// 		exit(0);
// 	}
// }
